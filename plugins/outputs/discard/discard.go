package discard

import (
	"github.com/influxdata/telegraf/plugins"
	"github.com/influxdata/telegraf/plugins/outputs"
)

type Discard struct{}

func (d *Discard) Connect() error                        { return nil }
func (d *Discard) Close() error                          { return nil }
func (d *Discard) SampleConfig() string                  { return "" }
func (d *Discard) Description() string                   { return "Send metrics to nowhere at all" }
func (d *Discard) Write(metrics []plugins.Metric) error { return nil }

func init() {
	outputs.Add("discard", func() plugins.Output { return &Discard{} })
}
