/*
Package bashcomp generates a bash completion source file for a command executable.

Example: (this file is in sample1/.)

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
        bashcomp.HandleBashCompletion()
    }

When this program is executed with -bash-completion flag, it prints a bash
source file and exits.  Source this output line this:

    $ . <(testcmd1 -bash-completion)

Then,

    $ testcmd1 -[TAB][TAB]

will show

    -foo  -test

*/
package bashcomp
