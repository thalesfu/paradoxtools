package pserialize

import (
	"bufio"
	"encoding"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/thalesfu/golangutils"
	"github.com/thalesfu/paradoxtools/utils"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func UnmarshalP[T any](content string) (t *T, ok bool) {
	firstLine, firstLinePosition := getFirstLine(content)

	if needRemove(firstLine) {
		content = content[firstLinePosition:]

		last := strings.LastIndex(content, "}")
		content = "{\n" + content[:last+1]
	} else {
		content = "{\n" + content + "\n}"
	}

	logPosition := 0
	var writer *bufio.Writer
	if logPosition > 0 {
		line := 1
		logWriter, deferFunc := golangutils.CreateLogWriter("punmarshal")
		writer = logWriter
		defer deferFunc()

		for i, c := range []byte(content) {
			writer.WriteByte(c)
			if c == '\n' {
				line++
				if err := writer.Flush(); err != nil {
					log.Fatal(err)
				}
			}

			if i == logPosition {
				endMessage := "\t<-- " + "at " + strconv.Itoa(i) + " line:" + strconv.Itoa(line) + " "
				_, err := writer.WriteString(endMessage)
				if err != nil {
					log.Fatal(err)
				}
				if err := writer.Flush(); err != nil {
					log.Fatal(err)
				}
				break
			}
		}
	}

	err := Unmarshal([]byte(content), &t)
	if err != nil {
		switch e := err.(type) {
		case *UnmarshalTypeError:
			log.Printf("解析%T类型P设文件失败 at %d error:\n%s", t, e.Offset, e.Error())
			return nil, false
		default:
			log.Printf("解析%T类型P设文件失败 error:\n%s", t, e.Error())
			return nil, false
		}

	}

	return t, true
}

func getFirstLine(content string) (string, int) {
	firstLineBytes := make([]byte, 0)
	cb := []byte(content)
	matchN := false
	matchT := false
	postion := 0

	for i, b := range cb {
		postion = i
		if b == '\n' {
			matchN = true
			continue
		} else if b == '\t' {
			matchT = true
			continue
		}

		if matchN && matchT {
			break
		}

		firstLineBytes = append(firstLineBytes, b)
	}

	return string(firstLineBytes), postion
}

func needRemove(line string) bool {
	if strings.HasPrefix(line, "CK2txt") {
		return true
	}

	return false
}

// Unmarshal parses the JSON-encoded data and stores the result
// in the value pointed to by v. If v is nil or not a pointer,
// Unmarshal returns an InvalidUnmarshalError.
//
// Unmarshal uses the inverse of the encodings that
// Marshal uses, allocating maps, slices, and pointers as necessary,
// with the following additional rules:
//
// To unmarshal JSON into a pointer, Unmarshal first handles the case of
// the JSON being the JSON literal null. In that case, Unmarshal sets
// the pointer to nil. Otherwise, Unmarshal unmarshals the JSON into
// the value pointed at by the pointer. If the pointer is nil, Unmarshal
// allocates a new value for it to point to.
//
// To unmarshal JSON into a value implementing the Unmarshaler interface,
// Unmarshal calls that value's UnmarshalJSON method, including
// when the input is a JSON null.
// Otherwise, if the value implements encoding.TextUnmarshaler
// and the input is a JSON quoted string, Unmarshal calls that value's
// UnmarshalText method with the unquoted form of the string.
//
// To unmarshal JSON into a struct, Unmarshal matches incoming object
// keys to the keys used by Marshal (either the struct field name or its tag),
// preferring an exact match but also accepting a case-insensitive match. By
// default, object keys which don't have a corresponding struct field are
// ignored (see Decoder.DisallowUnknownFields for an alternative).
//
// To unmarshal JSON into an interface value,
// Unmarshal stores one of these in the interface value:
//
//	bool, for JSON booleans
//	float64, for JSON numbers
//	string, for JSON strings
//	[]interface{}, for JSON arrays
//	map[string]interface{}, for JSON objects
//	nil for JSON null
//
// To unmarshal a JSON list into a slice, Unmarshal resets the slice length
// to zero and then appends each element to the slice.
// As a special case, to unmarshal an empty JSON list into a slice,
// Unmarshal replaces the slice with a new empty slice.
//
// To unmarshal a JSON list into a Go list, Unmarshal decodes
// JSON list elements into corresponding Go list elements.
// If the Go list is smaller than the JSON list,
// the additional JSON list elements are discarded.
// If the JSON list is smaller than the Go list,
// the additional Go list elements are set to zero values.
//
// To unmarshal a JSON object into a map, Unmarshal first establishes a map to
// use. If the map is nil, Unmarshal allocates a new map. Otherwise Unmarshal
// reuses the existing map, keeping existing entries. Unmarshal then stores
// key-value pairs from the JSON object into the map. The map's key type must
// either be any string type, an integer, implement json.Unmarshaler, or
// implement encoding.TextUnmarshaler.
//
// If the JSON-encoded data contain a syntax error, Unmarshal returns a SyntaxError.
//
// If a JSON value is not appropriate for a given target type,
// or if a JSON number overflows the target type, Unmarshal
// skips that field and completes the unmarshaling as best it can.
// If no more serious errors are encountered, Unmarshal returns
// an UnmarshalTypeError describing the earliest such error. In any
// case, it's not guaranteed that all the remaining fields following
// the problematic one will be unmarshaled into the target object.
//
// The JSON null value unmarshals into an interface, map, pointer, or slice
// by setting that Go value to nil. Because null is often used in JSON to mean
// “not present,” unmarshaling a JSON null into any other Go type has no effect
// on the value and produces no error.
//
// When unmarshaling quoted strings, invalid UTF-8 or
// invalid UTF-16 surrogate pairs are not treated as an error.
// Instead, they are replaced by the Unicode replacement
// character U+FFFD.
func Unmarshal(data []byte, v any) error {
	// Check for well-formedness.
	// Avoids filling out half a data structure
	// before discovering a JSON syntax error.
	var d decodeState
	//err := checkValid(data, &d.scan)
	//if err != nil {
	//	return err
	//}

	d.init(data)
	return d.unmarshal(v)
}

// Unmarshaler is the interface implemented by types
// that can unmarshal a JSON description of themselves.
// The input can be assumed to be a valid encoding of
// a JSON value. UnmarshalJSON must copy the JSON data
// if it wishes to retain the data after returning.
//
// By convention, to approximate the behavior of Unmarshal itself,
// Unmarshalers implement UnmarshalJSON([]byte("null")) as a no-op.
type Unmarshaler interface {
	UnmarshalP([]byte) error
}

// An UnmarshalTypeError describes a JSON value that was
// not appropriate for a value of a specific Go type.
type UnmarshalTypeError struct {
	Value  string       // description of JSON value - "bool", "list", "number -5"
	Type   reflect.Type // type of Go value it could not be assigned to
	Offset int64        // error occurred after reading Offset bytes
	Struct string       // name of the struct type containing the field
	Field  string       // the full path from root node to the field
}

func (e *UnmarshalTypeError) Error() string {
	if e.Struct != "" || e.Field != "" {
		return "json: cannot unmarshal " + e.Value + " into Go struct field " + e.Struct + "." + e.Field + " of type " + e.Type.String()
	}
	return "json: cannot unmarshal " + e.Value + " into Go value of type " + e.Type.String()
}

// An InvalidUnmarshalError describes an invalid argument passed to Unmarshal.
// (The argument to Unmarshal must be a non-nil pointer.)
type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "json: Unmarshal(nil)"
	}

	if e.Type.Kind() != reflect.Pointer {
		return "json: Unmarshal(non-pointer " + e.Type.String() + ")"
	}
	return "json: Unmarshal(nil " + e.Type.String() + ")"
}

