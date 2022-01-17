package main

import (
	"fmt"

	"example.com/agent"
	"example.com/issue"
)

func main() {
	agents := agent.Agent_input()

	fmt.Print(agents)

	issue.Issue_input(agents)
}
