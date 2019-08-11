package main

import (
	"fmt"
	"log"

	"github.com/koind/go-copy/file"
	flag "github.com/spf13/pflag"
)

var (
	fromPath string
	toPath   string
	offset   int64
	limit    int
)

func init() {
	flag.StringVarP(&fromPath, "from", "f", "", "Path to file for copy")
	flag.StringVarP(&toPath, "to", "t", "", "Put a copy of the file on the path")
	flag.Int64VarP(&offset, "offset", "o", 0, "Offset to copy")
	flag.IntVarP(&limit, "limit", "l", 0, "Limit for copying")
}

func main() {
	flag.Parse()

	if fromPath == "" || toPath == "" {
		log.Fatal("Specify the paths to copy the file")
	}

	if limit == 0 {
		log.Fatal("Indicate the limit and offset")
	}

	err := file.Copy(fromPath, toPath, offset, limit)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Сopy was successful")
	}
}
