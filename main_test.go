package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	f, err := os.Open("/Users/daqige/Code/PatentAdminPlat/config/banner")
	assert.NoError(t, err)
	bytes, err := ioutil.ReadAll(f)
	assert.NoError(t, err)
	fmt.Println(bytes)
	fmt.Printf("\n")
	for _, b := range bytes {
		fmt.Printf("%d, ", b)
	}
	fmt.Printf("\n")
	fmt.Println(string(bytes))
}
