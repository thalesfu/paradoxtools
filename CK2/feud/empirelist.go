package feud

import (
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/ck2utils"
	"github.com/thalesfu/paradoxtools/CK2/config"
	"github.com/thalesfu/paradoxtools/CK2/history/province"
	"github.com/thalesfu/paradoxtools/CK2/translations"
	"github.com/thalesfu/paradoxtools/segments"
)

var AllEmpires = map[string]Empire{}

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

func generateEmpire(segment *segments.Segment) Empire {
	empire := &BaseEmpire{}
	empire.TitleCode = segment.Name
	empire.Title = ck2utils.GetFeudName(segment.Name)
	empire.TitleName = translations.GetFeudName(segment.Name)
	for _, child := range segment.Subs {
		tryAppendKingdom(empire, child)
	}

	return empire
}

func tryAppendKingdom(e *BaseEmpire, segment *segments.Segment) {
	if ck2utils.IsKingdomSegment(segment) {
		if e.Kingdoms == nil {
			e.Kingdoms = make(map[string]Kingdom)
		}
		e.Kingdoms[segment.Name] = generateKingdom(e, segment)
	}
}

func generateKingdom(e *BaseEmpire, segment *segments.Segment) Kingdom {
	kingdom := &BaseKingdom{}
	kingdom.TitleCode = segment.Name
	kingdom.Title = ck2utils.GetFeudName(segment.Name)
	kingdom.TitleName = translations.GetFeudName(segment.Name)
	kingdom.Parent = e
	for _, child := range segment.Subs {
		tryAppendDuke(kingdom, child)
	}

	return kingdom
}

func tryAppendDuke(k *BaseKingdom, segment *segments.Segment) {
	if ck2utils.IsDukeSegment(segment) {
		if k.Dukes == nil {
			k.Dukes = make(map[string]Duke)
		}
		k.Dukes[segment.Name] = generateDuke(k, segment)
	}
}

func generateDuke(k *BaseKingdom, segment *segments.Segment) Duke {
	duke := &BaseDuke{}
	duke.TitleCode = segment.Name
	duke.Title = ck2utils.GetFeudName(segment.Name)
	duke.TitleName = translations.GetFeudName(segment.Name)
	duke.Parent = k
	for _, child := range segment.Subs {
		tryAppendCounty(duke, child)
	}

	return duke
}

func tryAppendCounty(d *BaseDuke, segment *segments.Segment) {
	if ck2utils.IsCountySegment(segment) {
		if d.Counties == nil {
			d.Counties = make(map[string]County)
		}
		d.Counties[segment.Name] = generateCounty(d, segment)
	}
}

func generateCounty(d *BaseDuke, segment *segments.Segment) County {
	county := &BaseCounty{}
	county.TitleCode = segment.Name
	county.Title = ck2utils.GetFeudName(segment.Name)
	county.TitleName = translations.GetFeudName(county.TitleCode)
	county.Parent = d
	if p, ok := province.ProvinceList[county.TitleCode]; ok {
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

func tryAppendBarony(c *BaseCounty, segment *segments.Segment) {
	if ck2utils.IsBaronySegment(segment) {
		if c.Baronies == nil {
			c.Baronies = make(map[string]Barony)
		}
		c.Baronies[segment.Name] = generateBarony(c, segment)
	}
}

func generateBarony(c *BaseCounty, segment *segments.Segment) Barony {
	barony := &BaseBarony{}
	barony.TitleCode = segment.Name
	barony.Title = ck2utils.GetFeudName(segment.Name)
	barony.TitleName = translations.GetFeudName(segment.Name)
	barony.Parent = c
	return barony
}
