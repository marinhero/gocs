/*
** my_ftp.go
** Author: Marin Alcaraz
** Mail   <marin.alcaraz@gmail.com>
** Started on  Mon Jan 12 17:22:29 2015 Marin Alcaraz
** Last update Wed Jan 14 12:34:27 2015 Marin Alcaraz
 */

package main

import (
	"fmt"
	"net"
	"os"
)

func connect(s string, p string) (con net.Conn, err error) {
	connectionString := s + ":" + string(p)
	fmt.Println("Attempting to connect...")
	conn, err := net.Dial("tcp", connectionString)
	if err != nil {
		return nil, fmt.Errorf("Dial error: %s", err)
	}
	return conn, nil
}

func main() {
	arguments := os.Args[1:]
	if len(arguments) != 2 {
		fmt.Println("Usage: goftp server port")
	} else {
		server := arguments[0]
		port := arguments[1]
		c, err := connect(server, port)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(interact(c))
	}
}
