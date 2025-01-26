package worker

import (
	"github.com/vflame6/astra-vdi-activity-analyzer/agent/http"
	"github.com/vflame6/astra-vdi-activity-analyzer/agent/utils"
)

type Agent struct {
	Config    utils.Config
	ServerURL string
	w         *Worker
	noSender  bool
}

func Ping(serverURL string) error {
	url := serverURL + "/api/ping"
	err := http.SendPingRequest(url)
	if err != nil {
		return err
	}
	return nil
}

func Register(serverURL, hostname, password string) (string, error) {
	url := serverURL + "/api/register"
	key, err := http.SendRegisterRequest(url, hostname, password)
	if err != nil {
		return "", err
	}
	return key, nil
}

func NewAgent(config *utils.Config, serverURL string, noSender bool) Agent {
	worker := NewWorker()
	return Agent{
		Config:    *config,
		ServerURL: serverURL,
		w:         worker,
		noSender:  noSender,
	}
}

func (a *Agent) Start() {
	if a.noSender {
		a.w.RunOfflineScreenshoter(a.Config.ClientName)
	} else {
		a.w.RunOnlineScreenshoter(a.ServerURL, a.Config.ClientName, a.Config.Key)
	}
}

func (a *Agent) Stop() {
	a.w.Shutdown()
}

func (a *Agent) HealthCheck() error {
	url := a.ServerURL + "/api/health"
	err := http.SendHealthCheckRequest(url, a.Config.ClientName, a.Config.Key)
	if err != nil {
		return err
	}

	return nil
}
