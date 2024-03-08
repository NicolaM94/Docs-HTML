package managers

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

type Document struct {
	Name     string
	Creation string
	Size     string
	Type     string
	Icon     string
}

// Parse the type of the file
func ParseFileType(name string) string {
	var out = ""
	for n := len(name) - 1; n >= 0; n-- {
		if name[n] == '.' {
			return out
		}
		out = string(name[n]) + out
	}
	return out
}

// Parses the weight of a file starting from its int64 bit count and returning a string with 6 carachters
// (. included) with a trailing bit type description
func ParseFileSize(wgt int64) string {
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

// Collects documents from the userfolder, used in data delivery after login
func CollectDocuments(userFolderName string) ([]Document, error) {
	// Init settings struct and links to the user folder
	// Normalize the path
	var docbasepath = Settings{}.Populate().DocBasePath
	if docbasepath[len(docbasepath)-1] != '/' {
		docbasepath = docbasepath + "/"
	}

	var collector []Document
	// Loops over the folder to collect documents
	filepath.WalkDir(docbasepath+userFolderName, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			temp := Document{}
			temp.Name = d.Name()

			info, err := d.Info()
			if err != nil {
				return err
			}
			temp.Creation = info.ModTime().Format("2006-01-02 15:04:05")
			temp.Type = ParseFileType(info.Name())
			temp.Size = ParseFileSize(info.Size())
			temp.Icon = "./fileicons/" + strings.ToLower(temp.Type) + ".png"
			collector = append(collector, temp)
		}
		return nil
	})
	return collector, nil
}
