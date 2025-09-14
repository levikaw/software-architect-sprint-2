package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	monolithUrl := os.Getenv("MONOLITH_URL")
	moviesUrl := os.Getenv("MOVIES_SERVICE_URL")
	eventsUrl := os.Getenv("EVENTS_SERVICE_URL")
	gradualMigration := os.Getenv("GRADUAL_MIGRATION")
	migrationPercent := os.Getenv("MOVIES_MIGRATION_PERCENT")

	for _, v := range []string{port, monolithUrl, moviesUrl, eventsUrl, gradualMigration} {
		if v == "" {
			log.Fatal("One of env vars not set (PORT, MONOLITH_URL,  MOVIES_SERVICE_URL, EVENTS_SERVICE_URL, GRADUAL_MIGRATION)")
		}
	}

	requestsCount := 0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var host string
		if r.URL.Path == "/api/movies" {
			requestsCount++
			host = getMoviesHost(requestsCount, monolithUrl, moviesUrl, gradualMigration, migrationPercent)
		} else {
			host = monolithUrl
		}

		target, err := url.Parse(host)
		if err != nil {
			log.Fatal(err)
		}
		proxy := httputil.NewSingleHostReverseProxy(target)

		r.URL.Host = target.Host
		r.URL.Scheme = target.Scheme
		r.Host = target.Host

		proxy.ServeHTTP(w, r)
	})

	fmt.Printf("Starting server at port %s", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
