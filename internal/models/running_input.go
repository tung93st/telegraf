package models

import (
	"fmt"
	"time"

	"github.com/influxdata/telegraf/plugins"
	"github.com/influxdata/telegraf/selfstat"
)

var GlobalMetricsGathered = selfstat.Register("agent", "metrics_gathered", map[string]string{})

type RunningInput struct {
	Input  plugins.Input
	Config *InputConfig

	trace       bool
	defaultTags map[string]string

	MetricsGathered selfstat.Stat
}

func NewRunningInput(
	input plugins.Input,
	config *InputConfig,
) *RunningInput {
	return &RunningInput{
		Input:  input,
		Config: config,
		MetricsGathered: selfstat.Register(
			"gather",
			"metrics_gathered",
			map[string]string{"input": config.Name},
		),
	}
}

// InputConfig containing a name, interval, and filter
type InputConfig struct {
	Name              string
	NameOverride      string
	MeasurementPrefix string
	MeasurementSuffix string
	Tags              map[string]string
	Filter            Filter
	Interval          time.Duration
}

func (r *RunningInput) Name() string {
	return "inputs." + r.Config.Name
}

// MakeMetric either returns a metric, or returns nil if the metric doesn't
// need to be created (because of filtering, an error, etc.)
func (r *RunningInput) MakeMetric(
	measurement string,
	fields map[string]interface{},
	tags map[string]string,
	mType plugins.ValueType,
	t time.Time,
) plugins.Metric {
	m := makemetric(
		measurement,
		fields,
		tags,
		r.Config.NameOverride,
		r.Config.MeasurementPrefix,
		r.Config.MeasurementSuffix,
		r.Config.Tags,
		r.defaultTags,
		r.Config.Filter,
		true,
		mType,
		t,
	)

	if r.trace && m != nil {
		fmt.Println("> " + m.String())
	}

	r.MetricsGathered.Incr(1)
	GlobalMetricsGathered.Incr(1)
	return m
}

func (r *RunningInput) Trace() bool {
	return r.trace
}

func (r *RunningInput) SetTrace(trace bool) {
	r.trace = trace
}

func (r *RunningInput) SetDefaultTags(tags map[string]string) {
	r.defaultTags = tags
}
