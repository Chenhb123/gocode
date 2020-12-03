package main

import (
	"log"
	"net"
	"time"
)

func main() {
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	conn.Close()
	log.Println("close ok")

	var buf = make([]byte, 32)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("read error:", err)
	} else {
		log.Printf("read % bytes, content is %s\n", n, string(buf[:n]))
	}

	n, err = conn.Write(buf)
	if err != nil {
		log.Println("write error:", err)
	} else {
		log.Printf("write % bytes, content is %s\n", n, string(buf[:n]))
	}

	time.Sleep(time.Second * 1000)
}

// func main() {
// 	log.Println("begin dial...")
// 	conn, err := net.Dial("tcp", ":8888")
// 	if err != nil {
// 		log.Println("dial error:", err)
// 		return
// 	}
// 	defer conn.Close()
// 	log.Println("dial ok")
// 	data := make([]byte, 65536)
// 	var total int
// 	for {
// 		conn.SetWriteDeadline(time.Now().Add(time.Microsecond * 5))
// 		n, err := conn.Write(data)
// 		if err != nil {
// 			total += n
// 			log.Printf("write %d bytes, error:%s\n", n, err)
// 			break
// 		}
// 		total += n
// 		log.Printf("write %d bytes this time, %d bytes in total\n", n, total)
// 	}

// 	log.Printf("write %d bytes in total\n", total)
// 	time.Sleep(time.Second * 10000)
// }

// func main() {
// 	log.Println("begin dial...")
// 	conn, err := net.Dial("tcp", ":8888")
// 	if err != nil {
// 		log.Println("dial error:", err)
// 		return
// 	}
// 	defer conn.Close()
// 	log.Println("dial ok")

// 	data := make([]byte, 65536)
// 	conn.Write(data)

// 	time.Sleep(time.Second * 10000)
// }

// func main() {
// 	// if len(os.Args) <= 1 {
// 	// 	fmt.Println("usage: go run client2.go YOUR_CONTENT")
// 	// 	return
// 	// }
// 	log.Println("begin dial...")
// 	conn, err := net.Dial("tcp", ":8888")
// 	if err != nil {
// 		log.Println("dial error:", err)
// 		return
// 	}
// 	defer conn.Close()
// 	log.Println("dial ok")

// 	time.Sleep(time.Second * 2)
// 	// data := os.Args[1]
// 	// conn.Write([]byte(data))
// 	// conn.Close()
// 	data := make([]byte, 65535)
// 	conn.Write(data)

// 	// time.Sleep(time.Second * 10000)
// }
