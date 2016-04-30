// Copyright 2015 The TCell Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tcell

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
)

// Terminfo represents a terminfo entry.  Note that we use friendly names
// in Go, but when we write out JSON, we use the same names as terminfo.
// The name, aliases and smous, rmous fields do not come from terminfo directly.
type Terminfo struct {
	Name         string   `json:"name"`
	Aliases      []string `json:"aliases,omitempty"`
	Columns      int      `json:"cols,omitempty"`   // cols
	Lines        int      `json:"lines,omitempty"`  // lines
	Colors       int      `json:"colors,omitempty"` // colors
	Bell         string   `json:"bell,omitempty"`   // bell
	Clear        string   `json:"clear,omitempty"`  // clear
	EnterCA      string   `json:"smcup,omitempty"`  // smcup
	ExitCA       string   `json:"rmcup,omitempty"`  // rmcup
	ShowCursor   string   `json:"cnorm,omitempty"`  // cnorm
	HideCursor   string   `json:"civis,omitempty"`  // civis
	AttrOff      string   `json:"sgr0,omitempty"`   // sgr0
	Underline    string   `json:"smul,omitempty"`   // smul
	Bold         string   `json:"bold,omitempty"`   // bold
	Blink        string   `json:"blink,omitempty"`  // blink
	Reverse      string   `json:"rev,omitempty"`    // rev
	Dim          string   `json:"dim,omitempty"`    // dim
	EnterKeypad  string   `json:"smkx,omitempty"`   // smkx
	ExitKeypad   string   `json:"rmkx,omitempty"`   // rmkx
	SetFg        string   `json:"setaf,omitempty"`  // setaf
	SetBg        string   `json:"setbg,omitempty"`  // setab
	SetCursor    string   `json:"cup,omitempty"`    // cup
	CursorBack1  string   `json:"cub1,omitempty"`   // cub1
	CursorUp1    string   `json:"cuu1,omitempty"`   // cuu1
	PadChar      string   `json:"pad,omitempty"`    // pad
	KeyBackspace string   `json:"kbs,omitempty"`    // kbs
	KeyF1        string   `json:"kf1,omitempty"`    // kf1
	KeyF2        string   `json:"kf2,omitempty"`    // kf2
	KeyF3        string   `json:"kf3,omitempty"`    // kf3
	KeyF4        string   `json:"kf4,omitempty"`    // kf4
	KeyF5        string   `json:"kf5,omitempty"`    // kf5
	KeyF6        string   `json:"kf6,omitempty"`    // kf6
	KeyF7        string   `json:"kf7,omitempty"`    // kf7
	KeyF8        string   `json:"kf8,omitempty"`    // kf8
	KeyF9        string   `json:"kf9,omitempty"`    // kf9
	KeyF10       string   `json:"kf10,omitempty"`   // kf10
	KeyF11       string   `json:"kf11,omitempty"`   // kf11
	KeyF12       string   `json:"kf12,omitempty"`   // kf12
	KeyF13       string   `json:"kf13,omitempty"`   // kf13
	KeyF14       string   `json:"kf14,omitempty"`   // kf14
	KeyF15       string   `json:"kf15,omitempty"`   // kf15
	KeyF16       string   `json:"kf16,omitempty"`   // kf16
	KeyF17       string   `json:"kf17,omitempty"`   // kf17
	KeyF18       string   `json:"kf18,omitempty"`   // kf18
	KeyF19       string   `json:"kf19,omitempty"`   // kf19
	KeyF20       string   `json:"kf20,omitempty"`   // kf20
	KeyF21       string   `json:"kf21,omitempty"`   // kf21
	KeyF22       string   `json:"kf22,omitempty"`   // kf22
	KeyF23       string   `json:"kf23,omitempty"`   // kf23
	KeyF24       string   `json:"kf24,omitempty"`   // kf24
	KeyF25       string   `json:"kf25,omitempty"`   // kf25
	KeyF26       string   `json:"kf26,omitempty"`   // kf26
	KeyF27       string   `json:"kf27,omitempty"`   // kf27
	KeyF28       string   `json:"kf28,omitempty"`   // kf28
	KeyF29       string   `json:"kf29,omitempty"`   // kf29
	KeyF30       string   `json:"kf30,omitempty"`   // kf30
	KeyF31       string   `json:"kf31,omitempty"`   // kf31
	KeyF32       string   `json:"kf32,omitempty"`   // kf32
	KeyF33       string   `json:"kf33,omitempty"`   // kf33
	KeyF34       string   `json:"kf34,omitempty"`   // kf34
	KeyF35       string   `json:"kf35,omitempty"`   // kf35
	KeyF36       string   `json:"kf36,omitempty"`   // kf36
	KeyF37       string   `json:"kf37,omitempty"`   // kf37
	KeyF38       string   `json:"kf38,omitempty"`   // kf38
	KeyF39       string   `json:"kf39,omitempty"`   // kf39
	KeyF40       string   `json:"kf40,omitempty"`   // kf40
	KeyF41       string   `json:"kf41,omitempty"`   // kf41
	KeyF42       string   `json:"kf42,omitempty"`   // kf42
	KeyF43       string   `json:"kf43,omitempty"`   // kf43
	KeyF44       string   `json:"kf44,omitempty"`   // kf44
	KeyF45       string   `json:"kf45,omitempty"`   // kf45
	KeyF46       string   `json:"kf46,omitempty"`   // kf46
	KeyF47       string   `json:"kf47,omitempty"`   // kf47
	KeyF48       string   `json:"kf48,omitempty"`   // kf48
	KeyF49       string   `json:"kf49,omitempty"`   // kf49
	KeyF50       string   `json:"kf50,omitempty"`   // kf50
	KeyF51       string   `json:"kf51,omitempty"`   // kf51
	KeyF52       string   `json:"kf52,omitempty"`   // kf52
	KeyF53       string   `json:"kf53,omitempty"`   // kf53
	KeyF54       string   `json:"kf54,omitempty"`   // kf54
	KeyF55       string   `json:"kf55,omitempty"`   // kf55
	KeyF56       string   `json:"kf56,omitempty"`   // kf56
	KeyF57       string   `json:"kf57,omitempty"`   // kf57
	KeyF58       string   `json:"kf58,omitempty"`   // kf58
	KeyF59       string   `json:"kf59,omitempty"`   // kf59
	KeyF60       string   `json:"kf60,omitempty"`   // kf60
	KeyF61       string   `json:"kf61,omitempty"`   // kf61
	KeyF62       string   `json:"kf62,omitempty"`   // kf62
	KeyF63       string   `json:"kf63,omitempty"`   // kf63
	KeyF64       string   `json:"kf64,omitempty"`   // kf64
	KeyInsert    string   `json:"kich,omitempty"`   // kich1
	KeyDelete    string   `json:"kdch,omitempty"`   // kdch1
	KeyHome      string   `json:"khome,omitempty"`  // khome
	KeyEnd       string   `json:"kend,omitempty"`   // kend
	KeyHelp      string   `json:"khlp,omitempty"`   // khlp
	KeyPgUp      string   `json:"kpp,omitempty"`    // kpp
	KeyPgDn      string   `json:"knp,omitempty"`    // knp
	KeyUp        string   `json:"kcuu1,omitempty"`  // kcuu1
	KeyDown      string   `json:"kcud1,omitempty"`  // kcud1
	KeyLeft      string   `json:"kcub1,omitempty"`  // kcub1
	KeyRight     string   `json:"kcuf1,omitempty"`  // kcuf1
	KeyBacktab   string   `json:"kcbt,omitempty"`   // kcbt
	KeyExit      string   `json:"kext,omitempty"`   // kext
	KeyClear     string   `json:"kclr,omitempty"`   // kclr
	KeyPrint     string   `json:"kprt,omitempty"`   // kprt
	KeyCancel    string   `json:"kcan,omitempty"`   // kcan
	Mouse        string   `json:"kmous,omitempty"`  // kmous
	MouseMode    string   `json:"XM,omitempty"`     // XM
	AltChars     string   `json:"acsc,omitempty"`   // acsc
	EnterAcs     string   `json:"smacs,omitempty"`  // smacs
	ExitAcs      string   `json:"rmacs,omitempty"`  // rmacs
	EnableAcs    string   `json:"enacs,omitempty"`  // enacs
	KeyShfRight  string   `json:"kRIT,omitempty"`   // kRIT
	KeyShfLeft   string   `json:"kLFT,omitempty"`   // kLFT
	KeyShfHome   string   `json:"kHOM,omitempty"`   // kHOM
	KeyShfEnd    string   `json:"kEND,omitempty"`   // kEND

	// These are non-standard extensions to terminfo.  This includes
	// true color support, and some additional keys.  Its kind of bizarre
	// that shifted variants of left and right exist, but not up and down.
	// Terminal support for these are going to vary amongst XTerm
	// emulations, so don't depend too much on them in your application.

	SetFgRGB     string `json:"_setfrgb,omitempty"` // setfrgb
	SetBgRGB     string `json:"_setbrgb,omitempty"` // setbrgb
	KeyShfUp     string `json:"_kscu1,omitempty"`   // shift-up
	KeyShfDown   string `json:"_kscud1,omitempty"`  // shift-down
	KeyCtrlUp    string `json:"_kccu1,omitempty"`   // ctrl-up
	KeyCtrlDown  string `json:"_kccud1,omitempty"`  // ctrl-left
	KeyCtrlRight string `json:"_kccuf1,omitempty"`  // ctrl-right
	KeyCtrlLeft  string `json:"_kccub1,omitempty"`  // ctrl-left
	KeyMetaUp    string `json:"_kmcu1,omitempty"`   // meta-up
	KeyMetaDown  string `json:"_kmcud1,omitempty"`  // meta-left
	KeyMetaRight string `json:"_kmcuf1,omitempty"`  // meta-right
	KeyMetaLeft  string `json:"_kmcub1,omitempty"`  // meta-left
	KeyAltUp     string `json:"_kacu1,omitempty"`   // alt-up
	KeyAltDown   string `json:"_kacud1,omitempty"`  // alt-left
	KeyAltRight  string `json:"_kacuf1,omitempty"`  // alt-right
	KeyAltLeft   string `json:"_kacub1,omitempty"`  // alt-left
	KeyCtrlHome  string `json:"_kchome,omitempty"`
	KeyCtrlEnd   string `json:"_kcend,omitempty"`
	KeyMetaHome  string `json:"_kmhome,omitempty"`
	KeyMetaEnd   string `json:"_kmend,omitempty"`
	KeyAltHome   string `json:"_kahome,omitempty"`
	KeyAltEnd    string `json:"_kaend,omitempty"`
}

