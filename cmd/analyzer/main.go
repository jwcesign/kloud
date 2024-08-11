package main

import (
	"os"

	pkgserver "k8s.io/apiserver/pkg/server"
	"k8s.io/component-base/cli"

	"github.com/jwcesign/kloud/cmd/analyzer/app"
)

func main() {
	ctx := pkgserver.SetupSignalContext()
	cmd := app.NewAnalyzer(ctx)
	code := cli.Run(cmd)
	os.Exit(code)
}
