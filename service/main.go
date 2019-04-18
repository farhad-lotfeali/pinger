package service

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//Service ...
// type Service interface {
// }

//PingResult ...
type PingResult struct {
	SiteID  int    `json:"site_id"`
	Elapsed int64  `json:"elapsed"`
	Date    string `json:"date"`
}

//ParsePingResult ...
func ParsePingResult(raw string) *PingResult {
	p := PingResult{}

	data := strings.Split(raw, "|")

	p.SiteID, _ = strconv.Atoi(data[0])
	p.Elapsed, _ = strconv.ParseInt(data[1], 10, 64)
	p.Date = data[2]

	return &p
}

//SaveResult ...
func SaveResult(f *os.File, result string) {
	if _, err := f.WriteString(result + "\n"); err != nil {
		panic(err)
	}
}

//CheckURL ...
func CheckURL(url string) (result bool, elapsedTime time.Duration) {
	start := trace()

	res, err := http.Get(url)
	if err != nil {
		return false, un(start)
	}

	defer res.Body.Close()
	return true, un(start)
}

//ReadSites ...
func ReadSites() [][]string {
	f, err := os.Open("./sites.csv")

	if err != nil {
		fmt.Println("the sites file can not found")
	}

	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll()

	if err != nil {
		panic(err)
	}

	return lines
}

//Ping ...
func Ping(now time.Time, sites [][]string) {
	f, err := os.OpenFile("./uptime.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, site := range sites {
		result, elapsed := CheckURL(site[1])
		fmt.Print(site[1] + " Result ")

		if result == false {
			fmt.Println(" Failed")
			SaveResult(f, site[0]+"|-1|"+now.Format("2006-01-02 03:04:05"))
			continue
		}

		fmt.Println(strconv.FormatInt(int64(elapsed/time.Millisecond), 10) + " ms")
		SaveResult(f, site[0]+"|"+strconv.FormatInt(int64(elapsed/time.Millisecond), 10)+"|"+now.Format("2006-01-02 03:04:05"))
	}

}

func trace() time.Time {
	return time.Now()
}

func un(startTime time.Time) time.Duration {
	endTime := time.Now()
	return endTime.Sub(startTime)
}
