// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/joy12825/gf.

package main

import (
	"context"

	"go.opentelemetry.io/otel/exporters/prometheus"

	"github.com/joy12825/gf/v2/frame/g"
	"github.com/joy12825/gf/v2/os/gctx"
	"github.com/joy12825/gf/v2/os/gmetric"
)

var (
	meter = gmetric.GetGlobalProvider().Meter(gmetric.MeterOption{
		Instrument:        "github.com/joy12825/gf/example/metric/basic",
		InstrumentVersion: "v1.0",
	})
	counter = meter.MustCounter(
		"goframe.metric.demo.counter",
		gmetric.MetricOption{
			Help: "This is a simple demo for Counter usage",
			Unit: "bytes",
			Attributes: gmetric.Attributes{
				gmetric.NewAttribute("const_attr_1", 1),
			},
		},
	)
	observableCounter = meter.MustObservableCounter(
		"goframe.metric.demo.observable_counter",
		gmetric.MetricOption{
			Help: "This is a simple demo for ObservableCounter usage",
			Unit: "%",
			Attributes: gmetric.Attributes{
				gmetric.NewAttribute("const_attr_4", 4),
			},
		},
	)
)

func main() {
	var ctx = gctx.New()

	// Callback for observable metrics.
	meter.MustRegisterCallback(func(ctx context.Context, obs gmetric.Observer) error {
		obs.Observe(observableCounter, 10, gmetric.Option{
			Attributes: gmetric.Attributes{
				gmetric.NewAttribute("dynamic_attr_1", 1),
			},
		})
		return nil
	}, observableCounter)

	// Prometheus exporter to export metrics as Prometheus format.
	exporter, err := prometheus.New(
		prometheus.WithoutCounterSuffixes(),
		prometheus.WithoutUnits(),
	)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	// OpenTelemetry provider.
	provider := otelmetric.MustProvider(
		otelmetric.WithReader(exporter),
		otelmetric.WithBuiltInMetrics(),
	)
	provider.SetAsGlobal()
	defer provider.Shutdown(ctx)

	// Counter.
	counter.Inc(ctx)
	counter.Add(ctx, 10, gmetric.Option{
		Attributes: gmetric.Attributes{
			gmetric.NewAttribute("dynamic_attr_2", 2),
		},
	})

	// HTTP Server for metrics exporting.
	otelmetric.StartPrometheusMetricsServer(8000, "/metrics")
}
