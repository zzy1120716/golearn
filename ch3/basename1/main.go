package main

// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c => c, a/b.c.go => b.c
func basename(s string) string {
	// Discard last '/' and everything before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func main() {
	println(basename("a"))
	println(basename("a.go"))
	println(basename("a/b/c"))
	println(basename("a/b.c.go"))
}