type stack []string

func (st stack) Push(v string) stack {
	return append(st, v)
}

func (st stack) Pop() (string, stack) {
	v := ""
	if len(st) > 0 {
		v = st[len(st)-1]
		st = st[:len(st)-1]
	}
	return v, st
}

func (st stack) PopInt() (int, stack) {
	v := ""
	v, st = st.Pop()
	i, _ := strconv.Atoi(v)
	return i, st
}

func (st stack) PopBool() (bool, stack) {
	v := ""
	v, st = st.Pop()
	if v == "1" {
		return true, st
	}
	return false, st
}

func (st stack) PushInt(i int) stack {
	return st.Push(strconv.Itoa(i))
}

func (st stack) PushBool(i bool) stack {
	if i {
		return st.Push("1")
	}
	return st.Push("0")
}

func nextch(s string, index int) (byte, int) {
	if index < len(s) {
		return s[index], index + 1
	}
	return 0, index
}

// static vars
var svars [26]string

// TParm takes a terminfo parameterized string, such as setaf or cup, and
// evaluates the string, and returns the result with the parameter
// applied.
func (t *Terminfo) TParm(s string, p ...int) string {
	var stk stack
	var a, b string
	var ai, bi int
	var ab bool
	var dvars [26]string
	var params [9]int

	buf := bytes.NewBufferString(s)
	out := &bytes.Buffer{}

	// make sure we always have 9 parameters -- makes it easier
	// later to skip checks
	for i := 0; i < len(params) && i < len(p); i++ {
		params[i] = p[i]
	}

	nest := 0

	for {

		ch, err := buf.ReadByte()
		if err != nil {
			break
		}

		if ch != '%' {
			out.WriteByte(ch)
			continue
		}

		ch, err = buf.ReadByte()
		if err != nil {
			// XXX Error
			break
		}

		switch ch {
		case '%': // quoted %
			out.WriteByte(ch)

		case 'i': // increment both parameters (ANSI cup support)
			params[0]++
			params[1]++

		case 'c', 's':
			// NB: these, and 'd' below are special cased for
			// efficiency.  They could be handled by the richer
			// format support below, less efficiently.
			a, stk = stk.Pop()
			out.WriteString(a)

		case 'd':
			ai, stk = stk.PopInt()
			out.WriteString(strconv.Itoa(ai))

		case '0', '1', '2', '3', '4', 'x', 'X', 'o', ':':
			f := "%"
			if ch == ':' {
				ch, _ = buf.ReadByte()
			}
			f += string(ch)
			for ch == '+' || ch == '-' || ch == '#' || ch == ' ' {
				f += string(ch)
			}
			for (ch >= '0' && ch <= '9') || ch == '.' {
				ch, _ = buf.ReadByte()
				f += string(ch)
			}
			switch ch {
			case 'd', 'x', 'X', 'o':
				ai, stk = stk.PopInt()
				out.WriteString(fmt.Sprintf(f, ai))
			case 'c', 's':
				a, stk = stk.Pop()
				out.WriteString(fmt.Sprintf(f, a))
			}

		case 'p': // push parameter
			ch, _ = buf.ReadByte()
			ai = int(ch - '1')
			if ai >= 0 && ai < len(params) {
				stk = stk.PushInt(params[ai])
			} else {
				stk = stk.PushInt(0)
			}

		case 'P': // pop & store variable
			ch, _ = buf.ReadByte()
			if ch >= 'A' && ch <= 'Z' {
				svars[int(ch-'A')], stk = stk.Pop()
			} else if ch >= 'a' && ch <= 'z' {
				dvars[int(ch-'a')], stk = stk.Pop()
			}

		case 'g': // recall & push variable
			ch, _ = buf.ReadByte()
			if ch >= 'A' && ch <= 'Z' {
				stk = stk.Push(svars[int(ch-'A')])
			} else if ch >= 'a' && ch <= 'z' {
				stk = stk.Push(dvars[int(ch-'a')])
			}

		case '\'': // push(char)
			ch, _ = buf.ReadByte()
			buf.ReadByte() // must be ' but we don't check
			stk = stk.Push(string(ch))

		case '{': // push(int)
			ai = 0
			ch, _ = buf.ReadByte()
			for ch >= '0' && ch <= '9' {
				ai *= 10
				ai += int(ch - '0')
				ch, _ = buf.ReadByte()
			}
			// ch must be '}' but no verification
			stk = stk.PushInt(ai)

		case 'l': // push(strlen(pop))
			a, stk = stk.Pop()
			stk = stk.PushInt(len(a))

		case '+':
			bi, stk = stk.PopInt()
			ai, stk = stk.PopInt()
			stk = stk.PushInt(ai + bi)

		case '-':
			bi, stk = stk.PopInt()
			ai, stk = stk.PopInt()
			stk = stk.PushInt(ai - bi)

		case '*':
			bi, stk = stk.PopInt()
			ai, stk = stk.PopInt()
			stk = stk.PushInt(ai * bi)

		case '/':
			bi, stk = stk.PopInt()
			ai, stk = stk.PopInt()
			if bi != 0 {
				stk = stk.PushInt(ai / bi)
			} else {
				stk = stk.PushInt(0)
			}

		case 'm': // push(pop mod pop)
			bi, stk = stk.PopInt()
			ai, stk = stk.PopInt()
			if bi != 0 {
				stk = stk.PushInt(ai % bi)
			} else {
				stk = stk.PushInt(0)
			}

		case '&': // AND
			bi, stk = stk.PopInt()
			ai, stk = stk.PopInt()
			stk = stk.PushInt(ai & bi)

		case '|': // OR
			bi, stk = stk.PopInt()
			ai, stk = stk.PopInt()
			stk = stk.PushInt(ai | bi)

		case '^': // XOR
			bi, stk = stk.PopInt()
			ai, stk = stk.PopInt()
			stk = stk.PushInt(ai ^ bi)

		case '~': // bit complement
			ai, stk = stk.PopInt()
			stk = stk.PushInt(ai ^ -1)

		case '!': // logical NOT
			ai, stk = stk.PopInt()
			stk = stk.PushBool(ai != 0)

		case '=': // numeric compare or string compare
			b, stk = stk.Pop()
			a, stk = stk.Pop()
			stk = stk.PushBool(a == b)

		case '>': // greater than, numeric
			bi, stk = stk.PopInt()
			ai, stk = stk.PopInt()
			stk = stk.PushBool(ai > bi)

		case '<': // less than, numeric
			bi, stk = stk.PopInt()
			ai, stk = stk.PopInt()
			stk = stk.PushBool(ai < bi)

		case '?': // start conditional

		case 't':
			ab, stk = stk.PopBool()
			if ab {
				// just keep going
				break
			}
			nest = 0
		ifloop:
			// this loop consumes everything until we hit our else,
			// or the end of the conditional
			for {
				ch, err = buf.ReadByte()
				if err != nil {
					break
				}
				if ch != '%' {
					continue
				}
				ch, _ = buf.ReadByte()
				switch ch {
				case ';':
					if nest == 0 {
						break ifloop
					}
					nest--
				case '?':
					nest++
				case 'e':
					if nest == 0 {
						break ifloop
					}
				}
			}

		case 'e':
			// if we got here, it means we didn't use the else
			// in the 't' case above, and we should skip until
			// the end of the conditional
			nest = 0
		elloop:
			for {
				ch, err = buf.ReadByte()
				if err != nil {
					break
				}
				if ch != '%' {
					continue
				}
				ch, _ = buf.ReadByte()
				switch ch {
				case ';':
					if nest == 0 {
						break elloop
					}
					nest--
				case '?':
					nest++
				}
			}

		case ';': // endif

		}
	}
	return out.String()
}

