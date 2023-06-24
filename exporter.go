package mackerelspeedtest

import (
	"time"

	"github.com/mackerelio/mackerel-client-go"
)

type MackerelExporter struct {
	client      *mackerel.Client
	serviceName string
}

func NewMackerelExporter(apiKey, serviceName string) *MackerelExporter {
	client := mackerel.NewClient(apiKey)
	return &MackerelExporter{client: client, serviceName: serviceName}
}

func (e *MackerelExporter) Export(metrics *SpeedTestMetrics) error {
	ts := metrics.Timestamp.Round(time.Minute).Unix()
	err := e.client.PostServiceMetricValues(e.serviceName, []*mackerel.MetricValue{
		{
			Name:  "speedtest.ping.latency",
			Time:  ts,
			Value: metrics.PingLatency.Seconds(),
		},
		{
			Name:  "speedtest.ping.jitter",
			Time:  ts,
			Value: metrics.PingJitter.Seconds(),
		},
		{
			Name:  "speedtest.bandwidth.download",
			Time:  ts,
			Value: metrics.DownloadBandwidthBps,
		},
		{
			Name:  "speedtest.bandwidth.upload",
			Time:  ts,
			Value: metrics.UploadBandwidthBps,
		},
		{
			Name:  "speedtest.packet.loss_percentage",
			Time:  ts,
			Value: metrics.PacketLossPercentage,
		},
	})
	return err
}
