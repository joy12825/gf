// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/joy12825/gf.

package otelmetric

import (
	"context"

	"go.opentelemetry.io/otel/metric"

	"github.com/joy12825/gf/v2/encoding/gjson"
	"github.com/joy12825/gf/v2/errors/gcode"
	"github.com/joy12825/gf/v2/errors/gerror"
	"github.com/joy12825/gf/v2/os/gmetric"
)

// localObservableUpDownCounterPerformer is an implementer for interface CounterPerformer.
type localObservableUpDownCounterPerformer struct {
	gmetric.ObservableMetric
	metric.Float64ObservableUpDownCounter
}

// newObservableUpDownCounterPerformer creates and returns a UpDownCounterPerformer that truly takes action to
// implement ObservableUpDownCounter.
func (l *localMeterPerformer) newObservableUpDownCounterPerformer(
	meter metric.Meter,
	metricName string,
	metricOption gmetric.MetricOption,
) (gmetric.ObservableUpDownCounterPerformer, error) {
	var (
		options = []metric.Float64ObservableUpDownCounterOption{
			metric.WithDescription(metricOption.Help),
			metric.WithUnit(metricOption.Unit),
		}
	)
	if metricOption.Callback != nil {
		callback := metric.WithFloat64Callback(func(ctx context.Context, observer metric.Float64Observer) error {
			return metricOption.Callback(ctx, l.newMetricObserver(metricOption, observer))
		})
		options = append(options, callback)
	}
	counter, err := meter.Float64ObservableUpDownCounter(metricName, options...)
	if err != nil {
		return nil, gerror.WrapCodef(
			gcode.CodeInternalError,
			err,
			`create Float64ObservableUpDownCounter "%s" failed with option: %s`,
			metricName, gjson.MustEncodeString(metricOption),
		)
	}
	return &localObservableUpDownCounterPerformer{
		Float64ObservableUpDownCounter: counter,
	}, nil
}
