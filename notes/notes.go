package notes

// frequency in hertz
const (
	C5  float64 = 523
	Db5 float64 = 554
	D5  float64 = 587
	Eb5 float64 = 622
	E5  float64 = 659
	F5  float64 = 698
	Gb5 float64 = 740
	G5  float64 = 784
	Ab5 float64 = 831
	A5  float64 = 880
	Bb5 float64 = 932
	B5  float64 = 987
	C6  float64 = 1046
)

var CharToNote map[rune]float64
var KeyMap map[string]string

func init() {
	CharToNote = make(map[rune]float64)
	KeyMap = make(map[string]string)

	CharToNote['a'] = C5
	CharToNote['w'] = Db5
	CharToNote['s'] = D5
	CharToNote['e'] = Eb5
	CharToNote['d'] = E5
	CharToNote['f'] = F5
	CharToNote['t'] = Gb5
	CharToNote['g'] = G5
	CharToNote['y'] = Ab5
	CharToNote['h'] = A5
	CharToNote['u'] = Bb5
	CharToNote['j'] = B5
	CharToNote['k'] = C6

	KeyMap["a"] = "C"
	KeyMap["w"] = "C#"
	KeyMap["s"] = "D"
	KeyMap["e"] = "D#"
	KeyMap["d"] = "E"
	KeyMap["f"] = "F"
	KeyMap["t"] = "F#"
	KeyMap["g"] = "G"
	KeyMap["y"] = "G#"
	KeyMap["h"] = "A"
	KeyMap["u"] = "A#"
	KeyMap["j"] = "B"
	KeyMap["k"] = "C"
}
