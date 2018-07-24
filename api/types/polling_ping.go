package types

// PollingPing stores Polling Ping data
type PollingPing struct {
	PendingCommands bool `json:"pending_commands" header:"PENDING_COMMANDS"`
}
