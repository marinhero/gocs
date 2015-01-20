/*
** server.go
** Author: Marin Alcaraz
** Mail   <marin.alcaraz@gmail.com>
** Started on  Thu Jan 15 12:08:34 2015 Marin Alcaraz
** Last update Tue Jan 20 18:29:40 2015 Marin Alcaraz
 */

package main

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
)

//Bad function! but I wanted to test the switch :P
func cmdExec(command string) (string, error) {
	c := command[0]
	switch c {
	case 49:
		ecmd := "ls"
		cmd := exec.Command("/bin/" + ecmd)
		fmt.Printf("==> Executing: %s\n",
			strings.Join(cmd.Args, " "))
		output, err := cmd.Output()
		return string(output) + "\r", err
	case 50:
		ecmd := "who"
		cmd := exec.Command("/usr/bin/" + ecmd)
		fmt.Printf("==> Executing: %s\n",
			strings.Join(cmd.Args, " "))
		output, err := cmd.Output()
		return string(output) + "\r", err
	case 51:
		ecmd := "pwd"
		cmd := exec.Command("/bin/" + ecmd)
		fmt.Printf("==> Executing: %s\n",
			strings.Join(cmd.Args, " "))
		output, err := cmd.Output()
		return string(output) + "\r", err
	}
	return "Undefined command", fmt.Errorf("Undefined command")
}

func handleConnection(conn net.Conn) error {
	request := make([]byte, 8)
	banner := fmt.Sprintf("Welcome!\nAvailable commands {1: ls, 2: who, 3: pwd}\n")
	_, err := conn.Write([]byte(banner))
	if err != nil {
		return fmt.Errorf("Error handleConnection: %s", err)
	}
	for {
		_, err := conn.Read(request)
		output, err := cmdExec(string(request))
		if err != nil {
			_, err = conn.Write([]byte(
				fmt.Sprintf("==> Error: %s\n", err.Error())))
		} else {
			_, err = conn.Write([]byte(output))
		}
		if err != nil {
			return fmt.Errorf("Error handleConnection: %s", err)
		}
	}
	conn.Close()
	return nil
}

func initServer(port string) error {
	connectionString := fmt.Sprintf("127.0.0.1:%s", port)
	ln, err := net.Listen("tcp", connectionString)
	if err != nil {
		return fmt.Errorf("Error at initServer: %s", err)
	}
	for {
		conn, err := ln.Accept()
		fmt.Println("Host connected: ", ln.Addr())
		if err != nil {
			return fmt.Errorf("Error at initServer: %s", err)
		}
		go handleConnection(conn)
	}
}
