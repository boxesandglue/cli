package main

import (
	"context"
	"fmt"
	"os"

	"github.com/boxesandglue/cli/risor/backend/bag"
	"github.com/boxesandglue/cli/risor/frontend"
	"github.com/risor-io/risor"
)

func dothings() error {
	data, err := os.ReadFile("main.rsr")
	if err != nil {
		return err
	}

	ctx := context.Background()

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	_, err = risor.Eval(ctx,
		string(data),
		risor.WithLocalImporter(wd),
		risor.WithGlobals(map[string]any{
			"frontend": frontend.Module(),
			"bag":      bag.Module(),
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
