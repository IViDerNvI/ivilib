package main

import (
	"os"

	ivilib "github.com/IViDerNvI/ivilib/cmd"
	_ "github.com/IViDerNvI/ivilib/internal"
)

func main() {
	ivilib.Main(os.Args)
}
