package main

import (
	"context"
	"maxblog-be-demo/internal/app"
	"maxblog-be-demo/internal/cmd/env"
)

func main() {
	config := env.LoadEnv()
	ctx := context.Background()
	app.Launch(
		ctx,
		app.SetConfigFile(*config),
	)
}
