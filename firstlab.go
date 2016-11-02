package main

import (
	"fmt"
	"net"
	"os"
	"io/ioutil"
)

const (
	CONNECTION_HOST = "10.62.0.189:8000"
	CONNECTION_TYPE = "tcp4"
)

func main() {
	conn, err := net.Dial(CONNECTION_TYPE, CONNECTION_HOST)
	checkError(err)

	request := []byte(fmt.Sprintf("GET /echo.php?message=testing HTTP/1.0\r\nHost: %s\r\n\r\n", CONNECTION_HOST))

	_, err = conn.Write(request)
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
