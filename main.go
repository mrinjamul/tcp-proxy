package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/mrinjamul/tcp-proxy/app"
	flag "github.com/spf13/pflag"
)

var (
	// AppName is the name of the application
	AppName = "tcp-proxy"
	// Author is the author of the application
	Author = "@mrinjamul"
	// Version is the version of the application
	Version = "dev"
	// CommitHash is the commit hash of the application
	CommitHash = "none"
	// BuildDate is the date of the build
	BuildDate = "unknown"
)

// flag variables
var (
	flagTarget     string
	flagTargetPort int
	flagListenPort string
	flagHelp       bool
	flagVersion    bool
)

func main() {

	// parse flags
	flag.Parse()

	// if user does not supply flags, print usage
	if flag.NFlag() == 0 {
		printUsage()
	}
	// print help
	if flagHelp {
		printUsage()
		os.Exit(0)
	}
	// print version info
	if flagVersion {
		printVersion()
		os.Exit(0)
	}
	// check if target and target is empty or not
	if flagTarget == "" || flagTargetPort == 0 {
		fmt.Println("Target host and port are required")
		os.Exit(1)
	}

	log.Println("Starting proxy server on port " + flagListenPort)
	// create proxy server
	listener, err := net.Listen("tcp", ":"+flagListenPort)
	if err != nil {
		log.Fatalln("unable to bind to port")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("unable to accept connection")
		}
		app.CreateProxy(conn, flagTarget, flagTargetPort)
	}
}

func init() {
	flag.StringVarP(&flagTarget, "target", "t", "", "Target host")
	flag.IntVarP(&flagTargetPort, "target-port", "p", 0, "Target port")
	flag.StringVarP(&flagListenPort, "listen-port", "l", "8000", "Listen port")
	flag.BoolVarP(&flagHelp, "help", "h", false, "print help")
	flag.BoolVarP(&flagVersion, "version", "v", false, "print version")
}

func printUsage() {
	fmt.Printf("Usage: %s [options]\n", AppName)
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(0)
}
func printVersion() {
	fmt.Println(AppName + " " + Version + "+" + CommitHash + " " + BuildDate + " by " + Author)
}
