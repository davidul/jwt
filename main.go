package main

import (
	"jwt/cmd"
)

func main() {

	err := cmd.Execute()
	if err != nil {
		return
	}
}
