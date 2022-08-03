package route

import (
	"awsh/pkg/prompt"
)

func Main() string {
	// 操作対象のAWSリソース群
	route := []string{"S3", "ECS"}
	var select_action string
	select_route := prompt.ChooseValueFromPromptItems("Select the resource to be operated", route)

	switch select_route {
	case "S3":
		actions := []string{
			"ListBuckets",
			"ListObjects",
			"GetObject",
			"DownloadObject",
		}

		select_action = prompt.ChooseValueFromPromptItems("Select an action", actions)
	case "ECS":
		// TODO: アクションを追加する
	}

	return select_action
}
