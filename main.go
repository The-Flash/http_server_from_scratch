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
		_, e := socket.Accept()
		log.Print()
		log.Print()
		log.Printf("Incoming connection")
		if e != nil {
			panic(e)
		}
	}
}
