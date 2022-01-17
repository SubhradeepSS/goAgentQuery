package agent

import (
	"fmt"
	"strings"
)

func Agent_input() []Agent {
	var no_agents int

	fmt.Printf("Enter the no of agents: ")
	fmt.Scan(&no_agents)

	agents := make([]Agent, no_agents)

	for idx := range agents {
		fmt.Printf("\nEnter data for agent %v\n", idx+1)

		var id int
		var is_available bool
		var available_since float64

		fmt.Printf("Enter id: ")
		fmt.Scan(&id)

		fmt.Printf("Enter is_available: ")
		fmt.Scan(&is_available)

		fmt.Printf("Enter available_since: ")
		fmt.Scan(&available_since)

		var roles string
		fmt.Printf("Enter the roles separated by comma: ")
		fmt.Scan(&roles)

		roles_splitted := strings.Split(roles, ",")

		agents[idx] = Agent{id, is_available, available_since, roles_splitted}
	}

	return agents
}
