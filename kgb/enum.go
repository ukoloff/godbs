package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/phf/go-queue/queue"
)

// GeoDet enumerates DBS files in GEODET folder
func GeoDet() chan string {
	c := make(chan string)

	q := queue.New()
	q.PushBack("C:/Geodet")

	go enumerate(q, c)

	return c
}

func enumerate(q *queue.Queue, c chan string) {
	for 0 != q.Len() {
		f, err := os.Open(q.PopFront().(string))
		if err != nil {
			panic(err)
		}

		for {
			files, _ := f.Readdir(1)
			if 0 == len(files) {
				break
			}
			file := files[0]
			filename := filepath.Join(f.Name(), file.Name())
			if file.IsDir() {
				q.PushBack(filename)
				continue
			}
			switch strings.ToLower(filepath.Ext(file.Name())) {
			case ".dbs":
				c <- filename
			}
		}
	}
	close(c)
}
