package tasks

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func TcpRead() {
	var sb strings.Builder
	ls, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	for {
		Conn, err := ls.Accept()
		o, err := os.OpenFile("test.txt", os.O_APPEND, 0644)
		if err != nil {
			panic(err)
		}
		lines := GetLinesChannel(Conn)
		for line := range lines {
			fmt.Print(line)
			sb.WriteString(line)
			o.WriteString(line)
		}

		ss := sb.String()
		request, err := RequestFromReader(strings.NewReader(ss))
		fmt.Println(request)
	}

}
