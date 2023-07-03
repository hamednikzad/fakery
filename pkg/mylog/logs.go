package mylog

import (
	"github.com/hamednikzad/fakery/pkg/color"
	"log"
)

func Println(message string, c color.Color) {
	//log.Printf("Red: %s %s None: %s %s", color.ColorGRN, "red string", color.ColorMAG, "colorless string")
	log.Printf("%s%s%s\n", c, message, color.None)
}
