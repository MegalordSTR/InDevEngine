package window

import (
	"github.com/gonutz/w32"
	"strconv"
)

// Key represents a key on the keyboard.
type Key int

// These are all available keyboard keys.
const (
	KeyA Key = 1 + iota
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
	KeyG
	KeyH
	KeyI
	KeyJ
	KeyK
	KeyL
	KeyM
	KeyN
	KeyO
	KeyP
	KeyQ
	KeyR
	KeyS
	KeyT
	KeyU
	KeyV
	KeyW
	KeyX
	KeyY
	KeyZ
	Key0
	Key1
	Key2
	Key3
	Key4
	Key5
	Key6
	Key7
	Key8
	Key9
	KeyNum0
	KeyNum1
	KeyNum2
	KeyNum3
	KeyNum4
	KeyNum5
	KeyNum6
	KeyNum7
	KeyNum8
	KeyNum9
	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyF13
	KeyF14
	KeyF15
	KeyF16
	KeyF17
	KeyF18
	KeyF19
	KeyF20
	KeyF21
	KeyF22
	KeyF23
	KeyF24
	KeyEnter
	KeyNumEnter
	KeyLeftControl
	KeyRightControl
	KeyLeftShift
	KeyRightShift
	KeyLeftAlt
	KeyRightAlt
	KeyLeft
	KeyRight
	KeyUp
	KeyDown
	KeyEscape
	KeySpace
	KeyBackspace
	KeyTab
	KeyHome
	KeyEnd
	KeyPageDown
	KeyPageUp
	KeyDelete
	KeyInsert
	KeyNumAdd
	KeyNumSubtract
	KeyNumMultiply
	KeyNumDivide
	KeyCapslock
	KeyPrint
	KeyPause

	// NOTE keyCount has to come last
	keyCount
)

func (k Key) String() string {
	switch k {
	case KeyA:
		return "A"
	case KeyB:
		return "B"
	case KeyC:
		return "C"
	case KeyD:
		return "D"
	case KeyE:
		return "E"
	case KeyF:
		return "F"
	case KeyG:
		return "G"
	case KeyH:
		return "H"
	case KeyI:
		return "I"
	case KeyJ:
		return "J"
	case KeyK:
		return "K"
	case KeyL:
		return "L"
	case KeyM:
		return "M"
	case KeyN:
		return "N"
	case KeyO:
		return "O"
	case KeyP:
		return "P"
	case KeyQ:
		return "Q"
	case KeyR:
		return "R"
	case KeyS:
		return "S"
	case KeyT:
		return "T"
	case KeyU:
		return "U"
	case KeyV:
		return "V"
	case KeyW:
		return "W"
	case KeyX:
		return "X"
	case KeyY:
		return "Y"
	case KeyZ:
		return "Z"
	case Key0:
		return "0"
	case Key1:
		return "1"
	case Key2:
		return "2"
	case Key3:
		return "3"
	case Key4:
		return "4"
	case Key5:
		return "5"
	case Key6:
		return "6"
	case Key7:
		return "7"
	case Key8:
		return "8"
	case Key9:
		return "9"
	case KeyNum0:
		return "Num0"
	case KeyNum1:
		return "Num1"
	case KeyNum2:
		return "Num2"
	case KeyNum3:
		return "Num3"
	case KeyNum4:
		return "Num4"
	case KeyNum5:
		return "Num5"
	case KeyNum6:
		return "Num6"
	case KeyNum7:
		return "Num7"
	case KeyNum8:
		return "Num8"
	case KeyNum9:
		return "Num9"
	case KeyF1:
		return "F1"
	case KeyF2:
		return "F2"
	case KeyF3:
		return "F3"
	case KeyF4:
		return "F4"
	case KeyF5:
		return "F5"
	case KeyF6:
		return "F6"
	case KeyF7:
		return "F7"
	case KeyF8:
		return "F8"
	case KeyF9:
		return "F9"
	case KeyF10:
		return "F10"
	case KeyF11:
		return "F11"
	case KeyF12:
		return "F12"
	case KeyF13:
		return "F13"
	case KeyF14:
		return "F14"
	case KeyF15:
		return "F15"
	case KeyF16:
		return "F16"
	case KeyF17:
		return "F17"
	case KeyF18:
		return "F18"
	case KeyF19:
		return "F19"
	case KeyF20:
		return "F20"
	case KeyF21:
		return "F21"
	case KeyF22:
		return "F22"
	case KeyF23:
		return "F23"
	case KeyF24:
		return "F24"
	case KeyEnter:
		return "Enter"
	case KeyNumEnter:
		return "NumEnter"
	case KeyLeftControl:
		return "LeftControl"
	case KeyRightControl:
		return "RightControl"
	case KeyLeftShift:
		return "LeftShift"
	case KeyRightShift:
		return "RightShift"
	case KeyLeftAlt:
		return "LeftAlt"
	case KeyRightAlt:
		return "RightAlt"
	case KeyLeft:
		return "Left"
	case KeyRight:
		return "Right"
	case KeyUp:
		return "Up"
	case KeyDown:
		return "Down"
	case KeyEscape:
		return "Escape"
	case KeySpace:
		return "Space"
	case KeyBackspace:
		return "Backspace"
	case KeyTab:
		return "Tab"
	case KeyHome:
		return "Home"
	case KeyEnd:
		return "End"
	case KeyPageDown:
		return "PageDown"
	case KeyPageUp:
		return "PageUp"
	case KeyDelete:
		return "Delete"
	case KeyInsert:
		return "Insert"
	case KeyNumAdd:
		return "NumAdd"
	case KeyNumSubtract:
		return "NumSubtract"
	case KeyNumMultiply:
		return "NumMultiply"
	case KeyNumDivide:
		return "NumDivide"
	case KeyCapslock:
		return "Capslock"
	case KeyPrint:
		return "Print"
	case KeyPause:
		return "Pause"
	default:
		return "Unknown key " + strconv.Itoa(int(k))
	}
}

