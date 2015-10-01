package lcddigits

func Parse(n int) string {
	switch n {
	case 0:
		return "._.\n|.|\n|_|"
	case 3:
		return "._.\n._|\n._|"
	case 5:
		return "._.\n|_.\n._|"
	case 6:
		return "._.\n|_.\n|_|"
	case 7:
		return "._.\n|.|\n..|"
	case 8:
		return "._.\n|_|\n|_|"
	case 9:
		return "._.\n|_|\n..|"
	}
	return ""
}
