package lcddigits

var (
	lcd_values = [][]byte{
		[]byte("._.|.||_|"),
		[]byte(".....|..|"),
		[]byte("._.._||_."),
		[]byte("._.._|._|"),
		[]byte("...|_|..|"),
		[]byte("._.|_.._|"),
		[]byte("._.|_.|_|"),
		[]byte("._.|.|..|"),
		[]byte("._.|_||_|"),
		[]byte("._.|_|..|"),
	}
)

func Parse(n int) string {
	return string(lcd_values[n])
}
