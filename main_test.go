package main

import (
	"os"
	"testing"

	require "github.com/alecthomas/assert/v2"
	"github.com/alecthomas/repr"
)

func TestExe(t *testing.T) {
	dat, err := os.ReadFile("example.paral")
	require.NoError(t, err)

	toml, err := paralParser.ParseString("", string(dat))
	require.NoError(t, err)
	repr.Println(toml)
}