// TPuts emits the string to the writer, but expands inline padding
// indications (of the form $<[delay]> where [delay] is msec) to
// a suitable number of padding characters (usually null bytes) based
// upon the supplied baud.  At high baud rates, more padding characters
// will be inserted.  All Terminfo based strings should be emitted using
// this function.
func (t *Terminfo) TPuts(w io.Writer, s string, baud int) {
	for {
		beg := strings.Index(s, "$<")
		if beg < 0 {
			// Most strings don't need padding, which is good news!
			io.WriteString(w, s)
			return
		}
		io.WriteString(w, s[:beg])
		s = s[beg+2:]
		end := strings.Index(s, ">")
		if end < 0 {
			// unterminated.. just emit bytes unadulterated
			io.WriteString(w, "$<"+s)
			return
		}
		val := s[:end]
		s = s[end+1:]
		padus := 0
		unit := 1000
		dot := false
	loop:
		for i := range val {
			switch val[i] {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				padus *= 10
				padus += int(val[i] - '0')
				if dot {
					unit *= 10
				}
			case '.':
				if !dot {
					dot = true
				} else {
					break loop
				}
			default:
				break loop
			}
		}
		cnt := int(((baud / 8) * padus) / unit)
		for cnt > 0 {
			io.WriteString(w, t.PadChar)
			cnt--
		}
	}
}

