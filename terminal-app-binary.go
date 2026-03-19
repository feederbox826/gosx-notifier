package gosxnotifier

import _ "embed"

//go:embed terminal-notifier.temp.zip
var zipFile []byte

// terminalnotifier returns raw, uncompressed file data.
func terminalnotifier() []byte {
	return zipFile
}
