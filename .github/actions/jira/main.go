package main

import (
	"fmt"
	"os"
)

func main() {
	jiraURL := os.Getenv("JIRA_URL")
	if jiraURL == "" {
		fmt.Println("JIRA_URL is not set")
		os.Exit(1)
	}

	// 打印 JIRA URL (可以替换为实际与 Jira API 交互的代码)
	fmt.Printf("The Jira base URL is: %s\n", jiraURL)

	// 此处可以实现更多逻辑与 Jira API 的交互
}
