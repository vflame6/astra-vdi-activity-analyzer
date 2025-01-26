package worker

import (
	"github.com/vflame6/astra-vdi-activity-analyzer/agent/capture"
	"github.com/vflame6/astra-vdi-activity-analyzer/agent/http"
	"log"
	"runtime"
	"time"
)

type Worker struct {
	id   int
	stop *chan bool
}

func NewWorker() *Worker {
	s := make(chan bool)
	w := &Worker{
		id:   1,
		stop: &s,
	}
	return w
}

func (w *Worker) RunOnlineScreenshoter(serverURL, hostname, secret string) {
	url := serverURL + "/api/screenshot/" + hostname

	go func() {
		for {
			select {
			case <-*w.stop:
				return
			default:
				filenames := capture.CaptureScreen(hostname)
				for _, filename := range filenames {
					err := http.SendScreenshot(url, filename, secret)
					if err != nil {
						log.Printf("Failed to send screenshot: %v\n", err)
					}
					err = capture.DeleteScreenshot(filename)
					if err != nil {
						log.Printf("Failed to delete screenshot: %v\n", err)
					}
				}
				time.Sleep(5 * time.Second)
			}
			runtime.Gosched()
		}
	}()
}

func (w *Worker) RunOfflineScreenshoter(hostname string) {
	go func() {
		for {
			select {
			case <-*w.stop:
				return
			default:
				filenames := capture.CaptureScreen(hostname)
				for _, filename := range filenames {
					log.Println("Got a screenshot: " + filename)
				}
				time.Sleep(5 * time.Second)
			}
			runtime.Gosched()
		}
	}()
}

func (w *Worker) Shutdown() {
	*w.stop <- true
	close(*w.stop)
}
