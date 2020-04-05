package api

import (
	"github.com/pkg/errors"

	"github.com/mxc-foundation/chirpstack-application-server/pkg/api/as"
	"github.com/mxc-foundation/chirpstack-application-server/pkg/api/external"
	"github.com/mxc-foundation/chirpstack-application-server/pkg/api/js"
	"github.com/mxc-foundation/chirpstack-application-server/pkg/config"
)

func Setup(conf config.Config) error {
	if err := as.Setup(conf); err != nil {
		return errors.Wrap(err, "setup application-server api error")
	}

	if err := external.Setup(conf); err != nil {
		return errors.Wrap(err, "setup external api error")
	}

	if err := js.Setup(conf); err != nil {
		return errors.Wrap(err, "setup join-server api error")
	}

	return nil
}
