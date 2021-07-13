package collector

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	
	"github.com/go-kit/log"
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
	// md_metrics, err := c.getAwsMetadata()

	return nil
}

func (c *awsmetadataCollector) getAwsMetadata() (map[string]string, error) {
	return nil, nil
}

// get scheduled events via instance metadata
func (c *awsmetadataCollector) getAwsScheduledEvents() (string, error) {
	mdURL := "http://169.254.169.254/latest/meta-data/events/maintenance/scheduled"

	resp, e := http.Get(mdURL)
	if e != nil {
		return "", e
	}
	defer resp.Body.Close()

	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return "", e
	}

	return string(body), nil
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

	nb, e := time.Parse(tformat, event.NotBefore)
	if e != nil {
		return [3]int{0, 0, 0}, e
	}
	na, e := time.Parse(tformat, event.NotAfter)
	if e != nil {
		return [3]int{0, 0, 0}, e
	}

	metrics[1] = int(nb.Unix())
	metrics[2] = int(na.Unix())

	return metrics, nil
}
