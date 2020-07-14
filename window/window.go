package window

import (
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

type Window interface {
	Create() error
	Close() error
	listenWindowMessages()
}

type window struct {
	width    int
	height   int
	handle   w32.HWND
	Finished chan interface{}
}

func NewWindow(width, height int) *window {
	return &window{
		width:    width,
		height:   height,
		Finished: make(chan interface{}),
	}
}

// Create and Run Windows window
func (w *window) Create() error {
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
		Right:  int32(w.width),
		Bottom: int32(w.height),
	}

	if w32.AdjustWindowRect(&windowSize, windowedStyle, false) {
		w.width = int(windowSize.Width())
		w.height = int(windowSize.Height())
	}

	windowClassName, err := syscall.UTF16PtrFromString("InDevWindowClass")
	if err != nil {
		return err
	}
	windowHandle := w32.CreateWindowEx(
		0,
		windowClassName,
		nil,
		windowedStyle,
		0,
		0,
		w.width,
		w.height,
		0,
		0,
		0,
		nil,
	)
	if windowHandle == 0 {
		return fmt.Errorf("CreateWindowEx failed")
	}

	if !w32.RegisterRawInputDevices(w32.RAWINPUTDEVICE{
		UsagePage: 0x01,
		Usage:     0x06,
		Target:    windowHandle,
	}) {
		return fmt.Errorf("RegisterRawInputDevices failed")
	}

	go w.listenWindowMessages()

	return nil
}

func (w *window) listenWindowMessages() {
	var msg w32.MSG
	w32.PeekMessage(&msg, w.handle, 0, 0, w32.PM_NOREMOVE)
	for msg.Message != w32.WM_QUIT {
		if w32.PeekMessage(&msg, w.handle, 0, 0, w32.PM_REMOVE) {
			w32.TranslateMessage(&msg)
			w32.DispatchMessage(&msg)
		}
	}
	close(w.Finished)
}

// TODO привязать все хендлеры к окну и вызывать процедуру Сlose из них
// Close window Handle (use only to manually close window)
func (w *window) Close() error {
	ok := w32.DestroyWindow(w.handle)
	if !ok {
		return fmt.Errorf("DestroyWindowError")
	}
	close(w.Finished)
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
					w32.DestroyWindow(hwnd)
				}
			}
		}
		return 1, ok
	default:
		return 1, false
	}
}
