package main

import (
	"flag"
	"fmt"
	"strings"
)

// recipients type
type recipients []map[string]string

func (rcpts *recipients) String() string {
	return fmt.Sprint(*rcpts)
}

func (rcpts *recipients) Set(value string) error {
	rcpt := strings.Split(value, ":")

	*rcpts = append(*rcpts, map[string]string{"name": rcpt[0], "email": rcpt[1]})

	return nil
}

func main() {
	var rcptsFlag recipients
	flag.Var(&rcptsFlag, "email", "space separated list of email recipients")
	flag.Parse()

	for i := range rcptsFlag {
		fmt.Println(rcptsFlag[i]["name"], rcptsFlag[i]["email"])
	}
}
