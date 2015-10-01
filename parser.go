package lcddigits

var (
	lcd_values = [][]byte{
		[]byte("._.\n|.|\n|_|"),
		[]byte("...\n..|\n..|"),
		[]byte("._.\n._|\n|_."),
		[]byte("._.\n._|\n._|"),
		[]byte("...\n|_|\n..|"),
		[]byte("._.\n|_.\n._|"),
		[]byte("._.\n|_.\n|_|"),
		[]byte("._.\n|.|\n..|"),
		[]byte("._.\n|_|\n|_|"),
		[]byte("._.\n|_|\n..|"),
	}
)

func Parse(n int) string {
	return string(lcd_values[n])
}
