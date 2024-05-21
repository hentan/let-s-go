package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	io.Copy(os.Stdout, conn)
	fmt.Println("\nDone")
}
