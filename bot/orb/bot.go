package orb

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/log"
	"github.com/spf13/viper"
)

var _orb Orb

func New(l log.Logger, v *viper.Viper) *Orb {
	_orb = Orb{Log: l, Config: v}
	return &_orb
}

type Orb struct {
	bot.Client
	Log    log.Logger
	Config *viper.Viper
}

func (o *Orb) SetupBot() {
	var err error
	if o.Client, err = disgo.New(
		o.Config.GetString("bot~token"),
		bot.WithGatewayConfigOpts(
			gateway.WithOS("mobile"),
			gateway.WithIntents(gateway.IntentsNonPrivileged),
			gateway.WithCompress(true),
		),
		bot.WithEventListeners(listeners),
		bot.WithLogger(o.Log),
	); err != nil {
		o.Log.Fatal("Client error: ", err)
	}
}

func (o *Orb) StartNLock() {
	ctx, c := context.WithTimeout(context.Background(), time.Second*10)
	defer c()
	defer func() {
		o.Close(ctx)
		o.Log.Info("Client closed, exiting program\n\tplease wait")
	}()
	if err := o.OpenGateway(ctx); err != nil {
		o.Log.Error("Gateway open step error: ", err)
	}
	k := make(chan os.Signal, 1)
	signal.Notify(k, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-k
}
