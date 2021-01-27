package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	conns   = new(sync.Map)
	message = make(chan pack)
)

type pack struct {
	sender string
	data   string
}

func signalFunc() error {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	<-c
	return errors.New("exit")
}

func writeFunc(ctx context.Context) func() error {
	return func() error {
		for {
			select {
			case <-ctx.Done():
				return errors.New("exit")
			case msg := <-message:
				go func(msg pack) {
					conns.Range(func(addr, value interface{}) bool {
						if addr.(string) != msg.sender {
							conn := value.(net.Conn)
							_, err := conn.Write([]byte(msg.data))
							if err != nil {
								log.Printf("send message to %s error: %v, with message: %s", addr, err, msg.data)
							}
						}
						return true
					})
				}(msg)
			}
		}
	}
}

func readFunc(ctx context.Context, conn net.Conn) func() error {
	return func() error {
		addr := conn.RemoteAddr().String()
		reader := bufio.NewReader(conn)

		for {
			select {
			case <-ctx.Done():
				return errors.New("exit")
			default:
				msg, err := reader.ReadString('\n')
				if err != nil {
					data := fmt.Sprintf("[%s] exit\n", addr)
					message <- pack{sender: addr, data: data}

					conns.Delete(addr)
					fmt.Print(data)
					return nil
				}
				data := fmt.Sprintf("[%s]: %s", addr, msg)
				fmt.Print(data)
				message <- pack{sender: addr, data: data}
			}
		}
	}
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	group, ctx := errgroup.WithContext(ctx)

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	group.Go(signalFunc)
	group.Go(writeFunc(ctx))
	group.Go(func() error {

		c := make(chan net.Conn)
		go func() {
			for {
				conn, err := listener.Accept()
				if err != nil {
					log.Printf("accept new conn failed: %v", err)
					continue
				}
				c <- conn
			}
		}()

		for {
			select {
			case <-ctx.Done():
				log.Println("server exit.")
				return errors.New("exit")
			case conn := <-c:
				addr := conn.RemoteAddr().String()
				conns.Store(addr, conn)
				group.Go(readFunc(ctx, conn))

				data := fmt.Sprintf("[%s] join\n", addr)
				message <- pack{sender: addr, data: data}

				fmt.Print(data)
			}
		}
	})

	log.Println("server run.")
	group.Wait()
}
