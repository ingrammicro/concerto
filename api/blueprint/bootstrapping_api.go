package blueprint

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
)


// BootstrappingService manages bootstrapping operations
type BootstrappingService struct {
	concertoService utils.ConcertoService
}

// NewBootstrappingService returns a bootstrapping service
func NewBootstrappingService(concertoService utils.ConcertoService) (*BootstrappingService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &BootstrappingService{
		concertoService: concertoService,
	}, nil

}

// GetBootstrappingConfiguration returns the list of policy files as a JSON response with the desired configuration changes
func (bs *BootstrappingService) GetBootstrappingConfiguration() (bootstrappingConfigurations *types.BootstrappingConfiguration, status int, err error) {
	log.Debug("GetBootstrappingConfiguration")

	data, status, err := bs.concertoService.Get("/blueprint/configuration")
	if err != nil {
		return nil, status, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, status, err
	}

	if err = json.Unmarshal(data, &bootstrappingConfigurations); err != nil {
		return nil, status, err
	}

	return bootstrappingConfigurations, status, nil
}

// ReportBootstrappingAppliedConfiguration
func (bs *BootstrappingService) ReportBootstrappingAppliedConfiguration(BootstrappingAppliedConfigurationVector *map[string]interface{}) (err error) {
	log.Debug("ReportBootstrappingAppliedConfiguration")

	data, status, err := bs.concertoService.Put("/blueprint/applied_configuration", BootstrappingAppliedConfigurationVector)
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// ReportBootstrappingLog reports a policy files application result
func (bs *BootstrappingService) ReportBootstrappingLog(BootstrappingContinuousReportVector *map[string]interface{}) (command *types.BootstrappingContinuousReport, status int, err error) {
	log.Debug("ReportBootstrappingLog")

	data, status, err := bs.concertoService.Post("/blueprint/bootstrap_logs", BootstrappingContinuousReportVector)
	if err != nil {
		return nil, status, err
	}

	if err = json.Unmarshal(data, &command); err != nil {
		return nil, status, err
	}

	return command, status, nil
}


//
func (bs *BootstrappingService) DownloadPolicyFile(url string, filePath string) (realFileName string, status int, err error) {
	log.Debug("DownloadPolicyFile")

	realFileName, status, err = bs.concertoService.GetFile(url, filePath)
	if err != nil {
		return realFileName, status, err
	}

	return realFileName, status, nil
}