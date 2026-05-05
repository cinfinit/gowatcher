//go:build dev

package gowatcher

import (
	"context"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

// Config holds watch configuration.
type Config struct {
	Dir       string
	Pattern   string
	OnReload  func()
	BuildArgs []string
	RunArgs   []string
}

// Watch starts file watching in dev mode.
func Watch(dir string) {
	cfg := Config{Dir: dir}
	watch(cfg)
}

// WatchWithConfig starts watching with custom config.
func WatchWithConfig(cfg Config) {
	watch(cfg)
}

func watch(cfg Config) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Printf("[gowatcher] failed to create watcher: %v", err)
		return
	}
	defer watcher.Close()

	err = filepath.Walk(cfg.Dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		log.Printf("[gowatcher] failed to walk dir: %v", err)
		return
	}

	log.Printf("[gowatcher] watching %s", cfg.Dir)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go reloader(ctx, cfg)

	debounce := time.NewTimer(0)
	<-debounce.C // drain

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Has(fsnotify.Write) || event.Has(fsnotify.Create) {
				if shouldReload(event.Name) {
					debounce.Reset(300 * time.Millisecond)
					<-debounce.C
					log.Printf("[gowatcher] reload triggered by %s", event.Name)
					cancel()
					ctx, cancel = context.WithCancel(context.Background())
					go reloader(ctx, cfg)
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Printf("[gowatcher] watcher error: %v", err)
		}
	}
}

func reloader(ctx context.Context, cfg Config) {
	if cfg.OnReload != nil {
		cfg.OnReload()
	}

	buildArgs := append([]string{"build", "-o", tempBinary()}, cfg.BuildArgs...)
	buildArgs = append(buildArgs, cfg.Dir)

	cmd := exec.CommandContext(ctx, "go", buildArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Printf("[gowatcher] build failed: %v", err)
		return
	}

	runArgs := append([]string{tempBinary()}, cfg.RunArgs...)
	runCmd := exec.CommandContext(ctx, runArgs[0], runArgs[1:]...)
	runCmd.Stdout = os.Stdout
	runCmd.Stderr = os.Stderr
	if err := runCmd.Start(); err != nil {
		log.Printf("[gowatcher] failed to start: %v", err)
		return
	}
	<-ctx.Done()
	runCmd.Process.Kill()
}

func shouldReload(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return ext == ".go" || ext == ".mod" || ext == ".sum"
}

func tempBinary() string {
	if runtime.GOOS == "windows" {
		return filepath.Join(os.TempDir(), "gowatcher_app.exe")
	}
	return filepath.Join(os.TempDir(), "gowatcher_app")
}
