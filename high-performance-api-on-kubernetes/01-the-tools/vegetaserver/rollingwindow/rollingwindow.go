package rollingwindow

import (
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

const (
	typeAdd      = "CMD_ADD"
	typeSnapshot = "CMD_SNAPSHOT"
)

// New creates a new Metrics
func New(size int) Metrics {

	r := Metrics{
		cmd:    make(chan cmd),
		window: make([]*vegeta.Result, size),
		next:   0,
	}

	go func() {
		for {
			c := <-r.cmd
			switch c.Type {
			case typeAdd:
				r.add(c.AddIn)
			case typeSnapshot:
				metrics := r.snapshot()
				c.SnapshotOut <- metrics
			}

		}
	}()

	return r
}

// Metrics represents the struture of a windowed metrics.
// Does not expose any values
type Metrics struct {
	cmd    chan cmd
	window []*vegeta.Result
	next   int
}

// Add adds a vegeta result to a rolling metrics window
func (r *Metrics) Add(result *vegeta.Result) {
	r.cmd <- cmd{
		Type:  typeAdd,
		AddIn: result,
	}
}

// Snapshot returns a clsoed vegeta snapshot
func (r Metrics) Snapshot() vegeta.Metrics {
	o := make(chan vegeta.Metrics)
	r.cmd <- cmd{
		Type:        typeSnapshot,
		SnapshotOut: o,
	}
	select {
	case metrics := <-o:
		return metrics
	case <-time.After(10 * time.Second):
		panic("Snapshot of RollingWindowMetrics failed")
	}

}

func (r *Metrics) add(result *vegeta.Result) {
	r.window[r.next] = result
	r.next = (r.next + 1) % len(r.window)
}

func (r Metrics) snapshot() vegeta.Metrics {
	metrics := new(vegeta.Metrics)
	for _, res := range r.window {
		if res == nil {
			break
		}
		metrics.Add(res)
	}
	metrics.Close()
	return *metrics
}

type cmd struct {
	Type        string
	AddIn       *vegeta.Result
	SnapshotOut chan vegeta.Metrics
}
