package tasks

import (
	"bytes"
	"io"
)

func GetLinesChannel(o io.ReadCloser) <-chan string {
	ch := make(chan string)
	str := ""

	go func() {
		defer close(ch)
		defer o.Close()
		for {
			byts := make([]byte, 8)
			_, err := o.Read(byts)
			if err != nil {
				//ch <- fmt.Sprintf("%s", err)
				//log.Fatal(err)
				break
			}
			if i := bytes.IndexByte(byts, 0); i != -1 {
				byts = byts[:i]
			}
			if i := bytes.IndexByte(byts, '\n'); i != -1 {
				str += string(byts[:i+1])
				ch <- str
				str = string(byts[i+1:])

			} else {
				str += string(byts)
			}
		}
		if str != "" {
			ch <- str
		}
	}()
	return ch
}
