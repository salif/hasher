package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"salif.eu/go/hasher"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln(r)
		}
	}()
	fmt.Print(getResult(os.Args[1:]))
}

func getResult(args []string) string {
	var argsc = len(args)
	if argsc == 0 {
		return hash(read(os.Stdin))
	} else if argsc == 1 {
		return hash(readFile(args[0]))
	} else if argsc == 3 {
		return verify(read(os.Stdin), args[0], args[1], args[2])
	} else if argsc == 4 {
		return verify(readFile(args[0]), args[1], args[2], args[3])
	} else {
		panic("Invalid arguments")
	}
}

func hash(password string) string {
	var hash, salt, version = hasher.Hash(string(password))
	return fmt.Sprintf("%s %s %d", hash, salt, version)
}

func verify(password string, hash string, salt string, v string) string {
	var version, err = strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%t", hasher.Verify(string(password), hash, salt, version))
}

func read(r io.Reader) string {
	var b, err = io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func readFile(filename string) string {
	var file, oErr = os.Open(filename)
	if oErr != nil {
		panic(oErr)
	}
	defer func() {
		if cErr := file.Close(); cErr != nil {
			panic(cErr)
		}
	}()
	return read(file)
}
