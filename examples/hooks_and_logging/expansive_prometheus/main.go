package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"
	"time"

	"github.com/tdx/franz-go/pkg/kgo"
)

var (
	debugPort   = flag.Int("debug-port", 9999, "localhost port that metrics can be curled from (path /metrics)")
	seedBrokers = flag.String("brokers", "localhost:9092", "comma delimited list of seed brokers")
	topic       = flag.String("topic", "test", "topic to consume for metric incrementing")
	produce     = flag.Bool("produce", false, "if true, rather than consume, produce to the topic once per second (value \"foo\")")
)

func main() {
	flag.Parse()

	metrics := NewMetrics("kgo")

	opts := []kgo.Opt{
		kgo.SeedBrokers(strings.Split(*seedBrokers, ",")...),
		kgo.WithHooks(metrics),
		kgo.DefaultProduceTopic(*topic),
		kgo.WithLogger(kgo.BasicLogger(os.Stderr, kgo.LogLevelInfo, func() string {
			return time.Now().Format("[2006-01-02 15:04:05.999] ")
		})),
	}
	if !*produce {
		opts = append(opts, kgo.ConsumeTopics(*topic))
	}

	cl, err := kgo.NewClient(opts...)
	if err != nil {
		panic(fmt.Sprintf("unable to create client: %v", err))
	}
	defer cl.Close()

	go func() {
		http.Handle("/metrics", metrics.Handler())
		log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", *debugPort), nil))
	}()

	if *produce {
		for range time.Tick(time.Second) {
			if err := cl.ProduceSync(context.Background(), kgo.StringRecord("foo")).FirstErr(); err != nil {
				panic(fmt.Sprintf("unable to produce: %v", err))
			}
		}
	} else {
		for {
			cl.PollFetches(context.Background()) // busy work...
		}
	}
}
