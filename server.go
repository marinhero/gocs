/*
** server.go
** Author: Marin Alcaraz
** Mail   <marin.alcaraz@gmail.com>
** Started on  Thu Jan 15 12:08:34 2015 Marin Alcaraz
** Last update Thu Jan 15 12:40:05 2015 Marin Alcaraz
 */

package main

import (
	"fmt"
	"net"
)

func initServer(port string) error {
	connectionString := fmt.Sprintf("127.0.0.1:%s", port)
	ln, err := net.Listen("tcp", connectionString)
	if err != nil {
		return fmt.Errorf("Error at initServer: %s", err)
	}
	for {
		_, err := ln.Accept()
		fmt.Println("Host connected: ", ln.Addr())
		if err != nil {
			return fmt.Errorf("Error at initServer: %s", err)
		}
		//go handleConnection(conn)
	}
}
