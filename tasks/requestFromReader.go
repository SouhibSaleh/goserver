package tasks

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type ParsingState string

const (
	doneState ParsingState = "DONE"
	initState ParsingState = "INIT"
)

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}
type Request struct {
	RequestLine
	ParsingState
}

func (r *Request) done() bool {
	return r.ParsingState == doneState
}
func (r *Request) parse(data []byte) (int, error) {

	fmt.Println(data)
	switch r.ParsingState {
	case initState:
		for i, arg := range data {
			if arg == 10 {
				r.ParsingState = doneState
				arg, _ := parseRequestLine(string(data[:i]))
				r.RequestLine = *arg
				return i + 1, nil
			}
		}
	case doneState:
		return 0, nil
	}
	return 0, nil
}

func parseRequestLine(requestLineString string) (r *RequestLine, e error) {
	requestLineString = strings.ReplaceAll(requestLineString, "\r", "")
	requestLineAtt := strings.Split(requestLineString, " ")

	if len(requestLineAtt) != 3 {
		return nil, fmt.Errorf("the request line is wrong ")
	}

	httpVersionConvert := strings.Split(requestLineAtt[2], "/")
	if len(httpVersionConvert) != 2 {
		return nil, fmt.Errorf("Wrong http version formal")
	}

	requestLine := &RequestLine{
		HttpVersion:   httpVersionConvert[1],
		RequestTarget: requestLineAtt[1],
		Method:        requestLineAtt[0],
	}

	return requestLine, nil
}
func parseRequest(s string) (*Request, error) {

	args := strings.Split(s, "\n")
	requestLineString := args[0]
	requestLine, err := parseRequestLine(requestLineString)

	if err != nil {
		return nil, err
	}

	request := &Request{
		RequestLine: *requestLine,
	}

	fmt.Println(request)
	return request, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {

	request := &Request{}
	request.ParsingState = initState
	buffer := make([]byte, 1024)
	bufLen := 0

	for !request.done() {
		n, err := reader.Read(buffer[bufLen:])
		if err != nil {
			if errors.Is(err, io.EOF) {
				request.ParsingState = doneState
				break
			}

			return nil, err
		}
		bufLen += n
		readN, err := request.parse(buffer[:bufLen])

		if err != nil {
			return nil, err

		}
		copy(buffer, buffer[readN:bufLen])
		bufLen -= readN

	}

	return request, nil
}
