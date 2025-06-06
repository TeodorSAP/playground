package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	FORMAT_JSON   = "json"
	FORMAT_RANDOM = "random"
	LETTERS       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// makePrintable fills buf with pseudo‑random printable ASCII.
func makePrintable(buf []byte) {
	for i := range buf {
		buf[i] = LETTERS[rand.Intn(len(LETTERS))]
	}
}

// makeJSON fills buf with a JSON object containing a timestamp and a message.
func makeJSON(buf []byte) {
	timestamp := time.Now().Format(time.RFC3339Nano)

	// Template for JSON with empty message
	template := `{"time": "%s", "body": "%s"}`
	// Calculate the overhead (fixed) length with empty message
	overhead := len(fmt.Sprintf(template, timestamp, ""))

	// If the buffer is too small for even the fixed part, fill as much as possible
	if len(buf) < overhead {
		copy(buf, fmt.Sprintf(template, timestamp, "")[:len(buf)])
		return
	}

	// Calculate max message length that fits in buf
	maxMsgLen := len(buf) - (overhead - 2) // -2 because %s for message is replaced by message itself

	msg := make([]byte, maxMsgLen)
	for i := range msg {
		msg[i] = LETTERS[rand.Intn(len(LETTERS))]
	}

	jsonStr := fmt.Sprintf(template, timestamp, string(msg))
	copy(buf, jsonStr)
	// Pad with spaces if needed
	for i := len(jsonStr); i < len(buf); i++ {
		buf[i] = ' '
	}
}

func main() {
	// --- CLI flags ----------------------------------------------------------
	bytesPerRec := flag.Int("b", 64, "bytes per log record")
	interval := flag.String("i", "1s", "interval between records (e.g. 250ms, 1s, 500µs)")
	rate := flag.String("r", "0", "rate of logs per second (e.g. 100 for 100 logs/sec)")
	format := flag.String("f", "json", "output format ('json' = json lines logs, 'random' = random printable ASCII)")
	flag.Parse()

	if *bytesPerRec <= 0 {
		fmt.Fprintln(os.Stderr, "-b must be > 0")
		os.Exit(1)
	}
	var period time.Duration
	if *rate != "0" {
		rateInt, err := strconv.Atoi(*rate)
		if err != nil || rateInt <= 0 {
			fmt.Fprintf(os.Stderr, "invalid -r value: %v\n", err)
			os.Exit(1)
		}
		period = time.Second / time.Duration(rateInt)
	} else {
		var err error
		period, err = time.ParseDuration(*interval)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid -i value: %v\n", err)
			os.Exit(1)
		}
	}

	// --- generator loop -----------------------------------------------------
	buf := make([]byte, *bytesPerRec)
	ticker := time.NewTicker(period)
	defer ticker.Stop()

	for {
		switch *format {
		case "json":
			makeJSON(buf)
		case "random":
			makePrintable(buf)
		}
		fmt.Printf("%s %s\n", time.Now().Format(time.RFC3339Nano), buf)
		<-ticker.C
	}
}
