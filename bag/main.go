package main

import (
	"context"
	"fmt"
	"os"

	rbag "github.com/boxesandglue/cli/risor/backend/bag"
	rnode "github.com/boxesandglue/cli/risor/backend/node"
	rfrontend "github.com/boxesandglue/cli/risor/frontend"
	"github.com/speedata/optionparser"
	rcxpath "github.com/speedata/risorcxpath"

	"github.com/risor-io/risor"
)

// Version is the version of the program.
var Version string

func dothings() error {
	defaults := map[string]string{
		"loglevel": "info",
	}
	op := optionparser.NewOptionParser()
	op.Banner = "bag - a frontend for boxes and glue"
	op.Coda = "\nUsage: bag [options] <filename>"
	op.On("--loglevel LVL", "Set the log level (debug, info, warn, error)", defaults)
	op.Command("version", "Print version and exit")
	if err := op.Parse(); err != nil {
		return err
	}
	var mainfile string
	for _, arg := range op.Extra {
		switch arg {
		case "version":
			fmt.Printf("bag version %s\n", Version)
			return nil
		default:
			mainfile = arg
		}
	}

	if mainfile == "" {
		return fmt.Errorf("usage: %s <filename>", os.Args[0])
	}
	data, err := os.ReadFile(mainfile)
	if err != nil {
		return err
	}

	ctx := context.Background()

	setupLog(defaults["loglevel"])

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	_, err = risor.Eval(ctx,
		string(data),
		risor.WithLocalImporter(wd),
		risor.WithConcurrency(),
		risor.WithGlobals(map[string]any{
			"frontend": rfrontend.Module(),
			"bag":      rbag.Module(),
			"node":     rnode.Module(),
			"cxpath":   rcxpath.Module(),
		}))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := dothings(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
