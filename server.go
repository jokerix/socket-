package main

import (
	"runtime"
	"net"
	"os"
	"bytes"
	"encoding/binary"
)

const (
	ConstHeader  = "msg>"
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS1= "10.10.10.22:9000"
	SERVER_ADDRESS2 = "10.10.10.22:9001"
	SERVER_ADDRESS3= "10.10.10.22:9002"
	SERVER_ADDRESS4= "10.10.10.22:9003"
	SERVER_ADDRESS5= "10.10.10.22:9004"
	SERVER_ADDRESS6= "10.10.10.22:9005"
	SERVER_ADDRESS7= "10.10.10.22:9006"
	SERVER_ADDRESS8= "10.10.10.22:9007"
	SERVER_ADDRESS9= "10.10.10.22:9008"
	SERVER_ADDRESS10= "10.10.10.22:9090"
)
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag := make(chan bool, 1)
     	go StartSocket(SERVER_NETWORK,SERVER_ADDRESS1,flag)
		go StartSocket(SERVER_NETWORK,SERVER_ADDRESS2,flag)
		go StartSocket(SERVER_NETWORK,SERVER_ADDRESS3,flag)
		go StartSocket(SERVER_NETWORK,SERVER_ADDRESS4,flag)
		go StartSocket(SERVER_NETWORK,SERVER_ADDRESS5,flag)
		go StartSocket(SERVER_NETWORK,SERVER_ADDRESS6,flag)
		go StartSocket(SERVER_NETWORK,SERVER_ADDRESS7,flag)
		go StartSocket(SERVER_NETWORK,SERVER_ADDRESS8,flag)
		go StartSocket(SERVER_NETWORK,SERVER_ADDRESS9,flag)
		go StartSocket(SERVER_NETWORK,SERVER_ADDRESS10,flag)
	<-flag
}
	var HEART_MSG []byte = []byte("pong" )
		func StartSocket(servernetwork string, serveraddress string, flag chan bool) {
				netListen, err := net.Listen(servernetwork, serveraddress)
				CheckError(err)
				defer func() {
					netListen.Close()
						flag <- true
						}()
			for {
				conn, err := netListen.Accept()
				if err != nil {
				continue
				}
				go handleConnection(conn)
			}
}
 	func CheckError(err error) {
		if err != nil {
			os.Exit(1)
		}
	}
	func handleConnection(conn net.Conn) {
			buffer := make([]byte, 1024)
		for {
			n, err := conn.Read(buffer)
			if err != nil {
				return
			}
			c := string(buffer[:n])     //
			c=string(buffer[:4])       //无需这行
				if c=="ping"{       //可以使用c!=""{
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