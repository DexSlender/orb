package orb

import "github.com/disgoorg/disgo/events"

var listeners = &events.ListenerAdapter{
    OnReady: func(event *events.Ready) {
        _orb.Log.Info("Logged in as: ", event.User.Tag())
    },
}