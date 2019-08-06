# go-copy

Library to copy files on Go.


## Installation

Run the following command from you terminal:

```bash
go get github.com/koind/go-copy
```

## Usage

Usage example.

```
./go-copy.go --from "./test.txt" --to "./test2.txt" --offset 0 --limit 1024
```

## Available Methods

The following methods are available:

##### koind/go-copy/file

```go
Copy(fromPath string, toPath string, offset int64, limit int) (bool, error)
```

## View command line options

Run the following command from you terminal:
```
./go-copy -help
```

Help information

```
Usage of ./go-copy:
  -f, --from string   Path to file for copy
      --help          Print help
  -l, --limit int     Limit for copying
  -o, --offset int    Offset to copy
  -t, --to string     Put a copy of the file on the path
```