package main

import (
	"fmt"
	"os"
	"strconv"

	"./password"
)

func main() {
	var length uint = 12
	var spec bool = true
	var args []string = os.Args

	for i, v := range args {
		if (v == "--l") || (v == "--length") {
			len, err := strconv.Atoi(args[i+1])
			if err != nil {
				fmt.Println(err)
			}

			length = uint(len)
		}

		if v == "--spec-disable" {
			spec = false
		}
	}

	fmt.Printf("Generated password: %s\n", password.GenerateRand(length, spec))
}
