package config

import "path"

const CKIIPath = "/Volumes/[C] Windows 11.hidden/Program Files (x86)/Steam/steamapps/common/Crusader Kings II"

var CK2ProvinceFold = path.Join(CKIIPath, "history/provinces")

var LandedTitleFile = path.Join(CKIIPath, "common/landed_titles/landed_titles.txt")
