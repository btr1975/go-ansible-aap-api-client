/*
Package connection provides a basic connection for Ansible AAP
*/
package connection

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net/http"
	"os"
)

// Connection is the basic connection
type Connection struct {
	BaseURL    string
	Username   string
	Password   string
	SSLVerify  bool
	APIVersion string
	Headers    map[string]string
	tlsConfig  *tls.Config
	transport  *http.Transport
}

// NewConnection creates a new connection
//
//	:param baseURL: The base URL of the AAP server
//	:param username: The username to use for authentication
//	:param password: The password to use for authentication
//	:param sslVerify: Whether to verify the SSL certificate
//	:param certPath: The path to the certificate to use for SSL verification
func NewConnection(baseURL string, username string, password string, sslVerify bool, certPath string) (*Connection, error) {
	tlsConfig := &tls.Config{}

	caPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, fmt.Errorf("failed to load system cert pool: %v", err)
	}

	tlsConfig.RootCAs = caPool

	if !sslVerify {
		tlsConfig.InsecureSkipVerify = true

	} else {
		tlsConfig.InsecureSkipVerify = false
		if certPath == "" {
			return nil, errors.New("certPath is required when sslVerify is true")
		}

		certData, err := os.ReadFile(certPath)
		if err != nil {
			return nil, err
		}

		clientCert, err := tls.X509KeyPair(certData, certData)
		if err != nil {
			return nil, fmt.Errorf("error parsing PEM into X509 key pair: %v", err)
		}

		tlsConfig.Certificates = []tls.Certificate{clientCert}

	}

	return &Connection{
		BaseURL:    baseURL,
		Username:   username,
		Password:   password,
		SSLVerify:  sslVerify,
		APIVersion: "/api/v2",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		tlsConfig: tlsConfig,
		transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}, nil
}

func (connection *Connection) Get(uri string, params map[string]string) (response *http.Response, err error) {
	client := &http.Client{
		Transport: connection.transport,
	}

	request, err := http.NewRequest("GET", connection.BaseURL+connection.APIVersion+uri, nil)

	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err = client.Do(request)

	return response, err

}
