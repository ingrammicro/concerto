package brownfield

import (
	"fmt"
	"runtime"

	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

func cmdConfigure(c *cli.Context) error {
	f := format.GetFormatter()
	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't read config", err)
	}
	cs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't set up connection to Concerto", err)
	}
	if !config.CurrentUserIsAdmin {
		if runtime.GOOS == "windows" {
			f.PrintFatal("Must run as administrator user", fmt.Errorf("running as non-administrator user"))
		} else {
			f.PrintFatal("Must run as super-user", fmt.Errorf("running as non-administrator user"))
		}
	}
	applyConcertoSettings(cs, f)
	configureConcertoFirewall(cs, f)
	return nil
}
