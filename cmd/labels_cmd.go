package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/labels"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/utils/format"
)

// WireUpLabel prepares common resources to send request to Concerto API
func WireUpLabel(c *cli.Context) (ds *labels.LabelService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = labels.NewLabelService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up label service", err)
	}

	return ds, f
}

// LabelList subcommand function
func LabelList(c *cli.Context) error {
	debugCmdFuncInfo(c)

	labelsSvc, formatter := WireUpLabel(c)
	labels, err := labelsSvc.GetLabelList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive labels data", err)
	}

	if err = formatter.PrintList(labels); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// LabelCreate subcommand function
func LabelCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)

	labelsSvc, formatter := WireUpLabel(c)
	checkRequiredFlags(c, []string{"name", "resource_type"}, formatter)
	label, err := labelsSvc.CreateLabel(utils.FlagConvertParams(c))
	if err != nil {
		formatter.PrintFatal("Couldn't create label", err)
	}
	if err = formatter.PrintItem(*label); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// LabelFiltering subcommand function receives a collection of references to labelable objects
// Evaluates the matching of assigned labels with the labels requested for filtering.
func LabelFiltering(c *cli.Context, items []types.Labelable, labelIDsByName map[string]string) []types.Labelable {
	debugCmdFuncInfo(c)

	if c.String("labels") != "" {
		_, formatter := WireUpLabel(c)
		labelNamesIn := LabelsUnifyInputNames(c.String("labels"), formatter)
		var filteringLabelIDs []string
		for _, name := range labelNamesIn {
			id := labelIDsByName[name]
			if id != "" {
				filteringLabelIDs = append(filteringLabelIDs, id)
			}
		}
		var result []types.Labelable
		for _, item := range items {
			if item.FilterByLabelIDs(filteringLabelIDs) {
				result = append(result, item)
			}
		}
		return result
	}

	return items
}

// LabelAssignNamesForIDs subcommand function receives a collection of references to labelables objects
// Resolves the Labels names associated to a each resource from given Labels ids, loading object with respective labels names
func LabelAssignNamesForIDs(c *cli.Context, items []types.Labelable, labelNamesByID map[string]string) {
	debugCmdFuncInfo(c)
	for _, labelable := range items {
		labelable.FillInLabelNames(labelNamesByID)
	}
}

// LabelLoadsMapping subcommand function retrieves the current label list in IMCO; then prepares two mapping structures (Name <-> ID and ID <-> Name)
func LabelLoadsMapping(c *cli.Context) (map[string]string, map[string]string) {
	debugCmdFuncInfo(c)

	labelsSvc, formatter := WireUpLabel(c)
	labels, err := labelsSvc.GetLabelList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive labels data", err)
	}

	labelIDsByName := make(map[string]string)
	labelNamesByID := make(map[string]string)

	for _, label := range labels {
		labelIDsByName[label.Name] = label.ID
		labelNamesByID[label.ID] = label.Name
	}
	return labelIDsByName, labelNamesByID
}

// LabelsUnifyInputNames subcommand function evaluates the received labels names (comma separated string).
// Validates, remove duplicates and resolves a slice with unique label names.
func LabelsUnifyInputNames(labelsNames string, formatter format.Formatter) []string {
	labelNamesIn := utils.RemoveDuplicates(strings.Split(labelsNames, ","))
	for _, c := range labelNamesIn {
		matched := regexp.MustCompile(`^[A-Za-z0-9 .\s_-]+$`).MatchString(c)
		if !matched {
			formatter.PrintFatal("Invalid label name ", fmt.Errorf("Invalid label format: %v (Labels would be indicated with their name, which must satisfy to be composed of spaces, underscores, dots, dashes and/or lower/upper -case alphanumeric characters-)", c))
		}
	}
	return labelNamesIn
}

// LabelResolution subcommand function retrieves a labels map(Name<->ID) based on label names received to be processed.
// The function evaluates the received labels names (comma separated string); with them, solves the assigned IDs for the given labels names.
// If the label name is not available in IMCO yet, it is created.
func LabelResolution(c *cli.Context, labelsNames string) []string {
	debugCmdFuncInfo(c)

	labelsSvc, formatter := WireUpLabel(c)
	labelNamesIn := LabelsUnifyInputNames(labelsNames, formatter)
	labelsMapNameToID, _ := LabelLoadsMapping(c)

	// Obtain output mapped labels Name<->ID; currently in IMCO platform as well as if creation is required
	labelsOutMap := make(map[string]string)
	for _, name := range labelNamesIn {
		// check if the label already exists in IMCO, creates it if it does not exist
		if labelsMapNameToID[name] == "" {
			labelPayload := make(map[string]interface{})
			labelPayload["name"] = name
			newLabel, err := labelsSvc.CreateLabel(&labelPayload)
			if err != nil {
				formatter.PrintFatal("Couldn't create label", err)
			}
			labelsOutMap[name] = newLabel.ID
		} else {
			labelsOutMap[name] = labelsMapNameToID[name]
		}
	}
	labelsIdsArr := make([]string, 0)
	for _, mp := range labelsOutMap {
		labelsIdsArr = append(labelsIdsArr, mp)
	}
	return labelsIdsArr
}

// LabelAdd subcommand function assigns a single label from a single labelable resource
func LabelAdd(c *cli.Context) error {
	debugCmdFuncInfo(c)

	labelsSvc, formatter := WireUpLabel(c)
	checkRequiredFlags(c, []string{"id", "label"}, formatter)

	labelsIdsArr := LabelResolution(c, c.String("label"))
	if len(labelsIdsArr) > 1 {
		formatter.PrintFatal("Too many label names. Please, Use only one label name", fmt.Errorf("Invalid parameter: %v - %v", c.String("label"), labelsIdsArr))
	}
	labelID := labelsIdsArr[0]

	resData := make(map[string]string)
	resData["id"] = c.String("id")
	resData["resource_type"] = c.String("resource_type")
	resourcesData := make([]interface{}, 0, 1)
	resourcesData = append(resourcesData, resData)

	labelIn := map[string]interface{}{
		"resources": resourcesData,
	}

	labeledResources, err := labelsSvc.AddLabel(&labelIn, labelID)
	if err != nil {
		formatter.PrintFatal("Couldn't add label data", err)
	}
	if err = formatter.PrintList(labeledResources); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// LabelRemove subcommand function de-assigns a single label from a single labelable resource
func LabelRemove(c *cli.Context) error {
	debugCmdFuncInfo(c)

	labelsSvc, formatter := WireUpLabel(c)
	checkRequiredFlags(c, []string{"id", "label"}, formatter)

	labelsMapNameToID, _ := LabelLoadsMapping(c)
	labelID := labelsMapNameToID[c.String("label")]

	err := labelsSvc.RemoveLabel(labelID, c.String("resource_type"), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't remove label", err)
	}
	return nil
}
