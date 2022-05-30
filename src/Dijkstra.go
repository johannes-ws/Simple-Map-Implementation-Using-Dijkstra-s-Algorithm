package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Graph struct {
	Nodes []string `json:"nodes"`
	Edges []Edge   `json:"edges"`
}

type Edge struct {
	Src    string `json:"src"`
	Dest   string `json:"dest"`
	Weight int    `json:"weight"`
}

func main() {
	fmt.Print("Masukkan nama file: ")
	var file string
	fmt.Scanln(&file)
	jsonFile, err := os.Open("C:\\Users\\johan\\OneDrive\\Documents\\GitHub\\Simple-Map-Implementation-Using-Dijkstra-s-Algorithm\\test\\" + file)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully Opened " + file)
		fmt.Println()
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	var Dijkstra Graph
	json.Unmarshal(byteValue, &Dijkstra)
	for i := 0; i < len(Dijkstra.Nodes); i++ {
		fmt.Println("Node ke-" + strconv.Itoa(i+1) + ": " + Dijkstra.Nodes[i])
	}
	fmt.Println()
	for i := 0; i < len(Dijkstra.Edges); i++ {
		fmt.Println("Edge ke-" + strconv.Itoa(i+1))
		fmt.Println("Src: " + Dijkstra.Edges[i].Src)
		fmt.Println("Dest: " + Dijkstra.Edges[i].Dest)
		fmt.Println("Weight: " + strconv.Itoa(Dijkstra.Edges[i].Weight))
		fmt.Println()
	}
}
