package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/salifm/hasher"
)

func main() {
	var f func(args []string) error
	args := len(os.Args)
	if args == 1 {
		f = repl
	} else if args == 2 {
		f = hash
	} else if args == 4 {
		f = verify
	} else {
		f = func(args []string) error {
			return fmt.Errorf("Invalid arguments")
		}
	}
	err := f(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "{\"error\": \"%s\"}\n", err.Error())
	}
}

func hash(args []string) error {
	return hashAndPrint(args[1])
}

func verify(args []string) error {
	result, err := hasher.Verify(args[1], args[2], args[3])
	if err != nil {
		return err
	}
	fmt.Printf("{\"match\": \"%v\"}\n", result)
	return nil
}

func repl(args []string) error {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("password: ")
		s := scanner.Scan()
		if !s {
			return nil
		}
		if scanner.Err() != nil {
			return scanner.Err()
		}
		err := hashAndPrint(scanner.Text())
		if err != nil {
			return err
		}
	}
}

func hashAndPrint(password string) error {
	hash, salt, err := hasher.Hash(password)
	if err != nil {
		return err
	}
	fmt.Printf("{\"hash\": \"%s\", \"salt\": \"%s\"}\n", hash, salt)
	return nil
}
