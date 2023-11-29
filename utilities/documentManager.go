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
	if len(stringed) == 1 {
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

// Collects documents as Document class from the given path
func CollectDocuments(fromPath string) (docs []Document) {

	err := filepath.WalkDir(fromPath, func(path string, d fs.DirEntry, err error) error {

		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				panic(err)
			}
			tempDoc := Document{}
			tempDoc.Name = info.Name()
			tempDoc.Date = info.ModTime().Format("dd-mm-yyyy")
			tempDoc.Size = parseWheight(info.Size())
			tempDoc.Type = getExtentionType(info.Name())
			tempDoc.Path = path

			docs = append(docs, tempDoc)
		}
		return nil
	})
	if err != nil {
		return nil
	}
	return nil
}
