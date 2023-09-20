package main

import (
	"os"

	"github.com/insolar/insconfig"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/insolar/observer-framework/configuration"
)

func main() {
	configs := map[string]interface{}{
		".artifacts/observer.yaml": configuration.Observer{},
	}

	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	for filePath, config := range configs {
		func() {
			f, err := os.Create(filePath)
			if err != nil {
				log.Fatal().Msg(errors.Wrapf(err, "failed to create config file %s", filePath).Error())
			}
			err = insconfig.NewYamlTemplater(config).TemplateTo(f)

			defer func() {
				err := f.Close()
				if err != nil {
					log.Fatal().Msg(errors.Wrapf(err, "failed to close config file %s", filePath).Error())
				}
			}()

			if err != nil {
				log.Fatal().Msg(errors.Wrapf(err, "failed to write config file %s", filePath).Error())
			}
		}()
	}
}
