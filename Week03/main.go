package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// there are two http server
	s1 := http.Server{Addr: ":8000"}
	s2 := http.Server{Addr: ":8001"}

	// new a stop chan
	stop := make(chan struct{})

	// new a errgroup
	group, _ := errgroup.WithContext(context.Background())

	// start server 1
	group.Go(func() error {
		go func() {
			<-stop
			s1.Shutdown(context.Background())
		}()
		return s1.ListenAndServe()
	})

	// start server 2
	group.Go(func() error {
		go func() {
			<-stop
			s2.Shutdown(context.Background())
		}()
		return s2.ListenAndServe()
	})

	// wait
	go func() {
		group.Wait()
		close(stop)
	}()

	// listen signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGKILL)
	for {
		switch <-c {
		case syscall.SIGINT, syscall.SIGTERM:
			s1.Shutdown(context.Background())
			time.Sleep(time.Second)
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
