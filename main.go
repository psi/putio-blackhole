package main

import (
	"fmt"
	"log"
	"path"

	"github.com/fsnotify/fsnotify"
)

func main() {
	cfg := LoadConfig()

	putioClient := NewPutIOClient(cfg.PutIOToken)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Create == fsnotify.Create {
					torrentUrl := fmt.Sprintf("%s/%s", cfg.BaseURL, path.Base(event.Name))
					err = AddTorrentToPutIO(torrentUrl, putioClient)

					if err != nil {
						log.Println("Failed to add to Put.IO:", torrentUrl)
						log.Println(err)
					} else {
						log.Println("Added file to Put.IO:", event.Name)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(cfg.WatchDir)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
