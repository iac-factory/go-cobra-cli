package example

import "gitlab.com/ethrgg/templates/go-cobra-cli/internal/types/output"

var (
	name   string        // an example named-argument
	format = output.JSON // setting to a constant provides a default, valid value
)
