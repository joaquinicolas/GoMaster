package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)


/**
Go constanst for working with time:
	15 --> parsing the hour
	04 --> parsing the minutes
	05 --> parsing the seconds
*/
func main() {
	var tm string
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s string\n",
			filepath.Base(os.Args[0]))

		os.Exit(1)
	}

	tm = os.Args[1]
	d, err := time.Parse("15:04:05", tm)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Hour(), d.Minute(), d.Second())
	} else {
		fmt.Println(err)
	}
}
