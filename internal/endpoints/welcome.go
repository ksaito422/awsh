package endpoints

import (
	"fmt"
)

// Welcome screen to be displayed after entering a command.
func Welcome() {
	fmt.Print("    ####       ###             ###      #####      ##         \n")
	fmt.Print("   ######       ##     ###     ##      ##   ##     ##         \n")
	fmt.Print("  ##    ##       ##   ## ##   ##       ##          ## ###     \n")
	fmt.Print(" ##      ##      ##  ##  ##  ##         #####      ###  ###   \n")
	fmt.Print("  ##    ###       ####    ####              ##     ##    ##   \n")
	fmt.Print("   ###### ##      ###     ###          ##   ##     ##    ##   \n")
	fmt.Print("    ####   ##      #       #            #####      ##    ##   \n")
}
