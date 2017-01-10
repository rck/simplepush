package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"simplepush"
)

var (
	flagK = flag.String("k", "", "Set simplepush.io `key`")
	flagP = flag.String("p", "", "Set `password`, if set send message encrypted")
	flagE = flag.String("e", "", "Set `event`")
	flagT = flag.String("t", "", "Set `title`")
	flagM = flag.String("m", "", "Set `message`")
)

var Program, Version string

func main() {
	Program = path.Base(os.Args[0])
	if Version == "" {
		Version = "Unknown Version"
	}

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s (%s)\n", Program, Version)
		fmt.Fprintf(os.Stderr, "Usage: %s -k key -m message [-t title] [-e event] [-p password]\n", Program)
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Parse()

	err := simplepush.Send(simplepush.Message{
		SimplePushKey: *flagK,
		Password:      *flagP,
		Title:         *flagT,
		Message:       *flagM,
		Event:         *flagE,
		Encrypt:       *flagP != "",
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
