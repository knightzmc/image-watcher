package main

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"time"
)

func main() {
	watchingDir := flag.String("dir", "./", "The directory to watch")
	flag.Parse()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Printf("Could not create watcher %s\n", err)
	}

	err = watcher.Add(*watchingDir)
	if err != nil {
		fmt.Printf("Could not add %s to watcher %s", &watchingDir, err)
	}

	for {
		event := <-watcher.Events
		if event.Op == fsnotify.Create {
			go func() {
				time.Sleep(time.Second)
				fmt.Println(event)
			}()
		}
	}
}
