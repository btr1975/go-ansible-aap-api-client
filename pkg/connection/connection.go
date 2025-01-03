/*
Package connection provides a basic connection for Ansible AAP
*/
package connection

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

// BasicConnection is the basic connection interface
type BasicConnection interface {
	Get(uri string, params map[string]string) (response *http.Response, err error)
	Post(uri string, data []byte) (response *http.Response, err error)
	Patch(uri string, data []byte) (response *http.Response, err error)
	Delete(uri string, data []byte) (response *http.Response, err error)
}

// Connection is the basic connection
type Connection struct {
	BaseURL    *url.URL
	Username   string
	Password   string
	SSLVerify  bool
	APIVersion string
	Headers    map[string]string
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
	baseURLParsed, err := url.Parse(baseURL)

	if err != nil {
		return nil, err
	}

	connection := &Connection{
		BaseURL:    baseURLParsed,
		Username:   username,
		Password:   password,
		SSLVerify:  sslVerify,
		APIVersion: "/api/v2",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	tlsConfig, err := connection.createTLSConfig(sslVerify, certPath)

	if err != nil {
		return nil, err
	}

	connection.transport = &http.Transport{TLSClientConfig: tlsConfig}

	return connection, nil
}

// createTLSConfig creates a TLS config
//
//	:param sslVerify: Whether to verify the SSL certificate
//	:param certPath: The path to the certificate to use for SSL verification
func (connection *Connection) createTLSConfig(sslVerify bool, certPath string) (*tls.Config, error) {
	tlsConfig := &tls.Config{}

	if !sslVerify {
		tlsConfig.InsecureSkipVerify = true

	} else {
		tlsConfig.InsecureSkipVerify = false

		caPool, err := x509.SystemCertPool()
		if err != nil {
			return nil, fmt.Errorf("failed to load system cert pool: %v", err)
		}

		tlsConfig.RootCAs = caPool

		if certPath != "" {
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

	}

	return tlsConfig, nil
}

// createRequest creates a new request
//
//	:param method: The HTTP method to use
//	:param finalURL: The final URL to use
//	:param data: The data to send
func (connection *Connection) createRequest(method string, finalURL string, data *bytes.Buffer) (*http.Request, error) {

	if data == nil {
		request, err := http.NewRequest(method, finalURL, nil)

		if err != nil {
			return nil, err
		}

		for key, value := range connection.Headers {
			request.Header.Set(key, value)
		}

		request.SetBasicAuth(connection.Username, connection.Password)

		return request, nil

	}

	request, err := http.NewRequest(method, finalURL, data)

	if err != nil {
		return nil, err
	}

	for key, value := range connection.Headers {
		request.Header.Set(key, value)
	}

	request.SetBasicAuth(connection.Username, connection.Password)

	return request, nil

}

// checkOK checks if the response is OK
//
//	:param response: The response to check
func (connection *Connection) checkOK(response *http.Response) bool {
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return true
	}

	return false
}

// Get performs a GET request
//
//	:param uri: The URI to use
//	:param params: The parameters to pass
func (connection *Connection) Get(uri string, params map[string]string) (response *http.Response, err error) {
	client := &http.Client{
		Transport: connection.transport,
	}

	if params != nil {
		q := connection.BaseURL.Query()
		for key, value := range params {
			q.Set(key, value)
		}
		connection.BaseURL.RawQuery = q.Encode()
	}

	finalURL := connection.BaseURL.JoinPath(connection.APIVersion, uri)

	request, err := connection.createRequest("GET", finalURL.String(), nil)

	if err != nil {
		return nil, err
	}

	response, err = client.Do(request)

	if err != nil {
		return nil, err
	}

	if !connection.checkOK(response) {
		return nil, fmt.Errorf("error GET response code %d", response.StatusCode)
	}

	return response, err

}

// Post performs a POST request
//
//	:param uri: The URI to use
//	:param data: The data to POST
func (connection *Connection) Post(uri string, data []byte) (response *http.Response, err error) {
	client := &http.Client{
		Transport: connection.transport,
	}

	finalURL := connection.BaseURL.JoinPath(connection.APIVersion, uri)

	request, err := connection.createRequest("POST", finalURL.String(), bytes.NewBuffer(data))

	if err != nil {
		return nil, err
	}

	response, err = client.Do(request)

	if err != nil {
		return nil, err
	}

	if !connection.checkOK(response) {
		return nil, fmt.Errorf("error POST response code %d", response.StatusCode)
	}

	return response, err
}

// Patch performs a PATCH request
//
//	:param uri: The URI to use
//	:param data: The data to PATCH
func (connection *Connection) Patch(uri string, data []byte) (response *http.Response, err error) {
	client := &http.Client{
		Transport: connection.transport,
	}

	finalURL := connection.BaseURL.JoinPath(connection.APIVersion, uri)

	request, err := connection.createRequest("PATCH", finalURL.String(), bytes.NewBuffer(data))

	if err != nil {
		return nil, err
	}

	response, err = client.Do(request)

	if err != nil {
		return nil, err
	}

	if !connection.checkOK(response) {
		return nil, fmt.Errorf("error PATCH response code %d", response.StatusCode)
	}

	return response, err
}

// Delete performs a DELETE request
//
//	:param uri: The URI to use
//	:param data: The data to DELETE
func (connection *Connection) Delete(uri string, data []byte) (response *http.Response, err error) {
	client := &http.Client{
		Transport: connection.transport,
	}

	finalURL := connection.BaseURL.JoinPath(connection.APIVersion, uri)

	request, err := connection.createRequest("DELETE", finalURL.String(), bytes.NewBuffer(data))

	if err != nil {
		return nil, err
	}

	response, err = client.Do(request)

	if err != nil {
		return nil, err
	}

	if !connection.checkOK(response) {
		return nil, fmt.Errorf("error DELETE response code %d", response.StatusCode)
	}

	return response, err
}
