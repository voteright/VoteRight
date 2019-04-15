package config

// Config represents a configuration for the primary voting server
type Config struct {
	ListenURL           string
	DatabaseFile        string
	Verification        bool
	VerificationServers []string
}
