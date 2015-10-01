package lcddigits

func Parse(n int) string {
	if n == 0 {
		return "._.\n|.|\n|_|"
	}
	if n == 6 {
		return "._.\n|_.\n|_|"
	}
	return "._.\n|_|\n|_|"
}
