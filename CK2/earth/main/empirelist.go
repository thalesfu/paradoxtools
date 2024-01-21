package main

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/ck2utils"
	"github.com/thalesfu/paradoxtools/CK2/config"
	"github.com/thalesfu/paradoxtools/CK2/feud"
	"github.com/thalesfu/paradoxtools/CK2/history/province"
	"github.com/thalesfu/paradoxtools/CK2/translations"
	"github.com/thalesfu/paradoxtools/segments"
)

var AllEmpires = map[string]feud.Empire{}

func init() {
	segments := segments.LoadSegments(config.LandedTitleFile)

	for _, segment := range segments {
		tryAppendEmpire(segment)
	}
}

func tryAppendEmpire(segment *segments.Segment) {
	if ck2utils.IsEmpireSegment(segment) {
		empire := generateEmpire(segment)
		AllEmpires[empire.GetTitleCode()] = empire
	}
}

func generateEmpire(segment *segments.Segment) feud.Empire {
	empire := &feud.BaseEmpire{}
	empire.TitleCode = segment.Name
	empire.Title = ck2utils.GetFeudName(segment.Name)
	empire.TitleName = translations.GetFeudName(segment.Name)
	for _, child := range segment.Subs {
		tryAppendKingdom(empire, child)
	}

	return empire
}

func tryAppendKingdom(e *feud.BaseEmpire, segment *segments.Segment) {
	if ck2utils.IsKingdomSegment(segment) {
		if e.Kingdoms == nil {
			e.Kingdoms = make(map[string]feud.Kingdom)
		}
		e.Kingdoms[segment.Name] = generateKingdom(e, segment)
	}
}

func generateKingdom(e *feud.BaseEmpire, segment *segments.Segment) feud.Kingdom {
	kingdom := &feud.BaseKingdom{}
	kingdom.TitleCode = segment.Name
	kingdom.Title = ck2utils.GetFeudName(segment.Name)
	kingdom.TitleName = translations.GetFeudName(segment.Name)
	kingdom.Parent = e
	for _, child := range segment.Subs {
		tryAppendDuke(kingdom, child)
	}

	return kingdom
}

func tryAppendDuke(k *feud.BaseKingdom, segment *segments.Segment) {
	if ck2utils.IsDukeSegment(segment) {
		if k.Dukes == nil {
			k.Dukes = make(map[string]feud.Duke)
		}
		k.Dukes[segment.Name] = generateDuke(k, segment)
	}
}

func generateDuke(k *feud.BaseKingdom, segment *segments.Segment) feud.Duke {
	duke := &feud.BaseDuke{}
	duke.TitleCode = segment.Name
	duke.Title = ck2utils.GetFeudName(segment.Name)
	duke.TitleName = translations.GetFeudName(segment.Name)
	duke.Parent = k
	for _, child := range segment.Subs {
		tryAppendCounty(duke, child)
	}

	return duke
}

func tryAppendCounty(d *feud.BaseDuke, segment *segments.Segment) {
	if ck2utils.IsCountySegment(segment) {
		if d.Counties == nil {
			d.Counties = make(map[string]feud.County)
		}
		d.Counties[segment.Name] = generateCounty(d, segment)
	}
}

func generateCounty(d *feud.BaseDuke, segment *segments.Segment) feud.County {
	county := &feud.BaseCounty{}
	county.TitleCode = segment.Name
	county.Title = ck2utils.GetFeudName(segment.Name)
	county.TitleName = translations.GetFeudName(county.TitleCode)
	county.Parent = d
	if p, ok := province.ProvinceCodeList[county.TitleCode]; ok {
		if p.GetID() == "" {
			fmt.Println("Province ID is empty:", county.TitleCode)
		}
		county.ID = p.GetID()
	} else {
		fmt.Println("Province not found:", county.TitleCode)
	}
	for _, child := range segment.Subs {
		tryAppendBarony(county, child)
	}

	return county
}

func tryAppendBarony(c *feud.BaseCounty, segment *segments.Segment) {
	if ck2utils.IsBaronySegment(segment) {
		if c.Baronies == nil {
			c.Baronies = make(map[string]feud.Barony)
		}
		c.Baronies[segment.Name] = generateBarony(c, segment)
	}
}

func generateBarony(c *feud.BaseCounty, segment *segments.Segment) feud.Barony {
	barony := &feud.BaseBarony{}
	barony.TitleCode = segment.Name
	barony.Title = ck2utils.GetFeudName(segment.Name)
	barony.TitleName = translations.GetFeudName(segment.Name)
	barony.Parent = c
	return barony
}
