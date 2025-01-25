package notes

// frequency in hertz
const (
	C5  uintptr = 523
	Db5 uintptr = 554
	D5  uintptr = 587
	Eb5 uintptr = 622
	E5  uintptr = 659
	F5  uintptr = 698
	Gb5 uintptr = 740
	G5  uintptr = 784
	Ab5 uintptr = 831
	A5  uintptr = 880
	Bb5 uintptr = 932
	B5  uintptr = 987
)

var CharToNote map[rune]uintptr

func init() {
	CharToNote = make(map[rune]uintptr)

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
}
