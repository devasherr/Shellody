package main

import (
	"bytes"
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/charmbracelet/lipgloss"
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

	whiteKeyStyle := lipgloss.NewStyle().
		Width(4).
		Height(10).
		Align(lipgloss.Center, lipgloss.Bottom).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240")).
		Background(lipgloss.Color("15"))

	blackKeyStyle := lipgloss.NewStyle().
		Width(2).
		Height(6).
		Align(lipgloss.Center, lipgloss.Bottom).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240")).
		Background(lipgloss.Color("0"))

	whiteKeyPressedStyle := whiteKeyStyle.Background(lipgloss.Color("220"))
	blackKeyPressedStyle := blackKeyStyle.Background(lipgloss.Color("220"))

	whiteKeys := []string{"a", "s", "d", "f", "g", "h", "j", "k"}
	blackKeys := []string{"w", "e", "", "t", "y", "u"}

	keyStyles := make(map[rune]lipgloss.Style)
	for i, key := range whiteKeys {
		keyStyles[rune(key[0])] = whiteKeyStyle
		if i < len(blackKeys) && blackKeys[i] != "" {
			keyStyles[rune(blackKeys[i][0])] = blackKeyStyle
		}
	}

	// State of each key
	pressedKeys := make(map[rune]bool)

	renderPiano := func() string {
		var pianoRow []string
		for i, key := range whiteKeys {
			style := whiteKeyStyle
			if pressedKeys[rune(key[0])] {
				style = whiteKeyPressedStyle
			}
			pianoRow = append(pianoRow, style.Render(notes.KeyMap[key]))

			if i < len(blackKeys) && blackKeys[i] != "" {
				style := blackKeyStyle
				if pressedKeys[rune(blackKeys[i][0])] {
					style = blackKeyPressedStyle
				}
				pianoRow = append(pianoRow, style.Render(notes.KeyMap[blackKeys[i]]))
			}
		}
		return lipgloss.JoinHorizontal(lipgloss.Top, pianoRow...)
	}

	fmt.Print("\033[H\033[2J")
	fmt.Print("\033[?25l")
	fmt.Println(renderPiano())

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

		if freq, ok := notes.CharToNote[char]; ok {
			wave := generateSineWave(freq, time.Millisecond*time.Duration(duration))
			player = ctx.NewPlayer(bytes.NewReader(wave))
			player.Play()

			pressedKeys[char] = true

			fmt.Print("\033[H")
			fmt.Println(renderPiano())

			time.Sleep(100 * time.Millisecond)

			pressedKeys[char] = false
			fmt.Print("\033[H")
			fmt.Println(renderPiano())
		}

		lock.Unlock()
	}
}
