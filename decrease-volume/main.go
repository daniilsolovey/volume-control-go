package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"

	"github.com/itchyny/volume-go"
)

const (
	showingTime = "300"
)

func main() {
	vol, err := volume.GetVolume()
	if err != nil {
		log.Fatalf("get volume failed: %+v", err)
	}
	fmt.Printf("current volume: %d\n", vol)

	if vol == 0 {
		message := "volume: " + strconv.Itoa(vol) + "%"
		notify := exec.Command("notify-send", "-t", showingTime, "info", message)
		err = notify.Run()
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	err = volume.SetVolume(vol - 10)
	if err != nil {
		log.Fatalf("set volume failed: %+v", err)
	}

	fmt.Printf("set volume success\n")

	message := "volume: " + strconv.Itoa(vol-10) + "%"
	notify := exec.Command("notify-send", "-t", showingTime, "info", message)
	err = notify.Run()
	if err != nil {
		log.Fatal(err)
	}
}
