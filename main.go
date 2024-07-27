package main

import (
	"context"

	"github.com/ruriazz/ocean-test/api"
	"github.com/ruriazz/ocean-test/pkg/manager"
)

func main() {
	ctx := context.Background()
	mgr := manager.New(ctx)

	api.Init(mgr)
	mgr.StartServer()
}
