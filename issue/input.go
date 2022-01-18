package issue

import (
	"fmt"

	"example.com/agent/agent"
	"example.com/agent/query"
)

func Issue_input(agents []agent.Agent) {
	var no_issues, query_mode int

	fmt.Print("\n\nEnter the no of issues: ")
	fmt.Scan(&no_issues)

	for no_issues > 0 {
		no_issues--

		var roles string
		fmt.Printf("Enter the roles of issue separated by comma: ")
		fmt.Scan(&roles)

		fmt.Printf("Enter the issue selection mode(1 for available, 2 for least busy(highest available since), 3 for random): ")
		fmt.Scan(&query_mode)

		query.Query(agents, roles, query_mode)
	}
}
