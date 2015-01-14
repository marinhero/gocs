/*
** my_ftp.go
** Author: Marin Alcaraz
** Mail   <marin.alcaraz@gmail.com>
** Started on  Mon Jan 12 17:22:29 2015 Marin Alcaraz
** Last update Wed Jan 14 17:32:12 2015 Marin Alcaraz
 */

package main

import (
	"fmt"
	"os"
)

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
		err = interact(c)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
