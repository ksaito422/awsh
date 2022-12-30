package endpoints

import (
	"awsh/pkg/prompt"
)

// Returns the name of the action for the resource being operated on.
func (r *Route) Operation() string {
	const promptDescription = "Select an action"
	// 操作対象のAWSリソース群
	or := []string{
		string(S3),
		string(ECS),
	}

	var action string
	resource := prompt.ChooseValueFromPromptItems("Select the resource to be operated", or)

	switch resource {
	case "S3":
		actions := []string{
			ListBuckets.String(),
			ListObjects.String(),
			DownloadObject.String(),
		}

		action = prompt.ChooseValueFromPromptItems(promptDescription, actions)
	case "ECS":
		actions := []string{
			ECS_EXEC.String(),
		}

		action = prompt.ChooseValueFromPromptItems(promptDescription, actions)
	}

	return action
}