func (d *decodeState) unmarshal(v any) error {
	d.target = v
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}

	d.scan.reset()
	d.scanWhile(scanSkipSpace, scanSkipComment)
	// We decode rv not rv.Elem because the Unmarshaler interface
	// test must be applied at the top level of the value.
	err := d.value(rv, nil)
	if err != nil {
		return d.addErrorContext(err)
	}

	if d.savedError != nil {
		return d.savedError
	}

	for _, f := range d.afterFunctions {
		err := f()

		if err != nil {
			return err
		}
	}

	return nil
}

// A Number represents a JSON number literal.
type Number string

// String returns the literal text of the number.
func (n Number) String() string { return string(n) }

// Float64 returns the number as a float64.
func (n Number) Float64() (float64, error) {
	return strconv.ParseFloat(string(n), 64)
}

// Int64 returns the number as an int64.
func (n Number) Int64() (int64, error) {
	return strconv.ParseInt(string(n), 10, 64)
}

// An errorContext provides context for type errors during decoding.
type errorContext struct {
	Struct     reflect.Type
	FieldStack []string
}

// decodeState represents the state while decoding a JSON value.
type decodeState struct {
	data                  []byte
	off                   int // next read offset in data
	opcode                int // last read result
	scan                  scanner
	errorContext          *errorContext
	savedError            error
	useNumber             bool
	disallowUnknownFields bool
	target                any
	afterFunctions        []func() error
}

// readIndex returns the position of the last byte read.
func (d *decodeState) readIndex() int {
	return d.off - 1
}

// phasePanicMsg is used as a panic message when we end up with something that
// shouldn't happen. It can indicate a bug in the JSON decoder, or that
// something is editing the data slice while the decoder executes.
const phasePanicMsg = "JSON decoder out of sync - data changing underfoot?"

func (d *decodeState) init(data []byte) *decodeState {
	d.data = data
	d.off = 0
	d.savedError = nil
	if d.errorContext != nil {
		d.errorContext.Struct = nil
		// Reuse the allocated space for the FieldStack slice.
		d.errorContext.FieldStack = d.errorContext.FieldStack[:0]
	}
	return d
}

// saveError saves the first err it is called with,
// for reporting at the end of the unmarshal.
func (d *decodeState) saveError(err error) {
	if d.savedError == nil {
		d.savedError = d.addErrorContext(err)
	}
}

// addErrorContext returns a new error enhanced with information from d.errorContext
func (d *decodeState) addErrorContext(err error) error {
	if d.errorContext != nil && (d.errorContext.Struct != nil || len(d.errorContext.FieldStack) > 0) {
		switch err := err.(type) {
		case *UnmarshalTypeError:
			err.Struct = d.errorContext.Struct.Name()
			err.Field = strings.Join(d.errorContext.FieldStack, ".")
		}
	}
	return err
}

// skip scans to the end of what was started.
func (d *decodeState) skip() {
	s, data, i := &d.scan, d.data, d.off
	depth := len(s.parseState)
	for {
		//if i == 210324 {
		//	fmt.Println()
		//}
		op := s.step(s, data[i])
		i++
		if len(s.parseState) < depth {
			d.off = i
			d.opcode = op
			return
		}
	}
}

// scanNext processes the byte at d.data[d.off].
func (d *decodeState) scanNext() {
	if d.off < len(d.data) {
		d.opcode = d.scan.step(&d.scan, d.data[d.off])
		d.off++
	} else {
		d.opcode = d.scan.eof()
		d.off = len(d.data) + 1 // mark processed EOF with len+1
	}
}

// scanWhile processes bytes in d.data[d.off:] until it
// receives a scan code not equal to op.
func (d *decodeState) scanWhile(ops ...int) {
	opMap := make(map[int]bool)
	for _, op := range ops {
		opMap[op] = true
	}
	s, data, i := &d.scan, d.data, d.off
	for i < len(data) {
		newOp := s.step(s, data[i])
		i++
		if !opMap[newOp] {
			d.opcode = newOp
			d.off = i
			return
		}
	}

	d.off = len(data) + 1 // mark processed EOF with len+1
	d.opcode = d.scan.eof()
}

