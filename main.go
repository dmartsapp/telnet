package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var (
	payload    bool
	iterations int
	udp        *bool
	timeout    int
	scheme     string
)

func init() {
	flag.IntVar(&iterations, "iterations", 1, "Number of times to check")
	flag.IntVar(&timeout, "timeout", 5, "Timeout in seconds to connect")
	flag.BoolVar(&payload, "download", false, "Check if payload can be downloaded")
	flag.StringVar(&scheme, "scheme", "https", "URL Scheme to use for web requests (http/https are only supported for now)")
	udp = flag.Bool("udp", false, "Use UDP instead of tcp to connect to endpoint")

	flag.Usage = func() {
		fmt.Println("Usage: " + os.Args[0] + " [options] <fqdn|IP> port")
		fmt.Println("options:")
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("Example (fqdn): " + os.Args[0] + " google.com 443")
		fmt.Println("Example (IP): " + os.Args[0] + " 10.10.10.10 443")
		fmt.Println("Example (fqdn with scheme): " + os.Args[0] + "-scheme https 10.10.10.10 443")
		os.Exit(0)
	}
}

func flagExists(flagname string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == flagname {
			found = true
		}
	})

	return found
}

func main() {

	flag.Parse()
	if len(flag.Args()) != 2 {
		flag.Usage()
	}

	if !*udp {
		if flagExists("scheme") {
			fmt.Println("Exists")
		}
	}

	// if len(flag.Args()) == 1 {
	// 	url, err := url.Parse(flag.Args()[0])
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		os.Exit(1)
	// 	}

	// 	port, _ := strconv.Atoi(url.Port())
	// 	if port == 0 {
	// 		switch url.Scheme {
	// 		case "http":
	// 			port = 80
	// 		case "https":
	// 			port = 443
	// 		default:
	// 			port = 443
	// 		}
	// 	}

	// 	resp, _ := http.Get(url.String())

	// 	fmt.Println(resp.StatusCode)

	// 	/*
	// 		placeholder for http or https check with or without payload download
	// 	*/
	// } else {
	// 	if !*udp {
	// 		addr := flag.Args()[0] + ":" + flag.Args()[1]
	// 		start := time.Now()
	// 		tcpaddr, err := net.ResolveTCPAddr("tcp", addr)
	// 		end := time.Now()
	// 		if err != nil {
	// 			fmt.Println(err.Error())
	// 			os.Exit(1)
	// 		}
	// 		fmt.Println("Resolved IP:Port of '" + addr + "' to " + tcpaddr.String() + " in " + strconv.Itoa(int(end.Sub(start).Milliseconds())) + "ms")
	// 		// var wg sync.WaitGroup
	// 		for i := 0; i < iterations; i++ {
	// 			// resp := dialNow("tcp", addr, timeout, &wg)
	// 			resp := dialNow("tcp", addr, timeout)
	// 			fmt.Println("TCP port checked successfully for '" + addr + "' in: " + strconv.Itoa(resp) + "ms")
	// 		}
	// 		// wg.Wait()

	// 	} else {
	// 		/*
	// 			Placeholder for UDP connection
	// 		*/
	// 		fmt.Println("UDP")
	// 	}
	// }

}

// func dialNow(protocol string, addressport string, timeout int, wg *sync.WaitGroup) int {
func dialNow(protocol string, addressport string, timeout int) int {
	start := time.Now()
	connect, err := net.DialTimeout(protocol, addressport, time.Duration(timeout)*time.Second)
	end := time.Now()
	if err != nil {
		fmt.Println(err)
		// wg.Done()
		os.Exit(1)
	}
	connect.Close()
	// wg.Done()
	return int((end.Sub(start)).Milliseconds())
}
