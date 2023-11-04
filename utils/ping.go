package handlers

import (
	"fmt"
	"github.com/go-ini/ini"
	"net/http"
	"os"
	"time"
)

func ping(website string) {
	cfg, err := ini.Load("/root/.ioson/.ioson.ini")
	if err != nil {
		fmt.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	ttl, err := cfg.Section("").Key("ttl").Int()
	if err != nil {
		fmt.Printf("Failed to read ttl value: %v", err)
		os.Exit(1)
	}

	for {
		resp, err := http.Get(website)
		if err != nil {
			fmt.Printf("Failed to ping website: %v", err)
		} else {
			fmt.Printf("Website %s is up and running!\n", website)
			resp.Body.Close()
		}
		time.Sleep(time.Duration(ttl) * time.Second)
	}
}