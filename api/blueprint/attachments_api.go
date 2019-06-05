package blueprint

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
)

// AttachmentService manages attachments operations
type AttachmentService struct {
	concertoService utils.ConcertoService
}

// NewAttachmentService returns a Concerto attachment service
func NewAttachmentService(concertoService utils.ConcertoService) (*AttachmentService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("must initialize ConcertoService before using it")
	}

	return &AttachmentService{
		concertoService: concertoService,
	}, nil
}

// GetAttachment returns a attachment by its ID
func (as *AttachmentService) GetAttachment(attachmentID string) (attachment *types.Attachment, err error) {
	log.Debug("GetAttachment")

	data, status, err := as.concertoService.Get(fmt.Sprintf("/blueprint/attachments/%s", attachmentID))
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

// DownloadAttachment gets an attachment file from given url saving file into given file path
func (as *AttachmentService) DownloadAttachment(url string, filePath string) (realFileName string, status int, err error) {
	log.Debug("DownloadAttachment")

	realFileName, status, err = as.concertoService.GetFile(url, filePath, false)
	if err != nil {
		return realFileName, status, err
	}

	return realFileName, status, nil
}

// DeleteAttachment deletes a attachment by its ID
func (as *AttachmentService) DeleteAttachment(attachmentID string) (err error) {
	log.Debug("DeleteAttachment")

	data, status, err := as.concertoService.Delete(fmt.Sprintf("/blueprint/attachments/%s", attachmentID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
