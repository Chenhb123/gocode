package main

import (
	"fmt"
	"log"
	"net"
)

// func handleConn(c net.Conn) {
// 	defer c.Close()
// 	for {
// 		// read from the connection
// 		// ... ...
// 		time.Sleep(5 * time.Second)
// 		// var buf = make([]byte, 10)
// 		var buf = make([]byte, 65535)
// 		log.Println("start to read from conn")
// 		c.SetReadDeadline(time.Now().Add(time.Microsecond * 10))
// 		n, err := c.Read(buf)
// 		if err != nil {
// 			log.Println("conn read error:", err)
// 			return
// 		}
// 		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))

// 		// write to the connection
// 		//... ...
// 	}
// }

// func handleConn(c net.Conn) {
// 	defer c.Close()
// 	for {
// 		// read from the connection
// 		time.Sleep(10 * time.Second)
// 		var buf = make([]byte, 65536)
// 		log.Println("start to read from conn")
// 		c.SetReadDeadline(time.Now().Add(time.Microsecond * 10))
// 		n, err := c.Read(buf)
// 		if err != nil {
// 			log.Printf("conn read %d bytes,  error: %s", n, err)
// 			if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
// 				continue
// 			}
// 			return
// 		}
// 		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
// 	}
// }

// func handleConn(c net.Conn) {
// 	defer c.Close()
// 	time.Sleep(time.Second * 10)
// 	for {
// 		// read from the connection
// 		time.Sleep(5 * time.Second)
// 		var buf = make([]byte, 60000)
// 		log.Println("start to read from conn")
// 		n, err := c.Read(buf)
// 		if err != nil {
// 			log.Printf("conn read %d bytes,  error: %s", n, err)
// 			if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
// 				continue
// 			}
// 		}

// 		// log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
// 		log.Printf("read %d bytes\n", n)
// 	}
// }

func handleConn(c net.Conn) {
	defer c.Close()

	// read from the connection
	var buf = make([]byte, 10)
	log.Println("start to read from conn")
	n, err := c.Read(buf)
	if err != nil {
		log.Println("conn read error:", err)
	} else {
		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}

	n, err = c.Write(buf)
	if err != nil {
		log.Println("conn write error:", err)
	} else {
		log.Printf("write %d bytes, content is %s\n", n, string(buf[:n]))
	}
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		// start a new goroutine to handle
		// the new connection.
		go handleConn(c)
	}
}
