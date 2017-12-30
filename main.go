package main

import (
	"bufio"
	"log"
	"net"

	"github.com/pkg/errors"
)

const (
	Port = ":3333"
)

func Open(addr string) (*bufio.ReadWriter, error) {
	log.Println("Dial " + addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, errors.Wrap(err, "Dialing of addr "+addr+" failed")
	}
	return bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)), nil
}

func Listen() (net.Listener, error) {
	toListen, err := net.Listen("tcp", Port)
	if err != nil {
		return nil, err
	}
	log.Println("Listening on ", toListen.Addr().String())
	for {
		log.Println("Accepting a connection request.")
		conn, err := toListen.Accept()
		if err != nil {
			log.Println("Failed to accept incoming connection request:", err)
			continue
		}
		log.Println("Handling incoming message")
		go HandleMessage(conn)
	}

}

func HandleMessage(conn net.Conn) {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	defer conn.Close()
	for {
		log.Print("Receive text: ")
		text, err := rw.ReadString('\n')
		if err != nil {
			// TODO Add error handling
		}
		// TODO complete reading tutorial for handleMessages
	}
}
