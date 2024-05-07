package main

import (
	"context"

	box "github.com/sagernet/sing-box"
)

func check() error {
	options, err := readConfigAndMerge()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	instance, err := box.New(box.Options{
		Context: ctx,
		Options: options,
	})
	if err == nil {
		instance.Close()
	}
	cancel()
	return err
}
