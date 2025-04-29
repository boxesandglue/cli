package main

import (
	"context"
	"fmt"
	"os"

	rbag "github.com/boxesandglue/cli/risor/backend/bag"
	rnode "github.com/boxesandglue/cli/risor/backend/node"

	"github.com/boxesandglue/cli/risor/frontend"
	"github.com/risor-io/risor"
)

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
			"frontend": frontend.Module(),
			"bag":      rbag.Module(),
			"node":     rnode.Module(),
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
