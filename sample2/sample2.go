package main

import (
	"flag"
	"github.com/omakoto/bashcomp"
)

var (
	test = flag.Bool("test", false, "test flag")
	foo  = flag.Bool("foo", false, "test flag 2")
)

func main() {
	bashcomp.HandleBashCompletionNoFiles("aaa", "aab", "abc", "bbb")
}
