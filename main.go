package main

import (
	"fmt"
	//	"time"
)

type Config struct {
	work_dir_path string
}

func main() {
	repo, err := NewRepository("/tmp/whatever/hello", "git@github.com:Primer42/test.git")
	if err != nil {
		fmt.Println("Got error:", err.Error())
		return
	}
	fmt.Println(repo)
	/*
		watcher_delay, err := time.ParseDuration("15s")
		if err != nil {
			fmt.Sprintln("Error parsing delay: %s", err.Error())
		}
		gw := GitWatcher{"remote", watcher_delay}
		err = gw.Watch()
	*/

}
