package qdrantURItoClient

import (
	"errors"
	"github.com/qdrant/go-client/qdrant"
	"net/url"
	"strconv"
)

func UriToClient(URI string) (*qdrant.Client, error) {
	//Простая функция которая из database url делает строку Qdrant Client.
	//
	// Формат database url: qdrant://[api_key]@][netloc][:port][/?param1=value1&...]
	//
	// Пример
	//  из "qdrant://1234567890@localhost:6333?UseTLS=1",
	//	результат: *qdrant.Client

	if len(URI) < 9 {
		return nil, errors.New("wrong uri'")
	}
	if URI[0:9] != "qdrant://" {
		return nil, errors.New("wrong protocol, support only 'qdrant://'")
	}
	UriObj, err := url.Parse(URI)
	if err != nil {
		return nil, err
	}
	if UriObj.Hostname() == "" {
		return nil, errors.New("Empty host")
	}
	apiKey := UriObj.User.Username()
	if UriObj.Hostname() == "" {
		return nil, errors.New("Empty host")
	}
	host := UriObj.Hostname()
	port := 6333
	if UriObj.Port() != "" {
		port, err = strconv.Atoi(UriObj.Port())
		if err != nil {
			return nil, err
		}
	}
	qs := UriObj.Query()
	UseTLS := false
	if qs.Get("UseTLS") != "" {
		UseTLS, err = strconv.ParseBool(qs.Get("UseTLS"))
		if err != nil {
			UseTLS = false
		}
	}

	client, err := qdrant.NewClient(&qdrant.Config{
		Host:   host,
		Port:   port,
		APIKey: apiKey,
		UseTLS: UseTLS, // uses default config with minimum TLS version set to 1.3
		// TLSConfig: &tls.Config{...},
		// GrpcOptions: []grpc.DialOption{},
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}
