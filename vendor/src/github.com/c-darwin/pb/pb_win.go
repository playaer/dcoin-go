// +build windows

package pb

import (
	"src/github.com/olekukonko/ts"
	"os"
)

var tty = os.Stdin

func terminalWidth() (int, error) {
	size, err := ts.GetSize()
	return size.Col(), err
}
