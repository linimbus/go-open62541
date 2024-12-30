package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	opcua "github.com/linimbus/go-open62541"
)

func displayNodeTree(cli *opcua.Client, prefix string, nodeTree *opcua.NodeTree) {

	nodes := make([]opcua.NodeInfo, 0)
	// for _, node := range nodeTree.SubNodes {
	// 	nodes = append(nodes, node.Node)
	// }

	nodes = append(nodes, nodeTree.Node)

	values, err := cli.ReadNodes(nodes)
	if err != nil {
		log.Fatal(err.Error())
	}

	value, err := cli.ReadNode(nodeTree.Node)
	if err == nil {
		if values[0].ToString() != value.ToString() {
			log.Printf("Error: %s %s\n", values[0].ToString(), value.ToString())
		}
		// log.Printf("%s NodeID: %s Index: %d Level: %d Value: %s\n", prefix, nodeTree.Node.NodeID, nodeTree.Node.Index, nodeTree.Level, value.ToString())
	}

	// } else {
	// 	log.Printf("%s NodeID: %s Index: %d Level: %d \n", prefix, nodeTree.Node.NodeID, nodeTree.Node.Index, nodeTree.Level)
	// }

	for _, subNode := range nodeTree.SubNodes {
		displayNodeTree(cli, prefix+"  ", subNode)
	}

}

func opcua_testing() {
	cli, err := opcua.NewClient("opc.tcp://192.168.3.22:53530/OPCUA/SimulationServer")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer cli.Close()

	for i := 0; i < 10000; i++ {

		nodeTree, err := cli.BrowseNode()
		if err != nil {
			log.Fatal(err.Error())
			break
		}

		displayNodeTree(cli, "", nodeTree)
	}
}

func main() {
	// for debug
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	opcua_testing()
	// wait for a long time

	log.Printf("wait for a long time\n")
	time.Sleep(time.Hour)
}