// TGoto returns a string suitable for addressing the cursor at the given
// row and column.  The origin 0, 0 is in the upper left corner of the screen.
func (t *Terminfo) TGoto(col, row int) string {
	return t.TParm(t.SetCursor, row, col)
}

// TColor returns a string corresponding to the given foreground and background
// colors.  Either fg or bg can be set to -1 to elide.
func (t *Terminfo) TColor(fg, bg Color) string {
	fi := int(fg)
	bi := int(bg)
	rv := ""
	// As a special case, we map bright colors to lower versions if the
	// color table only holds 8.  For the remaining 240 colors, the user
	// is out of luck.  Someday we could create a mapping table, but its
	// not worth it.
	if t.Colors == 8 {
		if fi > 7 && fi < 16 {
			fi -= 8
		}
		if bi > 7 && bi < 16 {
			bi -= 8
		}
	}
	if t.Colors > fi && fi >= 0 {
		rv += t.TParm(t.SetFg, fi)
	}
	if t.Colors > bi && bi >= 0 {
		rv += t.TParm(t.SetBg, bi)
	}
	return rv
}

var terminfos map[string]*Terminfo
var aliases map[string]string
var dblock sync.Mutex

func initDB() {
	if terminfos == nil {
		terminfos = make(map[string]*Terminfo)
	}
	if aliases == nil {
		aliases = make(map[string]string)
	}
}

