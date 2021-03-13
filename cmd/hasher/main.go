package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"salif.eu/go/hasher"
)

func main() {
	var f func(args []string) string
	var args = len(os.Args)
	if args == 1 {
		f = readAndHash
	} else if args == 2 {
		f = hash
	} else if args == 3 {
		f = readAndVerifyVL
	} else if args == 4 {
		f = readAndVerify
	} else if args == 5 {
		f = verify
	} else {
		f = func(args []string) string {
			panic("Invalid arguments")
		}
	}
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln(r)
		}
	}()
	var result = f(os.Args)
	fmt.Print(result)
}

func readAndHash(args []string) string {
	password, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	var hash, salt, version = hasher.Hash(string(password))
	return fmt.Sprintf("%s %s %d", hash, salt, version)
}

func hash(args []string) string {
	var hash, salt, version = hasher.Hash(args[1])
	return fmt.Sprintf("%s %s %d", hash, salt, version)
}

func readAndVerifyVL(args []string) string {
	var password, rErr = ioutil.ReadAll(os.Stdin)
	if rErr != nil {
		panic(rErr)
	}
	return fmt.Sprintf("%t", hasher.Verify(string(password), args[1], args[2], hasher.VERSION))
}

func readAndVerify(args []string) string {
	var password, rErr = ioutil.ReadAll(os.Stdin)
	if rErr != nil {
		panic(rErr)
	}
	var v, cErr = strconv.Atoi(args[3])
	if cErr != nil {
		panic(cErr)
	}
	return fmt.Sprintf("%t", hasher.Verify(string(password), args[1], args[2], v))
}

func verify(args []string) string {
	var v, err = strconv.Atoi(args[4])
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%t", hasher.Verify(args[1], args[2], args[3], v))
}
