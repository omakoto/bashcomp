package bashcomp

import (
	"flag"
	"os"
	"path"
	"strings"
	"text/template"
)

const (
	completionFlagName = "bash-completion"
)

var (
	bashCompleteEnabled = flag.Bool(completionFlagName, false, "dump bash completion source")
)

type values struct {
	Command    string
	Flags      string
	AllowFiles bool
	RawArgs    string
}

// Generate bash completion source file from command line flags for a command.
// This is for commands that accept file names as arguments.
func HandleBashCompletion() {
	handleBashCompletionWithOptions(true)
}

// Generate bash completion source file from command line flags for a command.
//
// This is for commands that take fixed set of strings as arguments, not file names.
func HandleBashCompletionNoFiles(rawArgs ...string) {
	handleBashCompletionWithOptions(false, rawArgs...)
}

func handleBashCompletionWithOptions(allowFiles bool, rawArgs ...string) {
	if !flag.Parsed() {
		flag.Parse()
	}
	if !*bashCompleteEnabled {
		return
	}
	templ, err := template.New("template").Parse(source)
	if err != nil {
		panic(err)
	}
	v := values{}
	v.Command = path.Base(os.Args[0])
	v.AllowFiles = allowFiles
	v.RawArgs = strings.Join(rawArgs, " ")

	v.Flags = ""
	flag.VisitAll(func(flag *flag.Flag) {
		if flag.Name != completionFlagName {
			v.Flags = v.Flags + " -" + flag.Name
		}
	})

	err = templ.Execute(os.Stdout, v)
	if err != nil {
		panic(err)
	}
	os.Exit(0)
}

var source = `
# Bash autocomplete script for the {{.Command}} command.
# Source it with the following command:
# . <({{.Command}} -bash-completion)
_{{ .Command }}_complete() {
  local cur="${COMP_WORDS[COMP_CWORD]}"

  local cand=""
  case "$cur" in
    "")
      {{ .Command }} -h >/dev/tty
      ;;
    -*)
      cand="{{.Flags}}"
      ;;
  esac
  if [ "x$cand" = "x" ] ; then
    COMPREPLY=(
        {{if.AllowFiles}}
        $(compgen -f -- ${cur})
        {{else}}
        $(compgen -W "{{.RawArgs}}" -- ${cur})
        {{end}}
        )
  else
    COMPREPLY=($(compgen -W "$cand" -- ${cur}))
  fi
}

complete -o filenames -o bashdefault -F _{{.Command}}_complete {{.Command}}
`
