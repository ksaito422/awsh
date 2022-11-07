package endpoints

import (
	"awsh/pkg/prompt"
)

// Returns the name of the action for the resource being operated on.
func Operation() string {
	const promptDescription = "Select an action"
	// 操作対象のAWSリソース群
	opeResource := []string{"S3", "ECS"}
	var action string
	resource := prompt.ChooseValueFromPromptItems("Select the resource to be operated", opeResource)

	switch resource {
	case "S3":
		opeActions := []string{
			"ListBuckets",
			"ListObjects",
			"DownloadObject",
		}

		action = prompt.ChooseValueFromPromptItems(promptDescription, opeActions)
	case "ECS":
		opeActions := []string{
			"StartECS",
			"ecs-exec",
			"StopECSTask",
		}

		action = prompt.ChooseValueFromPromptItems(promptDescription, opeActions)
	}

	return action
}
