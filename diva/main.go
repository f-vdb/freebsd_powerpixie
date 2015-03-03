package main

import (
	//"flag"
	//"fmt"
	"github.com/f-vdb/freebsd_powerpixie/ping"
	//"log"
	//"net/http"
	//"os/exec"
	//"strings"
	"sync"
	"time"
)

var (
	statusLock sync.Mutex
	// Play it safe: assume somebody is using the server
	// samba
	reachableUsers bool = true
)

// Runs infinitly as a goroutine, periodically pinging users
func pingUsers() {
	for {
		statusLock.Lock()
		// get hosts
		statusLock.Unlock()

		reachable := false
		result := make(chan *time.Duration)
		go ping.Ping("192.168.1.11", 5*time.Second, result)
		reachable = <-result != nil

		statusLock.Lock()
		reachableUsers = reachable
		statusLock.Unlock()
		time.Sleep(10 * time.Second)
	}
}

func main() {
	go pingUsers()
	//fmt.Printf("%t\n", reachableUsers)
	// go checkShutdown
}
