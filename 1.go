package main

import (
	"os"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		os.Exit(1)
	}

	os.Stdout.WriteString(time.String())
}
