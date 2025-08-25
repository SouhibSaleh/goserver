package tasks

import (
	"errors"
	"fmt"
	"github.com/SouhibSaleh/goserver/headers"
	"io"
	"strconv"
	"strings"
	"time"
)

type ParsingState string

const (
	initState    ParsingState = "INIT"
	headersState ParsingState = "HEADERS"
	bodyState    ParsingState = "BODY"
	doneState    ParsingState = "DONE"
	errState     ParsingState = "ERROR"
)

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}
type Request struct {
	RequestLine
	ParsingState
	*headers.Headers
	Body []byte
}

func NewRequest() *Request {
	return &Request{
		ParsingState: initState,
		Headers:      headers.NewHeaders(),
		Body:         make([]byte, 0),
	}

}

func (r *Request) done() bool {
	return r.ParsingState == doneState
}
func (r *Request) parse(data []byte) (int, error) {

	currendIndex := 0
	for {
		switch r.ParsingState {
		case initState:
			for i, arg := range data {
				if arg == 10 {
					r.ParsingState = headersState
					arg, err := parseRequestLine(string(data[:i]))
					if err != nil {
						r.ParsingState = errState
						return 0, err
					}
					r.RequestLine = *arg
					currendIndex = i + 1
					break
				}
			}
		case headersState:
			n, done, err := r.Headers.Parse(data[currendIndex:])
			if err != nil {
				r.ParsingState = errState
				return 0, err
			}
			if done {
				val := r.Headers.Get("content-length")
				currendIndex += 2
				if len(val) == 0 {
					r.ParsingState = doneState
					break
				}
				r.ParsingState = bodyState
				break
			}
			currendIndex += n
		case bodyState:
			contentBytes, err := strconv.Atoi(r.Headers.Get("content-length"))

			if err != nil {
				return 0, fmt.Errorf("Something wrong happened with content-length", err)
			}
			notFinishedByts := min(contentBytes-len(r.Body), len(data[currendIndex:]))

			fmt.Println(notFinishedByts)
			time.Sleep(time.Second)

			if notFinishedByts < 0 {
				r.ParsingState = errState
				return 0, fmt.Errorf("the number of actual bytes exceed the content-length")
			}

			r.Body = append(r.Body, data[currendIndex:currendIndex+notFinishedByts]...)
			currendIndex += notFinishedByts
			if len(r.Body) == contentBytes {
				r.ParsingState = doneState
				break
			} else if notFinishedByts == 0 {
				return 0, nil
			}

		case doneState:
			return 0, nil
		default:
			return 0, fmt.Errorf("Something wrong happened while parsing")
		}

	}
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

	return request, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	fmt.Println()
	request := NewRequest()

	buffer := make([]byte, 1024)
	bufLen := 0

	for !request.done() {
		n, err := reader.Read(buffer[bufLen:])

		if err != nil {
			if errors.Is(err, io.EOF) {
				if request.ParsingState != doneState {
					return nil, fmt.Errorf("something wrong happened")
				}
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
