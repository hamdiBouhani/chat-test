package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"

	"chat-test/server"
)

func main() {

	s, err := server.NewServer()
	if err != nil {
		err := errors.Wrap(err, "unable to create a new server")
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	if s != nil {
		err := s.Run()
		if err != nil {
			err := errors.Wrap(err, "unable to run server")
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(2)
		}
	} else {
		s.Log.Infoln("no rest server started")
	}
	fmt.Println("test")
}
