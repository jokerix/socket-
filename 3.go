package  main
import (
	"net"
	"fmt"
	"log"
	"bytes"
	"encoding/binary"
)
const (ConstHeader         = "msg>")
var HEART_MSG []byte = []byte("pong" )
func  main()  {
	for i:=8000;i<8010;i++{
		go listen(i)
	}
	select {}
}
func listen(port int) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(":%d start listen\n", port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("new connection from %v\n", conn.LocalAddr())
			go handlenConnection(conn)
		}
	}
}
func handlenConnection(conn net.Conn){
	fmt.Printf("start handle connection:%v\n",conn.LocalAddr())
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		c := string(buffer[:n])
		if c=="ping"{
			senderToClient(conn,HEART_MSG)
		}
		return
	}
}
func senderToClient(conn net.Conn, msg []byte)(count int,error error) {
	return conn.Write(Packet(msg))
}
func Packet(message []byte) []byte {
	return append(append([]byte(ConstHeader), IntToBytes(len(message))...), message...)
}
func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}