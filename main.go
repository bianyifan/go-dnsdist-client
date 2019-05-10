package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"./dnsdist"
)

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] \n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	var host = flag.String("host", "127.0.0.1:5199", "host:port to connect to")
	var key = flag.String("key", "", "shared secret for the console")
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 0 {
		usage()
		os.Exit(2)
	}
	dc, err := dnsdist.Dial(*host, *key)

	if err != nil {
		log.Fatalf("Failure dialing: %s", err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			_, _ = fmt.Fprintln(os.Stderr, err)
			continue
		}
		cmdString = strings.TrimSuffix(cmdString, "\n")
		resp, err := dc.Command(cmdString)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			continue
		}
		fmt.Println(resp)
	}
}
