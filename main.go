package main

import (
	"fmt"
	"github.com/SouhibSaleh/goserver/tasks"
	"os"
)

func main() {
	//tasks.ReadFromFile()
	//tasks.TestArea()
	o, err := os.Open("test.txt")
	if err != nil {
		return
	}
	lines := tasks.GetLinesChannel(o)
	for line := range lines {
		fmt.Println(line)
	}
}
