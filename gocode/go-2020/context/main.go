package main

import (
	"context"
	"fmt"
)

func f(ctx context.Context) {
	context.WithValue(ctx, "foo", -6)
}

func main() {
	ctx := context.TODO()
	// f(ctx)
	copy := context.WithValue(ctx, "foo", -6)
	fmt.Println(copy.Value("foo"))
}
