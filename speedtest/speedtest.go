package speedtest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type Result struct {
	Type      string
	Timestamp string
	Ping      struct {
		Jitter  float64
		Latency float64
	}
	Download struct {
		Bandwidth uint64
		Bytes     uint64
		Elapsed   uint64
	}
	Upload struct {
		Bandwidth uint64
		Bytes     uint64
		Elapsed   uint64
	}
	PacketLoss float64
	ISP        string
	Interface  struct {
		InternalIP string
		Name       string
		MacAddr    string
		IsVPN      bool
		ExternalIP string
	}
	Server struct {
		ID       uint64
		Host     string
		Port     uint64
		Name     string
		Location string
		Country  string
		IP       string
	}
	Result struct {
		ID        string
		URL       string
		Persisted bool
	}
}

type Runner struct {
	commandPath string
	serverID int64
}

func NewRunner(commandPath string, serviceID int64) *Runner {
	return &Runner{commandPath: commandPath, serverID: serviceID}
}

func (r *Runner) Validate() error {
	speedtestCmd := exec.Command(r.commandPath, "--version")
	out, err := speedtestCmd.Output()
	if err != nil {
		return err
	}
	if !strings.Contains(string(out), "Speedtest by Ookla") {
		return fmt.Errorf("%s is not an official speedtest CLI provided by Ookla", r.commandPath)
	}
	return nil
}

func (r *Runner) Run() (*Result, error) {
	speedtestCmd := exec.Command(r.commandPath, fmt.Sprintf("--server-id=%d", r.serverID), "--format=json")
	var stderr bytes.Buffer
	speedtestCmd.Stderr = &stderr
	out, err := speedtestCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("%w: %s", err, stderr.String())
	}
	result := &Result{}
	if err := json.Unmarshal(out, &result); err != nil {
		return nil, err
	}
	return result, nil
}
