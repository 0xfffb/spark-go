package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/hashicorp/yamux"
)

func clientStream(session *yamux.Session, name string) {

	stream, err := session.Open()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 3; i++ {
		n, err := stream.Write([]byte("hello " + name))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s %d bytes written\n", name, n)
		time.Sleep(time.Second)
	}
}

func main() {

	conn, err := net.Dial("tcp4", "localhost:3000")
	if err != nil {
		log.Fatalf("TCP dial: %s", err)
	}

	// Setup client side of yamux
	session, err := yamux.Client(conn, nil)
	if err != nil {
		log.Fatal(err)
	}

	go clientStream(session, "foo")
	go clientStream(session, "bar1")
	go clientStream(session, "bar2")
	go clientStream(session, "bar3")
	go clientStream(session, "bar4")
	go clientStream(session, "bar5")
	go clientStream(session, "bar6")
	clientStream(session, "zip")
}
