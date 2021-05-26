package collector

import (
	"testing"
)

func TestAwsMetadata(t *testing.T) {
	wantevent := map[string]string {
		"NotBefore" : "1 May 2021 22:00:00 GMT", 
		"NotAfter" : "2 May 2021 00:00:00 GMT",
		"State" : "active",
	}
	wantmetrics := map[string]int {
		"notbefore" : 1619906400,
		"notafter" : 1619913600,
		"state" : 1,
	}

	events, err := parseAwsScheduledEvents(`[ {
  "NotBefore" : "1 May 2021 22:00:00 GMT",
  "Code" : "system-reboot",
  "Description" : "scheduled reboot do-not-complete",
  "EventId" : "instance-event-0000a0aa0aa0a0aaa",
  "NotAfter" : "2 May 2021 00:00:00 GMT",
  "State" : "active"
} ]`)

	if err != nil {
		t.Fatal(err)
	}

	for _, event := range events {
		if event.State != wantevent["State"] {
			t.Fatalf("want event State %s, got %s", wantevent["State"], event.State)
		}
		if event.NotBefore != wantevent["NotBefore"] {
			t.Fatalf("want event NotBefore %s, got %s", wantevent["NotBefore"], event.NotBefore)
		}
		if event.NotAfter != wantevent["NotAfter"] {
			t.Fatalf("want event NotAfter %s, got %s", wantevent["NotAfter"], event.NotAfter)
		}

		eventMetrics, err := parseAwsScheduledEventMetrics(event)

		if err != nil {
			t.Fatal(err)
		}

		if eventMetrics[0] != wantmetrics["state"] {
			t.Fatalf("want metric state %d, got %d", wantmetrics["state"], eventMetrics[0])
		}
		if eventMetrics[1] != wantmetrics["notbefore"] {
			t.Fatalf("want metric notbefore %d, got %d", wantmetrics["notbefore"], eventMetrics[0])
		}
		if eventMetrics[2] != wantmetrics["notafter"] {
			t.Fatalf("want metric notafter %d, got %d", wantmetrics["notafter"], eventMetrics[0])
		}
	}

}
