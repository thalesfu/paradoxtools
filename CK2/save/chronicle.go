package save

import "github.com/thalesfu/paradoxtools/utils/pserialize"

type ChronicleEntry struct {
	Text            string `paradox_field:"text" paradox_text:"escaped" json:"text,omitempty"`
	Picture         string `paradox_field:"picture" json:"picture,omitempty"`
	Portrait        int    `paradox_field:"portrait" json:"portrait,omitempty"`
	PortraitCulture string `paradox_field:"portrait_culture" json:"portrait_culture,omitempty"`
	Type            string `paradox_field:"type" json:"type,omitempty"`
}

type ChronicleChapter struct {
	Chronicles []*ChronicleEntry `paradox_field:"chronicle_entry" paradox_type:"list" json:"chronicle_entry,omitempty"`
	Year       int               `paradox_field:"year" json:"year,omitempty"`
}

type Chronicle struct {
	ChronicleChapters []*ChronicleChapter `paradox_field:"chronicle_chapter" paradox_type:"list" json:"chronicle_chapter,omitempty"`
	Character         int                 `paradox_field:"character" json:"character,omitempty"`
}

type ChronicleCollection struct {
	Chronicle         *Chronicle       `paradox_field:"chronicle" json:"chronicle,omitempty"`
	ChroniclePosition int              `paradox_field:"chronicle_position" json:"chronicle_position,omitempty"`
	ChapterPosition   int              `paradox_field:"chapter_position" json:"chapter_position,omitempty"`
	EntryPosition     int              `paradox_field:"entry_position" json:"entry_position,omitempty"`
	ChronicleIconLit  pserialize.PBool `paradox_field:"chronicle_icon_lit" json:"chronicle_icon_lit,omitempty"`
}

type CharacterPlayerData struct {
	ChronicleCollection         *ChronicleCollection `paradox_field:"chronicle_collection" json:"chronicle_collection,omitempty"`
	SocietyShowInterestCooldown pserialize.Year      `paradox_field:"society_show_interest_cooldown" json:"society_show_interest_cooldown,omitempty"`
	JoinSocietyCooldown         pserialize.Year      `paradox_field:"join_society_cooldown" json:"join_society_cooldown,omitempty"`
	Telws                       int                  `paradox_field:"telws" json:"telws,omitempty"`
	Telbc                       int                  `paradox_field:"telbc" json:"telbc,omitempty"`
	Telde                       int                  `paradox_field:"telde" json:"telde,omitempty"`
	Telld                       pserialize.PBool     `paradox_field:"telld" json:"telld,omitempty"`
	Telsc                       pserialize.PBool     `paradox_field:"telsc" json:"telsc,omitempty"`
}
