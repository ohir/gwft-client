package client

import (
	proto "example.com/mP"
)

const me string = "/mC Client"

func Say() string {
	return me[3:] + " (" + Vstamp + ") " +
		" said: tinker here.\n\tThen: " +
		proto.Say()
}
