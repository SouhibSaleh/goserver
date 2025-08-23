package main

import (
	"github.com/SouhibSaleh/goserver/tasks"
)

func main() {
	//tasks.ReadFromFile()

	//tasks.TestArea()

	//o, err := os.Open("test.txt")
	//if err != nil {
	//	return
	//}
	//lines := tasks.GetLinesChannel(o)
	//for line := range lines {
	//	fmt.Println(line)
	//}

	go tasks.TcpRead()
	//request, err := tasks.RequestFromReader(strings.NewReader("GET / HTTP/1.1\nHost: localhost:42069\nUser-Agent: curl/7.81.0\nAccept: */*\n\n"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(request)
	//fmt.Printf("%#v", request)
	for {
	}
}
