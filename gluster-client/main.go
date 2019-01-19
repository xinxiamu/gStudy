package main

import "github.com/gluster/gogfapi/gfapi"

func init() {
	println(">>>init1")
}

func main() {
	print("main:", "abc")
	vol := &gfapi.Volume{}
	if err := vol.Init("models", "192.168.200.10"); err != nil {
		// handle error
	}

	if err := vol.Mount(); err != nil {
		// handle error
	}
	defer vol.Unmount()

	f, err := vol.Create("testfile")
	if err != nil {
		// handle error
	}
	defer f.Close()

	if _, err := f.Write([]byte("hello中国人张牧天")); err != nil {
		// handle error
	}

	return
}
