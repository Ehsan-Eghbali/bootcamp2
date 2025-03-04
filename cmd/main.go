package main

import (
	"bootcamp2/bootstrap"
	"context"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	bootstrap.Initiate(ctx)
}
