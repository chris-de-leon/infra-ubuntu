package version

import "fmt"

func WithoutPrefix() string {
	return "1.0.3"
}

func WithPrefix() string {
	return fmt.Sprintf("v%s", WithoutPrefix())
}
