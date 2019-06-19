package blueprint

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
)

// ScriptService manages scripts operations
type ScriptService struct {
	concertoService utils.ConcertoService
}

// NewScriptService returns a Concerto script service
func NewScriptService(concertoService utils.ConcertoService) (*ScriptService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &ScriptService{
		concertoService: concertoService,
	}, nil
}

// GetScriptList returns the list of scripts as an array of Scripts
func (sc *ScriptService) GetScriptList() (scripts []*types.Script, err error) {
	log.Debug("GetScriptsList")

	data, status, err := sc.concertoService.Get("/blueprint/scripts")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &scripts); err != nil {
		return nil, err
	}

	return scripts, nil
}

// GetScript returns a script by its ID
func (sc *ScriptService) GetScript(scriptID string) (script *types.Script, err error) {
	log.Debug("GetScript")

	data, status, err := sc.concertoService.Get(fmt.Sprintf("/blueprint/scripts/%s", scriptID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &script); err != nil {
		return nil, err
	}

	return script, nil
}

// CreateScript creates a script
func (sc *ScriptService) CreateScript(scriptVector *map[string]interface{}) (script *types.Script, err error) {
	log.Debug("CreateScript")

	data, status, err := sc.concertoService.Post("/blueprint/scripts", scriptVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &script); err != nil {
		return nil, err
	}

	return script, nil
}

// UpdateScript updates a script by its ID
func (sc *ScriptService) UpdateScript(scriptVector *map[string]interface{}, scriptID string) (script *types.Script, err error) {
	log.Debug("UpdateScript")

	data, status, err := sc.concertoService.Put(fmt.Sprintf("/blueprint/scripts/%s", scriptID), scriptVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &script); err != nil {
		return nil, err
	}

	return script, nil
}

// DeleteScript deletes a script by its ID
func (sc *ScriptService) DeleteScript(scriptID string) (err error) {
	log.Debug("DeleteScript")

	data, status, err := sc.concertoService.Delete(fmt.Sprintf("/blueprint/scripts/%s", scriptID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// AddScriptAttachment adds an attachment to script by its ID
func (sc *ScriptService) AddScriptAttachment(attachmentIn *map[string]interface{}, scriptID string) (script *types.Attachment, err error) {
	log.Debug("AddScriptAttachment")

	data, status, err := sc.concertoService.Post(fmt.Sprintf("/blueprint/scripts/%s/attachments", scriptID), attachmentIn)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &script); err != nil {
		return nil, err
	}

	return script, nil
}

// UploadScriptAttachment uploads an attachment file
func (sc *ScriptService) UploadScriptAttachment(sourceFilePath string, targetURL string) error {
	log.Debug("UploadScriptAttachment")

	data, status, err := sc.concertoService.PutFile(sourceFilePath, targetURL)
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// UploadedScriptAttachment sets "uploaded" status to the attachment by its ID
func (sc *ScriptService) UploadedScriptAttachment(attachmentVector *map[string]interface{}, attachmentID string) (attachment *types.Attachment, err error) {
	log.Debug("UploadedScriptAttachment")

	data, status, err := sc.concertoService.Put(fmt.Sprintf("/blueprint/attachments/%s/uploaded", attachmentID), attachmentVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &attachment); err != nil {
		return nil, err
	}

	return attachment, nil
}

// ListScriptAttachments returns the list of Attachments for a given script ID
func (sc *ScriptService) ListScriptAttachments(scriptID string) (attachments []*types.Attachment, err error) {
	log.Debug("ListScriptAttachments")

	data, status, err := sc.concertoService.Get(fmt.Sprintf("/blueprint/scripts/%s/attachments", scriptID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &attachments); err != nil {
		return nil, err
	}

	return attachments, nil
}
