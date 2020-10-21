// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/observer-framework/blob/master/LICENSE.md.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/insolar/insconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/insolar/observer-framework/configuration"
)

var stop = make(chan os.Signal, 1)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	cfg := &configuration.Observer{}
	params := insconfig.Params{
		EnvPrefix:        "observer",
		ConfigPathGetter: &insconfig.DefaultPathGetter{},
	}
	insConfigurator := insconfig.New(params)
	if err := insConfigurator.Load(cfg); err != nil {
		panic(err)
	}
	fmt.Println("Starts with configuration:\n", insConfigurator.ToYaml(cfg))

	pfefe := NewProfefe(cfg.Profefe, "observer")
	err := pfefe.Start()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	defer func() {
		err := pfefe.Stop()
		if err != nil {
			log.Error().Msg(err.Error())
		}
	}()

	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-stop:
		log.Info().Msg("gracefully stopping by signal")
	}
}
