package prompt

import (
	"log"

	"github.com/manifoldco/promptui"
)

// Invokes a selective prompt that expands an array of
// arguments and returns the selected value.
func ChooseValueFromPromptItems(l string, i []string) string {
	prompt := promptui.Select{
		Label: l,
		Items: i,
	}
	_, v, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	return v
}
