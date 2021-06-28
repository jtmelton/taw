package domain

// Options struct represents the cli options passed to the taw tool
type Options struct {
	InputDirectory string
	OutputFile     string
}

// ExtensionCount struct represents an extension and its' associated count
type ExtensionCount struct {
	Extension string `json:"extension"`
	Count     int    `json:"count"`
}

// ExtensionCounts struct represents an array of extension counts
type ExtensionCounts struct {
	ExtensionCounts []ExtensionCount `json:"extension-counts"`
}
