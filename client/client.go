package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	fmt.Println("client start.")

	err := start()
	if err != nil {
		fmt.Errorf("%s", err)
	}

	fmt.Println("client end.")

}

func start() error {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Fprintf(conn, "Hello, Socket Connection !")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return err
	}
	fmt.Println(status)

	return nil
}
