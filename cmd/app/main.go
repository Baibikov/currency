package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main()  {
	if err := app(); err != nil {
		logrus.Panic(err)

		os.Exit(1)
	}
}