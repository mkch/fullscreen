package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"syscall"
	"unsafe"

	"github.com/mkch/fullscreen/win32"
)

const defaultTitle = "Bad Kids"

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), `Make a window fullscreen.
Usage: %s [title]
The default title is "%v".`, os.Args[0], defaultTitle)
		flag.PrintDefaults()
	}
	flag.Parse()
	log.Default().SetFlags(0)
}

func main() {
	title := defaultTitle
	args := flag.Args()
	if len(args) > 1 {
		log.Fatal("To many arguments. At most 1 is allowed.")
	} else if len(args) == 1 {
		title = args[0]
	}

	titleW, err := syscall.UTF16FromString(title)
	if err != nil {
		log.Panic(win32.Error("UTF16FromString", err))
	}

	hwnd, _, err := win32.FindWindowW.Call(0, uintptr(unsafe.Pointer(&titleW[0])))
	if hwnd == 0 {
		if err.(syscall.Errno) != syscall.Errno(0) {
			log.Panic(win32.Error("FindWindowW", err))
		}
		log.Fatalf("No such window: %v", title)
	}

	style, _, err := win32.GetWindowLongW.Call(hwnd, _uintptr(win32.GWL_STYLE))
	if style == 0 && err.(syscall.Errno) != syscall.Errno(0) {
		log.Panic(win32.Error("GetWindowLongW", err))
	}

	// Remove border and catpion.
	style = style &^ win32.WS_BORDER &^ win32.WS_CAPTION &^ win32.WS_SIZEBOX
	ret, _, err := win32.SetWindowLongW.Call(hwnd, _uintptr(win32.GWL_STYLE), style)
	if ret == 0 {
		log.Panic(win32.Error("SetWindowLongW", err))
	}

	// Make it topmost.
	ret, _, err = win32.SetWindowPos.Call(
		hwnd,
		_uintptr(win32.HWND_TOPMOST),
		0, 0, 0, 0, win32.SWP_NOSIZE,
	)
	if ret == 0 {
		log.Panic(win32.Error("SetWindowLongW", err))
	}

	// Maximize it.
	_, _, err = win32.ShowWindow.Call(hwnd, win32.SW_MAXIMIZE)
	if err.(syscall.Errno) != syscall.Errno(0) {
		log.Panic(win32.Error("ShowWindow", err))
	}

	log.Println("Done")
}

// _uintptr converts an int n to uintptr. Only useful when convertint a negative const.
func _uintptr(n int) uintptr {
	return uintptr(n)
}
