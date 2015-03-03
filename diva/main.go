package main

import (
	//"flag"
	"fmt"
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
	//who string
)

// Runs infinitly as a goroutine, periodically pinging users
func pingUsers() {
	for {
		statusLock.Lock()
		// get hosts
		statusLock.Unlock()

		reachable := false
		result := make(chan *time.Duration)
		go ping.Ping("192.168.0.1", 1*time.Second, result)
		reachable = <-result != nil
		//who = <-result != nil

		statusLock.Lock()
		reachableUsers = reachable
		//fmt.Printf("%t\n", reachableUsers)
		statusLock.Unlock()
		time.Sleep(2 * time.Second)
	}
}

func main() {
	go pingUsers()
	for i := 0; i < 10000; i++ {
		fmt.Printf("----------- %t\n", reachableUsers)
		time.Sleep(2 * time.Second)
	}

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
	// go checkShutdown
}
