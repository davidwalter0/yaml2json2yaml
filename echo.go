package main

import (
	"io"
	"log"
	"net"
)

func echo(con net.Conn) {
	_, err := io.Copy(con, con)
	if err != nil {
		log.Print(err)
	}
	err = con.Close()
	if err != nil {
		log.Print(err)
	}
}

func main() {
	log.SetFlags(log.Lshortfile)
	// ln, err := net.Listen("tcp", ":9999")
	ln, err := net.Listen("tcp", ":7")
	if err != nil {
		log.Fatal(err)
	}
	for {
		con, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go echo(con)
	}
	err = ln.Close()
	if err != nil {
		log.Fatal(err)
	}
}

/*
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile)
	con, err := net.Dial("tcp", ":9999")
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(con, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	if tcpcon, ok := con.(*net.TCPConn); ok {
		tcpcon.CloseWrite()
	}
	_, err = io.Copy(os.Stdout, con)
	if err != nil {
		log.Fatal(err)
	}
	err = con.Close()
	if err != nil {
		log.Fatal(err)
	}
}
*/
