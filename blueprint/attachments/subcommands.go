package attachments

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

// SubCommands returns attachments commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "show",
			Usage:  "Shows information about a specific attachment",
			Action: cmd.AttachmentShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Attachment Id",
				},
			},
		},
		{
			Name:   "download",
			Usage:  "Downloads an attachment",
			Action: cmd.AttachmentDownload,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Attachment Id",
				},
				cli.StringFlag{
					Name:  "filepath",
					Usage: "path and file name to download attachment file, i.e: --filename /folder-path/filename.ext",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes an attachment",
			Action: cmd.AttachmentDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Attachment Id",
				},
			},
		},
	}
}
