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
	LETTERS       = "abcdefghijklmnopqrstuvwxyzABCDEFHIJKLMNOPQRSTUVWXYZ0123456789"
)

var DEBUG_TIME = 10 * time.Second

// makePrintable fills buf with pseudo‑random printable ASCII.
func makePrintable(buf []byte) {
	for i := range buf {
		buf[i] = LETTERS[rand.Intn(len(LETTERS))]
	}
}

// makeJSON fills buf with a JSON object containing a timestamp and a message.
func makeJSON(buf []byte) int {
	timestamp := time.Now().Format(time.RFC3339Nano)
	template := `{"time": "%s", "body": "%s"}`

	// Minimum valid JSON length (empty message)
	minLen := len(fmt.Sprintf(template, timestamp, ""))

	if len(buf) < minLen {
		fmt.Fprintf(os.Stderr, "Error: -b must be at least %d for valid JSON output\n", minLen)
		os.Exit(1)
	}

	// Find the largest message length that fits in buf
	maxMsgLen := len(buf) - (minLen - 2)
	if maxMsgLen < 0 {
		maxMsgLen = 0
	}
	var jsonStr string
	for {
		msg := make([]byte, maxMsgLen)
		for i := range msg {
			msg[i] = LETTERS[rand.Intn(len(LETTERS))]
		}
		jsonStr = fmt.Sprintf(template, timestamp, string(msg))
		if len(jsonStr) <= len(buf) {
			break
		}
		maxMsgLen--
	}

	copy(buf, jsonStr)
	// Pad with spaces if needed
	for i := len(jsonStr); i < len(buf); i++ {
		buf[i] = ' '
	}
	return len(jsonStr)
}

func main() {
	// --- CLI flags ----------------------------------------------------------
	bytesPerRec := flag.Int("b", 64, "bytes per log record")
	interval := flag.String("i", "1s", "interval between records (e.g. 250ms, 1s, 500µs)")
	rate := flag.String("r", "0", "rate of logs per second (e.g. 100 for 100 logs/sec)")
	format := flag.String("f", "json", "output format ('json' = json lines logs, 'random' = random printable ASCII)")
	flag.Parse()

	// Calculate minimum required size for JSON
	if *format == FORMAT_JSON {
		timestamp := time.Now().Format(time.RFC3339Nano)
		template := `{"time": "%s", "body": "%s"}`
		minLen := len(fmt.Sprintf(template, timestamp, ""))

		if *bytesPerRec < minLen {
			fmt.Fprintf(os.Stderr, "-b must be at least %d for valid JSON output\n", minLen)
			os.Exit(1)
		}
	}

	// Validate size
	if *bytesPerRec <= 0 {
		fmt.Fprintln(os.Stderr, "-b must be > 0")
		os.Exit(1)
	}

	// Validate interval and rate
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

	logRecordsCount := 0
	tickerDebug := time.NewTicker(DEBUG_TIME)
	defer tickerDebug.Stop()

	for {
		select {
		case <-tickerDebug.C:
			fmt.Printf("DEBUG: %d log records printed last 10 sec\n", logRecordsCount)
			logRecordsCount = 0
		case <-ticker.C:
			switch *format {
			case "json":
				jsonLen := makeJSON(buf)
				fmt.Printf("%s\n", string(buf[:jsonLen]))
			case "random":
				makePrintable(buf)
				fmt.Printf("%s %s\n", time.Now().Format(time.RFC3339Nano), buf)
			}
			logRecordsCount++
		}
	}
}
