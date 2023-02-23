package ck2utils

import (
	"github.com/thalesfu/paradoxtools/segments"
)

func IsEmpireSegment(s *segments.Segment) bool {
	if s.Parent != nil {
		return false
	}

	return IsEmpireString(s.Name)
}

func IsKingdomSegment(s *segments.Segment) bool {
	if s.Parent != nil && !IsEmpireSegment(s.Parent) {
		return false
	}

	return IsKingdomString(s.Name)
}

func IsDukeSegment(s *segments.Segment) bool {
	if s.Parent != nil && !IsKingdomSegment(s.Parent) {
		return false
	}

	return IsDukeString(s.Name)
}

func IsCountySegment(s *segments.Segment) bool {
	if s.Parent != nil && !IsDukeSegment(s.Parent) {
		return false
	}

	return IsCountyString(s.Name)
}

func IsBaronySegment(s *segments.Segment) bool {

	if s.Parent != nil && !IsCountySegment(s.Parent) {
		return false
	}

	return IsBaronyString(s.Name)
}

func IsFeudSegment(s *segments.Segment) bool {
	return IsEmpireSegment(s) || IsKingdomSegment(s) || IsDukeSegment(s) || IsCountySegment(s) || IsBaronySegment(s)
}
