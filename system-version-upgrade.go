package main

import (
	"fmt"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func systemUpdate(c *cli.Context) error {
	//add logic so it runs

	// sudo dnf upgrade --refresh
	// sudo dnf system-upgrade download --releasever=43
	// sudo dnf offline reboot

	fmt.Println("works")
	return logic()
}

func logic() error {
	//RUNING sudo dnf upgrade --refresh
	//prepere commands to run
	cmd := exec.Command("sudo", "dnf", "upgrade", "--refresh")
	//run and throw error if there is one)
	output, err := cmd.Output()
	if err != nil {
		return err
	}
	//OutPuts the output from the console after the command was ran
	fmt.Println(string(output))

	fmt.Println(string("Step 1 is compleated"))

	//RUNING sudo dnf system-upgrade download --releasever=43
	//prepere commands to run
	cmd2 := exec.Command("sudo", "dnf", "system-upgrade", "download", "--releasever=43")
	//run and throw error if there is one)
	output2, err2 := cmd2.Output()
	if err2 != nil {
		return err
	}
	//OutPuts the output from the console after the command was ran
	fmt.Println(string(output2))

	fmt.Println(string("Step 2 is compleated"))

	fmt.Println(string("Warning system might restart"))

	//RUNING sudo dnf offline reboot
	//prepere commands to run
	cmd3 := exec.Command("sudo", "dnf", "offline", "reboot")
	//run and throw error if there is one)
	output3, err3 := cmd3.Output()
	if err3 != nil {
		return err
	}
	//OutPuts the output from the console after the command was ran
	fmt.Println(string(output3))

	fmt.Println(string("System is succsesfully updated"))

	return nil
}
