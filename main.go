package main

import (
	"fmt"

	"github.com/devasherr/terminal_music/notes"
	"github.com/eiannone/keyboard"
	"golang.org/x/sys/windows"
)

func main() {
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	beepProc := kernel32.NewProc("Beep")

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press ESC to quit")
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		beepProc.Call(notes.CharToNote[char], uintptr(300))

		if key == keyboard.KeyEsc {
			break
		}
	}
}
