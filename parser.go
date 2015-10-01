package lcddigits

func Parse(n int) string {
	switch n {
	case 0:
		return "._.|.||_|"
	case 1:
		return ".....|..|"
	case 2:
		return "._.._||_."
	case 3:
		return "._.._|._|"
	case 5:
		return "._.|_.._|"
	case 6:
		return "._.|_.|_|"
	case 7:
		return "._.|.|..|"
	case 8:
		return "._.|_||_|"
	case 9:
		return "._.|_|..|"
	}
	return ""
}
