package main

import (
	"log"
	"strconv"
)

func getMoviesHost(requestsCount int, monolithUrl string, moviesUrl string, gradualMigration string, migrationPercentEnv string) string {

	migrationPercent, convErr := strconv.Atoi(migrationPercentEnv)

	if convErr != nil {
		log.Fatal("couldn't get MOVIES_MIGRATION_PERCENT")
	}

	serverUrl := monolithUrl
	if gradualMigration == "true" {
		if requestsCount == 0 {
			serverUrl = moviesUrl
		} else {
			if requestsCount%(100/migrationPercent) == 0 {
				serverUrl = moviesUrl
			}
		}
	}

	return serverUrl
}
