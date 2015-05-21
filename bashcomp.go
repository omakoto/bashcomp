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

func HandleBashCompletion() {
	HandleBashCompletionWithOptions(true)
}

func HandleBashCompletionWithOptions(allowFiles bool, rawArgs ...string) {
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
# . <({{.Command}} bash-completion)
_{{ .Command }}_complete() {
  local cur="${COMP_WORDS[COMP_CWORD]}"

  local cand=""
  case "$cur" in
    -*)
      cand="{{.Flags}}"
      ;;
  esac
  cand=$cand" {{.RawArgs}}"
  if [ "x$cand" = "x" ] ; then
    COMPREPLY=(
        {{if.AllowFiles}}
        $(compgen -f -- ${cur})
        {{end}}
        )
  else
    COMPREPLY=($(compgen -W "$cand" -- ${cur}))
  fi
}

complete -o filenames -o bashdefault -F _{{.Command}}_complete {{.Command}}
`

// func main() {
// 	for _, e := range os.Environ() {
// 		pair := strings.Split(e, "=")
// 		if strings.HasPrefix(pair[0], "COMP_") {
// 			fmt.Fprintf(os.Stderr, "%s=%q\n", pair[0], pair[1])
// 		}
// 	}
// }

// Example:
// COMP_TYPE="9"
// COMP_LINE="abc def"
// COMP_POINT="7"
// COMP_KEY="9"
