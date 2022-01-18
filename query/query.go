package query

import (
	"fmt"
	"math/rand"
	"strings"

	"example.com/agent/agent"
)

func contains(roles []string, role string) bool {
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
}

func Query(agents []agent.Agent, roles string, query_mode int) {
	roles_splitted := strings.Split(roles, ",")

	var matching_agents []agent.Agent

	for _, agent_ := range agents {
		if !agent_.Is_available {
			continue
		}

		is_matching := true
		for _, role := range roles_splitted {
			if !contains(agent_.Roles, role) {
				is_matching = false
				break
			}
		}

		if is_matching {
			matching_agents = append(matching_agents, agent_)
		}
	}

	if len(matching_agents) == 0 {
		fmt.Println("No matching agents found")
		return
	}

	switch query_mode {
	case 1:
		fmt.Println(matching_agents)

	case 2:
		var req_agent agent.Agent
		max_available_since := matching_agents[0].Available_since

		for i := 1; i < len(matching_agents); i++ {
			if matching_agents[i].Available_since > max_available_since {
				max_available_since = matching_agents[i].Available_since
				req_agent = matching_agents[i]
			}
		}

		fmt.Println(req_agent)

	case 3:
		random_idx := rand.Intn(len(matching_agents))
		fmt.Println(matching_agents[random_idx])
	}
}
