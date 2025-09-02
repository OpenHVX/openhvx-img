package main

import (
	"flag"
	"fmt"
	"log"
	"openhvx-img/images"
)

func main() {
	root := flag.String("root", ".", "Root directory of images repository")
	output := flag.String("output", "./_index/images.json", "Output JSON index file")
	pretty := flag.Bool("pretty", false, "Pretty-print JSON")
	flag.Parse()

	idx, err := images.BuildIndex(*root)
	if err != nil {
		log.Fatalf("scan error: %v", err)
	}

	if err := images.WriteIndex(idx, *output, *pretty); err != nil {
		log.Fatalf("write error: %v", err)
	}

	fmt.Printf("Indexed %d images -> %s\n", len(idx.Images), *output)
}
