package stringutil

func FullName(f, l string) (string, int) {
	full := f + " " + l
	len := len(full)
	return full, len
}

func FullNameNakedReturn(f, l string) (full string, length int) {
	full = f + " " + l
	length = len(full)
	return
}
