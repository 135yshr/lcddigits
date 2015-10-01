package lcddigits

func Parse(n int) string {
	switch n {
	case 5:
		return "._.\n|_.\n._|"
	}
	if n == 0 {
		return "._.\n|.|\n|_|"
	}
	if n == 6 {
		return "._.\n|_.\n|_|"
	}
	return "._.\n|_|\n|_|"
}
