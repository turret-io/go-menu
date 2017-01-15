package menu

import (
	"bytes"
	"fmt"
	"testing"
)

const cmd = "cmd"
const desc = "this is a command"

func getOpts() []CommandOption {
	return []CommandOption{
		CommandOption{cmd, desc, func(...string) error {
			fmt.Println("in function")
			return nil
		}},
	}
}

func getMenuOpts() MenuOptions {
	return NewMenuOptions("", 0, "")
}

// Test menu creation with default options
func TestMenuOptions(t *testing.T) {
	cmdOpts := getOpts()
	menuOpts := MenuOptions{}

	menu := NewMenu(cmdOpts, menuOpts)

	if menu.Commands[0].Command != cmd {
		t.Error("Command is not set")
	}

	if menu.Commands[0].Description != desc {
		t.Error("Description is not set")
	}

	if menu.Options.Prompt != "> " {
		t.Error("Unexpected prompt")
	}

	if menu.Options.MenuLength != 100 {
		t.Error("Unexpected menu length")
	}

	if menu.Options.MenuCommand != "menu" {
		t.Error("Unexpected MenuCommand")
	}
}

// Test that the menu struct is created
func TestSimpleMenu(t *testing.T) {
	cmdOpts := getOpts()
	menuOpts := getMenuOpts()

	menu := NewMenu(cmdOpts, menuOpts)

	if menu.Commands[0].Command != cmd {
		t.Error("Command is not set")
	}

	if menu.Commands[0].Description != desc {
		t.Error("Description is not set")
	}

	if menu.Options.Prompt != "> " {
		t.Error("Unexpected prompt")
	}

	if menu.Options.MenuLength != 100 {
		t.Error("Unexpected menu length")
	}
}

// Run a simple test on the menu using junk as input
func TestJunkInput(t *testing.T) {
	cmdOpts := getOpts()
	menuOpts := getMenuOpts()

	menu := NewMenu(cmdOpts, menuOpts)

	input := bytes.NewReader([]byte("blah\n"))
	menu.start(input)
}

// Run a simple test using good data as input
func TestGoodInput(t *testing.T) {
	cmdOpts := getOpts()
	menuOpts := getMenuOpts()

	menu := NewMenu(cmdOpts, menuOpts)

	input := bytes.NewReader([]byte(fmt.Sprintf("%s\n", cmd)))
	menu.start(input)
}
