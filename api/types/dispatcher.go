package types

type ScriptCharacterization struct {
	Order      int               `json:"execution_order"`
	UUID       string            `json:"uuid"`
	Script     DispatcherScript  `json:"script"`
	Parameters map[string]string `json:"parameter_values"`
}

type DispatcherScript struct {
	Code            string   `json:"code"`
	UUID            string   `json:"uuid"`
	AttachmentPaths []string `json:"attachment_paths"`
}

type ScriptConclusion struct {
	UUID       string `json:"script_characterization_id"`
	Output     string `json:"output"`
	ExitCode   int    `json:"exit_code"`
	StartedAt  string `json:"started_at"`
	FinishedAt string `json:"finished_at"`
}
