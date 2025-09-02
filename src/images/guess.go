package images

import "strings"

func GuessOS(name string) string {
	n := strings.ToLower(name)
	switch {
	case strings.Contains(n, "ubuntu"), strings.Contains(n, "debian"), strings.Contains(n, "rocky"),
		strings.Contains(n, "almalinux"), strings.Contains(n, "centos"):
		return "linux"
	case strings.Contains(n, "win"), strings.Contains(n, "windows"):
		return "windows"
	default:
		return "unknown"
	}
}

func GuessArch(name string) string {
	n := strings.ToLower(name)
	switch {
	case strings.Contains(n, "arm64"), strings.Contains(n, "aarch64"):
		return "arm64"
	case strings.Contains(n, "amd64"), strings.Contains(n, "x64"), strings.Contains(n, "x86_64"):
		return "x86_64"
	default:
		return "x86_64" // défaut raisonnable
	}
}

func GuessGen(name string) int {
	n := strings.ToLower(name)
	if strings.Contains(n, "gen1") {
		return 1
	}
	return 2
}
