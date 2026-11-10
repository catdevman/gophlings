package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"time"
)

func main() {
	only := flag.Int("n", -1, "run only the exercise with this number")
	from := flag.Int("from", -1, "start at this exercise number")
	watch := flag.Bool("watch", false, "rerun the current exercise when its files change")
	list := flag.Bool("list", false, "list the exercises")
	flag.Parse()

	root, err := moduleRoot()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	exercises, err := discover(root)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if *list {
		for _, ex := range exercises {
			fmt.Printf("%2d  %s\n", ex.n, ex.dir)
		}
		return
	}

	selected := filter(exercises, *only, *from)
	if len(selected) == 0 {
		fmt.Fprintln(os.Stderr, "no exercise matched")
		os.Exit(1)
	}

	fmt.Println("Gophlings")
	fmt.Println("Learn Go by fixing small programs. The run stops at the first failure.")
	fmt.Println()

	for i, ex := range selected {
		for {
			fmt.Printf("==> [%d/%d] %s\n", i+1, len(selected), ex.dir)
			if test(root, ex.dir) {
				fmt.Println("pass")
				fmt.Println()
				break
			}
			fmt.Printf("\nFix %s and run again.\n", ex.dir)
			if !*watch {
				os.Exit(1)
			}
			fmt.Println("watching for changes...")
			if err := waitForChange(filepath.Join(root, ex.dir)); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			fmt.Println()
		}
	}
	fmt.Println("Every exercise passes. Nice work.")
}

type exercise struct {
	n   int
	dir string
}

var dirPattern = regexp.MustCompile(`^(\d+)_`)

func discover(root string) ([]exercise, error) {
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}
	var out []exercise
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		match := dirPattern.FindStringSubmatch(entry.Name())
		if match == nil {
			continue
		}
		n, err := strconv.Atoi(match[1])
		if err != nil {
			continue
		}
		out = append(out, exercise{n: n, dir: entry.Name()})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].n < out[j].n })
	return out, nil
}

func filter(all []exercise, only, from int) []exercise {
	if only >= 0 {
		for _, ex := range all {
			if ex.n == only {
				return []exercise{ex}
			}
		}
		return nil
	}
	if from < 0 {
		return all
	}
	var out []exercise
	for _, ex := range all {
		if ex.n >= from {
			out = append(out, ex)
		}
	}
	return out
}

func test(root, dir string) bool {
	cmd := exec.Command("go", "test", "-count=1", "-timeout", "30s", "-race", "./"+dir)
	cmd.Dir = root
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run() == nil
}

func moduleRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("go.mod not found")
		}
		dir = parent
	}
}

func waitForChange(dir string) error {
	before, err := snapshot(dir)
	if err != nil {
		return err
	}
	for {
		time.Sleep(300 * time.Millisecond)
		after, err := snapshot(dir)
		if err != nil {
			return err
		}
		if !sameSnapshot(before, after) {
			return nil
		}
	}
}

func snapshot(dir string) (map[string]int64, error) {
	out := map[string]int64{}
	err := filepath.WalkDir(dir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil || entry.IsDir() {
			return err
		}
		info, err := entry.Info()
		if err != nil {
			return err
		}
		out[path] = info.ModTime().UnixNano() + info.Size()
		return nil
	})
	return out, err
}

func sameSnapshot(a, b map[string]int64) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