func rawInputToKey(kb w32.RAWKEYBOARD) (key Key, down bool) {
	virtualKey := kb.VKey
	scanCode := kb.MakeCode
	flags := kb.Flags

	down = flags&w32.RI_KEY_BREAK == 0

	if virtualKey == 255 {
		// discard "fake keys" which are part of an escaped sequence
		return 0, down
	} else if virtualKey == w32.VK_SHIFT {
		virtualKey = uint16(w32.MapVirtualKey(
			uint(scanCode),
			w32.MAPVK_VSC_TO_VK_EX,
		))
	}

	isE0 := (flags & w32.RI_KEY_E0) != 0

	switch virtualKey {
	case w32.VK_CONTROL:
		if isE0 {
			return KeyRightControl, down
		} else {
			return KeyLeftControl, down
		}
	case w32.VK_MENU:
		if isE0 {
			return KeyRightAlt, down
		} else {
			return KeyLeftAlt, down
		}
	case w32.VK_RETURN:
		if isE0 {
			return KeyNumEnter, down
		}
	case w32.VK_INSERT:
		if !isE0 {
			return KeyNum0, down
		}
	case w32.VK_HOME:
		if !isE0 {
			return KeyNum7, down
		}
	case w32.VK_END:
		if !isE0 {
			return KeyNum1, down
		}
	case w32.VK_PRIOR:
		if !isE0 {
			return KeyNum9, down
		}
	case w32.VK_NEXT:
		if !isE0 {
			return KeyNum3, down
		}
	case w32.VK_LEFT:
		if !isE0 {
			return KeyNum4, down
		}
	case w32.VK_RIGHT:
		if !isE0 {
			return KeyNum6, down
		}
	case w32.VK_UP:
		if !isE0 {
			return KeyNum8, down
		}
	case w32.VK_DOWN:
		if !isE0 {
			return KeyNum2, down
		}
	case w32.VK_CLEAR:
		if !isE0 {
			return KeyNum5, down
		}
	}

	if virtualKey >= 'A' && virtualKey <= 'Z' {
		return KeyA + Key(virtualKey-'A'), down
	} else if virtualKey >= '0' && virtualKey <= '9' {
		return Key0 + Key(virtualKey-'0'), down
	} else if virtualKey >= w32.VK_NUMPAD0 && virtualKey <= w32.VK_NUMPAD9 {
		return KeyNum0 + Key(virtualKey-w32.VK_NUMPAD0), down
	} else if virtualKey >= w32.VK_F1 && virtualKey <= w32.VK_F24 {
		return KeyF1 + Key(virtualKey-w32.VK_F1), down
	} else {
		switch virtualKey {
		case w32.VK_RETURN:
			return KeyEnter, down
		case w32.VK_LEFT:
			return KeyLeft, down
		case w32.VK_RIGHT:
			return KeyRight, down
		case w32.VK_UP:
			return KeyUp, down
		case w32.VK_DOWN:
			return KeyDown, down
		case w32.VK_ESCAPE:
			return KeyEscape, down
		case w32.VK_SPACE:
			return KeySpace, down
		case w32.VK_BACK:
			return KeyBackspace, down
		case w32.VK_TAB:
			return KeyTab, down
		case w32.VK_HOME:
			return KeyHome, down
		case w32.VK_END:
			return KeyEnd, down
		case w32.VK_NEXT:
			return KeyPageDown, down
		case w32.VK_PRIOR:
			return KeyPageUp, down
		case w32.VK_DELETE:
			return KeyDelete, down
		case w32.VK_INSERT:
			return KeyInsert, down
		case w32.VK_LSHIFT:
			return KeyLeftShift, down
		case w32.VK_RSHIFT:
			return KeyRightShift, down
		case w32.VK_PRINT:
			return KeyPrint, down
		case w32.VK_PAUSE:
			return KeyPause, down
		case w32.VK_CAPITAL:
			return KeyCapslock, down
		case w32.VK_MULTIPLY:
			return KeyNumMultiply, down
		case w32.VK_ADD:
			return KeyNumAdd, down
		case w32.VK_SUBTRACT:
			return KeyNumSubtract, down
		case w32.VK_DIVIDE:
			return KeyNumDivide, down
		}
	}

	return Key(0), false
}
