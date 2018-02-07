package types

type PollingPing struct {
	PendingCommands bool `json:"pending_commands" header:"PENDING_COMMANDS"`
}
