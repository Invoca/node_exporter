package collector

import (
	"time"
	"encoding/json"
	
	"github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type sheduled_event struct {
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
	registerCollector("awsmetadata", defaultEnabled, NewAwsmetadataCollector)
}

// NewAwsmetadataCollector returns a new Collector exposing AWS metadata stats
func NewAwsmetadataCollector(logger log.Logger) (Collector, error) {
	return &awsmetadataCollector{logger}, nil
}

func (c *awsmetadataCollector) Update(ch chan<- prometheus.Metric) error {
	// md_metrics, err := c.getAwsMetadata()

	return nil
}

func (c *awsmetadataCollector) getAwsMetadata() (map[string]string, error) {
	return nil, nil
}

func (c *awsmetadataCollector) getAwsScheduledEvents() (map[string]string, error) {
	return nil, nil
}

// takes an array of json objects in string format and returns populated structs
func parseAwsScheduledEvents(data string) ([]sheduled_event, error) {
	return nil, nil
}

// returns metrics in the order {active, notbefore, notafter}
func parseAwsScheduledEventMetrics(event sheduled_event) ([3]int, error) {
/*	if event.State == "active" {
		
	}
*/
	return [3]int{0, 0, 0}, nil
}
