package tasks

import "io"

func GetLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)
	return ch
}
