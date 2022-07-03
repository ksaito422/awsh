package prompt

import (
	"log"

	"github.com/manifoldco/promptui"
)

func ChooseValueFromPrompt(l string, d string) string {
	prompt := promptui.Prompt{
		Label:   l,
		Default: d,
	}
	v, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	return v
}

