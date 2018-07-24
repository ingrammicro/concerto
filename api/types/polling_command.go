package types

// PollingCommand stores Polling Command data
type PollingCommand struct {
	ID       string `json:"id" header:"ID"`
	Script   string `json:"script" header:"SCRIPT"`
	Stdout   string `json:"stdout" header:"STDOUT"`
	Stderr   string `json:"stderr" header:"STDERR"`
	ExitCode int    `json:"exit_code" header:"EXIT_CODE"`
}
