package main

import (
	"log"

	opcua "github.com/linimbus/go-open62541"
)

func displayNodeTree(cli *opcua.Client, prefix string, nodeTree *opcua.NodeTree) {

	value, err := cli.ReadNode(nodeTree.Node)
	if err == nil {
		log.Printf("%s NodeID: %s Index: %d Level: %d Value: %s\n", prefix, nodeTree.Node.NodeID, nodeTree.Node.Index, nodeTree.Level, value.ToString())
	} else {
		log.Printf("%s NodeID: %s Index: %d Level: %d \n", prefix, nodeTree.Node.NodeID, nodeTree.Node.Index, nodeTree.Level)
	}

	for _, subNode := range nodeTree.SubNodes {
		displayNodeTree(cli, prefix+"  ", subNode)
	}

}

func main() {
	cli, err := opcua.NewClient("opc.tcp://192.168.3.22:53530/OPCUA/SimulationServer")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer cli.Close()

	nodeTree, err := cli.BrowseNode()
	if err != nil {
		log.Fatal(err.Error())
	}

	displayNodeTree(cli, "", nodeTree)
}
