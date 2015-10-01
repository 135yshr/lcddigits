package lcddigits

func Parse(n int) (ret []byte) {
	switch n {
	case 0:
		ret = []byte("._.|.||_|")
	case 1:
		ret = []byte(".....|..|")
	case 2:
		ret = []byte("._.._||_.")
	case 3:
		ret = []byte("._.._|._|")
	case 5:
		ret = []byte("._.|_.._|")
	case 6:
		ret = []byte("._.|_.|_|")
	case 7:
		ret = []byte("._.|.|..|")
	case 8:
		ret = []byte("._.|_||_|")
	case 9:
		ret = []byte("._.|_|..|")
	}
	return
}
