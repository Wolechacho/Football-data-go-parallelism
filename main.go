package main

import "fmt"

func main() {
	resp := GetFootballData(competition, year, 1)
	AddTeamToMap(resp.FootballMatches)

	go AllocateJobs(resp.TotalPages)
	teams := CreateWorkerThread(workerPoolSize)

	fmt.Println(teams)
}
