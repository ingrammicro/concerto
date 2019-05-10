package templates

import (
	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/cmd"
)

// SubCommands returns templates commands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all available templates",
			Action: cmd.TemplateList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label as a query filter",
				},
			},
		},
		{
			Name:   "show",
			Usage:  "Shows information about a specific template",
			Action: cmd.TemplateShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Template Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a new template.",
			Action: cmd.TemplateCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the template",
				},
				cli.StringFlag{
					Name:  "generic-image-id",
					Usage: "Identifier of the OS image that the template builds on",
				},
				cli.StringFlag{
					Name:  "run-list",
					Usage: "A list of comma separated cookbook recipes that is run on the servers at start-up, i.e: --run-list imco::client,1password,wordpress",
				},
				cli.StringFlag{
					Name:  "cookbook-versions",
					Usage: "The cookbook versions used to configure the service recipes in the run-list, i.e: --cookbook-versions \"imco:3.0.3,1password~>1.3.0,wordpress:0.1.0\" \n\tCookbook version format: [NAME<OPERATOR>VERSION] \n\tSupported Operators:\n\t\tChef supermarket cookbook '~>','=','>=','>','<','<='\n\t\tUploaded cookbook ':'",
				},
				cli.StringFlag{
					Name:  "configuration-attributes",
					Usage: "The attributes used to configure the service recipes in the run-list, as a json formatted parameter",
				},
				cli.StringFlag{
					Name:  "configuration-attributes-from-file",
					Usage: "The attributes used to configure the service recipes in the run-list, from file or STDIN, as a json formatted parameter. \n\tFrom file: --configuration-attributes-from-file attrs.json \n\tFrom STDIN: --configuration-attributes-from-file -",
				},
				cli.StringFlag{
					Name:  "labels",
					Usage: "A list of comma separated label names to be associated with template",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing template",
			Action: cmd.TemplateUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Template Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the template",
				},
				cli.StringFlag{
					Name:  "run-list",
					Usage: "A list of comma separated cookbook recipes that is run on the servers at start-up, i.e: --run-list imco::client,1password,wordpress",
				},
				cli.StringFlag{
					Name:  "cookbook-versions",
					Usage: "The cookbook versions used to configure the service recipes in the run-list, i.e: --cookbook-versions \"imco:3.0.3,1password~>1.3.0,wordpress:0.1.0\" \n\tCookbook version format: [NAME<OPERATOR>VERSION] \n\tSupported Operators:\n\t\tChef supermarket cookbook '~>','=','>=','>','<','<='\n\t\tUploaded cookbook ':'",
				},
				cli.StringFlag{
					Name:  "configuration-attributes",
					Usage: "The attributes used to configure the service recipes in the run-list, as a json formatted parameter",
				},
				cli.StringFlag{
					Name:  "configuration-attributes-from-file",
					Usage: "The attributes used to configure the service recipes in the run-list, from file or STDIN, as a json formatted parameter. \n\tFrom file: --configuration-attributes-from-file attrs.json \n\tFrom STDIN: --configuration-attributes-from-file -",
				},
			},
		},
		{
			Name:   "compile",
			Usage:  "Compiles an existing template",
			Action: cmd.TemplateCompile,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Template Id",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a template",
			Action: cmd.TemplateDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Template Id",
				},
			},
		},
		{
			Name:   "list-template-scripts",
			Usage:  "Shows the script characterisations of a template",
			Action: cmd.TemplateScriptList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "template-id",
					Usage: "Template Id",
				},
				cli.StringFlag{
					Name:  "type",
					Usage: "Must be \"operational\", \"boot\" or \"shutdown\"",
				},
			},
		},
		{
			Name:   "show-template-script",
			Usage:  "Shows information about a specific script characterisation",
			Action: cmd.TemplateScriptShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "template-id",
					Usage: "Template Id",
				},
				cli.StringFlag{
					Name:  "id",
					Usage: "Script Id",
				},
			},
		},
		{
			Name:   "create-template-script",
			Usage:  "Creates a new script characterisation for a template and appends it to the list of script characterisations of the same type.",
			Action: cmd.TemplateScriptCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "template-id",
					Usage: "Template Id",
				},
				cli.StringFlag{
					Name:  "type",
					Usage: "Must be \"operational\", \"boot\" or \"shutdown\"",
				},
				cli.StringFlag{
					Name:  "script-id",
					Usage: "Identifier for the script that is parameterised by the script characterisation",
				},
				cli.StringFlag{
					Name:  "parameter-values",
					Usage: "A map that assigns a value to each script parameter, as a json formatted parameter; i.e: '{\"param1\":\"val1\",\"param2\":\"val2\"}'",
				},
				cli.StringFlag{
					Name:  "parameter-values-from-file",
					Usage: "A map that assigns a value to each script parameter, from file or STDIN, as a json formatted parameter. \n\tFrom file: --parameter-values-from-file params.json \n\tFrom STDIN: --parameter-values-from-file -",
				},
			},
		},
		{
			Name:   "update-template-script",
			Usage:  "Updates an existing script characterisation for a template.",
			Action: cmd.TemplateScriptUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "template-id",
					Usage: "Template Id",
				},
				cli.StringFlag{
					Name:  "id",
					Usage: "Identifier for the template-script that is parameterised by the script characterisation",
				},
				cli.StringFlag{
					Name:  "parameter-values",
					Usage: "A map that assigns a value to each script parameter, as a json formatted parameter; i.e: '{\"param1\":\"val1\",\"param2\":\"val2\"}'",
				},
				cli.StringFlag{
					Name:  "parameter-values-from-file",
					Usage: "A map that assigns a value to each script parameter, from file or STDIN, as a json formatted parameter. \n\tFrom file: --parameter-values-from-file params.json \n\tFrom STDIN: --parameter-values-from-file -",
				},
			},
		},
		{
			Name:   "reorder-template-scripts",
			Usage:  "Reorders the scripts of the template and type specified according to the provided order, changing their execution order as corresponds.",
			Action: cmd.TemplateScriptReorder,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "template-id",
					Usage: "Template Id",
				},
				cli.StringFlag{
					Name:  "type",
					Usage: "Must be \"operational\", \"boot\", or \"shutdown\"",
				},
				cli.StringFlag{
					Name:  "script-ids",
					Usage: "A list of comma separated scripts ids that must contain all the ids of scripts of the given template and type in the desired execution order",
				},
			},
		},
		{
			Name:   "delete-template-script",
			Usage:  "Removes a parametrized script from a template",
			Action: cmd.TemplateScriptDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "template-id",
					Usage: "Template Id",
				},
				cli.StringFlag{
					Name:  "id",
					Usage: "Identifier for the template-script that is parameterised by the script characterisation",
				},
			},
		},
		{
			Name:   "list-template-servers",
			Usage:  "Returns information about the servers that use a specific template. ",
			Action: cmd.TemplateServersList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "template-id",
					Usage: "Template Id",
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
					Usage: "Template Id",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Label name",
				},
				cli.StringFlag{
					Name:   "resource-type",
					Usage:  "Resource Type",
					Value:  "template",
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
					Usage: "Template Id",
				},
				cli.StringFlag{
					Name:  "label",
					Usage: "Label name",
				},
				cli.StringFlag{
					Name:   "resource-type",
					Usage:  "Resource Type",
					Value:  "template",
					Hidden: true,
				},
			},
		},
	}
}
