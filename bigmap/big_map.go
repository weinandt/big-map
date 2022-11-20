package big_map

import (
	"log"
	"syscall"
)

func mmap() {
	bytes, err := syscall.Mmap(-1, 0, 1000000, syscall.PROT_WRITE|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_SHARED)
	if err != nil {
		log.Fatal("Problem with mmap", err)
	}

	log.Println("length: ", len(bytes))
	log.Println(bytes[0])
}

func Test() {
	log.Println("in Test")
	mmap()
	log.Println("done withe test.")
}
