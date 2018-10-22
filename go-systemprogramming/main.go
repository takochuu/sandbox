package main

import (
	"io"
	"net"
	"os"

	"github.com/acidlemon/go-dumper"
)

func main() {
	netConn()
	multiWrite()
	gzip()
}

func netConn() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))
	io.Copy(os.Stdout, conn)
}

func multiWrite() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout)
	dump.Dump(writer)
	io.WriteString(writer, "io.MultiWriter example\n")
}

func gzip() {
	file, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(file)
	writer.Header.Name = "test.txt"
	writer.Write([]byte("gzip.Writer example\n"))
	writer.Close()
}
