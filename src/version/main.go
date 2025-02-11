package version

import "fmt"

func WithoutPrefix() string {
	return "1.0.6"
}

func WithPrefix() string {
	return fmt.Sprintf("v%s", WithoutPrefix())
}
