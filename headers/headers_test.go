package headers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHeadersParsing1(t *testing.T) {
	header := NewHeaders()
	data := []byte("Host: localhost:42069\r\n\r\n")
	n, done, err := header.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, header)
	assert.Equal(t, "localhost:42069", header.Get("Host"))
	assert.Equal(t, 23, n)
	assert.False(t, done)

}

func TestHeadersParsing2(t *testing.T) {
	header := NewHeaders()
	data := []byte("       Host : localhost:42069       \r\n\r\n")
	n, done, err := header.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)
}

func TestHeadersParsing3(t *testing.T) {
	header := NewHeaders()
	data := []byte("       Host: localhost:42069       \r\nMEOW:  Meow_mewe\r\n")
	header.Parse(data)
	assert.Equal(t, header.Get("MEOW"), "Meow_mewe")
}
func TestHeadersParsing4(t *testing.T) {
	header := NewHeaders()
	data := []byte("Host: localhost:42069\r\nMEOW:  Meow_mewe\r\nHost: localhost:4269\r\n\r\n")

	header.Parse(data)
	assert.Equal(t, header.Get("MEOW"), "Meow_mewe")
	fmt.Println(header.Get("Host"))
}
