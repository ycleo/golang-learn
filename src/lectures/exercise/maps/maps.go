//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
func displayServerInfo(serverStatus map[string]int) {
	fmt.Println("There are", len(serverStatus), "servers.")
	fmt.Println()
	status := make(map[int]int)

	for server, stat := range serverStatus {

		switch stat {
		case Online:
			status[Online] += 1
			fmt.Println(server, "is online")
		case Offline:
			status[Offline] += 1
			fmt.Println(server, "is offline")
		case Maintenance:
			status[Maintenance] += 1
			fmt.Println(server, "is under maintenance")
		case Retired:
			status[Retired] += 1
			fmt.Println(server, "is retired")
		default:
			panic("unhandled server status")
		}
	}

	fmt.Println()
	fmt.Println(status[Online], "servers are online.")
	fmt.Println(status[Offline], "servers are offline.")
	fmt.Println(status[Maintenance], "servers are under maintenance")
	fmt.Println(status[Retired], "servers are retired.")
}

func maintainAllServers(serverStatus map[string]int) {
	for key, _ := range serverStatus {
		serverStatus[key] = Maintenance
	}
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	//* Create a map using the server names as the key and the server status
	//  as the value
	//* Set all of the server statuses to `Online` when creating the map

	serverStatus := make(map[string]int)

	for _, server := range servers {
		serverStatus[server] = Online
	}

	//* After creating the map, perform the following actions:
	//  - call display server info function
	displayServerInfo(serverStatus)
	fmt.Println("==================================")
	//  - change server status of `darkstar` to `Retired`
	serverStatus["darkstar"] = Retired
	//  - change server status of `aiur` to `Offline`
	serverStatus["aiur"] = Offline
	//  - call display server info function
	displayServerInfo(serverStatus)
	fmt.Println("==================================")
	//  - change server status of all servers to `Maintenance`
	maintainAllServers(serverStatus)
	//  - call display server info function
	displayServerInfo(serverStatus)

}
