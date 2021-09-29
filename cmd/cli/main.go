package main

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/Insulince/jtone/cmd/cli/cli"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jtone := cli.New()

	if err := jtone.ExecuteContext(ctx); err != nil {
		fmt.Println(errors.Wrap(err, "error"))
	}
}
