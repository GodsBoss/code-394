package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/GodsBoss/code-394/pkg/code394"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: code394 <file>")
		os.Exit(0)
	}
	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file '%s': %+v\n", filename, err)
		os.Exit(1)
	}
	var problem code394.Problem
	err = json.Unmarshal(data, &problem)
	if err != nil {
		fmt.Printf("Error unmarshaling from JSON: %+v\n", err)
		os.Exit(1)
	}
	solution := problem.Solve()
	if solution == nil {
		fmt.Println("No solution found.")
	}
	fmt.Printf("Solution: %s\n", strings.Join(solution, " "))
}
