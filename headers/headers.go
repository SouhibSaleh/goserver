package headers

import (
	"bytes"
	"fmt"
	"strings"
)

type Headers struct {
	headers map[string]string
}

func (h *Headers) Set(key string, value string) {
	key = strings.ToLower(key)
	val, ok := h.headers[strings.ToLower(key)]
	if !ok {
		h.headers[key] = value
	} else {
		fmt.Println(key, fmt.Sprintf("%s,%s", val, value))
		h.headers[key] = fmt.Sprintf("%s,%s", val, value)
	}
}

func (h *Headers) Get(key string) string {
	return h.headers[strings.ToLower(key)]
}

func isToken(token string) bool {
	valid := strings.Join(
		strings.Split("!,#,$,%,&,',*,+,-,.,^,_,`,|,~", ","), "")

	flag := false
	for _, arg := range token {
		arg := string(arg)
		if arg >= "a" && arg <= "z" ||
			arg >= "A" && arg <= "Z" ||
			arg >= "0" && arg <= "9" ||
			strings.Contains(valid, arg) {
			flag = true
		}
		if flag == false {
			return false
		}
	}

	return true
}

var EOL = []byte("\r\n")

func headerParser(data []byte) (string, string, error) {
	splitedHeader := bytes.SplitN(data, []byte(":"), 2)
	if len(splitedHeader) != 2 {
		return "", "", fmt.Errorf("The Header Is Invalid")
	}
	name := bytes.TrimLeft(splitedHeader[0], " ")
	value := bytes.TrimSpace(splitedHeader[1])

	if bytes.Contains(name, []byte(" ")) || len(name) == 0 {
		return "", "", fmt.Errorf("The Header name Is Invalid")
	}
	return string(name), string(value), nil
}

func (h Headers) Parse(data []byte) (n int, done bool, err error) {
	n = 0
	done = true
	isToken("")
	for {
		indexEOL := bytes.Index(data[n:], EOL)
		if indexEOL == -1 {
			break
		}
		if indexEOL == 0 {
			done = false
			break
		}
		name, value, err := headerParser(data[n : n+indexEOL])
		if err != nil {
			return 0, false, err
		}
		n += indexEOL + len(EOL)
		if !isToken(name) {
			fmt.Println("InValid Header Name")
			break
		}
		h.Set(name, value)
	}

	return n, done, err
}
func NewHeaders() *Headers {
	return &Headers{
		headers: make(map[string]string)}
}