// rescanLiteral is similar to scanWhile(scanContinue), but it specialises the
// common case where we're decoding a literal. The decoder scans the input
// twice, once for syntax errors and to check the length of the value, and the
// second to perform the decoding.
//
// Only in the second step do we use decodeState to tokenize literals, so we
// know there aren't any syntax errors. We can take advantage of that knowledge,
// and scan a literal's bytes much more quickly.
func (d *decodeState) rescanLiteral() {
	data, i := d.data, d.off

	if data[i-1] == '"' { // string
	Outer:
		for ; i < len(data); i++ {
			switch data[i] {
			case '\\':
				i++ // escaped char
			case '"':
				i++ // tokenize the closing quote too
				break Outer
			}
		}
	} else if (data[i-1] >= 'a' && data[i-1] <= 'z') ||
		(data[i-1] >= 'A' && data[i-1] <= 'Z') ||
		(data[i-1] >= '0' && data[i-1] <= '9') ||
		(130 < data[i-1] && data[i-1] <= 255) ||
		data[i-1] == '-' {
	Outer2:
		for ; i < len(data); i++ {
			c := data[i]
			if ('0' <= c && c <= '9') ||
				('a' <= c && c <= 'z') ||
				('A' <= c && c <= 'Z') ||
				c == '_' ||
				c == '.' ||
				c == '-' ||
				c == '\'' ||
				(130 <= c && c <= 255) {
				// continue
			} else {
				break Outer2
			}
		}
	}

	if i < len(data) {
		d.opcode = stateEndValue(&d.scan, data[i])
	} else {
		d.opcode = scanEnd
	}
	d.off = i + 1
}

func (d *decodeState) MapValue(v reflect.Value, keyName string, f *field) error {
	if v.IsValid() {
		// Check for unmarshaler.
		u, ut, pv := indirect(v, false)
		if u != nil {
			start := d.readIndex()
			d.skip()
			return u.UnmarshalP(d.data[start:d.off])
		}
		if ut != nil {
			d.saveError(&UnmarshalTypeError{Value: "list", Type: v.Type(), Offset: int64(d.off)})
			d.skip()
			return nil
		}
		v = pv
		t := v.Type()

		if v.IsNil() {
			v.Set(reflect.MakeMap(t))
		}

		var mapElem reflect.Value
		elemType := t.Elem()
		if !mapElem.IsValid() {
			mapElem = reflect.New(elemType).Elem()
		} else {
			mapElem.Set(reflect.Zero(elemType))
		}

		if err := d.value(mapElem, f); err != nil {
			return err
		}

		kt := t.Key()
		var kv reflect.Value
		switch {
		case kt.Kind() == reflect.String:
			kv = reflect.ValueOf(keyName).Convert(kt)
		default:
			switch kt.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				n, err := strconv.ParseInt(keyName, 10, 64)
				if err != nil || reflect.Zero(kt).OverflowInt(n) {
					d.saveError(&UnmarshalTypeError{Value: "number " + keyName, Type: kt})
					break
				}
				kv = reflect.ValueOf(n).Convert(kt)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				n, err := strconv.ParseUint(keyName, 10, 64)
				if err != nil || reflect.Zero(kt).OverflowUint(n) {
					d.saveError(&UnmarshalTypeError{Value: "number " + keyName, Type: kt})
					break
				}
				kv = reflect.ValueOf(n).Convert(kt)
			default:
				panic("json: Unexpected key type") // should never occur
			}
		}
		if kv.IsValid() {
			v.SetMapIndex(kv, mapElem)
			_, _, subvv := indirect(mapElem, false)
			if subvv.Kind() == reflect.Struct {
				for i := 0; i < subvv.NumField(); i++ {
					if subvv.Type().Field(i).Tag.Get("paradox_type") == "map_key" {
						subvv.Field(i).Set(kv)
						continue
					}
					if subvv.Type().Field(i).Tag.Get("paradox_type") == "map_index" {
						subvv.Field(i).Set(reflect.ValueOf(v.Len()))
						continue
					}
				}
			}
		}

	} else {
		d.skip()
	}
	d.scanNext()

	return nil
}

func (d *decodeState) mapValueValue(v reflect.Value, f *field) error {
	if v.IsValid() {
		// Check for unmarshaler.
		u, ut, pv := indirect(v, false)
		if u != nil {
			start := d.readIndex()
			d.skip()
			return u.UnmarshalP(d.data[start:d.off])
		}
		if ut != nil {
			d.saveError(&UnmarshalTypeError{Value: "map_value", Type: v.Type(), Offset: int64(d.off)})
			d.skip()
			return nil
		}
		v = pv

		c := d.data[d.off-1]

		if isSpace(c) {
			d.scanWhile(scanSkipSpace, scanSkipComment)
			c = d.data[d.off]
		}

		if c == '{' {
			if err := d.value(v, f); err != nil {
				return err
			}
		} else {
			start := d.readIndex()
			d.rescanLiteral()

			name := string(d.data[start:d.readIndex()])
			d.afterFunctions = append(d.afterFunctions, func() error {
				m, ok := findFieldMapValue(d.target, f.mapName)

				if ok {
					key := m.Type().Key()

					kv := reflect.ValueOf(name).Convert(key)

					mv := m.MapIndex(kv)

					if mv.IsNil() {
						return errors.New("not found map value " + string(name) + " in " + f.mapName)
					}

					_, _, rmv := indirect(mv, false)
					v.Set(rmv)

					return nil
				}

				return errors.New("not found map " + f.mapName)
			})
		}
	} else {
		d.skip()
	}
	d.scanNext()

	return nil
}

func findFieldMapValue(value any, path string) (reflect.Value, bool) {
	_, _, v := indirect(reflect.ValueOf(value), false)

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		t := v.Type().Field(i)
		fieldName := t.Tag.Get("paradox_field")
		if fieldName == path {
			return f, true
		}
	}

	return reflect.Value{}, false
}

