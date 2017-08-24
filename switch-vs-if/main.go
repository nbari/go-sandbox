package main

func A(s string) string {
	switch s {
	case "alpha":
		return s
	case "betta":
		return s
	case "gamma":
		return s
	default:
		return s
	}
}

func B(s string) string {
	if s == "alpha" {
		return s
	} else if s == "betta" {
		return s
	} else if s == "gamma" {
		return s
	} else {
		return s
	}
}
