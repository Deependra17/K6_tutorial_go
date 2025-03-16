package main

import (
	"fmt"
)

func main() {
	currentBranch := getCurrentBranch()

	if currentBranch != "main" {
		fmt.Println("Pushing code to the main branch...")
		pushToMainBranch()
	} else {
		fmt.Println("Welcome to the performance testing tool k6..")
	}
}

func getCurrentBranch() string {
	return "feature-branch"
}
func pushToMainBranch() {
	fmt.Println("Checking out to main branch...")
	fmt.Println("Pushing code to the remote main branch...")

}
