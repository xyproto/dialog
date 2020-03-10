package dialog

import (
	"bytes"
	"os"
	"os/exec"
	"sort"
	"strconv"
)

// Dialog represents a message box, menu or similar text widget.
// Also contains the path to the dialog executable.
type Dialog struct {
	width  int
	height int
	path   string
}

// New will create a new Dialog struct
func New(width, height int) *Dialog {
	return &Dialog{width, height, "/usr/bin/dialog"}
}

// SetPath will set the full path to the dialog utility (eg. "/usr/bin/dialog")
func (d *Dialog) SetPath(path string) {
	d.path = path
}

// MsgBox can display a message box with the given text to the user
func (d *Dialog) MsgBox(text string) error {
	cmd := exec.Command(d.path, "--msgbox", text, strconv.Itoa(d.height), strconv.Itoa(d.width))
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

// Menu can present a menu for the user
// Takes:
// * A message for the user
// * The desired height (counted in lines)
// * A map of menu itmes (item label -> item value)
// Returns the chosen item value.
func (d *Dialog) Menu(text string, menuheight int, menuItemPairs map[string]string) (string, error) {
	args := []string{"--menu", text, strconv.Itoa(d.height), strconv.Itoa(d.width), strconv.Itoa(menuheight)}

	// Sort the keys
	keys := make([]string, len(menuItemPairs))
	i := 0
	for key := range menuItemPairs {
		keys[i] = key
		i++
	}
	sort.Strings(keys)

	// Append the menu entries in sorted order
	for _, key := range keys {
		args = append(args, key)
		args = append(args, menuItemPairs[key])
	}

	// Run the dialog tool
	cmd := exec.Command(d.path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	var out bytes.Buffer
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}

	// Return the output from the dialog tool
	return out.String(), nil
}

// YesNo presents a yes/no choice to the user. Returns true if "yes".
func (d *Dialog) YesNo(text string) (bool, error) {
	cmd := exec.Command(d.path, "--yesno", text, strconv.Itoa(d.height), strconv.Itoa(d.width))
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	if err := cmd.Start(); err != nil {
		return false, err
	}
	return cmd.Wait() == nil, nil
}