func (d *decodeState) fieldListValue(v reflect.Value, f *field) error {

	if v.IsValid() {
		// Check for unmarshaler.
		u, ut, pv := indirect(v, false)
		if u != nil {
			start := d.readIndex()
			d.skip()
			return u.UnmarshalP(d.data[start:d.off])
		}
		if ut != nil {
			d.saveError(&UnmarshalTypeError{Value: "field_list", Type: v.Type(), Offset: int64(d.off)})
			d.skip()
			return nil
		}
		v = pv

		c := d.data[d.off-1]

		if isSpace(c) {
			d.scanWhile(scanSkipSpace, scanSkipComment)
			c = d.data[d.off]
		}

		if c != '{' {
			v.Set(reflect.MakeSlice(v.Type(), 1, 4))

			start := d.readIndex()
			d.rescanLiteral()

			if err := d.literalStore(d.data[start:d.readIndex()], v.Index(v.Len()-1), false, f); err != nil {
				return err
			}
		} else {

			i := 0
			for {
				// Look ahead for ] - can only happen on first iteration.
				d.scanWhile(scanSkipSpace, scanSkipComment)
				if d.opcode == scanEndObject {
					break
				}

				if v.IsNil() {
					v.Set(reflect.MakeSlice(v.Type(), 1, 4))
				} else {
					if v.Len() == v.Cap() {
						newCap := v.Cap() + v.Cap()/2
						newSplice := reflect.MakeSlice(v.Type(), v.Len()+1, newCap)
						reflect.Copy(newSplice, v)
						v.Set(newSplice)
					} else {
						v.SetLen(v.Len() + 1)
					}

				}

				if err := d.value(v.Index(v.Len()-1), f); err != nil {
					return err
				}

				if d.opcode == scanEndObject {
					break
				}

				i++
			}
		}

	} else {
		d.skip()
	}
	d.scanNext()

	return nil
}

func (d *decodeState) listValue(v reflect.Value, f *field) error {
	if v.IsValid() {
		// Check for unmarshaler.
		u, ut, pv := indirect(v, false)
		if u != nil {
			start := d.readIndex()
			d.skip()
			return u.UnmarshalP(d.data[start:d.off])
		}
		if ut != nil {
			d.saveError(&UnmarshalTypeError{Value: "list", Type: v.Type(), Offset: int64(d.off)})
			d.skip()
			return nil
		}
		v = pv

		if v.IsNil() {
			v.Set(reflect.MakeSlice(v.Type(), 1, 4))
		} else {
			if v.Len() == v.Cap() {
				newCap := v.Cap() + v.Cap()/2
				newSplice := reflect.MakeSlice(v.Type(), v.Len()+1, newCap)
				reflect.Copy(newSplice, v)
				v.Set(newSplice)
			} else {
				v.SetLen(v.Len() + 1)
			}

		}

		if err := d.value(v.Index(v.Len()-1), f); err != nil {
			return err
		}

	} else {
		d.skip()
	}
	d.scanNext()

	return nil
}

func (d *decodeState) Entity(v reflect.Value, fff *field) error {
	if d.opcode == scanBeginObject {
		if v.IsValid() {
			if err := d.object(v); err != nil {
				return err
			}
		} else {
			d.skip()
		}
		d.scanNext()
		return nil
	}

	// Check for unmarshaler.
	u, ut, pv := indirect(v, false)
	if u != nil {
		start := d.readIndex()
		d.skip()
		return u.UnmarshalP(d.data[start:d.off])
	}
	if ut != nil {
		d.saveError(&UnmarshalTypeError{Value: "object", Type: v.Type(), Offset: int64(d.off)})
		d.skip()
		return nil
	}
	v = pv
	t := v.Type()

	if v.Kind() == reflect.Interface && v.NumMethod() == 0 {
		oi := d.objectInterface()
		v.Set(reflect.ValueOf(oi))
		return nil
	}

	var fields structFields

	if v.Kind() == reflect.Struct {
		fields = cachedTypeFields(t)
	} else {
		d.saveError(&UnmarshalTypeError{Value: "object", Type: t, Offset: int64(d.off)})
		d.skip()
		return nil
	}

	var f *field

	if i, ok := fields.nameIndex[fff.defaultField]; ok {
		// Found an exact name match.
		f = &fields.list[i]
	}

	var subv reflect.Value

	if f != nil {
		subv = v
		for _, i := range f.index {
			if subv.Kind() == reflect.Pointer {
				if subv.IsNil() {
					// If a struct embeds a pointer to an unexported type,
					// it is not possible to set a newly allocated value
					// since the field is unexported.
					//
					// See https://golang.org/issue/21357
					if !subv.CanSet() {
						d.saveError(fmt.Errorf("json: cannot set embedded pointer to unexported struct: %v", subv.Type().Elem()))
						// Invalidate subv to ensure d.value(subv) skips over
						// the JSON value without assigning it to subv.
						subv = reflect.Value{}
						break
					}
					subv.Set(reflect.New(subv.Type().Elem()))
				}
				subv = subv.Elem()
			}
			subv = subv.Field(i)
		}
		if d.errorContext == nil {
			d.errorContext = new(errorContext)
		}
		d.errorContext.FieldStack = append(d.errorContext.FieldStack, f.name)
		d.errorContext.Struct = t
	} else if d.disallowUnknownFields {
		d.saveError(fmt.Errorf("json: unknown field %q", fff.defaultField))
	}

	// All bytes inside literal return scanContinue op code.
	start := d.readIndex()
	d.rescanLiteral()

	if v.IsValid() {
		if err := d.literalStore(d.data[start:d.readIndex()], subv, false, f); err != nil {
			return err
		}
	}

	return nil
}

// value consumes a JSON value from d.data[d.off-1:], decoding into v, and
// reads the following byte ahead. If v is invalid, the value is discarded.
// The first byte of the value has been read already.
func (d *decodeState) value(v reflect.Value, f *field) error {
	switch d.opcode {
	default:
		panic(phasePanicMsg)

	case scanBeginArray:
		if v.IsValid() {
			if err := d.array(v, f); err != nil {
				return err
			}
		} else {
			d.skip()
		}
		d.scanNext()

	case scanBeginObject:
		if v.IsValid() {
			if err := d.object(v); err != nil {
				return err
			}
		} else {
			d.skip()
		}
		d.scanNext()

	case scanBeginLiteral:
		// All bytes inside literal return scanContinue op code.
		start := d.readIndex()
		d.rescanLiteral()

		if v.IsValid() {
			if err := d.literalStore(d.data[start:d.readIndex()], v, false, f); err != nil {
				return err
			}
		}
	}
	return nil
}

