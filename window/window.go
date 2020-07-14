package window

import (
	"errors"
	"fmt"
	"github.com/gonutz/w32"
	"indev-engine/w32api"
	"syscall"
	"unsafe"
)

const (
	windowedStyle = w32.WS_CAPTION | w32.WS_SYSMENU | w32.WS_VISIBLE //w32.WS_OVERLAPPED |
)

type inputHandler func(w32.HWND, uint32, uintptr, uintptr) (code uintptr, ok bool)

func CreateWindow(width, height int) error {

	wndClassName, err := syscall.UTF16PtrFromString("InDevWindowClass")
	if err != nil {
		return err
	}
	class := w32.WNDCLASSEX{
		WndProc:   syscall.NewCallback(handleMessages(handleInput)), // Обработчики событий окна
		Cursor:    w32.LoadCursor(0, (*uint16)(unsafe.Pointer(uintptr(w32.IDC_ARROW)))),
		ClassName: wndClassName,
	}

	atom := w32.RegisterClassEx(&class)
	if atom == 0 {
		return w32api.RegisterClassExError{}
	}

	windowSize := w32.RECT{
		Left:   0,
		Top:    0,
		Right:  int32(width),
		Bottom: int32(height),
	}
	// NOTE MSDN says you cannot pass WS_OVERLAPPED to this function but it
	// seems to work (on XP and Windows 8.1 at least) in conjuntion with the
	// other flags
	if w32.AdjustWindowRect(&windowSize, windowedStyle, false) {
		width = int(windowSize.Width())
		height = int(windowSize.Height())
	}

	windowClassName, err := syscall.UTF16PtrFromString("InDevWindowClass")
	if err != nil {
		return err
	}
	window := w32.CreateWindowEx(
		0,
		windowClassName,
		nil,
		windowedStyle,
		0,
		0,
		width,
		height,
		0,
		0,
		0,
		nil,
	)
	if window == 0 {
		return errors.New("CreateWindowEx failed")
	}
	defer w32.DestroyWindow(window)

	if !w32.RegisterRawInputDevices(w32.RAWINPUTDEVICE{
		UsagePage: 0x01,
		Usage:     0x06,
		Target:    window,
	}) {
		return fmt.Errorf("RegisterRawInputDevices failed")
	}

	var msg w32.MSG
	w32.PeekMessage(&msg, 0, 0, 0, w32.PM_NOREMOVE)

	for msg.Message != w32.WM_QUIT { // TODO: добавить проверку на главный цикл
		if w32.PeekMessage(&msg, 0, 0, 0, w32.PM_REMOVE) {
			w32.TranslateMessage(&msg)
			w32.DispatchMessage(&msg)
		}
	}

	return nil
}

// Последовательная обработка сигналов через массив обработчиков
func handleMessages(inputHandlers ...inputHandler) func(hwnd w32.HWND, uMsg uint32, wParam, lParam uintptr) uintptr {
	return func(hwnd w32.HWND, uMsg uint32, wParam, lParam uintptr) uintptr {
		for _, handler := range inputHandlers {
			code, ok := handler(hwnd, uMsg, wParam, lParam)
			if ok {
				return code
			}
		}

		return w32api.DefaultWindowProcedure(hwnd, uMsg, wParam, lParam) // Возможно стоит этот обработчик вынести в передаваемый массив
	}
}

func handleInput(hwnd w32.HWND, uMsg uint32, wParam, lParam uintptr) (code uintptr, ok bool) {
	switch uMsg {
	case w32.WM_INPUT:
		raw, ok := w32.GetRawInputData(w32.HRAWINPUT(lParam), w32.RID_INPUT)
		if !ok {
			return 1, ok
		}
		if raw.Header.Type != w32.RIM_TYPEKEYBOARD {
			return 1, ok
		}
		key, down := rawInputToKey(raw.GetKeyboard())
		if key != 0 {
			if down {
				// NOTE Заглушка на выход
				if key == KeyQ {
					w32.PostQuitMessage(0)
				}
			}
		}
		return 1, ok
	default:
		return 1, false
	}
}
