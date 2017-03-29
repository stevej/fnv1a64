package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"log"
	"os"
)

// Read from stdin, output the FNV-1a 64-bit hash value as a decimal number.
func main() {
	reader := bufio.NewReader(os.Stdin)
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Panicf("could not read bytes from stdin %s", err)
	}

	hasher := fnv.New64a()
	_, err = hasher.Write(bytes)
	if err != nil {
		log.Panicf("could not hash values %s", err)
	}

	sum := hasher.Sum64()
	fmt.Printf("%d\n", sum)
}
