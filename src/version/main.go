package version

import "fmt"

func WithoutPrefix() string {
	return "1.1.2"
}

func WithPrefix() string {
	return fmt.Sprintf("v%s", WithoutPrefix())
}
