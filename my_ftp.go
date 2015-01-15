/*
** my_ftp.go
** Author: Marin Alcaraz
** Mail   <marin.alcaraz@gmail.com>
** Started on  Mon Jan 12 17:22:29 2015 Marin Alcaraz
** Last update Thu Jan 15 12:14:33 2015 Marin Alcaraz
 */

package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) != 2 {
		fmt.Println("Usage: goftp [-s] [IP] PORT")
	} else {
		if arguments[0] == "-s" {
			port := arguments[1]
			err := initServer(port)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			server := arguments[0]
			port := arguments[1]
			err := initClient(server, port)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
