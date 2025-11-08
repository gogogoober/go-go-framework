package inputservice

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GoGoKey = ebiten.Key

var KeyMap = map[string]GoGoKey{
	// Letters
	"a": ebiten.KeyA, "b": ebiten.KeyB, "c": ebiten.KeyC, "d": ebiten.KeyD,
	"e": ebiten.KeyE, "f": ebiten.KeyF, "g": ebiten.KeyG, "h": ebiten.KeyH,
	"i": ebiten.KeyI, "j": ebiten.KeyJ, "k": ebiten.KeyK, "l": ebiten.KeyL,
	"m": ebiten.KeyM, "n": ebiten.KeyN, "o": ebiten.KeyO, "p": ebiten.KeyP,
	"q": ebiten.KeyQ, "r": ebiten.KeyR, "s": ebiten.KeyS, "t": ebiten.KeyT,
	"u": ebiten.KeyU, "v": ebiten.KeyV, "w": ebiten.KeyW, "x": ebiten.KeyX,
	"y": ebiten.KeyY, "z": ebiten.KeyZ,

	// Number row
	"1": ebiten.Key1, "2": ebiten.Key2, "3": ebiten.Key3, "4": ebiten.Key4,
	"5": ebiten.Key5, "6": ebiten.Key6, "7": ebiten.Key7, "8": ebiten.Key8,
	"9": ebiten.Key9, "0": ebiten.Key0,

	// Punctuation on main row
	"-": ebiten.KeyMinus, "_": ebiten.KeyMinus, // shift for "_"
	"=": ebiten.KeyEqual, "+": ebiten.KeyEqual, // shift for "+"
	"[": ebiten.KeyLeftBracket, "{": ebiten.KeyLeftBracket, // shift for "{"
	"]": ebiten.KeyRightBracket, "}": ebiten.KeyRightBracket, // shift for "}"
	"\\": ebiten.KeyBackslash, "|": ebiten.KeyBackslash, // shift for "|"
	";": ebiten.KeySemicolon, ":": ebiten.KeySemicolon, // shift for ":"
	"'": ebiten.KeyApostrophe, "\"": ebiten.KeyApostrophe, // shift for `"`
	"`": ebiten.KeyGraveAccent, "~": ebiten.KeyGraveAccent, // shift for "~"
	",": ebiten.KeyComma, "<": ebiten.KeyComma, // shift for "<"
	".": ebiten.KeyPeriod, ">": ebiten.KeyPeriod, // shift for ">"
	"/": ebiten.KeySlash, "?": ebiten.KeySlash, // shift for "?"

	// Navigation / control
	"space": ebiten.KeySpace,
	"enter": ebiten.KeyEnter, "return": ebiten.KeyEnter,
	"esc": ebiten.KeyEscape, "escape": ebiten.KeyEscape,
	"backspace": ebiten.KeyBackspace,
	"tab":       ebiten.KeyTab,
	"capslock":  ebiten.KeyCapsLock,

	// Arrows
	"up": ebiten.KeyArrowUp, "arrowup": ebiten.KeyArrowUp,
	"down": ebiten.KeyArrowDown, "arrowdown": ebiten.KeyArrowDown,
	"left": ebiten.KeyArrowLeft, "arrowleft": ebiten.KeyArrowLeft,
	"right": ebiten.KeyArrowRight, "arrowright": ebiten.KeyArrowRight,

	"pagedown": ebiten.KeyPageDown,
	"home":     ebiten.KeyHome,
	"end":      ebiten.KeyEnd,
	"insert":   ebiten.KeyInsert,
	"delete":   ebiten.KeyDelete,

	// Function keys
	"f1": ebiten.KeyF1, "f2": ebiten.KeyF2, "f3": ebiten.KeyF3, "f4": ebiten.KeyF4,
	"f5": ebiten.KeyF5, "f6": ebiten.KeyF6, "f7": ebiten.KeyF7, "f8": ebiten.KeyF8,
	"f9": ebiten.KeyF9, "f10": ebiten.KeyF10, "f11": ebiten.KeyF11, "f12": ebiten.KeyF12,

	// Modifiers (Mac + Windows equivalents)
	"shift": ebiten.KeyShift, "lshift": ebiten.KeyShift, "rshift": ebiten.KeyShift,
	"ctrl": ebiten.KeyControl, "control": ebiten.KeyControl,
	"alt": ebiten.KeyAlt, "option": ebiten.KeyAlt, // Mac Option
	"cmd": ebiten.KeyMeta, "command": ebiten.KeyMeta, // Mac Command
	"meta": ebiten.KeyMeta, "super": ebiten.KeyMeta,
	"win": ebiten.KeyMeta, "windows": ebiten.KeyMeta, // Windows key
}

func IsKeyStringPressed(keys ...string) bool {
	for _, key := range keys {
		if k, ok := KeyMap[key]; ok {
			return ebiten.IsKeyPressed(k)
		}
	}
	return false
}

func IsKeyStringJustPressed(keys ...string) bool {
	for _, key := range keys {
		if k, ok := KeyMap[key]; ok {
			return inpututil.IsKeyJustPressed(k)
		}
	}
	return false
}

func IsKeyPressed(key GoGoKey) bool {
	return ebiten.IsKeyPressed(key)
}
