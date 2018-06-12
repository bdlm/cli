package cli

import (
	"io"
	"os"
	"os/signal"
)

func New(sigHandler func(sig os.Signal)) *App {
	app := &App{
		Author:       "Bedlam Software <bdlm@webbedlam.com>",
		Copyright:    "2018, Bedlam Software",
		Name:         "example-app",
		Version:      "1.0",
		HandleSignal: sigHandler,
		signal:       make(chan os.Signal, 1),
	}
	signal.Notify(app.signal)
	go func() {
		for sig := range app.signal {
			app.HandleSignal(sig)
		}
	}()
	return app
}

type App struct {
	Author       string
	Copyright    string
	HandleSignal func(sig os.Signal)
	Help         string
	Name         string
	Version      string
	Commands     []*App

	StdErr io.Writer
	StdOut io.Writer
	StdIn  io.Reader

	signal chan os.Signal
}

func (app *App) exit(code int) {
	os.Exit(code)
}

func (app *App) usage() {

}
