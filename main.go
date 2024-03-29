package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/The-Flash/http_server_from_scratch/simplenet"
)

func main() {
	ipFlag := flag.String("ip_addr", "127.0.0.1", "The IP address to use")
	portFlag := flag.Int("port", 8080, "The port to use.")
	flag.Parse()

	ip := simplenet.ParseIP(*ipFlag)
	fmt.Println(ip)
	port := *portFlag
	socket, err := simplenet.NewNetSocket(ip, port)
	if err != nil {
		panic(err)
	}
	defer socket.Close()
	log.Print("===============")
	log.Print("Server Started!")
	log.Print("===============")
	log.Print()
	log.Printf("addr: http://%s:%d", ip, port)

	for {
		rw, e := socket.Accept()
		log.Print()
		log.Print()
		log.Printf("Incoming connection")
		if e != nil {
			panic(e)
		}
		// Read request
		log.Print("Reading request")
		req, err := simplenet.ParseRequest(rw)
		log.Print("request: ", req)
		if err != nil {
			panic(err)
		}

		// Write response
		log.Print("Writing response")
		simplenet.WriteString(rw, "HTTP/1.1 200 OK\r\n"+
			"Content-Type: text/html; charset=utf-8\r\n"+
			"Content-Length: 20\r\n"+
			"\r\n"+
			"<h1>hello world</h1>")
		if err != nil {
			log.Print(err.Error())
			continue
		}
	}
}
