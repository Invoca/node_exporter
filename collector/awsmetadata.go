package collector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
)

type scheduledEvent struct {
  Code         string  `json:"code"`
  State        string  `json:"state"`
  Description  string  `json:"description"`
  EventId      string  `json:"eventid"`
  NotBefore    string  `json:"notbefore"`
  NotAfter     string  `json:"notafter"`
}

type awsmetadataCollector struct {
	logger log.Logger
}

func init() {
	registerCollector("awsmetadata", defaultEnabled, newAwsmetdataCollector)
}

// NewUnameCollector returns new unameCollector.
func newAwsmetdataCollector(logger log.Logger) (Collector, error) {
	return &unameCollector{logger}, nil
}

func (c *awsmetadataCollector) Update(ch chan<- prometheus.Metric) error {
	metrics, err := c.getAwsMetadata()
	if err != nil {
		return fmt.Errorf("couldn't get scheduled events from instance metadata: %w", err)
	}

	for i, metric := range metrics {

	}

	return nil
}

func (c *awsmetadataCollector) getAwsMetadata() ([][3]int, error) {
	metrics := [][3]int{}
	eventsMetadata, err := c.getAwsScheduledEvents()
	if err != nil {
		return nil, err
	}

	events, err := parseAwsScheduledEvents(eventsMetadata)
	if err != nil {
		return nil, err
	}

	for _, event := range events {
		eventMetrics, err := parseAwsScheduledEventMetrics(event)
		if err != nil {
			return nil, err
		}

		metrics = append(metrics, eventMetrics)
	}

	return metrics, nil
}

// get scheduled events via instance metadata
func (c *awsmetadataCollector) getAwsScheduledEvents() (string, error) {
	mdURL := "http://169.254.169.254/latest/meta-data/events/maintenance/scheduled"

	resp, err := http.Get(mdURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	mdEvents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(mdEvents), nil
}

// takes an array of json objects in string format and returns populated structs
func parseAwsScheduledEvents(data string) ([]scheduledEvent, error) {
	res := []scheduledEvent{}
	json.Unmarshal([]byte(data), &res)

	return res, nil
}

// returns metrics in the order {active, notbefore, notafter}
func parseAwsScheduledEventMetrics(event scheduledEvent) ([3]int, error) {
	var metrics [3]int
	tformat := "_2 Jan 2006 15:04:05 GMT"

	if event.State == "active" {
		metrics[0] = 1
	} else {
		metrics[0] = 0
	}

	nb, err := time.Parse(tformat, event.NotBefore)
	if err != nil {
		return [3]int{0, 0, 0}, err
	}
	na, err := time.Parse(tformat, event.NotAfter)
	if err != nil {
		return [3]int{0, 0, 0}, err
	}

	metrics[1] = int(nb.Unix())
	metrics[2] = int(na.Unix())

	return metrics, nil
}
