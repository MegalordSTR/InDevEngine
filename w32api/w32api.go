package w32api

import (
	"github.com/gonutz/w32"
)

var b = false

// Window Procedure (WndProc)
//
// HWND hwnd,        // handle to window
// UINT uMsg,        // message identifier
// WPARAM wParam,    // first message parameter
// LPARAM lParam)    // second message parameter
func DefaultWindowProcedure(hwnd w32.HWND, uMsg uint32, wParam, lParam uintptr) uintptr {
	switch uMsg {
	case w32.WM_DESTROY:
		w32.PostQuitMessage(0)
		return 1
	default:
		return w32.DefWindowProc(hwnd, uMsg, wParam, lParam)
	}
}
