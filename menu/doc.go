// Copyright 2015 TD Internet Solutions, LLC. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
The go-menu package provides a library to build simple,
interactive, command line menus in Go.

Installation:

		import "github.com/turret-io/go-menu/menu"

Example:

		func cmd1(args ...string) error {
			// Do something
			fmt.Println("Output of cmd1")
		}

		func cmd2(args ...string) error {
			//Do something
			fmt.Println("Output of cmd2")
		}

		func main() {
			commandOptions := []menu.CommandOption{
				menu.CommandOption{"command1", "Runs command1", cmd1},
				menu.CommandOption{"command2", "Runs command2", cmd2},
			}

			menuOptions := menu.NewMenuOptions("'menu' for help > ", 0)

			menu := menu.NewMenu(commandOptions, menuOptions)
			menu.Start()
		}

Notes:

Typing "exit" or "quit" at the prompt will exit the program.

Typing "menu" will display the menu.
*/
package menu
