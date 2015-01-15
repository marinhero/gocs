# Goftp
A small FTP implementation written during my Hacker School Winter 2015 batch.

## Overview

This is a thin implementation of the FTP protocol, client and server side.
The goal of the project *wasn't* emulate the FTP protocol, instead I focussed on learning some caveats of the Go language.

## Execution

	./myftp -s 21 //Will launch the process in server mode
	./myftp 127.0.0.1 21 //Will launch the process in client mode

### Available commands

The following commands are available for execution

 * ls
 * pwd
 * who
