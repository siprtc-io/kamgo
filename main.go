package main

import (
	"flag"
	"fmt"
	"kamgo/modules/log"
	"kamgo/router"
	"os"
)

const (
	DefaultConfFilePath = "config/config.toml"
)

var (
	confFilePath string
	help         bool
)

func init() {
	flag.StringVar(&confFilePath, "c", DefaultConfFilePath, "Config Path")
	flag.StringVar(&confFilePath, "config", DefaultConfFilePath, "Config Path")
	flag.BoolVar(&help, "h", false, "Show help message")
	flag.BoolVar(&help, "help", false, "Show help message")
	flag.Parse()
	flag.Usage = usage
}

func usage() {
	s := `kamgo : a Application for Kamailio Endpoint Management
		Usage: kamgo [Options...]
		Options:
    		-c,  -config=<path>           Config path of the site. Default is config/config.toml.
		Other options:
    		-h,  -help                  Show help message.
		`
	fmt.Printf(s)
	os.Exit(0)
}

func main() {

	log.Info("Application for Managing Kamailio Endpoint")

	if help {
		usage()
		return
	}
	log.Debugf("run with conf:%s", confFilePath)
	router.RunSubdomains(confFilePath)
}
