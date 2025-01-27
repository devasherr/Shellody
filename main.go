package main

import (
	"bytes"
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/devasherr/terminal_music/notes"
	"github.com/eiannone/keyboard"
	"github.com/hajimehoshi/oto/v2"
)

const sampleRate = 44100

func generateSineWave(freq float64, duration time.Duration) []byte {
	numSamples := int(float64(sampleRate) * duration.Seconds())
	wave := make([]byte, numSamples*2) // 16-bit PCM (2 bytes per sample)

	for i := 0; i < numSamples; i++ {
		sample := int16(math.Sin(2*math.Pi*freq*float64(i)/float64(sampleRate)) * 32767)
		wave[2*i] = byte(sample & 0xff)          // Low byte
		wave[2*i+1] = byte((sample >> 8) & 0xff) // High byte
	}

	return wave
}

func main() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	ctx, ready, err := oto.NewContext(sampleRate, 1, 2)
	if err != nil {
		panic(err)
	}

	<-ready
	var lock sync.Mutex
	var player oto.Player
	duration := 300

	fmt.Println("Press ESC to quit")
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyEsc {
			break
		}

		lock.Lock()
		if player != nil {
			player.Close()
		}

		wave := generateSineWave(notes.CharToNote[char], time.Millisecond*time.Duration(duration))
		player = ctx.NewPlayer(bytes.NewReader(wave))
		player.Play()

		lock.Unlock()
	}
}
