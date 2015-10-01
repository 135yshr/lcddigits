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

func Parse(n int) (ret []byte) {
	switch n {
	case 0:
		ret = lcd_values[0]
	case 1:
		ret = lcd_values[1]
	case 2:
		ret = lcd_values[2]
	case 3:
		ret = lcd_values[3]
	case 4:
		ret = lcd_values[4]
	case 5:
		ret = lcd_values[5]
	case 6:
		ret = lcd_values[6]
	case 7:
		ret = lcd_values[7]
	case 8:
		ret = lcd_values[8]
	case 9:
		ret = lcd_values[9]
	}
	return
}
