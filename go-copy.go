package main

import (
	"fmt"
	"log"

	"github.com/koind/go-copy/file"
	flag "github.com/spf13/pflag"
)

var (
	helpFlag bool
	fromPath string
	toPath   string
	offset   int64
	limit    int
)

func init() {
	flag.BoolVar(&helpFlag, "help", false, "Print help")
	flag.StringVarP(&fromPath, "from", "f", "", "Path to file for copy")
	flag.StringVarP(&toPath, "to", "t", "", "Put a copy of the file on the path")
	flag.Int64VarP(&offset, "offset", "o", 0, "Offset to copy")
	flag.IntVarP(&limit, "limit", "l", 0, "Limit for copying")
}

func main() {
	flag.Parse()

	if helpFlag {
		printDefaults()
		return
	}

	if fromPath == "" || toPath == "" {
		log.Fatal("Specify the paths to copy the file")
	}

	if limit == 0 {
		log.Fatal("Indicate the limit and offset")
	}

	isCopied, err := file.Copy(fromPath, toPath, offset, limit)
	if err != nil {
		fmt.Println(err)
		return
	}

	if isCopied {
		fmt.Println("Сopy was successful")
	}
}

// Displays help information
func printDefaults() {
	fmt.Println("Usage: go-copy <options>")
	fmt.Println("Options:")
	flag.VisitAll(func(flag *flag.Flag) {
		fmt.Println("\t-"+flag.Name, "\t", flag.Usage, "(Default "+flag.DefValue+")")
	})
}