type unquotedValue struct{}

// valueQuoted is like value but decodes a
// quoted string literal or literal null into an interface value.
// If it finds anything other than a quoted string literal or null,
// valueQuoted returns unquotedValue{}.
func (d *decodeState) valueQuoted() any {
	switch d.opcode {
	default:
		panic(phasePanicMsg)

	case scanBeginArray, scanBeginObject:
		d.skip()
		d.scanNext()

	case scanBeginLiteral:
		v := d.literalInterface()
		switch v.(type) {
		case nil, string:
			return v
		}
	}
	return unquotedValue{}
}

// indirect walks down v allocating pointers as needed,
// until it gets to a non-pointer.
// If it encounters an Unmarshaler, indirect stops and returns that.
// If decodingNull is true, indirect stops at the first settable pointer so it
// can be set to nil.
func indirect(v reflect.Value, decodingNull bool) (Unmarshaler, encoding.TextUnmarshaler, reflect.Value) {
	// Issue #24153 indicates that it is generally not a guaranteed property
	// that you may round-trip a reflect.Value by calling Value.Addr().Elem()
	// and expect the value to still be settable for values derived from
	// unexported embedded struct fields.
	//
	// The logic below effectively does this when it first addresses the value
	// (to satisfy possible pointer methods) and continues to dereference
	// subsequent pointers as necessary.
	//
	// After the first round-trip, we set v back to the original value to
	// preserve the original RW flags contained in reflect.Value.
	v0 := v
	haveAddr := false

	// If v is a named type and is addressable,
	// start with its address, so that if the type has pointer methods,
	// we find them.
	if v.Kind() != reflect.Pointer && v.Type().Name() != "" && v.CanAddr() {
		haveAddr = true
		v = v.Addr()
	}
	for {
		// Load value from interface, but only if the result will be
		// usefully addressable.
		if v.Kind() == reflect.Interface && !v.IsNil() {
			e := v.Elem()
			if e.Kind() == reflect.Pointer && !e.IsNil() && (!decodingNull || e.Elem().Kind() == reflect.Pointer) {
				haveAddr = false
				v = e
				continue
			}
		}

		if v.Kind() != reflect.Pointer {
			break
		}

		if decodingNull && v.CanSet() {
			break
		}

		// Prevent infinite loop if v is an interface pointing to its own address:
		//     var v interface{}
		//     v = &v
		if v.Elem().Kind() == reflect.Interface && v.Elem().Elem() == v {
			v = v.Elem()
			break
		}
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		if v.Type().NumMethod() > 0 && v.CanInterface() {
			if u, ok := v.Interface().(Unmarshaler); ok {
				return u, nil, reflect.Value{}
			}
			if !decodingNull {
				if u, ok := v.Interface().(encoding.TextUnmarshaler); ok {
					return nil, u, reflect.Value{}
				}
			}
		}

		if haveAddr {
			v = v0 // restore original value after round-trip Value.Addr().Elem()
			haveAddr = false
		} else {
			v = v.Elem()
		}
	}
	return nil, nil, v
}

// array consumes an array from d.data[d.off-1:], decoding into v.
// The first byte of the array ('[') has been read already.
func (d *decodeState) array(v reflect.Value, f *field) error {
	// Check for unmarshaler.
	u, ut, pv := indirect(v, false)
	if u != nil {
		start := d.readIndex()
		d.skip()
		return u.UnmarshalP(d.data[start:d.off])
	}
	if ut != nil {
		d.saveError(&UnmarshalTypeError{Value: "list", Type: v.Type(), Offset: int64(d.off)})
		d.skip()
		return nil
	}
	v = pv

	// Check type of target.
	switch v.Kind() {
	case reflect.Interface:
		if v.NumMethod() == 0 {
			// Decoding into nil interface? Switch to non-reflect code.
			ai := d.arrayInterface()
			v.Set(reflect.ValueOf(ai))
			return nil
		}
		// Otherwise it's invalid.
		fallthrough
	default:
		d.saveError(&UnmarshalTypeError{Value: "list", Type: v.Type(), Offset: int64(d.off)})
		d.skip()
		return nil
	case reflect.Array, reflect.Slice:
		break
	}

	i := 0
	for {
		// Look ahead for ] - can only happen on first iteration.
		d.scanWhile(scanSkipSpace, scanSkipComment)
		if d.opcode == scanEndArray {
			break
		}

		// Get element of list, growing if necessary.
		if v.Kind() == reflect.Slice {
			// Grow slice if necessary
			if i >= v.Cap() {
				newcap := v.Cap() + v.Cap()/2
				if newcap < 4 {
					newcap = 4
				}
				newv := reflect.MakeSlice(v.Type(), v.Len(), newcap)
				reflect.Copy(newv, v)
				v.Set(newv)
			}
			if i >= v.Len() {
				v.SetLen(i + 1)
			}
		}

		if i < v.Len() {
			// Decode into element.
			if err := d.value(v.Index(i), f); err != nil {
				return err
			}
		} else {
			// Ran out of fixed list: skip.
			if err := d.value(reflect.Value{}, f); err != nil {
				return err
			}
		}
		i++

		// Next token must be , or ].
		if d.opcode == scanSkipSpace {
			d.scanWhile(scanSkipSpace, scanSkipComment)
		}
		if d.opcode == scanEndArray {
			break
		}
		if d.opcode != scanArrayValue {
			panic(phasePanicMsg)
		}
	}

	if i < v.Len() {
		if v.Kind() == reflect.Array {
			// Array. Zero the rest.
			z := reflect.Zero(v.Type().Elem())
			for ; i < v.Len(); i++ {
				v.Index(i).Set(z)
			}
		} else {
			v.SetLen(i)
		}
	}
	if i == 0 && v.Kind() == reflect.Slice {
		v.Set(reflect.MakeSlice(v.Type(), 0, 0))
	}
	return nil
}

