package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	_"./getchar"
)

import "azul3d.org/keyboard.v1"

type Command struct {
	body string
}

type Server struct {
	address string
	conn    net.Conn
}

func (s *Server) Connect() {
	conn, err := net.Dial("tcp", s.address)
	if err != nil {
		log.Fatal("Unable to connect: ", err)
	}
	s.conn = conn
}

func (s *Server) Send(c *Command) {
	fmt.Fprintf(s.conn, c.body+"\n")
	fmt.Println("Sent command to Server: ", c.body)
}

func NewServer(address string) Server {
	s := Server{
		address: address,
		conn:    nil,
	}

	s.Connect()

	return s
}

func main() {
	flag.Parse()
	//address := flag.Arg(0)
	//server := NewServer(address)

	for {

		watcher := keyboard.NewWatcher()
		// Query for the map containing information about all keys
		status := watcher.States()
		left := status[keyboard.ArrowLeft]
		if left == keyboard.Down {
			fmt.Println("Left Key")
		
		}
		esc := status[keyboard.Escape]
		if esc == keyboard.Down {
			os.Exit(0)		
 //else {
		}
	//		fmt.Println("NO")
		

		//ascii, keyCode, err := getchar.GetChar()
		// Exit on error
		///
		//if err != nil {
		//	log.Fatal("Error encountered: ", err)
		//		os.Exit(0)
		//	// Exit on user interrupt
		//} else if ascii == 3 || ascii == 4 || ascii == 26 || ascii == 27 {
		//	fmt.Println("Exiting...")
		//	os.Exit(0)
		//} else {
		//	// Handle known input
		//	if keyCode == 37 {
		//		command := Command{body: "left"} // TODO: Make Enums or something
		//		server.Send(&command)
		//	} else if keyCode == 38 {
		//		command := Command{body: "up"} // TODO: Make Enums or something
		//		server.Send(&command)
		//	} else if keyCode == 39 {
		//		command := Command{body: "right"} // TODO: Make Enums or something
		//		server.Send(&command)
		//	} else if keyCode == 40 {
		//		command := Command{body: "down"} // TODO: Make Enums or something
		//		server.Send(&command)
		//	}
		//}
	}
}
