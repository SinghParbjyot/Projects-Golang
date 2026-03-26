package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func Start1() {

}
func getTime(w http.ResponseWriter, r *http.Request) {
	// Initialize a map to hold our response data
	response := make(map[string]string, 0)

	// Extract the "tz" query parameter from the URL (e.g., /time?tz=America/New_York)
	tz := r.URL.Query().Get("tz")

	// Split the string by commas to support multiple timezones in one request
	timezones := strings.Split(tz, ",")

	// Case 1: Single timezone requested (or parameter is empty)
	if len(timezones) <= 1 {
		// Attempt to load the location data for the given timezone string
		loc, err := time.LoadLocation(tz)
		if err != nil {
			// If the timezone is invalid, return a 404 Not Found
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf("invalid timezone %s", tz)))
		} else {
			// Format the current time for that location and add to response
			response["current_time"] = time.Now().In(loc).String()

			// Set the response header to JSON and encode the map
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}
	} else {
		// Case 2: Multiple timezones requested (comma-separated)
		for _, tzdb := range timezones {
			loc, err := time.LoadLocation(tzdb)
			if err != nil {
				// If any single timezone in the list is invalid, fail the whole request
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(fmt.Sprintf("invalid timezone %s in input", tzdb)))
				return
			}

			// Get the current time in the specific location and store it using the TZ name as the key
			now := time.Now().In(loc)
			response[tzdb] = now.String()
		}

		// Send back the full map of timezones as a JSON object
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
