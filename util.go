package main

import (
	"bufio"
	"net/url"
	"net/http"
	"crypto/tls"
)

// IsGzip returns true buffered Reader has the gzip magic.
func IsGzip(b *bufio.Reader) (bool, error) {
	return CheckBytes(b, []byte{0x1f, 0x8b})
}

// CheckBytes peeks at a buffered stream and checks if the first read bytes match.
func CheckBytes(b *bufio.Reader, buf []byte) (bool, error) {

	m, err := b.Peek(len(buf))
	if err != nil {
		return false, err
	}
	for i := range buf {
		if m[i] != buf[i] {
			return false, nil
		}
	}
	return true, nil
}


func GoProxy(){
	url_i := url.URL{}
	url_proxy, _ := url_i.Parse(proxy_addr)

	transport := http.Transport{}
	transport.Proxy = http.ProxyURL(url_proxy)// set proxy
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} //set ssl

	client := &http.Client{}
	client.Transport = transport
	resp, err := client.Get("http://example.com") // do request through proxy

	// other method
	//proxyUrl, err := url.Parse("87.236.233.92:8080")
	//httpClient := &http.Client { Transport: &http.Transport { Proxy: http.ProxyURL(proxyUrl) } }
}

