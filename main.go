/**
 * Program dir-tree-print prints the size and crc32 of each file in a directory
 * recursively.
 */
package main

import (
	"flag"
	"fmt"
	"hash/crc32"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/golang/glog"
)

var (
	dir = flag.String("dir", ".", "Directory to print info about.")
)

func main() {
	flag.Parse()
	if err := mainErr(); err != nil {
		glog.Errorf("error running program: %v", err)
		os.Exit(1)
	}
}

func mainErr() error {
	messages := []string{}
	err := filepath.Walk(*dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			hash, err := fingerprint(path)
			if err != nil {
				return err
			}
			messages = append(messages, fmt.Sprintf("%s: %s", strings.TrimPrefix(path, *dir), hash))
			return nil
		})
	if err != nil {
		return err
	}
	sort.Strings(messages)
	fmt.Printf("%s\n", strings.Join(messages, "\n"))
	return nil
}

func fingerprint(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "error", nil
	}
	return fmt.Sprintf("%d bytes; crc32 %d", len(data), crc32.ChecksumIEEE(data)), nil
}
