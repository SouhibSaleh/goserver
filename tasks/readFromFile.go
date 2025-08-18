package tasks

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func ReadFromFile() {
	o, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)

	}
	defer o.Close()

	// we can use strings.Builder instead normal string
	// so we don't have to reallocate
	str := ""

	for {
		byts := make([]byte, 8)
		_, err := o.Read(byts)
		if err != nil {
			fmt.Println(err)
			break
		}

		if i := bytes.IndexByte(byts, '\n'); i != -1 {
			str += string(byts[:i+1])
			fmt.Print(str)
			str = string(byts[i+1:])

		} else {
			str += string(byts)
		}
	}
	if str != "" {
		fmt.Print(str)
	}

}
