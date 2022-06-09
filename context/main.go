package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, 1, "sultan")
	speek(ctx)
}

func speek(ctx context.Context) {
	fmt.Println("1=", ctx.Value(1))

	newkey := context.WithValue(ctx, 2, "ram")
	quite(newkey)
	fmt.Println("1=", ctx.Value(1))
}

func quite(ctx context.Context) {
	fmt.Println("2=", ctx.Value(2))
}
