package save

type Nick struct {
	Name     string `paradox_field:"name" paradox_text:"escaped" json:"name,omitempty"`
	NickName string `paradox_field:"nickname" json:"nickname,omitempty"`
}
