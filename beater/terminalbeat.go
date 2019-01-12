package beater

import (
	"fmt"
	"time"

	"bufio"
	"os"
	"os/exec"
	"strings"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/live-wire/terminalbeat/config"
)

// Terminalbeat configuration.
type Terminalbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

type exitCode struct{}

// New creates an instance of terminalbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Terminalbeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts terminalbeat.
func (bt *Terminalbeat) Run(b *beat.Beat) error {
	logp.Info("terminalbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ch := make(chan string)
	chx := make(chan int)

	go runCommand(bt.config.Command, &ch)
	go listenForExit(&chx)
	// for {
	//     select {
	//         case msg := <-ch:
	//             fmt.Println("Message found:", msg)
	//         case <-chx:
	//             fmt.Println("Exit code entered")
	//             os.Exit(0)
	//     }
	// }

	ticker := time.NewTicker(bt.config.Period)
	counter := 1
	for {
		select {
		case <-bt.done:
			return nil
		case msg := <-ch:
			fmt.Println("Message found:", msg)
			event := beat.Event{
				Timestamp: time.Now(),
				Fields: common.MapStr{
					"type":    b.Info.Name,
					"counter": counter,
					"msg": msg,
					"command": bt.config.Command,
				},
			}
			bt.client.Publish(event)
			logp.Info("Event sent")
			counter++
        case <-chx:
            fmt.Println("Exit code entered")
            bt.done <- exitCode{}
            // os.Exit(0)
		case <-ticker.C:
		}
	}
}

// Stop stops terminalbeat.
func (bt *Terminalbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}

// Listens for exit code in stdin
func listenForExit(chx *chan int) {
    scanner := bufio.NewScanner(os.Stdin)
    inp := ""
    fmt.Println("Enter 0 to exit:")
    scanner.Scan()
    inp += scanner.Text()
    if inp == "0" {
        *chx <- 0
    } else {
        fmt.Println("Unrecognized input", inp)
        listenForExit(chx)
    }
    if scanner.Err() != nil {
        // handle error.
    }
}

// Runs a command and captures the stdout logs for the same
func runCommand(cmdName string, ch *chan string) {
    fmt.Println("Running command [", cmdName,"]")
    cmdArgs := strings.Fields(cmdName)
    cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
    stdout, _ := cmd.StdoutPipe()
    cmd.Start()
    num := 1
    for {
        r := bufio.NewReader(stdout)
        line, _ := r.ReadString('\n')
        if string(line) != "" {
            *(ch) <- string(line)
        }
        num = num + 1
    }
}
