package main

import (
    "github.com/Redeltaz/gofm/cmd"
    "github.com/Redeltaz/gofm/pkg/controller"
)

func main() {
    controller := controller.InitController()

    cmd.Root(controller)
}
