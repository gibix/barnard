package main

import (
	"crypto/tls"

	"github.com/bmmcginty/barnard/config"
	"github.com/bmmcginty/barnard/gumble/gumble"
	"github.com/bmmcginty/barnard/gumble/gumbleopenal"
	"github.com/bmmcginty/barnard/uiterm"
)

type Barnard struct {
	Config     *gumble.Config
	UserConfig *config.Config
	Hotkeys    *config.Hotkeys
	Client     *gumble.Client

	Address   string
	TLSConfig tls.Config

	Stream    *gumbleopenal.Stream
	Tx        bool
	Connected bool

	Ui              *uiterm.Ui
	UiOutput        uiterm.Textview
	UiInput         uiterm.Textbox
	UiStatus        uiterm.Label
	UiTree          uiterm.Tree
	UiInputStatus   uiterm.Label
	SelectedChannel *gumble.Channel
	selectedUser    *gumble.User

	notifyChannel chan []string

	exitStatus  int
	exitMessage string
}