var nullLiteral = []byte("null")
var textUnmarshalerType = reflect.TypeOf((*encoding.TextUnmarshaler)(nil)).Elem()

// object consumes an object from d.data[d.off-1:], decoding into v.
// The first byte ('{') of the object has been read already.
func (d *decodeState) object(v reflect.Value) error {
	// Check for unmarshaler.
	u, ut, pv := indirect(v, false)
	if u != nil {
		start := d.readIndex()
		d.skip()
		return u.UnmarshalP(d.data[start:d.off])
	}
	if ut != nil {
		d.saveError(&UnmarshalTypeError{Value: "object", Type: v.Type(), Offset: int64(d.off)})
		d.skip()
		return nil
	}
	v = pv
	t := v.Type()

	// Decoding into nil interface? Switch to non-reflect code.
	if v.Kind() == reflect.Interface && v.NumMethod() == 0 {
		oi := d.objectInterface()
		v.Set(reflect.ValueOf(oi))
		return nil
	}

	var fields structFields

	// Check type of target:
	//   struct or
	//   map[T1]T2 where T1 is string, an integer type,
	//             or an encoding.TextUnmarshaler
	switch v.Kind() {
	case reflect.Map:
		// Map key must either have string kind, have an integer kind,
		// or be an encoding.TextUnmarshaler.
		switch t.Key().Kind() {
		case reflect.String,
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		default:
			if !reflect.PointerTo(t.Key()).Implements(textUnmarshalerType) {
				d.saveError(&UnmarshalTypeError{Value: "object", Type: t, Offset: int64(d.off)})
				d.skip()
				return nil
			}
		}
		if v.IsNil() {
			v.Set(reflect.MakeMap(t))
		}
	case reflect.Struct:
		fields = cachedTypeFields(t)
		// ok
	default:
		d.saveError(&UnmarshalTypeError{Value: "object", Type: t, Offset: int64(d.off)})
		d.skip()
		return nil
	}

	var mapElem reflect.Value
	var origErrorContext errorContext
	if d.errorContext != nil {
		origErrorContext = *d.errorContext
	}

	for {
		// Read opening " of string key or closing }.
		d.scanWhile(scanSkipSpace, scanSkipComment)
		if d.opcode == scanEndObject {
			// closing } - can only happen on first iteration.
			break
		}
		if d.opcode != scanBeginLiteral {
			panic(phasePanicMsg)
		}

		// Read key.
		start := d.readIndex()
		d.rescanLiteral()
		item := d.data[start:d.readIndex()]
		key, ok := unquoteBytes(item)
		keyName := string(key)
		//fmt.Println(keyName)
		//if keyName == "ITA" {
		//	fmt.Printf("Debug point %s\n", keyName)
		//}
		if !ok {
			log.Printf("json: invalid key: %q", keyName)
			panic(phasePanicMsg)
		}

		// Figure out field corresponding to key.
		var subv reflect.Value
		destring := false // whether the value is wrapped in a string to be decoded first

		var f *field

		if v.Kind() == reflect.Map {
			elemType := t.Elem()
			if !mapElem.IsValid() {
				mapElem = reflect.New(elemType).Elem()
			} else {
				mapElem.Set(reflect.Zero(elemType))
			}
			subv = mapElem
		} else {
			if i, ok := fields.nameIndex[string(key)]; ok {
				// Found an exact name match.
				f = &fields.list[i]
			} else {
				// Fall back to the expensive case-insensitive
				// linear search.
				for i := range fields.list {
					ff := &fields.list[i]
					if ff.equalFold(ff.nameBytes, key) {
						f = ff
						break
					}
					if ff.IsMatch(key) {
						f = ff
						break
					}
				}
			}
			if f != nil {
				subv = v
				destring = f.quoted
				for _, i := range f.index {
					if subv.Kind() == reflect.Pointer {
						if subv.IsNil() {
							// If a struct embeds a pointer to an unexported type,
							// it is not possible to set a newly allocated value
							// since the field is unexported.
							//
							// See https://golang.org/issue/21357
							if !subv.CanSet() {
								d.saveError(fmt.Errorf("json: cannot set embedded pointer to unexported struct: %v", subv.Type().Elem()))
								// Invalidate subv to ensure d.value(subv) skips over
								// the JSON value without assigning it to subv.
								subv = reflect.Value{}
								destring = false
								break
							}
							subv.Set(reflect.New(subv.Type().Elem()))
						}
						subv = subv.Elem()
					}
					subv = subv.Field(i)
				}
				if d.errorContext == nil {
					d.errorContext = new(errorContext)
				}
				d.errorContext.FieldStack = append(d.errorContext.FieldStack, f.name)
				d.errorContext.Struct = t
			} else if d.disallowUnknownFields {
				d.saveError(fmt.Errorf("json: unknown field %q", key))
			}
		}

		// Read : before value.
		if d.opcode == scanSkipSpace {
			d.scanWhile(scanSkipSpace, scanSkipComment)
		}
		if d.opcode != scanObjectKey {
			panic(phasePanicMsg)
		}
		d.scanWhile(scanSkipSpace, scanSkipComment)

		if destring {
			switch qv := d.valueQuoted().(type) {
			case nil:
				if err := d.literalStore(nullLiteral, subv, false, f); err != nil {
					return err
				}
			case string:
				if err := d.literalStore([]byte(qv), subv, true, f); err != nil {
					return err
				}
			default:
				d.saveError(fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal unquoted value into %v", subv.Type()))
			}
		} else {
			if f != nil && f.IsEntity() {
				if err := d.Entity(subv, f); err != nil {
					return err
				}
			} else if f != nil && f.IsMap() {
				if err := d.MapValue(subv, keyName, f); err != nil {
					return err
				}
			} else if f != nil && f.IsMapValue() {
				if err := d.mapValueValue(subv, f); err != nil {
					return err
				}
			} else if f != nil && f.IsFieldList() {
				if err := d.fieldListValue(subv, f); err != nil {
					return err
				}
			} else if f != nil && f.IsList() {
				if err := d.listValue(subv, f); err != nil {
					return err
				}
			} else {
				if err := d.value(subv, f); err != nil {
					return err
				}
			}
		}

		// Write value back to map;
		// if using struct, subv points into struct already.
		if v.Kind() == reflect.Map {
			kt := t.Key()
			var kv reflect.Value
			switch {
			case reflect.PointerTo(kt).Implements(textUnmarshalerType):
				kv = reflect.New(kt)
				if err := d.literalStore(item, kv, true, f); err != nil {
					return err
				}
				kv = kv.Elem()
			case kt.Kind() == reflect.String:
				kv = reflect.ValueOf(key).Convert(kt)
			default:
				switch kt.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					s := string(key)
					n, err := strconv.ParseInt(s, 10, 64)
					if err != nil || reflect.Zero(kt).OverflowInt(n) {
						d.saveError(&UnmarshalTypeError{Value: "number " + s, Type: kt, Offset: int64(start + 1)})
						break
					}
					kv = reflect.ValueOf(n).Convert(kt)
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
					s := string(key)
					n, err := strconv.ParseUint(s, 10, 64)
					if err != nil || reflect.Zero(kt).OverflowUint(n) {
						d.saveError(&UnmarshalTypeError{Value: "number " + s, Type: kt, Offset: int64(start + 1)})
						break
					}
					kv = reflect.ValueOf(n).Convert(kt)
				default:
					panic("json: Unexpected key type") // should never occur
				}
			}
			if kv.IsValid() {
				v.SetMapIndex(kv, subv)
				_, _, subvv := indirect(subv, false)
				if subvv.Kind() == reflect.Struct {
					for i := 0; i < subvv.NumField(); i++ {
						if subvv.Type().Field(i).Tag.Get("paradox_type") == "map_key" {
							subvv.Field(i).Set(kv)
							continue
						}
						if subvv.Type().Field(i).Tag.Get("paradox_type") == "map_index" {
							subvv.Field(i).Set(reflect.ValueOf(v.Len()))
							continue
						}
					}
				}
			}
		}

		if d.errorContext != nil {
			// Reset errorContext to its original state.
			// Keep the same underlying list for FieldStack, to reuse the
			// space and avoid unnecessary allocs.
			d.errorContext.FieldStack = d.errorContext.FieldStack[:len(origErrorContext.FieldStack)]
			d.errorContext.Struct = origErrorContext.Struct
		}
		if d.opcode == scanEndObject {
			break
		}
	}
	return nil
}

