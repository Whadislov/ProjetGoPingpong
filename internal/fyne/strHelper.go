package myapp

func strHelper(list []string) string {
	str := ""
	for _, word := range list {
		str += word + ", "
	}
	// Remove extra ", "
	str = str[:len(str)-2]
	return str
}
