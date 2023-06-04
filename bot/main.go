package main

import (
	"github.com/dexslender/orb/bot/orb"
	"github.com/dexslender/orb/bot/util"
	"github.com/disgoorg/log"
)

func main() {
	l := log.New(log.Ltime | log.Lshortfile)
	c := util.SetupConfig(l)

	l.Info("Starting...")
	o := orb.New(l, c)

	o.SetupBot()

	o.StartNLock()
}
