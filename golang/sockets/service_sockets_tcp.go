package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":8080"
	add, err := net.ResolveTCPAddr("tcp4", service)
	checkErr(err)
	lis, err := net.ListenTCP("tcp", add)

	/**
	目标相同，使用不同的方法
	service := ":8080"
	lis,err := net.Listen("tcp",service)
	**/

	checkErr(err)

	for {
		con, err := lis.Accept()
		if err != nil {
			continue
		}

		//多线程
		go timereturn(con)
	}
}

func timereturn(con net.Conn) {
	defer con.Close()
	daytime := time.Now().String()
	con.Write([]byte(daytime))
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}
