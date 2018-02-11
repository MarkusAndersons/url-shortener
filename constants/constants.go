// Package constants holds all constant values used in the application
package constants

const (
	// Port is the default port number to use if none specified
	Port string = "8080"
	// Hostname is the host to run the server on
	Hostname string = "127.0.0.1"
	// ShortLen is the number of characters for the short URL to be
	// smaller numbers will increase the chance of collisions
	ShortLen int = 6
)
