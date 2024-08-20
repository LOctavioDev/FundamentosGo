package main

import (
	"image/color"
	"time"
)

// FILE TYPE
const (
	fileRegular int = iota
	fileDirectory
	fileExecutable
	fileCompress
	fileImage
	fileLink
)

// FILE EXTENSION

const (
	exe  = ".exe"
	deb  = ".deb"
	zip  = ".zip"
	gz   = ".gz"
	tar  = ".tar"
	rar  = ".rar"
	png  = ".png"
	jpg  = ".jpg"
	jpeg = "jpeg"
	gif  = ".gif"
)

// ESTRUCTURA DE ARCHIVOS

type File struct {
	name             string
	fileType         int
	isDir            bool
	isHidden         bool
	userName         string
	size             int64
	modificationTime time.Time
	mode             string
}

type styleFileType struct {
	icon   string
	color  string
	symbol string
}

// MAPA PARA LOS TIPOS DE ARCHIVO
var mapStyleByFileType = map[int]styleFileType{
	fileRegular:    {icon: "📄"},
	fileDirectory:  {icon: "📁", color: "BLUE", symbol: "/"},
	fileExecutable: {icon: "⏏️", color: "GREEN", symbol: "*"},
	fileCompress:   {icon: "📦", color: "RED"},
	fileImage:      {icon: "📷", color: "MAGENTA"},
	fileLink:       {icon: "🔗", color: "CYAN"},
}
