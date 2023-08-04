package main

import (
	"flag"
)

func scan(folder string) {
	print("scan")
}

func stats(email string) {
	print("stats")
}

func main() {
	var folder string
	var email string
	flag.StringVar(&folder, "folder", "", "Folder to scan")
	flag.StringVar(&email, "email", "", "Email to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}
	stats(email)
}
