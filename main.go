package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

//const node string = "[fe80::e046:9aff:fe4e:912e%wlp1s0]:33123"

//const node string = "[fe80::e046:9aff:fe4e:912e%enp2s0]:33123"

//const node string = "[fe80::1e8f:814e:9731:dec6%enp2s0]:33123"

const node string = "[::1]:33123"
const (
	dump = "dump\n"
)

func testConnection() {
	conn, err := net.Dial("tcp6", node)
	if err != nil {
		log.Println("node ", err)
		return
	}
	defer conn.Close()

	fmt.Fprintf(conn, dump)
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Println("no")
			break
		}
		log.Println(message)
		if message == "ok" || message == "bad" || message == "no" {
			break
		}
	}
}

func main() {
	go testConnection()

}
