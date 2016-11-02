package main

import (
	"fmt"
	"net"
	"os"
	"io/ioutil"
)

const (
	CONNECTION_HOST = "10.62.0.189"
	CONNECTION_TYPE = "tcp4"
)

func main() {
	var host = CONNECTION_HOST + ":" + port()
	conn, err := net.Dial(CONNECTION_TYPE, host)
	checkError(err)

	request := []byte(fmt.Sprintf("GET /echo.php?message=testing HTTP/1.0\r\nHost: %s\r\n\r\n", host))

	_, err = conn.Write(request)
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))
	os.Exit(0)
}

func port() string {
	if len(os.Args) > 1 && os.Args[1] != "" {
		return os.Args[1]
	}
	return "8000"
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
