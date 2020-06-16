package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	"github.com/docker/machine/drivers/generic"
)

func main() {
	plugin.RegisterDriver(generic.NewDriver("", ""))
}
