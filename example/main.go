package main

import (
	opcua "github.com/linimbus/go-open62541"
)

func main() {
	cli, _ := opcua.NewClient()
	cli.Close()

}
