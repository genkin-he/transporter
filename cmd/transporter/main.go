package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	_ "transporter/adaptor/all"
	_ "transporter/function/all"
	"transporter/log"
)

const (
	defaultPipelineFile = "pipeline.js"
)

var version = "dev" // set by release script

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, "USAGE\n")
	_, _ = fmt.Fprintf(os.Stderr, "  %s <command> [flags]\n", os.Args[0])
	_, _ = fmt.Fprintf(os.Stderr, "\n")
	_, _ = fmt.Fprintf(os.Stderr, "COMMANDS\n")
	_, _ = fmt.Fprintf(os.Stderr, "  run       run pipeline loaded from a file\n")
	_, _ = fmt.Fprintf(os.Stderr, "  test      display the compiled nodes without starting a pipeline\n")
	_, _ = fmt.Fprintf(os.Stderr, "  about     show information about available adaptors\n")
	_, _ = fmt.Fprintf(os.Stderr, "  init      initialize a config and pipeline file based from provided adaptors\n")
	_, _ = fmt.Fprintf(os.Stderr, "  xlog      manage the commit log\n")
	_, _ = fmt.Fprintf(os.Stderr, "  offset    manage the offset for sinks\n")
	_, _ = fmt.Fprintf(os.Stderr, "\n")
	_, _ = fmt.Fprintf(os.Stderr, "VERSION\n")
	_, _ = fmt.Fprintf(os.Stderr, "  %s\n", version)
	_, _ = fmt.Fprintf(os.Stderr, "\n")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	var run func([]string) error
	switch strings.ToLower(os.Args[1]) {
	case "run":
		run = runRun
	case "test":
		run = runTest
	case "about":
		run = runAbout
	case "init":
		run = runInit
	case "xlog":
		run = runXlog
	case "offset":
		run = runOffset
	default:
		usage()
		os.Exit(1)
	}

	if err := run(os.Args[2:]); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func baseFlagSet(setName string) *flag.FlagSet {
	cmdFlags := flag.NewFlagSet(setName, flag.ExitOnError)
	log.AddFlags(cmdFlags)
	return cmdFlags
}

func usageFor(fs *flag.FlagSet, short string) func() {
	return func() {
		_, _ = fmt.Fprintf(os.Stderr, "USAGE\n")
		_, _ = fmt.Fprintf(os.Stderr, "  %s\n", short)
		_, _ = fmt.Fprintf(os.Stderr, "\n")
		_, _ = fmt.Fprintf(os.Stderr, "FLAGS\n")
		w := tabwriter.NewWriter(os.Stderr, 0, 2, 2, ' ', 0)
		fs.VisitAll(func(f *flag.Flag) {
			_, _ = fmt.Fprintf(w, "\t-%s %s\t%s\n", f.Name, f.DefValue, f.Usage)
		})
		_ = w.Flush()
		_, _ = fmt.Fprintf(os.Stderr, "\n")
	}
}
