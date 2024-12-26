package hoi3

type FileLocation struct {
	BaseDirectory string
	DLCDirectory  string
	ModDirectory  string
	SaveFile      string
}

const BaseDirectory = "/Users/thalesfu/Windows/steam/steamapps/common/Hearts of Iron 3"
const DLCDirectory = "/Users/thalesfu/Windows/steam/steamapps/common/Hearts of Iron 3/tfh"
const ModDirectory = "/Users/thalesfu/Windows/steam/steamapps/common/Hearts of Iron 3/tfh/mod/ChineseLeaders"

func NewFileLocation(baseDirectory, dlcDirectory, modDirectory, saveFile string) *FileLocation {
	return &FileLocation{BaseDirectory: baseDirectory, DLCDirectory: dlcDirectory, ModDirectory: modDirectory, SaveFile: saveFile}
}

func NewFileLocationWithDefault(saveFile string) *FileLocation {
	return &FileLocation{BaseDirectory: BaseDirectory, DLCDirectory: DLCDirectory, ModDirectory: ModDirectory, SaveFile: saveFile}
}

func DefaultFileLocation() *FileLocation {
	return &FileLocation{BaseDirectory: BaseDirectory, DLCDirectory: DLCDirectory, ModDirectory: ModDirectory}
}
