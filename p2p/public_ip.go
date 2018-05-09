package p2p

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

var defaultServices = []string{
	"http://members.3322.org/dyndns/getip",
	"http://ifconfig.me/",
	"http://icanhazip.com/",
	"http://ifconfig.io/ip",
	"http://ident.me/",
	"http://whatismyip.akamai.com/",
	"http://myip.dnsomatic.com/",
	"http://diagnostic.opendns.com/myip",
	"http://myexternalip.com/raw",
}

type IpResult struct {
	Success bool
	Ip      string
	Error   error
}

var timeout time.Duration

func GetIP(services []string, to time.Duration) *IpResult {

	if services == nil || len(services) == 0 {
		services = defaultServices
	}
	if to == 0 {
		to = time.Duration(10)
	}
	timeout = to

	count := len(services)
	done := make(chan *IpResult)
	for k := range services {
		go ipAddress(services[k], done)
	}
	for {
		select {
		case result := <-done:
			if result.Success {
				return result
			} else {
				count--
				if count == 0 {
					result.Error = errors.New("All services doesn't available.")
					return result
				}
			}
			continue
		case <-time.After(time.Second * timeout):
			return &IpResult{false, "", errors.New("Timed out")}
		}
	}
}

func ipAddress(service string, done chan<- *IpResult) {

	timeout := time.Duration(time.Second * timeout)
	client := http.Client{Timeout: timeout}
	resp, err := client.Get(service)

	if err != nil {
		sendResult(&IpResult{false, "", errors.New("Time out")}, done)
		return
	}

	if err == nil {

		defer resp.Body.Close()

		address, err := ioutil.ReadAll(resp.Body)
		ip := fmt.Sprintf("%s", strings.TrimSpace(string(address)))
		if err == nil && net.ParseIP(ip) != nil {
			sendResult(&IpResult{true, ip, nil}, done)
			return
		}
	}
	sendResult(&IpResult{false, "", errors.New("Unable to talk with a service")}, done)
}

func sendResult(result *IpResult, done chan<- *IpResult) {
	select {
	case done <- result:
		return
	default:
		return
	}
}
