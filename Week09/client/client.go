package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	go func() {
		reader := bufio.NewReader(conn)
		for {
			msg, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalln("connect server error")
			}
			fmt.Printf("\n@< %s", msg)
			fmt.Print("[input]:")
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	log.Println("welcome, you can exit with q!")

	for {
		fmt.Print("[input]:")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("read input error: %v\n", err)
			continue
		}

		if input == "q!\n" || input == "q!\r\n" {
			log.Fatalln("exit")
		}

		_, err = conn.Write([]byte(input))
		if err != nil {
			log.Printf("send message to server error: %v", err)
		}
		fmt.Printf("@> %s", input)
	}
}
