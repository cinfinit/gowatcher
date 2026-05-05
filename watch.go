//go:build !dev

package gowatcher

// Watch is a no-op in non-dev builds.
func Watch(dir string) {}