// convertNumber converts the number literal s to a float64 or a Number
// depending on the setting of d.useNumber.
func (d *decodeState) convertNumber(s string) (any, error) {
	if d.useNumber {
		return Number(s), nil
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil, &UnmarshalTypeError{Value: "number " + s, Type: reflect.TypeOf(0.0), Offset: int64(d.off)}
	}
	return f, nil
}

var numberType = reflect.TypeOf(Number(""))

// literalStore decodes a literal stored in item into v.
//
// fromQuoted indicates whether this literal came from unwrapping a
// string from the ",string" struct tag option. this is used only to
// produce more helpful error messages.
func (d *decodeState) literalStore(item []byte, v reflect.Value, fromQuoted bool, f *field) error {
	// Check for unmarshaler.
	if len(item) == 0 {
		//Empty string given
		d.saveError(fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type()))
		return nil
	}
	isNull := item[0] == 'n' // null
	u, ut, pv := indirect(v, isNull)
	if u != nil {
		return u.UnmarshalP(item)
	}
	if ut != nil {
		if item[0] != '"' {
			if fromQuoted {
				d.saveError(fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type()))
				return nil
			}
			val := "number"
			switch item[0] {
			case 'n':
				val = "null"
			case 't', 'f':
				val = "bool"
			}
			d.saveError(&UnmarshalTypeError{Value: val, Type: v.Type(), Offset: int64(d.readIndex())})
			return nil
		}
		s, ok := unquoteBytes(item)
		if !ok {
			if fromQuoted {
				return fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type())
			}
			panic(phasePanicMsg)
		}
		return ut.UnmarshalText(s)
	}

	v = pv

	switch c := item[0]; c {
	case '"': // string
		s, ok := unquoteBytes(item)
		if !ok {
			s, ok = unquoteBytes(item)
			if fromQuoted {
				return fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type())
			}
			panic(phasePanicMsg)
		}
		switch v.Kind() {
		default:
			d.saveError(&UnmarshalTypeError{Value: "string", Type: v.Type(), Offset: int64(d.readIndex())})
		case reflect.Slice:
			if v.Type().Elem().Kind() != reflect.Uint8 {
				d.saveError(&UnmarshalTypeError{Value: "string", Type: v.Type(), Offset: int64(d.readIndex())})
				break
			}
			b := make([]byte, base64.StdEncoding.DecodedLen(len(s)))
			n, err := base64.StdEncoding.Decode(b, s)
			if err != nil {
				d.saveError(err)
				break
			}
			v.SetBytes(b[:n])
		case reflect.String:
			if v.Type() == numberType && !isValidNumber(string(s)) {
				return fmt.Errorf("json: invalid number literal, trying to unmarshal %q into Number", item)
			}
			if f.isEscapedText {
				text, _ := utils.DecodeEscapedText(s)
				v.SetString(text)
			} else {
				v.SetString(string(s))
			}
		case reflect.Interface:
			if v.NumMethod() == 0 {
				v.Set(reflect.ValueOf(string(s)))
			} else {
				d.saveError(&UnmarshalTypeError{Value: "string", Type: v.Type(), Offset: int64(d.readIndex())})
			}
		}

	default: // number
		if !('0' <= c && c <= '9' || 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_' || c == '.' || c == '-' || (130 < c && c <= 255)) {
			if fromQuoted {
				return fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type())
			}
			panic(phasePanicMsg)
		}
		s := string(item)
		switch v.Kind() {
		default:
			if v.Kind() == reflect.String {
				// s must be a valid number, because it's
				// already been tokenized.
				v.SetString(s)
				break
			}
			if fromQuoted {
				return fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal %q into %v", item, v.Type())
			}
			d.saveError(&UnmarshalTypeError{Value: "number", Type: v.Type(), Offset: int64(d.readIndex())})
		case reflect.Interface:
			n, err := d.convertNumber(s)
			if err != nil {
				d.saveError(err)
				break
			}
			if v.NumMethod() != 0 {
				d.saveError(&UnmarshalTypeError{Value: "number", Type: v.Type(), Offset: int64(d.readIndex())})
				break
			}
			v.Set(reflect.ValueOf(n))

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			n, err := strconv.ParseInt(s, 10, 64)
			if err != nil || v.OverflowInt(n) {
				d.saveError(&UnmarshalTypeError{Value: "number " + s, Type: v.Type(), Offset: int64(d.readIndex())})
				break
			}
			v.SetInt(n)

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			n, err := strconv.ParseUint(s, 10, 64)
			if err != nil || v.OverflowUint(n) {
				d.saveError(&UnmarshalTypeError{Value: "number " + s, Type: v.Type(), Offset: int64(d.readIndex())})
				break
			}
			v.SetUint(n)

		case reflect.Float32, reflect.Float64:
			n, err := strconv.ParseFloat(s, v.Type().Bits())
			if err != nil || v.OverflowFloat(n) {
				d.saveError(&UnmarshalTypeError{Value: "number " + s, Type: v.Type(), Offset: int64(d.readIndex())})
				break
			}
			v.SetFloat(n)

		case reflect.Bool:
			b, err := strconv.ParseBool(s)
			if err != nil {
				d.saveError(&UnmarshalTypeError{Value: "bool " + s, Type: v.Type(), Offset: int64(d.readIndex())})
				break
			}
			v.SetBool(b)
		}
	}
	return nil
}

