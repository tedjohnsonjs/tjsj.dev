package webserver

import (
	"net/http"
	"html/template"
	"fmt"
	"errors"
)

type homeResponseData struct {
	SplashText string
	ProjectsList []template.HTML
	RecentPostsList []template.HTML
}

type statsResponseData struct {
	StatsLists [2][]string
	StatsStartDatetime string
}

// General method for executing HTML templates by name
func (server *Server) executeHTMLTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	template := server.content.GetHTMLTemplate(templateName)
	if template == nil {
		server.errChan <- errors.New("HTML template '" + templateName + "' not found!")
		return
	}

	template.Execute(w, data)
}

// Respond with the HTML template "home.html"
func (server *Server) homeResponse(w http.ResponseWriter, r *http.Request) {
	// TODO: get list of recent posts metadata
	data := homeResponseData {
		server.content.GetRandomSplashText(),
		nil,
		nil,
	}

	server.executeHTMLTemplate(w, "home.html", data)
}

// Respond with a list of projects in the HTML template "projects.html"
func (server *Server) projectsResponse(w http.ResponseWriter, r *http.Request) {
	// TODO: get list of projects metadata
	//server.executeHTMLTemplate(w, "projects.html", nil)
	server.errorResponse(w, r, http.StatusNotFound)
}

// Respond with the project page of the id given in the URL
func (server *Server) projectResponse(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//id := vars["id"]
	// TODO: find the project
	//server.executeHTMLTemplate(w, "project.html", nil)
	server.errorResponse(w, r, http.StatusNotFound)
}

// Respond with a list of posts in the HTML template "posts.html"
func (server *Server) postsResponse(w http.ResponseWriter, r *http.Request) {
	// TODO: get list of posts metadata
	server.executeHTMLTemplate(w, "posts.html", nil)
	// server.errorResponse(w, r, http.StatusNotFound)
}

// Respond with the post page of the id given in the URL
func (server *Server) postResponse(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//id := vars["id"]
	// TODO: find the post
	//server.executeHTMLTemplate(w, "post.html", nil)
	server.errorResponse(w, r, http.StatusNotFound)
}

// Respond with the statistics page
func (server *Server) statsResponse(w http.ResponseWriter, r *http.Request) {
	statsLists := [2][]string{}

	// Total data transferred
	totalCompressedDataTransferred, totalUncompressedDataTransferred := server.stats.GetTotalDataTransferred()
	compressionRatio := float64(totalCompressedDataTransferred) / float64(totalUncompressedDataTransferred)

	// Hit counters
	hitCounters, hitCountersOrder := server.stats.GetHitCounters()
	totalHitCounts := 0
	topURLsList := []string{}
	for _, url := range hitCountersOrder {
		totalHitCounts += hitCounters[url]
		if len(topURLsList) < 10 {
			topURLsList = append(
				topURLsList,
				fmt.Sprintf("| %d hits: %s\n", hitCounters[url], url))
		}
	}

	// Referrer counters
	referrerCounters, referrerCountersOrder := server.stats.GetReferrerCounters()
	topReferrersList := []string{}
	for _, referrer := range referrerCountersOrder {
		if len(topReferrersList) < 10 {
			topReferrersList = append(
				topReferrersList,
				fmt.Sprintf("| %d referees: %s\n", referrerCounters[referrer], referrer))
		}
	}

	// Response code counters
	responseCodeCounters, responseCodeCountersOrder := server.stats.GetResponseCodeCounters()
	topResponseCodesList := []string{}
	for _, responseCode := range responseCodeCountersOrder {
		if len(topResponseCodesList) < 10 {
			topResponseCodesList = append(
				topResponseCodesList,
				fmt.Sprintf(
					"| %d responses: %d - %s\n",
					responseCodeCounters[responseCode],
					responseCode,
					http.StatusText(responseCode)))
		}
	}

	// Request stats list
	statsLists[0] = make([]string, 0, 7 + len(topURLsList) + len(topReferrersList) + len(topResponseCodesList))
	statsLists[0] = append(statsLists[0], fmt.Sprintf("Total website hits: %d", totalHitCounts))
	statsLists[0] = append(statsLists[0], "")
	statsLists[0] = append(statsLists[0], fmt.Sprintf("Top %d URLs:", len(topURLsList)))
	statsLists[0] = append(statsLists[0], topURLsList...)
	statsLists[0] = append(statsLists[0], "")
	statsLists[0] = append(statsLists[0], fmt.Sprintf("Top %d referrers:", len(topReferrersList)))
	statsLists[0] = append(statsLists[0], topReferrersList...)
	statsLists[0] = append(statsLists[0], "")
	statsLists[0] = append(statsLists[0], fmt.Sprintf("Top %d response codes:", len(topResponseCodesList)))
	statsLists[0] = append(statsLists[0], topResponseCodesList...)

	// System stats list
	sysstats := server.stats.GetSystemStats()
	statsLists[1] = []string {
		"Server: " + sysstats.Hostname,
		"Platform: " + sysstats.OS + "/" + sysstats.Arch,
		"",
		fmt.Sprintf("CPU usage: %d%% (%d cores)", sysstats.CPUUsage, sysstats.CPUCount),
		fmt.Sprintf("RAM usage: %s (%s available)",
			bytesToHumanReadable(sysstats.RAMUsage),
			bytesToHumanReadable(sysstats.RAMAvailable)),
		fmt.Sprintf("Active Goroutines: %d", sysstats.GoroutineCount),
		"",
		fmt.Sprintf("Total transferred: %s", bytesToHumanReadable(totalCompressedDataTransferred)),
		fmt.Sprintf("Before compression: %s", bytesToHumanReadable(totalUncompressedDataTransferred)),
		fmt.Sprintf("Avg. compression ratio: %f", compressionRatio),
		"",
		"Server Uptime: " + server.stats.GetUptime(),
		"System Uptime: " + sysstats.Uptime,
	}

	// Execute HTML template response
	server.executeHTMLTemplate(w, "stats.html", statsResponseData {
		statsLists,
		server.stats.GetStatsStartDatetime(),
	})
}

// Respond with the error page with an appropriate message
func (server *Server) errorResponse(w http.ResponseWriter, r *http.Request, code int) {
	w.WriteHeader(code)
	data := struct { Code int; Message string } { code, http.StatusText(code) }
	server.executeHTMLTemplate(w, "error.html", data)
}

// Convert bytes to an approprate magnitude
func bytesToHumanReadable(bytes uint64) string {
	if bytes < 1000 {
		return fmt.Sprintf("%dB", bytes);
	}

	magnitude := 0
	prefixes := []byte{'k','M','G','T','P','E'}
	for bytes >= 999950 && magnitude < len(prefixes) {
		bytes /= 1000;
		magnitude++
	}

	return fmt.Sprintf("%.1f%cB", float64(bytes) / 1000.0, prefixes[magnitude]);
}
