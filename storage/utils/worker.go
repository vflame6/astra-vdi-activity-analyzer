package utils

import (
	"github.com/vflame6/astra-vdi-activity-analyzer/storage/filesystem"
	"github.com/vflame6/astra-vdi-activity-analyzer/storage/http"
	"log"
	"runtime"
	"time"
)

type Worker struct {
	stop    *chan bool
	address string
}

func NewWorker(address string) *Worker {
	s := make(chan bool)
	w := &Worker{
		stop:    &s,
		address: address,
	}
	return w
}

func (w *Worker) Start() {
	go func() {
		for {
			select {
			case <-*w.stop:
				return
			default:
				filenames, err := filesystem.ListScreenshots()
				if err != nil {
					log.Println(err)
				}
				for _, filename := range filenames {
					err := http.SendScreenshot(filename, "http://"+w.address+"/api/screenshot")
					if err != nil {
						log.Println(err)
					}
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
