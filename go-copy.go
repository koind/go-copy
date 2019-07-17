package main

import (
	"flag"

	"github.com/koind/go-copy/file"
)

var (
	fromPath string
	toPath   string
	offset   int
	limit    int
)

func init() {
	flag.StringVar(&fromPath, "from", "", "")
	flag.StringVar(&toPath, "to", "", "")
	flag.IntVar(&offset, "offset", 0, "")
	flag.IntVar(&limit, "limit", 0, "")
}

func main() {
	flag.Parse()

	file.Copy(fromPath, toPath, offset, limit)
}
