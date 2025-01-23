package util

import "fmt"

func PrintBanner() {
	banner := `
           __ _ _              
  _____ __/ _(_) |_ __ _ _ ___ 
 / -_) \ /  _| | | '_ \ '_/ _ \
 \___/_\_\_| |_|_| .__/_| \___/
                 |_|           
`
	fmt.Printf("%v\nVersion: %v (%v) - %v - %v\n\n", banner, Version, GitCommit, Author, BuildDate)
}
