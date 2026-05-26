package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func systemUpdate(c *cli.Context) error {
	//add logic so it runs

	// sudo dnf upgrade --refresh
	// sudo dnf system-upgrade download --releasever=43
	// sudo dnf offline reboot	

	fmt.Println("works")
	return nil
}
