package simplenet

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type request struct {
	method string // GET, POST, etc.
	header MIMEHeader
	body   []byte
	uri    string // The raw URI from the request
	proto  string // "HTTP/1.1"
}

func ParseRequest(c *netSocket) (*request, error) {
	b := bufio.NewReader(*c)
	tp := NewReader(b) // need replace
	req := new(request)

	// First line: parse "GET /index.html HTTP/1.0"
	var s string
	s, _ = tp.ReadLine() // need replace
	sp := strings.Split(s, " ")
	req.method, req.uri, req.proto = sp[0], sp[1], sp[2]

	// Parse headers
	mimeHeader, _ := tp.ReadMIMEHeader() // need replace
	fmt.Println(mimeHeader)
	req.header = mimeHeader

	// Parse body
	if req.method == "GET" || req.method == "HEAD" {
		return req, nil
	}
	if len(req.header["Content-Length"]) == 0 {
		return nil, errors.New("no content length")
	}
	length, err := strconv.Atoi(req.header["Content-Length"][0])
	if err != nil {
		return nil, err
	}
	body := make([]byte, length)
	if _, err = io.ReadFull(b, body); err != nil {
		return nil, err
	}
	req.body = body
	return req, nil
}
