package hoi3

type FileLocation struct {
	BaseDirectory string
	ModDirectory  string
	SaveFile      string
}

const baseDirectory = "/Users/thalesfu/Windows/steam/steamapps/common/Hearts of Iron 3"
const modDirectory = "/Users/thalesfu/Windows/steam/steamapps/common/Hearts of Iron 3/tfh/mod/ChineseLeaders"

func NewFileLocation(baseDirectory, modDirectory, saveFile string) *FileLocation {
	return &FileLocation{BaseDirectory: baseDirectory, ModDirectory: modDirectory, SaveFile: saveFile}
}

func NewFileLocationWithDefault(saveFile string) *FileLocation {
	return &FileLocation{BaseDirectory: baseDirectory, ModDirectory: modDirectory, SaveFile: saveFile}
}
