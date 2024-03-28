package snaparser_server

import (
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

// Flags for configuring the snaparser server using command-line arguments.
var flags = []cli.Flag{
	altsrc.NewStringSliceFlag(
		&cli.StringSliceFlag{
			Name:    "allowed-origins",
			Aliases: []string{"o"},
			Usage:   "allow `ORIGIN(S)`",
			Value:   nil,
		},
	),
	altsrc.NewStringSliceFlag(
		&cli.StringSliceFlag{
			Name:    "allowed-methods",
			Aliases: []string{"m"},
			Usage:   "allow `METHOD(S)`",
			Value:   cli.NewStringSlice("POST"),
		},
	),
	altsrc.NewIntFlag(
		&cli.IntFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Usage:   "listen on `PORT`",
			Value:   8888,
		},
	),
	altsrc.NewStringFlag(
		&cli.StringFlag{
			Name:    "endpoint",
			Aliases: []string{"e"},
			Usage:   "upload endpoint `PATH`",
			Value:   "/upload",
		},
	),
	altsrc.NewStringFlag(
		&cli.StringFlag{
			Name:    "limiter",
			Aliases: []string{"i"},
			Usage:   "HTTP rate limiter type (none, lenient, normal, strict)",
			Value:   "none",
		},
	),
	altsrc.NewBoolFlag(
		&cli.BoolFlag{
			Name:    "log",
			Aliases: []string{"g"},
			Usage:   "enable logging",
			Value:   false,
		},
	),
	altsrc.NewBoolFlag(
		&cli.BoolFlag{
			Name:    "http",
			Aliases: []string{"t"},
			Usage:   "use HTTP instead of HTTPS",
			Value:   false,
		},
	),
	altsrc.NewPathFlag(
		&cli.PathFlag{
			Name:    "key-file",
			Aliases: []string{"k"},
			Usage:   "SSL secret key file `PATH`",
		},
	),
	altsrc.NewPathFlag(
		&cli.PathFlag{
			Name:    "cert-file",
			Aliases: []string{"c"},
			Usage:   "SSL certificate file `PATH`",
		},
	),
	&cli.StringFlag{
		Name:    "load",
		Aliases: []string{"l"},
		Usage:   "load TOML configuration from `FILE`",
	},
}

// action is the action to be executed when the snaparser server is run.
var action = func(ctx *cli.Context) error {
	return run(
		&RunConfig{
			allowedOrigins: ctx.StringSlice("allowed-origins"),
			allowedMethods: ctx.StringSlice("allowed-methods"),
			keyFilePath:    ctx.Path("key-file"),
			certFilePath:   ctx.Path("cert-file"),
			endpoint:       ctx.String("endpoint"),
			port:           ctx.Int("port"),
			limiterType:    limiterType(ctx.String("limiter")),
			useHTTP:        ctx.Bool("http"),
			enableLog:      ctx.Bool("log"),
		},
	)
}

// App represents the CLI application for running the snaparser server.
var App = &cli.App{
	Name:    "snaparser_server",
	Suggest: true,
	Version: "0.0.1",
	Authors: []*cli.Author{{Name: "vanillaiice", Email: "vanillaiice1@proton.me"}},
	Usage:   "backend server for snaparser",
	Flags:   flags,
	Before:  altsrc.InitInputSourceWithContext(flags, altsrc.NewTomlSourceFromFlagFunc("load")),
	Action:  action,
}
