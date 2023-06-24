package mackerelspeedtest

import (
	"time"

	"mgithub.com/Arthur1/mackerel-speedtest/speedtest"
)

type SpeedTestMetrics struct {
	Timestamp time.Time
	PingLatency time.Duration
	PingJitter time.Duration
	DownloadBandwidthBps uint64
	UploadBandwidthBps uint64
	PacketLossPercentage float64
}

type SpeedTestExecutor struct {
	commandPath string
	serverID int64
}

func NewSpeedTestExecutor(commandPath string, serverID int64) *SpeedTestExecutor {
	if commandPath == "" {
		commandPath = "speedtest"
	}
	return &SpeedTestExecutor{commandPath: commandPath, serverID: serverID}
}

func (e *SpeedTestExecutor) Execute() (*SpeedTestMetrics, error) {
	runner := speedtest.NewRunner(e.commandPath, e.serverID)
	if err := runner.Validate(); err != nil {
		return nil, err
	}
	result, err := runner.Run()
	if err != nil {
		return nil, err
	}
	ts, err := time.Parse(time.RFC3339, result.Timestamp)
	if err != nil {
		return nil, err
	}
	metrics := &SpeedTestMetrics{
		Timestamp: ts,
		PingLatency: time.Duration(result.Ping.Latency * float64(time.Millisecond)),
		PingJitter: time.Duration(result.Ping.Jitter * float64(time.Millisecond)),
		DownloadBandwidthBps: result.Download.Bandwidth * 8, // bytes/sec -> bits/sec
		UploadBandwidthBps: result.Upload.Bandwidth * 8, // bytes/sec -> bits/sec
		PacketLossPercentage: result.PacketLoss,
	}
	return metrics, nil
}
