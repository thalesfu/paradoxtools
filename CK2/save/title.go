package save

import "github.com/thalesfu/paradoxtools/utils/pserialize"

type Title struct {
	ID                            string                     `paradox_field:"id" paradox_type:"map_key" json:"id,omitempty"`
	Title                         string                     `paradox_field:"title" paradox_text:"escaped" json:"title,omitempty"`
	Holder                        int                        `paradox_field:"holder" json:"holder,omitempty"`
	Gender                        string                     `paradox_field:"gender" json:"gender,omitempty"`
	Laws                          []string                   `paradox_field:"law" paradox_type:"list" json:"law,omitempty"`
	CoatOfArms                    *CoatOfArms                `paradox_field:"coat_of_arms" json:"coat_of_arms,omitempty"`
	LastChange                    pserialize.Year            `paradox_field:"last_change" json:"last_change,omitempty"`
	Previous                      []int                      `paradox_field:"previous" paradox_type:"field_list" json:"previous,omitempty"`
	Histories                     map[string]*TitleHistory   `paradox_field:"history" json:"history,omitempty"`
	Active                        pserialize.PBool           `paradox_field:"active" json:"active,omitempty"`
	Adjective                     string                     `paradox_field:"adjective" json:"adjective,omitempty"`
	Adventurer                    pserialize.PBool           `paradox_field:"adventurer" json:"adventurer,omitempty"`
	AllowsMatrilinealMarriage     pserialize.PBool           `paradox_field:"allows_matrilineal_marriage" json:"allows_matrilineal_marriage,omitempty"`
	Armies                        []*ArmyID                  `paradox_field:"army" paradox_type:"list" json:"army,omitempty"`
	ArmySizePercentage            float32                    `paradox_field:"army_size_percentage" json:"army_size_percentage,omitempty"`
	AssimilatingLiege             *Title                     `paradox_field:"assimilating_liege" paradox_type:"entity" paradox_default_field:"id" json:"assimilating_liege,omitempty"`
	Dynasty                       int                        `paradox_field:"dynasty" json:"dynasty,omitempty"`
	BaseTitle                     *Title                     `paradox_field:"base_title" paradox_type:"entity" paradox_default_field:"id" json:"base_title,omitempty"`
	CampaignFund                  int                        `paradox_field:"campaign_fund" json:"campaign_fund,omitempty"`
	IsCustom                      pserialize.PBool           `paradox_field:"is_custom" json:"is_custom,omitempty"`
	IsDynamic                     pserialize.PBool           `paradox_field:"is_dynamic" json:"is_dynamic,omitempty"`
	Dynamic                       pserialize.PBool           `paradox_field:"dynamic" json:"dynamic,omitempty"`
	CannotCancelVote              pserialize.PBool           `paradox_field:"cannot_cancel_vote" json:"cannot_cancel_vote,omitempty"`
	Capital                       int                        `paradox_field:"capital" json:"capital,omitempty"`
	CoaDynasty                    int                        `paradox_field:"coa_dynasty" json:"coa_dynasty,omitempty"`
	Color                         []int                      `paradox_field:"color" paradox_type:"field_list" json:"color,omitempty"`
	ConquestCulture               string                     `paradox_field:"conquest_culture" json:"conquest_culture,omitempty"`
	CouncilVoting                 []string                   `paradox_field:"council_voting" paradox_type:"field_list" json:"council_voting,omitempty"`
	CustomGraphics                pserialize.PBool           `paradox_field:"custom_graphics" json:"custom_graphics,omitempty"`
	CustomName                    pserialize.PBool           `paradox_field:"custom_name" json:"custom_name,omitempty"`
	DeJureAssYears                int                        `paradox_field:"de_jure_ass_years" json:"de_jure_ass_years,omitempty"`
	DeJureLiege                   *Title                     `paradox_field:"de_jure_liege" paradox_type:"entity" paradox_default_field:"id" json:"base_titlede_jure_liege,omitempty"`
	DeJureLawChange               pserialize.Year            `paradox_field:"de_jure_law_change" json:"de_jure_law_change,omitempty"`
	DeJureLawChanges              int                        `paradox_field:"de_jure_law_changes" json:"de_jure_law_changes,omitempty"`
	DeJureLawChanger              int                        `paradox_field:"de_jure_law_changer" json:"de_jure_law_changer,omitempty"`
	ET                            string                     `paradox_field:"et" json:"et,omitempty"`
	Nominations                   []*Nomination              `paradox_field:"nomination" paradox_type:"list" json:"nomination,omitempty"`
	Flags                         map[string]pserialize.Year `paradox_field:"flags" json:"flags,omitempty"`
	Foa                           string                     `paradox_field:"foa" paradox_text:"escaped" json:"foa,omitempty"`
	Grant                         pserialize.PBool           `paradox_field:"grant" json:"grant,omitempty"`
	HoldingDynasty                int                        `paradox_field:"holding_dynasty" json:"holding_dynasty,omitempty"`
	Infamy                        *Infamy                    `paradox_field:"infamy" json:"infamy,omitempty"`
	Landless                      pserialize.PBool           `paradox_field:"landless" json:"landless,omitempty"`
	LawChangeTimeout              pserialize.Year            `paradox_field:"law_change_timeout" json:"law_change_timeout,omitempty"`
	LawVotes                      []*LawVote                 `paradox_field:"law_vote" paradox_type:"list" json:"law_vote,omitempty"`
	LawVoteDate                   pserialize.Year            `paradox_field:"law_vote_date" json:"law_vote_date,omitempty"`
	Liege                         *Title                     `paradox_field:"liege" paradox_type:"entity" paradox_default_field:"id" json:"liege,omitempty"`
	MajorRevolt                   pserialize.PBool           `paradox_field:"major_revolt" json:"major_revolt,omitempty"`
	Mercenary                     pserialize.PBool           `paradox_field:"mercenary" json:"mercenary,omitempty"`
	MercenaryType                 *MercenaryType             `paradox_field:"mercenary_type" json:"mercenary_type,omitempty"`
	Name                          string                     `paradox_field:"name" paradox_text:"escaped" json:"name,omitempty"`
	Nomad                         pserialize.PBool           `paradox_field:"nomad" json:"nomad,omitempty"`
	NormalLawChange               pserialize.Year            `paradox_field:"normal_law_change" json:"normal_law_change,omitempty"`
	NormalLawChanger              int                        `paradox_field:"normal_law_changer" json:"normal_law_changer,omitempty"`
	Pentarch                      string                     `paradox_field:"pentarch" json:"pentarch,omitempty"`
	Rebels                        pserialize.PBool           `paradox_field:"rebels" json:"rebels,omitempty"`
	RejectedLawChanges            []*RejectedLawChanges      `paradox_field:"rejected_law_changes" paradox_type:"list" json:"rejected_law_changes,omitempty"`
	ReplaceCaptainOnDeath         pserialize.PBool           `paradox_field:"replace_captain_on_death" json:"replace_captain_on_death,omitempty"`
	SetAllowFreeInfidelRevokation pserialize.PBool           `paradox_field:"set_allow_free_infidel_revokation" json:"set_allow_free_infidel_revokation,omitempty"`
	SetAllowTitleRevokation       pserialize.PBool           `paradox_field:"set_allow_title_revokation" json:"set_allow_title_revokation,omitempty"`
	SetAllowViceRoyalties         int                        `paradox_field:"set_allow_vice_royalties" json:"set_allow_vice_royalties,omitempty"`
	SetCoa                        int                        `paradox_field:"set_coa" json:"set_coa,omitempty"`
	SetInvestiture                string                     `paradox_field:"set_investiture" json:"set_investiture,omitempty"`
	SetProtectedInheritance       pserialize.PBool           `paradox_field:"set_protected_inheritance" json:"set_protected_inheritance,omitempty"`
	Settlement                    *Settlement                `paradox_field:"settlement" json:"settlement,omitempty"`
	SiphonsIncomeToCreator        float32                    `paradox_field:"siphons_income_to_creator" json:"siphons_income_to_creator,omitempty"`
	SuccLawChanger                int                        `paradox_field:"succ_law_changer" json:"succ_law_changer,omitempty"`
	Succession                    string                     `paradox_field:"succession" json:"succession,omitempty"`
	SuccessionElectors            []int                      `paradox_field:"succession_electors" paradox_type:"field_list" json:"succession_electors,omitempty"`
	Temporary                     pserialize.PBool           `paradox_field:"temporary" json:"temporary,omitempty"`
	TitleFemale                   string                     `paradox_field:"title_female" paradox_text:"escaped" json:"title_female,omitempty"`
	UsurpDate                     pserialize.Year            `paradox_field:"usurp_date" json:"usurp_date,omitempty"`
	ViceRoyalty                   pserialize.PBool           `paradox_field:"vice_royalty" json:"vice_royalty,omitempty"`
	ViceRoyaltyRevokation         pserialize.PBool           `paradox_field:"vice_royalty_revokation" json:"vice_royalty_revokation,omitempty"`
}

type TitleHistory struct {
	Holder *TitleHistoryHolder `paradox_field:"holder" paradox_type:"entity" paradox_default_field:"who" json:"holder,omitempty"`
}

type TitleHistoryHolder struct {
	Who  string `paradox_field:"who" json:"who,omitempty"`
	Type string `paradox_field:"type" json:"type,omitempty"`
}
