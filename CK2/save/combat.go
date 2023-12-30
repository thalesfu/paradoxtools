package save

type Combat struct {
	SiegeCombat []*CombatDetail `paradox_field:"siege_combat" paradox_type:"list" json:"siege_combat,omitempty"`
	LandCombat  []*CombatDetail `paradox_field:"land_combat" paradox_type:"list" json:"land_combat,omitempty"`
}

type CombatDetail struct {
	Attackers   []int           `paradox_field:"attackers" paradox_type:"list" json:"attackers,omitempty"`
	Defenders   []int           `paradox_field:"defenders" paradox_type:"list" json:"defenders,omitempty"`
	Location    int             `paradox_field:"location" json:"location,omitempty"`
	Day         int             `paradox_field:"day" json:"day,omitempty"`
	Attacker    *CombatAttacker `paradox_field:"attacker" json:"attacker,omitempty"`
	Defender    *CombatDefender `paradox_field:"defender" json:"defender,omitempty"`
	Event       int             `paradox_field:"event" json:"event,omitempty"`
	Adjacencies int             `paradox_field:"adjacencies" json:"adjacencies,omitempty"`
	Terrain     string          `paradox_field:"terrain" json:"terrain,omitempty"`
}

type Unit struct {
	ID   int `paradox_field:"id" json:"id,omitempty"`
	Type int `paradox_field:"type" json:"type,omitempty"`
}

type CombatDirection struct {
	Leader     int         `paradox_field:"leader" json:"leader,omitempty"`
	LastLeader int         `paradox_field:"last_leader" json:"last_leader,omitempty"`
	Losses     *ArmyDetail `paradox_field:"losses" json:"losses,omitempty"`
	Target     int         `paradox_field:"target" json:"target,omitempty"`
	SubUnit    []int       `paradox_field:"sub_unit" paradox_type:"field_list" json:"sub_unit,omitempty"`
	Tactic     int         `paradox_field:"tactic" json:"tactic,omitempty"`
	TacticDay  int         `paradox_field:"tactic_day" json:"tactic_day,omitempty"`
	Phase      int         `paradox_field:"phase" json:"phase,omitempty"`
}

type CombatAttacker struct {
	Unit        []*Unit          `paradox_field:"unit" paradox_type:"list" json:"unit,omitempty"`
	FlankLeft   *CombatDirection `paradox_field:"flank_left" json:"flank_left,omitempty"`
	FlankCenter *CombatDirection `paradox_field:"flank_center" json:"flank_center,omitempty"`
	FlankRight  *CombatDirection `paradox_field:"flank_right" json:"flank_right,omitempty"`
}

type CombatDefender struct {
	FlankLeft   *CombatDirection `paradox_field:"flank_left" json:"flank_left,omitempty"`
	FlankCenter *CombatDirection `paradox_field:"flank_center" json:"flank_center,omitempty"`
	FlankRight  *CombatDirection `paradox_field:"flank_right" json:"flank_right,omitempty"`
}
