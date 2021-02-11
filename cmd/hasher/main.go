package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/salif/hasher"
)

func main() {
	var f func(args []string) error
	args := len(os.Args)
	if args == 1 {
		f = readAndHash
	} else if args == 2 {
		f = hash
	} else if args == 3 {
		f = readAndVerify
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

func readAndHash(args []string) error {
	password, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	return hashAndPrint(string(password))
}

func verify(args []string) error {
	result, err := hasher.Verify(args[1], args[2], args[3])
	if err != nil {
		return err
	}
	fmt.Printf("{\"match\": \"%v\"}\n", result)
	return nil
}

func readAndVerify(args []string) error {
	password, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	return verifyAndPrint(string(password), args[1], args[2])
}

func hashAndPrint(password string) error {
	hash, salt, err := hasher.Hash(password)
	if err != nil {
		return err
	}
	fmt.Printf("{\"hash\": \"%s\", \"salt\": \"%s\"}\n", hash, salt)
	return nil
}

func verifyAndPrint(password string, hash string, salt string) error {
	result, err := hasher.Verify(password, hash, salt)
	if err != nil {
		return err
	}
	fmt.Printf("{\"match\": \"%v\"}\n", result)
	return nil
}
