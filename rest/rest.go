package rest

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func RunHttpRest() {
	fmt.Println("Start")
	http.HandleFunc("/uptime", getUptime)

	_ = http.ListenAndServe(":3000", nil)
	fmt.Println("Finish")
}

type PingResult struct {
	SiteID  int    `json:"site_id"`
	Elapsed int64  `json:"elapsed"`
	Date    string `json:"date"`
}

func ParsePingResult(raw string) *PingResult {
	p := PingResult{}

	data := strings.Split(raw, "|")

	p.SiteID, _ = strconv.Atoi(data[0])
	p.Elapsed, _ = strconv.ParseInt(data[1], 10, 64)
	p.Date = data[2]

	return &p
}

func getUptime(writer http.ResponseWriter, request *http.Request) {
	id := request.FormValue("id")
	if id == "" {
		_, _ = fmt.Fprintln(writer, "Your not Selected Site ID")
		return
	}

	f, err := os.Open("../uptime.log")

	if err != nil {
		_, _ = fmt.Fprintln(writer, "Error in read Log File")
		return
	}

	scanner := bufio.NewScanner(f)

	empty := true

	var output []PingResult
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), id) {
			output = append(output, *ParsePingResult(scanner.Text()))
			//_, _ = fmt.Fprintln(writer, ParsePingResult(scanner.Text()))
			empty = false
		}
	}

	if empty {
		_, _ = fmt.Fprintln(writer, "Empty Result")
	}

	_ = json.NewEncoder(writer).Encode(output)
}
