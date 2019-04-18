package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/farhad-lotfeali/pinger/service"
	"github.com/spf13/cobra"
)

//NewPinger ...
func NewPinger() *cobra.Command {
	return pingerCmd
}

// pingerCmd represents the pinger command
var pingerCmd = &cobra.Command{
	Use:   "pinger",
	Short: "run pinge sites in backgroun",
	Run:   runPinger,
}

func runPinger(cmd *cobra.Command, args []string) {
	fmt.Println("run the ping scheduler")
	done := make(chan bool)

	sites := service.ReadSites()

	pingerTimer := time.NewTicker(2 * time.Second)
	readTiemr := time.NewTicker(5 * time.Minute)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			now := <-pingerTimer.C
			service.Ping(now, sites)
		}
	}()

	go func() {
		for {
			<-readTiemr.C
			log.Printf("read sites ")
			sites = service.ReadSites()
		}
	}()

	go func() {
		<-sigs
		fmt.Printf("stop the pinger by Ctrl+C")
		pingerTimer.Stop()
		readTiemr.Stop()
		done <- true
	}()
	<-done
}
