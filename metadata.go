package ec2meta

// This package can get metadata for EC2
// It is referenced to https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html

import (
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	metaEndpoint       = "169.254.169.254"
	metaEndpointFormat = "http://169.254.169.254/latest/meta-data/"
)

var (
	httpClient *http.Client

	ErrUnAvailableService = errors.New("[err] this machine isn't used to metadata service, maybe its machine not in AWS.")
)

// CanLookup is to check whether this machine can access a server serving metadata in AWS.
func CanLookup() bool {
	_, err := net.LookupAddr(metaEndpoint)
	if err != nil {
		return false
	}
	return true
}

// Hostname is to get hostname from this instance.
func Hostname() ([]string, error) {
	if !CanLookup() {
		return []string{}, ErrUnAvailableService
	}
	return callAPI(metaEndpointFormat + "hostname")
}

// LocalHostname is to get local hostname from this instance.
func LocalHostname() ([]string, error) {
	if !CanLookup() {
		return []string{}, ErrUnAvailableService
	}
	return callAPI(metaEndpointFormat + "local-hostname")
}

// LocalIPV4 is to get local ipv4 from this instance.
func LocalIPV4() ([]string, error) {
	if !CanLookup() {
		return []string{}, ErrUnAvailableService
	}
	return callAPI(metaEndpointFormat + "local-ipv4")
}

// Mac is to get mac address from this instance.
func Mac() ([]string, error) {
	if !CanLookup() {
		return []string{}, ErrUnAvailableService
	}
	return callAPI(metaEndpointFormat + "mac")
}

// PublicHostname is to get public hostname from this instance.
func PublicHostname() ([]string, error) {
	if !CanLookup() {
		return []string{}, ErrUnAvailableService
	}
	return callAPI(metaEndpointFormat + "public-hostname")
}

// PublicIPV4 is to get public ipv4 from this instance.
func PublicIPV4() ([]string, error) {
	if !CanLookup() {
		return []string{}, ErrUnAvailableService
	}
	return callAPI(metaEndpointFormat + "public-ipv4")
}

func callAPI(url string) ([]string, error) {
	ret, err := httpClient.Get(url)
	if ret != nil {
		defer ret.Body.Close()
	}
	if err != nil {
		return []string{}, err
	}

	body, err := ioutil.ReadAll(ret.Body)
	if err != nil {
		return []string{}, err
	}

	return strings.Split(string(body), "\n"), nil
}

func init() {
	httpClient = &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   3 * time.Second,
				KeepAlive: 10 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout:   3 * time.Second,
			ExpectContinueTimeout: 3 * time.Second,
			ResponseHeaderTimeout: 3 * time.Second,
			IdleConnTimeout:       10 * time.Minute,
		},
		Timeout: 10 * time.Second,
	}
}
