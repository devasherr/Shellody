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
	CharToNote['s'] = Db5
	CharToNote['d'] = D5
	CharToNote['f'] = Eb5
	CharToNote['g'] = E5
	CharToNote['h'] = F5
	CharToNote['j'] = Gb5
	CharToNote['k'] = G5
	CharToNote['l'] = Ab5
	CharToNote['i'] = A5
	CharToNote['o'] = Bb5
	CharToNote['p'] = B5
}
