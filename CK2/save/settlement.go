package save

import "github.com/thalesfu/paradoxtools/utils/pserialize"

type Settlement struct {
	Type         string           `paradox_field:"type" json:"type,omitempty"`
	BuildTime    pserialize.Year  `paradox_field:"build_time" json:"build_time,omitempty"`
	FpMansion1   pserialize.PBool `paradox_field:"fp_mansion_1" paradox_type:"field" json:"fp_mansion_1,omitempty"`
	FpMansion2   pserialize.PBool `paradox_field:"fp_mansion_2" paradox_type:"field" json:"fp_mansion_2,omitempty"`
	FpMansion3   pserialize.PBool `paradox_field:"fp_mansion_3" paradox_type:"field" json:"fp_mansion_3,omitempty"`
	FpMansion4   pserialize.PBool `paradox_field:"fp_mansion_4" paradox_type:"field" json:"fp_mansion_4,omitempty"`
	FpBarracks1  pserialize.PBool `paradox_field:"fp_barracks_1" paradox_type:"field" json:"fp_barracks_1,omitempty"`
	FpBarracks2  pserialize.PBool `paradox_field:"fp_barracks_2" paradox_type:"field" json:"fp_barracks_2,omitempty"`
	FpBarracks3  pserialize.PBool `paradox_field:"fp_barracks_3" paradox_type:"field" json:"fp_barracks_3,omitempty"`
	FpStable1    pserialize.PBool `paradox_field:"fp_stable_1" paradox_type:"field" json:"fp_stable_1,omitempty"`
	FpStable2    pserialize.PBool `paradox_field:"fp_stable_2" paradox_type:"field" json:"fp_stable_2,omitempty"`
	FpStable3    pserialize.PBool `paradox_field:"fp_stable_3" paradox_type:"field" json:"fp_stable_3,omitempty"`
	FpBowyer1    pserialize.PBool `paradox_field:"fp_bowyer_1" paradox_type:"field" json:"fp_bowyer_1,omitempty"`
	FpBowyer2    pserialize.PBool `paradox_field:"fp_bowyer_2" paradox_type:"field" json:"fp_bowyer_2,omitempty"`
	FpBowyer3    pserialize.PBool `paradox_field:"fp_bowyer_3" paradox_type:"field" json:"fp_bowyer_3,omitempty"`
	FpShipyard1  pserialize.PBool `paradox_field:"fp_shipyard_1" paradox_type:"field" json:"fp_shipyard_1,omitempty"`
	FpShipyard2  pserialize.PBool `paradox_field:"fp_shipyard_2" paradox_type:"field" json:"fp_shipyard_2,omitempty"`
	FpShipyard3  pserialize.PBool `paradox_field:"fp_shipyard_3" paradox_type:"field" json:"fp_shipyard_3,omitempty"`
	FpWarehouse1 pserialize.PBool `paradox_field:"fp_warehouse_1" paradox_type:"field" json:"fp_warehouse_1,omitempty"`
	FpWarehouse2 pserialize.PBool `paradox_field:"fp_warehouse_2" paradox_type:"field" json:"fp_warehouse_2,omitempty"`
	FpWarehouse3 pserialize.PBool `paradox_field:"fp_warehouse_3" paradox_type:"field" json:"fp_warehouse_3,omitempty"`
	FpGarden1    pserialize.PBool `paradox_field:"fp_garden_1" paradox_type:"field" json:"fp_garden_1,omitempty"`
	FpGarden2    pserialize.PBool `paradox_field:"fp_garden_2" paradox_type:"field" json:"fp_garden_2,omitempty"`
	FpGarden3    pserialize.PBool `paradox_field:"fp_garden_3" paradox_type:"field" json:"fp_garden_3,omitempty"`
	FpCellar1    pserialize.PBool `paradox_field:"fp_cellar_1" paradox_type:"field" json:"fp_cellar_1,omitempty"`
	FpCellar2    pserialize.PBool `paradox_field:"fp_cellar_2" paradox_type:"field" json:"fp_cellar_2,omitempty"`
	FpCellar3    pserialize.PBool `paradox_field:"fp_cellar_3" paradox_type:"field" json:"fp_cellar_3,omitempty"`
	FpShrine1    pserialize.PBool `paradox_field:"fp_shrine_1" paradox_type:"field" json:"fp_shrine_1,omitempty"`
	FpShrine2    pserialize.PBool `paradox_field:"fp_shrine_2" paradox_type:"field" json:"fp_shrine_2,omitempty"`
	FpShrine3    pserialize.PBool `paradox_field:"fp_shrine_3" paradox_type:"field" json:"fp_shrine_3,omitempty"`
	FpVault1     pserialize.PBool `paradox_field:"fp_vault_1" paradox_type:"field" json:"fp_vault_1,omitempty"`
	FpVault2     pserialize.PBool `paradox_field:"fp_vault_2" paradox_type:"field" json:"fp_vault_2,omitempty"`
	FpVault3     pserialize.PBool `paradox_field:"fp_vault_3" paradox_type:"field" json:"fp_vault_3,omitempty"`
	Levy         *ArmyDetail      `paradox_field:"levy" json:"levy,omitempty"`
}

type SettlementConstruction struct {
	StartDate pserialize.Year `paradox_field:"start_date" json:"start_date,omitempty"`
	Date      pserialize.Year `paradox_field:"date" json:"date,omitempty"`
	Progress  float32         `paradox_field:"progress" json:"progress,omitempty"`
	Days      int             `paradox_field:"days" json:"days,omitempty"`
	Location  int             `paradox_field:"location" json:"location,omitempty"`
	Province  int             `paradox_field:"province" json:"province,omitempty"`
	Character int             `paradox_field:"character" json:"character,omitempty"`
	Type      string          `paradox_field:"type" json:"type,omitempty"`
}
