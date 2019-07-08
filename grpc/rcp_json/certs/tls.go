package certs

import (
	"crypto/tls"
	"golang.org/x/net/http2"
	"io/ioutil"
	"log"
	"path"
	"runtime"
)

func GetTLSConfig() (*tls.Config, error) {
	_, filename, _, ok1 := runtime.Caller(0)
	if !ok1 {
		panic("No caller information")
	}
	dir := path.Dir(filename)


	var certKeyPair *tls.Certificate
	cert, err := ioutil.ReadFile(path.Join(dir, "server.pem"))
	if err != nil {
		return nil, err
	}

	key, err := ioutil.ReadFile(path.Join(dir, "server.key"))
	if err != nil {
		return nil, err
	}

	pair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		log.Println("TLS KeyPair err: %v\n", err)
		return nil, err
	}

	certKeyPair = &pair

	return &tls.Config{
		Certificates: []tls.Certificate{*certKeyPair},
		NextProtos:   []string{http2.NextProtoTLS},
	}, nil
}
