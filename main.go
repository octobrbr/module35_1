package main

import (
	"log"
	"net"
	"time"

	"github.com/jboursiquot/go-proverbs"
)

const addr = "0.0.0.0:12345"

const proto = "tcp4"

func main() {
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {

	for {
		time.Sleep(3 * time.Second)
		conn.Write([]byte(proverbs.Random().Saying + "\n"))
	}

}
