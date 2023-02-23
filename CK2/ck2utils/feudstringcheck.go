package ck2utils

import "regexp"

var rxEmpire = regexp.MustCompile(`^(e_)([a-z]+)$`)
var rxKingdom = regexp.MustCompile(`^(k_)([a-z]+)$`)
var rxDuke = regexp.MustCompile(`^(d_)([a-z]+)$`)
var rxCount = regexp.MustCompile(`^(c_)([a-z]+)$`)
var rxBaron = regexp.MustCompile(`^(b_)([a-z]+)$`)
var rxFeud = regexp.MustCompile(`^([ekdcb]_)([a-z]+)$`)
var rxEmpireAdj = regexp.MustCompile(`^(e_)([a-z]+)_adj$`)
var rxKingdomAdj = regexp.MustCompile(`^(k_)([a-z]+)_adj$`)
var rxDukeAdj = regexp.MustCompile(`^(d_)([a-z]+)_adj$`)
var rxCountAdj = regexp.MustCompile(`^(c_)([a-z]+)_adj$`)
var rxBaronAdj = regexp.MustCompile(`^(b_)([a-z]+)_adj$`)
var rxFeudAdj = regexp.MustCompile(`^([ekdcb]_)([a-z]+)_adj$`)

func IsEmpireString(s string) bool {
	return rxEmpire.MatchString(s)
}

func IsKingdomString(s string) bool {
	return rxKingdom.MatchString(s)
}

func IsDukeString(s string) bool {
	return rxDuke.MatchString(s)
}

func IsCountyString(s string) bool {
	return rxCount.MatchString(s)
}

func IsBaronyString(s string) bool {
	return rxBaron.MatchString(s)
}

func IsFeudString(s string) bool {
	return rxFeud.MatchString(s)
}

func IsEmpireAdjString(s string) bool {
	return rxEmpireAdj.MatchString(s)
}

func IsKingdomAdjString(s string) bool {
	return rxKingdomAdj.MatchString(s)
}

func IsDukeAdjString(s string) bool {
	return rxDukeAdj.MatchString(s)
}

func IsCountyAdjString(s string) bool {
	return rxCountAdj.MatchString(s)
}

func IsBaronyAdjString(s string) bool {
	return rxBaronAdj.MatchString(s)
}

func IsFeudAdjString(s string) bool {
	return rxFeudAdj.MatchString(s)
}

func GetFeudName(s string) string {
	matchString := rxFeud.MatchString(s)
	if matchString {
		return rxFeud.FindStringSubmatch(s)[2]
	}

	return ""
}
