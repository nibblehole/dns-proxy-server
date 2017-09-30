package main

import (
	. "github.com/mageddo/dns-proxy-server/log"
	"github.com/miekg/dns"
	"os"
	"net"
	
)

// reference https://miek.nl/2014/August/16/go-dns-package/
func main(){


	config, _ := dns.ClientConfigFromFile("/etc/resolv.conf")
	c := new(dns.Client)

	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(os.Args[1]), dns.TypeA) // CAN BE A, AAA, MX, etc.
	m.RecursionDesired = true

	r, _, err := c.Exchange(m, net.JoinHostPort(config.Servers[0], config.Port)) // server and port to ask

	// if the answer not be returned
	if r == nil {
		LOGGER.Fatalf("**** error: %s", err.Error())
	}

	// what the code of the return message ?
	if r.Rcode != dns.RcodeSuccess {
		LOGGER.Fatalf(" *** invalid answer name %s after MX query for %s", os.Args[1], os.Args[1])
	}

	// looping through the anwsers
	for _, a := range r.Answer {
		LOGGER.Infof("%v", a)
	}

}