// AddTerminfo can be called to register a new Terminfo entry.
func AddTerminfo(t *Terminfo) {
	dblock.Lock()
	initDB()
	terminfos[t.Name] = t
	for _, x := range t.Aliases {
		terminfos[x] = t
	}
	dblock.Unlock()
}

func loadFromFile(fname string, term string) (*Terminfo, error) {
	f, e := os.Open(fname)
	if e != nil {
		return nil, e
	}
	d := json.NewDecoder(f)
	for {
		t := &Terminfo{}
		if e := d.Decode(t); e != nil {
			if e == io.EOF {
				return nil, ErrTermNotFound
			}
			return nil, e
		}
		if t.Name == term {
			return t, nil
		}
	}
}

// LookupTerminfo attemps to find a definition for the named $TERM.
// It first looks in the builtin database, which should cover just about
// everyone.  If it can't find one there, then it will attempt to read
// one from the JSON file located in either $TCELLDB, $HOME/.tcelldb
// or in this package's source directory as database.json).
func LookupTerminfo(name string) (*Terminfo, error) {
	dblock.Lock()
	initDB()
	t := terminfos[name]
	dblock.Unlock()

	if t == nil {
		// Load the database located here.  Its expected that TCELLSDB
		// points either to a single JSON file.
		if pth := os.Getenv("TCELLDB"); pth != "" {
			t, _ = loadFromFile(pth, name)
		}
		if t == nil {
			if pth := os.Getenv("HOME"); pth != "" {
				fname := path.Join(pth, ".tcelldb")
				t, _ = loadFromFile(fname, name)
			}
		}
		if t == nil {
			gopath := strings.Split(os.Getenv("GOPATH"),
				string(os.PathListSeparator))
			fname := path.Join("src",
				"github.com", "gdamore", "tcell",
				"database.json")
			for _, pth := range gopath {
				t, _ = loadFromFile(path.Join(pth, fname), name)
				if t != nil {
					break
				}
			}
		}
		if t != nil {
			dblock.Lock()
			terminfos[name] = t
			dblock.Unlock()
		}
	}
	if t == nil {
		return nil, ErrTermNotFound
	}
	return t, nil
}
