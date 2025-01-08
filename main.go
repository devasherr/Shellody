package main

import (
	"fmt"

	"github.com/devasherr/terminal_music/notes"
	"golang.org/x/sys/windows"
)

func main() {
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	beepProc := kernel32.NewProc("Beep")

	// Duration in milliseconds
	duration := 300

	keyNotes := []uintptr{notes.C5, notes.Db5, notes.D5, notes.Eb5, notes.E5, notes.F5, notes.Gb5, notes.G5, notes.Ab5, notes.A5, notes.Bb5, notes.B5}

	for _, key := range keyNotes {
		ret, _, err := beepProc.Call(key, uintptr(duration))
		if ret == 0 {
			fmt.Printf("Error calling Beep: %v\n", err)
			return
		}
	}
}
