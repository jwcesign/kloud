package main

import (
	pkgserver "k8s.io/apiserver/pkg/server"
)

func main() {
	ctx := pkgserver.SetupSignalContext()

}
