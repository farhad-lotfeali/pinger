package pinger

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

var done chan bool

// RunPinger run ping from sites in background
func RunPinger() {
	done = make(chan bool)

	fmt.Println("Start")

	sites := readSites()
	counter := 0
	timer := time.NewTicker(2 * time.Second)

	//read all files from site folder

	// create range and run goroutine
	go func() {
		counter++
		if counter == 100 {
			sites = readSites()
			counter = 0
		}

		f, err := os.OpenFile("./uptime.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			done <- true
			panic(err)
		}
		defer f.Close()

		for {
			now := <-timer.C

			for _, site := range sites {
				result, elapsed := checkURL(site[1])
				fmt.Print(site[1] + " Result ")

				if result == false {
					fmt.Println(" Failed")
					saveResult(f, site[0]+"|-1|"+now.Format("2006-01-02 03:04:05"))
					continue
				}

				fmt.Println(strconv.FormatInt(int64(elapsed/time.Millisecond), 10) + " ms")
				saveResult(f, site[0]+"|"+strconv.FormatInt(int64(elapsed/time.Millisecond), 10)+"|"+now.Format("2006-01-02 03:04:05"))
			}
		}
	}()

	<-done
}

func saveResult(f *os.File, result string) {

	if _, err := f.WriteString(result + "\n"); err != nil {
		done <- true
		panic(err)
	}
}

func checkURL(url string) (result bool, elapsedTime time.Duration) {
	start := trace()

	res, err := http.Get(url)
	if err != nil {
		return false, un(start)
	}

	defer res.Body.Close()
	return true, un(start)
}

func trace() time.Time {
	return time.Now()
}

func un(startTime time.Time) time.Duration {
	endTime := time.Now()
	return endTime.Sub(startTime)
}

func readSites() [][]string {
	f, err := os.Open("./sites.csv")

	if err != nil {
		fmt.Println("the sites file can not found")
		done <- true
	}

	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll()

	if err != nil {
		panic(err)
	}

	return lines
}
