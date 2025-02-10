package version

import "fmt"

func WithoutPrefix() string {
	return "1.0.0"
}

func WithPrefix() string {
	return fmt.Sprintf("v%s", WithoutPrefix())
}
