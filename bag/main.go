package main

import (
	"context"
	"fmt"
	"os"

	rbag "github.com/boxesandglue/cli/risor/backend/bag"
	rnode "github.com/boxesandglue/cli/risor/backend/node"
	rfrontend "github.com/boxesandglue/cli/risor/frontend"
	rcxpath "github.com/speedata/risorcxpath"

	"github.com/risor-io/risor"
)

// Version is the version of the program.
const Version = "0.1.0"

func dothings() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("usage: %s <filename>", os.Args[0])
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		return err
	}

	ctx := context.Background()
	setupLog()

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	_, err = risor.Eval(ctx,
		string(data),
		risor.WithLocalImporter(wd),
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
