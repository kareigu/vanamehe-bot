package utils

import "fmt"

const VideoCount = 5

var Videos [VideoCount]string

const (
	HUUTTIS    = iota
	JÄTTIS     = iota
	KÄÄNNÖS    = iota
	KIINNOSTAA = iota
	HOMO       = iota
)

var videoNames = []string{
	"huuttis.mp4",
	"j%C3%A4ttis.mp4",
	"käännös.mp4",
	"kiinnostaa.mp4",
	"homo.mp4",
}

func init() {
	for i, name := range videoNames {
		slice := Videos[0:i]
		slice = append(slice, fmt.Sprintf("https://karei.dev/files/vanamehe/%s", name))
	}
}
