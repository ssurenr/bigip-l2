package main

import (
	"encoding/json"
	"fmt"
	"github.com/ssurenr/bigip-l2/pkg/cccl"
	"io/ioutil"
	"os"
)

func main() {
	var net cccl.CcclNet
	data := readFile("samples/arp.json")
	err := json.Unmarshal(data, &net)
	if err != nil {
		fmt.Println("Error decoding JSON", err)
	}

	formatString := "%-20v %-15v %-v\n"
	fmt.Printf(formatString, "Name", "IP Address", "Mac Address")
	for _, entry := range net.Arps {
		fmt.Printf(formatString, entry.Name, entry.IpAddress, entry.MacAddress)
	}
}

// returns file content in byte array
func readFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error Opening file", err)
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error Reading content:", err)
	}

	return content
}
