package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/farhad-lotfeali/pinger/service"
	"github.com/spf13/cobra"
)

//NewServer ...
func NewServer() *cobra.Command {
	return serverCmd
}

// restCmd represents the rest command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run restfull server for add sites and see statics :)",
	Run:   runServer,
}

func runServer(cmd *cobra.Command, args []string) {
	fmt.Println("run rest server ")

	http.HandleFunc("/uptime", getUptime)

	log.Fatal(http.ListenAndServe(":3000", nil))

}

func getUptime(writer http.ResponseWriter, request *http.Request) {
	id := request.FormValue("id")
	if id == "" {
		_, _ = fmt.Fprintln(writer, "Your not Selected Site ID")
		return
	}

	f, err := os.Open("./uptime.log")

	if err != nil {
		dir, _ := os.Getwd()
		_, _ = fmt.Fprintf(writer, "Error in read Log File %s", dir)
		return
	}

	scanner := bufio.NewScanner(f)

	empty := true

	var output []service.PingResult
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), id) {
			output = append(output, *service.ParsePingResult(scanner.Text()))
			//_, _ = fmt.Fprintln(writer, ParsePingResult(scanner.Text()))
			empty = false
		}
	}

	if empty {
		_, _ = fmt.Fprintln(writer, "Empty Result")
	}

	_ = json.NewEncoder(writer).Encode(output)
}
