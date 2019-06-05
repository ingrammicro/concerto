package volumes

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

// SubCommands returns volumes commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all existing volumes",
			Action: cmd.VolumeList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "server-id",
					Usage: "Identifier of a server to return only the volumes that are attached with that server",
				},
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label as a query filter",
				},
			},
		},
		{
			Name:   "show",
			Usage:  "Shows information about the volume identified by the given id",
			Action: cmd.VolumeShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Volume Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a new volume",
			Action: cmd.VolumeCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the volume",
				},
				cli.IntFlag{
					Name:  "size",
					Usage: "Size for the volume, in GB",
				},
				cli.StringFlag{
					Name:  "cloud-account-id",
					Usage: "Identifier of the cloud account in which the volume is",
				},
				cli.StringFlag{
					Name:  "storage-plan-id",
					Usage: "Identifier of the storage plan on which the volume is based",
				},
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label names to be associated with volume",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing volume identified by the given id",
			Action: cmd.VolumeUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Volume Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the volume",
				},
			},
		},
		{
			Name:   "attach",
			Usage:  "Attaches the volume to server",
			Action: cmd.VolumeAttach,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Volume Id",
				},
				cli.StringFlag{
					Name:  "server-id",
					Usage: "Identifier of the server to attach the volume",
				},
			},
		},
		{
			Name:   "detach",
			Usage:  "Detaches a volume from server",
			Action: cmd.VolumeDetach,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Volume Id",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a volume",
			Action: cmd.VolumeDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Volume Id",
				},
			},
		},
		{
			Name:   "discard",
			Usage:  "Discards a volume but does not delete it from the cloud provider",
			Action: cmd.VolumeDiscard,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Volume Id",
				},
			},
		},
		{
			Name:   "add-label",
			Usage:  "This action assigns a single label from a single labelable resource",
			Action: cmd.LabelAdd,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Volume Id",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Label name",
				},
				cli.StringFlag{
					Name:   "resource-type",
					Usage:  "Resource Type",
					Value:  "volume",
					Hidden: true,
				},
			},
		},
		{
			Name:   "remove-label",
			Usage:  "This action unassigns a single label from a single labelable resource",
			Action: cmd.LabelRemove,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Volume Id",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Label name",
				},
				cli.StringFlag{
					Name:   "resource-type",
					Usage:  "Resource Type",
					Value:  "volume",
					Hidden: true,
				},
			},
		},
	}
}
