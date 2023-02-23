package segments

import (
	"bufio"
	"fmt"
	"github.com/thalesfu/paradoxtools"
	"io"
	"os"
	"strings"
)

const (
	SegmentStatusStart = iota
	SegmentStatusWaitingStart
	SegmentStatusWaitingEnd
	SegmentStatusStartEnd
	SegmentStatusEnd
)

type SegmentType struct {
	Name       string
	IsStartEnd func(byte) bool
	IsEnded    func(byte) bool
}

var SegmentTypeVariable = &SegmentType{Name: "Variable", IsStartEnd: paradoxtools.IsWhitespace, IsEnded: paradoxtools.IsNewlineEnd}
var SegmentTypeString = &SegmentType{Name: "String", IsStartEnd: paradoxtools.IsQuote, IsEnded: paradoxtools.IsNewlineEnd}
var SegmentTypeObject = &SegmentType{Name: "Object", IsStartEnd: paradoxtools.IsRightParentheses, IsEnded: paradoxtools.IsNewlineEnd}
var SegmentTypeComment = &SegmentType{Name: "Comment", IsStartEnd: paradoxtools.IsSharp, IsEnded: paradoxtools.IsNewlineEnd}

type Segment struct {
	Data       []byte
	Words      strings.Builder
	Content    strings.Builder
	Name       string
	Value      string
	StartIndex int64
	EndIndex   int64
	Subs       []*Segment
	Parent     *Segment
	Status     int
	LastChar   string
	Type       *SegmentType
}

func (s *Segment) CreateSub(index int64) *Segment {
	sub := &Segment{Parent: s, StartIndex: index}
	s.Subs = append(s.Subs, sub)
	return sub
}

func (s *Segment) CreateNext() *Segment {
	return &Segment{}
}

func (s *Segment) Read(byte byte, index int64) *Segment {
	s.JustRead(byte)
	if s.Parent != nil {
		s.Parent.JustRead(byte)
	}

	return s.SetStatus(byte, index)
}

func (s *Segment) JustRead(byte byte) {
	s.Data = append(s.Data, byte)
	char := string(byte)
	s.Words.WriteString(char)
	if !paradoxtools.IsWhitespace(byte) && !paradoxtools.IsLeftParentheses(byte) && !paradoxtools.IsRightParentheses(byte) {
		s.Content.WriteString(char)
	}
	s.LastChar = char
}

func (s *Segment) SetStatus(b byte, index int64) *Segment {
	if s.Type == SegmentTypeComment {
		if paradoxtools.IsNewlineEnd(b) {

			s.Status = SegmentStatusEnd

			fmtprint(s)

			se := s

			for se.Parent != nil && se.Parent.Status == SegmentStatusStartEnd {
				se.Parent.Status = SegmentStatusEnd
				fmtprint(se)
				se = se.Parent
			}

			if se.Parent != nil {
				return se.Parent
			}

			return se
		}

		return s
	}

	if paradoxtools.IsSharp(b) {
		if s.Name == "" && s.Content.String() == "#" {
			s.Type = SegmentTypeComment
			return s
		}

		sub := s.CreateSub(index)
		sub.Type = SegmentTypeComment
		sub.Read(b, index)
		return sub
	}

	if s.Status == SegmentStatusStart {
		if s.Type == nil {
			if s.Parent != nil && s.Parent.Type.IsStartEnd(b) {
				return s.Parent.SetStatus(b, index)
			}
		}

		if paradoxtools.IsEqual(b) {
			s.Status = SegmentStatusWaitingStart
			s.Name = strings.TrimRight(s.Content.String(), "=")
			s.Content.Reset()
			return s
		}
	}

	if s.Status == SegmentStatusWaitingStart {
		if paradoxtools.IsLeftParentheses(b) {
			s.Status = SegmentStatusWaitingEnd
			s.Type = SegmentTypeObject
			return s
		}

		if paradoxtools.IsAlphanumeric(b) {
			s.Status = SegmentStatusWaitingEnd
			s.Type = SegmentTypeVariable
			return s
		}

		if paradoxtools.IsQuote(b) {
			s.Status = SegmentStatusWaitingEnd
			s.Type = SegmentTypeString
			return s
		}

		if paradoxtools.IsSharp(b) {
			s.Status = SegmentStatusWaitingEnd
			s.Type = SegmentTypeComment
			return s
		}

		return s
	}

	if s.Status == SegmentStatusWaitingEnd {

		if s.Type == SegmentTypeObject && paradoxtools.IsAlphanumeric(b) {
			sub := s.CreateSub(index)
			sub.JustRead(b)
			sub.SetStatus(b, index)
			return sub
		}

		if s.Type.IsStartEnd(b) {
			s.Status = SegmentStatusStartEnd
			s.Value = s.Content.String()
		}

		return s
	}

	if s.Status == SegmentStatusStartEnd {
		if s.Type.IsEnded(b) {
			s.Status = SegmentStatusEnd
			fmtprint(s)
			s.EndIndex = index

			se := s

			for se.Parent != nil && se.Parent.Status == SegmentStatusStartEnd && se.Parent.Type.IsEnded(b) {
				se.Parent.Status = SegmentStatusEnd
				fmtprint(se)
				se = se.Parent
			}

			if se.Parent != nil {
				return se.Parent
			}

			return se
		}

		if s.Parent != nil && s.Parent.Type.IsStartEnd(b) {
			s.Status = SegmentStatusEnd
			fmtprint(s)
			return s.Parent.SetStatus(b, index)
		}

	}

	return s
}

func fmtprint(s *Segment) {
	se := s
	msg := fmt.Sprintf("%s [%s]", se.Name, se.Type.Name)
	for se.Parent != nil {
		msg = fmt.Sprintf("%s.%s", se.Parent.Name, msg)
		se = se.Parent
	}
	fmt.Sprintln(msg)
}

func LoadSegments(path string) []*Segment {
	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var index int64 = -1
	var segments []*Segment
	currentSegment := &Segment{}

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		index++

		currentSegment = currentSegment.Read(b, index)

		if currentSegment.Status == SegmentStatusEnd && currentSegment.Parent == nil {
			segments = append(segments, currentSegment)
			currentSegment = &Segment{}
		}
	}

	return segments
}
