package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func systemUpdate(c *cli.Context) error {
	//add logic so it runs

	// sudo dnf upgrade --refresh
	// sudo dnf system-upgrade download --releasever=43
	// sudo dnf offline reboot

	//fmt.Println("works")
	return logic()
}

func logic() error {
	//RUNING sudo dnf upgrade --refresh
	//prepere commands to run
	fmt.Println(string("preparoing to run step1"))
	cmd := exec.Command("sudo", "dnf", "upgrade", "--refresh")
	fmt.Println(string("running and throwing step1"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	// Running and executing
	err := cmd.Run()
	//if no error continue
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			fmt.Printf("Exit code: %d\n", exitErr.ExitCode())
		}
		fmt.Printf("Full error: %v\n", err)
		return err
	}

	fmt.Println(string("Step 1 is compleated"))

	//RUNING sudo dnf system-upgrade download --releasever=43
	//prepere commands to run
	cmd2 := exec.Command("sudo", "dnf", "system-upgrade", "download", "--releasever=43", "-y")
	//run and throw error if there is one)
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr
	cmd2.Stdin = os.Stdin
	cmd2.Env = append(os.Environ(), "LANG=en_US.UTF-8")

	// Running and executing
	err = cmd2.Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			fmt.Printf("Exit code: %d\n", exitErr.ExitCode())
		}
		fmt.Printf("Full error: %v\n", err)
		return err
	}
	// output2, err2 := cmd2.CombinedOutput()
	// if err2 != nil {
	// 	return err
	// }
	// //OutPuts the output from the console after the command was ran
	// fmt.Println(string(output2))

	fmt.Println(string("Step 2 is compleated"))

	fmt.Println(string("Warning system might restart"))

	//RUNING sudo dnf offline reboot
	//prepere commands to run
	cmd3 := exec.Command("sudo", "dnf", "offline", "reboot")
	//run and throw error if there is one)
	cmd3.Stdout = os.Stdout
	cmd3.Stderr = os.Stderr
	cmd3.Stdin = os.Stdin
	cmd3.Env = append(os.Environ(), "LANG=en_US.UTF-8")

	// Running and executing
	err = cmd3.Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			fmt.Printf("Exit code: %d\n", exitErr.ExitCode())
		}
		fmt.Printf("Full error: %v\n", err)
		return err
	}
	//run and throw error if there is one)
	// output3, err3 := cmd3.CombinedOutput()
	// if err3 != nil {
	// 	return err
	// }
	// //OutPuts the output from the console after the command was ran
	// fmt.Println(string(output3))

	fmt.Println(string("System is succsesfully updated"))

	return nil
}
