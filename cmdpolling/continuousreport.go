package cmdpolling

import (
	"errors"
	"fmt"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// Polling Continuous Report Run, parameterized values
const (
	RetriesNumber        = 5
	RetriesFactor        = 3
	DefaultThresholdTime = 10
)

func cmdContinuousReportRun(c *cli.Context) error {
	log.Debug("cmdContinuousReportRun")

	formatter := format.GetFormatter()
	pollingSvc := cmd.WireUpPolling(c)

	// cli command argument
	var cmdArg string
	if c.Args().Present() {
		cmdArg = c.Args().First()
	} else {
		formatter.PrintFatal("argument missing", errors.New("a script or command is required"))
	}

	// cli command threshold flag
	thresholdTime := c.Int("time")
	if !(thresholdTime > 0) {
		thresholdTime = DefaultThresholdTime
	}
	log.Debug("Time threshold:", thresholdTime)

	// Custom method for chunks processing
	fn := func(chunk string) error {
		log.Debug("sendChunks")
		err := retry(RetriesNumber, time.Second, func() error {
			log.Debug("Sending: ", chunk)

			commandIn := map[string]interface{}{
				"stdout": chunk,
			}

			_, statusCode, err := pollingSvc.ReportBootstrapLog(&commandIn)
			switch {
			// 0<100 error cases??
			case statusCode == 0:
				return fmt.Errorf("communication error %v %v", statusCode, err)
			case statusCode >= 500:
				return fmt.Errorf("server error %v %v", statusCode, err)
			case statusCode >= 400:
				return fmt.Errorf("client error %v %v", statusCode, err)
			default:
				return nil
			}
		})

		if err != nil {
			return fmt.Errorf("cannot send the chunk data, %v", err)
		}
		return nil
	}

	exitCode, err := utils.RunContinuousCmd(fn, cmdArg, thresholdTime)
	if err != nil {
		formatter.PrintFatal("cannot process continuous report command", err)
	}

	log.Info("completed: ", exitCode)
	os.Exit(exitCode)
	return nil
}

func retry(attempts int, sleep time.Duration, fn func() error) error {
	log.Debug("retry")

	if err := fn(); err != nil {
		if attempts--; attempts > 0 {
			log.Debug("Waiting to retry: ", sleep)
			time.Sleep(sleep)
			return retry(attempts, RetriesFactor*sleep, fn)
		}
		return err
	}
	return nil
}
