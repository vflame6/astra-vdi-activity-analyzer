package worker

import (
	"github.com/vflame6/astra-vdi-activity-analyzer/agent/http"
	"github.com/vflame6/astra-vdi-activity-analyzer/agent/utils"
)

type Agent struct {
	Config    utils.Config
	ServerURL string
	w         *Worker
}

func Ping(url string) error {
	err := http.SendPingRequest(url)
	if err != nil {
		return err
	}
	return nil
}

func Register(url, hostname, password string) (string, error) {
	key, err := http.SendRegisterRequest(url, hostname, password)
	if err != nil {
		return "", err
	}
	return key, nil
}

func NewAgent(config *utils.Config, serverURL string) Agent {
	worker := NewWorker()
	return Agent{
		Config:    *config,
		ServerURL: serverURL,
		w:         worker,
	}
}

func (a *Agent) Start() {
	a.w.RunScreenshoter(a.ServerURL, a.Config.ClientName, a.Config.Key)
}

func (a *Agent) Stop() {
	a.w.Shutdown()
}

func (a *Agent) HealthCheck() error {
	err := http.SendHealthCheckRequest(a.Config.Address, a.Config.ClientName, a.Config.Key)
	if err != nil {
		return err
	}

	return nil
}
