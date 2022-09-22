package main

import (
	"context"
	"flag"
	"maxblog-be-demo/internal/app"
)

func main() {
	config := flag.String("c", "env/dev.yaml", "配置文件")
	flag.Parse()
	ctx := context.Background()
	app.Launch(
		ctx,
		app.SetConfigFile(*config),
	)
}
