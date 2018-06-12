package cli_test

import (
	"testing"

	"github.com/bdlm/cli"
)

func TestCli(t *testing.T) {
	app := &cli.App{
		Author:    "Bedlam Software <bdlm@webbedlam.com>",
		Copyright: "2018, Bedlam Software",
		Name:      "example-app",
		Version:   "1.0",
	}

	if nil == app {
		t.Errorf("app is nil")
	}
}
