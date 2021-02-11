package server

import (
	"fmt"
	"log"
	"net/http"
)

// Info contains data required run Server.
type Info struct {
	Hostname  string `yaml:"Hostname"`
	UseHTTP   bool   `yaml:"UseHTTP"`
	UseHTTPS  bool   `yaml:"UseHTTPS"`
	HTTPPort  uint16 `yaml:"HTTPPort"`
	HTTPSPort uint16 `yaml:"HTTPSPort"`
	CertFile  string `yaml:"CertFile"`
	KeyFile   string `yaml:"KeyFile"`
}

// Run starts HTTP and HTTPs listeners.
func Run(handlers http.Handler, info Info) {
	done := make(chan bool)

	// Run HTTP service.
	if info.UseHTTP {
		go func() {
			addr := address(info.Hostname, info.HTTPPort)
			log.Printf("Running HTTP %s", addr)
			log.Fatal(http.ListenAndServe(addr, handlers))
		}()
	}

	// Run HTTP TLS service.
	if info.UseHTTPS {
		go func() {
			addr := address(info.Hostname, info.HTTPSPort)
			log.Printf("Running HTTPs %s", addr)
			log.Fatal(http.ListenAndServeTLS(addr,
				info.CertFile,
				info.KeyFile,
				handlers,
			))
		}()
	}

	<-done
}

// address returns "hostname:port" string
func address(hostname string, port uint16) string {
	return fmt.Sprintf("%s:%d", hostname, port)
}
