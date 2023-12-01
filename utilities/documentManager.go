package utilities

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

// Class used to blueprint docs collected to pass to doc viewer
type Document struct {
	Name string
	Size string
	Type string
	Date string
	Path string
	Icon string
}

// Class used to pass data to html
type DataToPass struct {
	Email string
	Data  []Document
}

func InString(searchfor, in string) bool {
	if len(searchfor) > len(in) {
		return false
	}
	for n := 0; n <= len(in)-len(searchfor); n++ {
		if in[n:n+len(searchfor)] == searchfor {
			return true
		}
	}
	return false
}

// Function to check if test is in arrayOfString
func inArray(test string, arrayOfString []string) bool {
	for a := range arrayOfString {
		if arrayOfString[a] == test {
			return true
		}
	}
	return false
}

// Returns the extension type of the given name file
func getExtentionType(name string) string {
	splitted := strings.Split(name, ".")
	return splitted[len(splitted)-1]
}

func parseWheight(wgt int64) string {
	if wgt == 0 {
		return fmt.Sprint(0) + "Kb"
	}
	res := float32(wgt) / 102.4
	frmt := 1
	for res >= 999 {
		res = res / 102.4
		frmt++
	}
	stringed := strings.SplitAfter(fmt.Sprintf("%v", res), ".")
	out := stringed[0][:len(stringed[0])-1] + "."
	if len(stringed) <= 1 {
		out += "00"
	} else {
		out += stringed[1][:2]
	}
	switch frmt {
	case 1:
		out += "Kb"
	case 2:
		out += "Mb"
	case 3:
		out += "Gb"
	case 4:
		out += "Tb"
	case 5:
		out += "Pb"
	case 6:
		out += "Zb"
	}
	return out
}

// Returns the icon path to set it in the data field
func SetIcon(name string) (out string) {
	parts := strings.SplitAfter(name, ".")
	filetype := strings.ToUpper(parts[len(parts)-1])
	base := "/public/icons/"
	if inArray(filetype, []string{"7Z", "RAR", "CPIO", "LBR", "ISO", "DMG", "JAR"}) {
		out = base + "rar.png"
	} else if inArray(filetype, []string{"CSV", "XLSX"}) {
		out = base + "excellico.png"
	} else if inArray(filetype, []string{"MP3,M4A,FLAC,WAV,WMA,AAC"}) {
		out = base + "musicico.png"
	} else if filetype == "PDF" {
		out = base + "pdfico.png"
	} else if filetype == "PPTX" {
		out = base + "powerpico.png"
	} else if inArray(filetype, []string{"PNG", "JPG", "GIF", "SVG"}) {
		out = base + "images.png"
	} else if filetype == "TXT" {
		out = base + "txtico.png"
	} else if inArray(filetype, []string{"MP4", "MOV", "WMV", "AVI", "AVCHD", "FLV", "F4V", "SWF", "MKV", "WEBM", "HTML5", "MPEG"}) {
		out = base + "videoico.png"
	} else {
		out = base + "document-round.svg"
	}
	return out
}

// Collects documents as Document class from the given path
func CollectDocuments(fromPath string) (docs []Document, err error) {
	err = filepath.WalkDir(fromPath, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				panic(err)
			}
			tempDoc := Document{}
			tempDoc.Name = info.Name()
			tempDoc.Date = info.ModTime().Format("dd-mm-yyyy")
			tempDoc.Size = string(info.Size())
			tempDoc.Type = getExtentionType(info.Name())
			tempDoc.Path = path
			tempDoc.Icon = SetIcon(info.Name())

			docs = append(docs, tempDoc)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return
}
