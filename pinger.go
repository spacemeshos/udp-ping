package pinger

import (
	"net"
	"time"

	//"bufio"
	"fmt"
	"html"
	"log"

	//"log"
	//"net"
	"net/http"
	"strings"
)

// GetIP gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		// see https://en.wikipedia.org/wiki/X-Forwarded-For#Format
		ips := strings.Split(forwarded, ", ")
		return ips[0]
	}

	hosts := strings.Split(r.RemoteAddr, ":")
	return hosts[0]
}

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func Pinger(w http.ResponseWriter, r *http.Request) {

	remote := getIP(r)
	ports, ok := r.URL.Query()["p"]
	if !ok || len(ports[0]) < 1 {
		_, _ = fmt.Fprint(w, html.EscapeString("missing required p param"))
		return
	}

	port := ports[0]
	log.Printf("UDP address to connect to " + remote + ":" + port)
	conn, err := net.DialTimeout("udp", remote+":"+port, time.Duration(10)*time.Second)
	if err != nil {
		log.Printf("error connecting %v", err)
		_, _ = fmt.Fprint(w, html.EscapeString("Error connecting: %v"), err)
		return
	}

	_, err = fmt.Fprintf(conn, "spacemesh")
	if err != nil {
		_, _ = fmt.Fprint(w, html.EscapeString("error writing to conn: %v"), err)
		return
	}

	// read response
	/*
	p := make([]byte, 2048)
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		log.Printf("Response: %s\n", p)
	} else {
		log.Printf("Error: %v\n", err)
	}*/

	_ = conn.Close()
	_, _ = fmt.Fprint(w, html.EscapeString("Pinged " + remote+":"+port))
}
