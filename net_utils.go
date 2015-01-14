/*
** net_utils.go
** Author: Marin Alcaraz
** Mail   <marin.alcaraz@gmail.com>
** Started on  Wed Jan 14 12:20:17 2015 Marin Alcaraz
** Last update Wed Jan 14 12:25:46 2015 Marin Alcaraz
 */

package main

import (
	"fmt"
	"net"
)

func interact(con net.Conn) error {
	var message string
	fmt.Println("Conected")
	for {
		fmt.Print("$> ")
		_, err := fmt.Scanln(&message)
		if message == "quit" {
			break
		}
		_, err = con.Write([]byte(message))
		if err != nil {
			return fmt.Errorf("Write error: %s", err)
		}
	}
	defer fmt.Println("Connection terminated")
	return con.Close()
}
