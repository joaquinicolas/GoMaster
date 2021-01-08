package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

/**
Go constanst for working with dates:
	02 --> parsing the day
	January --> parsing the month
	2006 --> parsing the year
	**/
func main() {
	var dt string
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s string\n",
			filepath.Base(os.Args[0]))

		return
	}

	dt = os.Args[1]
	d, err := time.Parse("02 January 2006", dt)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Day(), d.Month(), d.Year())
	} else {
		fmt.Println(err)
	}
}
