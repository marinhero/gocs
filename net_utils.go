/*
** net_utils.go
** Author: Marin Alcaraz
** Mail   <marin.alcaraz@gmail.com>
** Started on  Wed Jan 14 12:20:17 2015 Marin Alcaraz
** Last update Wed Jan 14 17:54:26 2015 Marin Alcaraz
 */

package main

import (
	"fmt"
	"net"
)

func connect(s string, p string) (con net.Conn, err error) {
	connectionString := s + ":" + string(p)
	conn, err := net.Dial("tcp", connectionString)
	if err != nil {
		return nil, fmt.Errorf("Dial error: %s", err)
	}
	return conn, nil
}

func interact(con net.Conn) error {
	var request string
	var err error

	readed := 0
	response := make([]byte, 8)

	for {
		//Read from server
		readed, err = con.Read(response)
		fmt.Print(string(response))
		if readed != len(response) {
			//Read request from keyboard
			fmt.Print("ftp> ")
			_, err = fmt.Scanf("%s", &request)
			if request == "quit" {
				break
			}
			//Send to server
			_, err = con.Write([]byte(request))
			//Valid terminated string
			_, err = con.Write([]byte("\r\n"))
		}
		if err != nil {
			return fmt.Errorf("Interact error: %s", err)
		}
		response = make([]byte, 8)
	}
	defer fmt.Println("Connection terminated")
	return con.Close()
}
