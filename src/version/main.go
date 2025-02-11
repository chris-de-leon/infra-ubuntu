package version

import "fmt"

func WithoutPrefix() string {
	return "1.0.5"
}

func WithPrefix() string {
	return fmt.Sprintf("v%s", WithoutPrefix())
}