// The xxxInterface routines build up a value to be stored
// in an empty interface. They are not strictly necessary,
// but they avoid the weight of reflection in this common case.

// valueInterface is like value but returns interface{}
func (d *decodeState) valueInterface() (val any) {
	switch d.opcode {
	default:
		panic(phasePanicMsg)
	case scanBeginArray:
		val = d.arrayInterface()
		d.scanNext()
	case scanBeginObject:
		val = d.objectInterface()
		d.scanNext()
	case scanBeginLiteral:
		val = d.literalInterface()
	}
	return
}

// arrayInterface is like array but returns []interface{}.
func (d *decodeState) arrayInterface() []any {
	var v = make([]any, 0)
	for {
		// Look ahead for ] - can only happen on first iteration.
		d.scanWhile(scanSkipSpace, scanSkipComment)
		if d.opcode == scanEndArray {
			break
		}

		v = append(v, d.valueInterface())

		// Next token must be , or ].
		if d.opcode == scanSkipSpace {
			d.scanWhile(scanSkipSpace, scanSkipComment)
		}
		if d.opcode == scanEndArray {
			break
		}
		if d.opcode != scanArrayValue {
			panic(phasePanicMsg)
		}
	}
	return v
}

// objectInterface is like object but returns map[string]interface{}.
func (d *decodeState) objectInterface() map[string]any {
	m := make(map[string]any)
	for {
		// Read opening " of string key or closing }.
		d.scanWhile(scanSkipSpace, scanSkipComment)
		if d.opcode == scanEndObject {
			// closing } - can only happen on first iteration.
			break
		}
		if d.opcode != scanBeginLiteral {
			panic(phasePanicMsg)
		}

		// Read string key.
		start := d.readIndex()
		d.rescanLiteral()
		item := d.data[start:d.readIndex()]
		key, ok := unquote(item)
		if !ok {
			panic(phasePanicMsg)
		}

		// Read : before value.
		if d.opcode == scanSkipSpace {
			d.scanWhile(scanSkipSpace, scanSkipComment)
		}
		if d.opcode != scanObjectKey {
			panic(phasePanicMsg)
		}
		d.scanWhile(scanSkipSpace, scanSkipComment)

		// Read value.
		m[key] = d.valueInterface()

		// Next token must be , or }.
		if d.opcode == scanSkipSpace {
			d.scanWhile(scanSkipSpace, scanSkipComment)
		}
		if d.opcode == scanEndObject {
			break
		}
		if d.opcode != scanObjectValue {
			panic(phasePanicMsg)
		}
	}
	return m
}

// literalInterface consumes and returns a literal from d.data[d.off-1:] and
// it reads the following byte ahead. The first byte of the literal has been
// read already (that's how the caller knows it's a literal).
func (d *decodeState) literalInterface() any {
	// All bytes inside literal return scanContinue op code.
	start := d.readIndex()
	d.rescanLiteral()

	item := d.data[start:d.readIndex()]

	switch c := item[0]; c {
	case 'n': // null
		return nil

	case 't', 'f': // true, false
		return c == 't'

	case '"': // string
		s, ok := unquote(item)
		if !ok {
			panic(phasePanicMsg)
		}
		return s

	default: // number
		if c != '-' && (c < '0' || c > '9') {
			panic(phasePanicMsg)
		}
		n, err := d.convertNumber(string(item))
		if err != nil {
			d.saveError(err)
		}
		return n
	}
}

// getu4 decodes \uXXXX from the beginning of s, returning the hex value,
// or it returns -1.
func getu4(s []byte) rune {
	if len(s) < 6 || s[0] != '\\' || s[1] != 'u' {
		return -1
	}
	var r rune
	for _, c := range s[2:6] {
		switch {
		case '0' <= c && c <= '9':
			c = c - '0'
		case 'a' <= c && c <= 'f':
			c = c - 'a' + 10
		case 'A' <= c && c <= 'F':
			c = c - 'A' + 10
		default:
			return -1
		}
		r = r*16 + rune(c)
	}
	return r
}

// unquote converts a quoted JSON string literal s into an actual string t.
// The rules are different than for Go, so cannot use strconv.Unquote.
func unquote(s []byte) (t string, ok bool) {
	s, ok = unquoteBytes(s)
	t = string(s)
	return
}

func unquoteBytes(s []byte) (t []byte, ok bool) {
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}

	return s, true
}
