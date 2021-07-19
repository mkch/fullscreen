package win32

import (
	"fmt"
	"syscall"
)

const (
	SM_CXSCREEN uintptr = 0
	SM_CYSCREEN uintptr = 1

	WS_POPUP   uintptr = 0x80000000
	WS_CHILD   uintptr = 0x40000000
	WS_BORDER  uintptr = 0x00800000
	WS_CAPTION uintptr = 0x00C00000
	WS_SIZEBOX uintptr = 0x00040000

	SWP_SHOWWINDOW uintptr = 0x0040
	SWP_NOSIZE     uintptr = 0x0001

	GWL_STYLE    int = -16
	HWND_TOPMOST int = -1

	SW_MAXIMIZE uintptr = 3
)

type POINT struct {
	X, Y int
}

var (
	ShowWindow,
	FindWindowW,
	GetSystemMetrics,
	SetWindowPos,
	SetWindowLongW,
	GetWindowLongW *syscall.LazyProc
)

func init() {
	user32 := syscall.NewLazyDLL("user32.dll")
	ShowWindow = user32.NewProc("ShowWindow")
	FindWindowW = user32.NewProc("FindWindowW")
	GetSystemMetrics = user32.NewProc("GetSystemMetrics")
	SetWindowLongW = user32.NewProc("SetWindowLongW")
	GetWindowLongW = user32.NewProc("GetWindowLongW")
	SetWindowPos = user32.NewProc("SetWindowPos")
}

func Error(f string, err error) error {
	return fmt.Errorf("%s: %w", f, err)
}
