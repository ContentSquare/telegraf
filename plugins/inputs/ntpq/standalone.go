package ntpq

import (
	"github.com/influxdata/telegraf"
	"time"
	"github.com/influxdata/telegraf/metric"
	"github.com/influxdata/telegraf/agent"
	"github.com/influxdata/telegraf/plugins/outputs/file"
	"os"
	"github.com/influxdata/telegraf/plugins/serializers"
)

type StandaloneMetricMaker struct {
}

func (tm *StandaloneMetricMaker) Name() string {
	return "Standalone"
}
func (tm *StandaloneMetricMaker) MakeMetric(
	measurement string,
	fields map[string]interface{},
	tags map[string]string,
	mType telegraf.ValueType,
	t time.Time,
) telegraf.Metric {
	switch mType {
	case telegraf.Untyped:
		if m, err := metric.New(measurement, tags, fields, t); err == nil {
			return m
		}
	case telegraf.Counter:
		if m, err := metric.New(measurement, tags, fields, t, telegraf.Counter); err == nil {
			return m
		}
	case telegraf.Gauge:
		if m, err := metric.New(measurement, tags, fields, t, telegraf.Gauge); err == nil {
			return m
		}
	}
	return nil
}

func Standalone() {
	fh, err := os.Open("/dev/stdout")
	if err != nil {
		panic(err)
	}

	f := file.File{
		Files: []string{fh.Name()},
	}
	s, _ := serializers.NewInfluxSerializer()
	f.SetSerializer(s)
	f.Connect()

	chanMetrics := make(chan telegraf.Metric, 10)

	acc := agent.NewAccumulator(&StandaloneMetricMaker{}, chanMetrics)

	n := &NTPQ{}
	n.runQ = n.runq
	n.Gather(acc)

	close(chanMetrics)

	for metrics := range chanMetrics {
		f.Write([]telegraf.Metric{metrics})
	}
}
