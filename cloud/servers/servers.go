package servers

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/codegangsta/cli"
	"github.com/ingrammicro/concerto/api/types"
	"github.com/ingrammicro/concerto/utils"
	"github.com/ingrammicro/concerto/webservice"
)

func cmdExecuteScript(c *cli.Context) error {
	utils.FlagsRequired(c, []string{"server_id", "script_id"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	res, code, err := webservice.Put(fmt.Sprintf("/v1/cloud/servers/%s/operational_scripts/%s/execute", c.String("server_id"), c.String("script_id")), nil)
	utils.CheckError(err)
	utils.CheckReturnCode(code, res)

	var event types.Event
	err = json.Unmarshal(res, &event)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tTIMESTAMP\tLEVEL\tHEADER\tDESCRIPTION\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", event.ID, event.Timestamp, event.Level, event.Header, event.Description)

	w.Flush()
	return nil
}
