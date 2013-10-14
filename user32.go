// Copyright 2010-2012 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package w32

import (
	// #include <WTypes.h>
	"C"
	"errors"
	"fmt"
	"syscall"
	"unsafe"
)

var (
	moduser32 = syscall.NewLazyDLL("user32.dll")

	procRegisterClassEx               = moduser32.NewProc("RegisterClassExW")
	procLoadIcon                      = moduser32.NewProc("LoadIconW")
	procLoadCursor                    = moduser32.NewProc("LoadCursorW")
	procShowWindow                    = moduser32.NewProc("ShowWindow")
	procUpdateWindow                  = moduser32.NewProc("UpdateWindow")
	procCreateWindowEx                = moduser32.NewProc("CreateWindowExW")
	procAdjustWindowRect              = moduser32.NewProc("AdjustWindowRect")
	procAdjustWindowRectEx            = moduser32.NewProc("AdjustWindowRectEx")
	procDestroyWindow                 = moduser32.NewProc("DestroyWindow")
	procDefWindowProc                 = moduser32.NewProc("DefWindowProcW")
	procDefDlgProc                    = moduser32.NewProc("DefDlgProcW")
	procPostQuitMessage               = moduser32.NewProc("PostQuitMessage")
	procGetMessage                    = moduser32.NewProc("GetMessageW")
	procTranslateMessage              = moduser32.NewProc("TranslateMessage")
	procDispatchMessage               = moduser32.NewProc("DispatchMessageW")
	procSendMessage                   = moduser32.NewProc("SendMessageW")
	procPostMessage                   = moduser32.NewProc("PostMessageW")
	procWaitMessage                   = moduser32.NewProc("WaitMessage")
	procSetWindowText                 = moduser32.NewProc("SetWindowTextW")
	procGetWindowTextLength           = moduser32.NewProc("GetWindowTextLengthW")
	procGetWindowText                 = moduser32.NewProc("GetWindowTextW")
	procGetWindowRect                 = moduser32.NewProc("GetWindowRect")
	procMoveWindow                    = moduser32.NewProc("MoveWindow")
	procScreenToClient                = moduser32.NewProc("ScreenToClient")
	procCallWindowProc                = moduser32.NewProc("CallWindowProcW")
	procSetWindowLong                 = moduser32.NewProc("SetWindowLongW")
	procSetWindowLongPtr              = moduser32.NewProc("SetWindowLongW")
	procGetWindowLong                 = moduser32.NewProc("GetWindowLongW")
	procGetWindowLongPtr              = moduser32.NewProc("GetWindowLongW")
	procEnableWindow                  = moduser32.NewProc("EnableWindow")
	procIsWindowEnabled               = moduser32.NewProc("IsWindowEnabled")
	procIsWindowVisible               = moduser32.NewProc("IsWindowVisible")
	procSetFocus                      = moduser32.NewProc("SetFocus")
	procInvalidateRect                = moduser32.NewProc("InvalidateRect")
	procGetClientRect                 = moduser32.NewProc("GetClientRect")
	procGetDC                         = moduser32.NewProc("GetDC")
	procReleaseDC                     = moduser32.NewProc("ReleaseDC")
	procSetCapture                    = moduser32.NewProc("SetCapture")
	procReleaseCapture                = moduser32.NewProc("ReleaseCapture")
	procGetWindowThreadProcessId      = moduser32.NewProc("GetWindowThreadProcessId")
	procMessageBox                    = moduser32.NewProc("MessageBoxW")
	procGetSystemMetrics              = moduser32.NewProc("GetSystemMetrics")
	procCopyRect                      = moduser32.NewProc("CopyRect")
	procEqualRect                     = moduser32.NewProc("EqualRect")
	procInflateRect                   = moduser32.NewProc("InflateRect")
	procIntersectRect                 = moduser32.NewProc("IntersectRect")
	procIsRectEmpty                   = moduser32.NewProc("IsRectEmpty")
	procOffsetRect                    = moduser32.NewProc("OffsetRect")
	procPtInRect                      = moduser32.NewProc("PtInRect")
	procSetRect                       = moduser32.NewProc("SetRect")
	procSetRectEmpty                  = moduser32.NewProc("SetRectEmpty")
	procSubtractRect                  = moduser32.NewProc("SubtractRect")
	procUnionRect                     = moduser32.NewProc("UnionRect")
	procCreateDialogParam             = moduser32.NewProc("CreateDialogParamW")
	procDialogBoxParam                = moduser32.NewProc("DialogBoxParamW")
	procGetDlgItem                    = moduser32.NewProc("GetDlgItem")
	procDrawIcon                      = moduser32.NewProc("DrawIcon")
	procClientToScreen                = moduser32.NewProc("ClientToScreen")
	procIsDialogMessage               = moduser32.NewProc("IsDialogMessageW")
	procIsWindow                      = moduser32.NewProc("IsWindow")
	procEndDialog                     = moduser32.NewProc("EndDialog")
	procPeekMessage                   = moduser32.NewProc("PeekMessageW")
	procTranslateAccelerator          = moduser32.NewProc("TranslateAcceleratorW")
	procSetWindowPos                  = moduser32.NewProc("SetWindowPos")
	procFillRect                      = moduser32.NewProc("FillRect")
	procDrawText                      = moduser32.NewProc("DrawTextW")
	procAddClipboardFormatListener    = moduser32.NewProc("AddClipboardFormatListener")
	procRemoveClipboardFormatListener = moduser32.NewProc("RemoveClipboardFormatListener")
	procOpenClipboard                 = moduser32.NewProc("OpenClipboard")
	procCloseClipboard                = moduser32.NewProc("CloseClipboard")
	procEnumClipboardFormats          = moduser32.NewProc("EnumClipboardFormats")
	procGetClipboardData              = moduser32.NewProc("GetClipboardData")
	procSetClipboardData              = moduser32.NewProc("SetClipboardData")
	procEmptyClipboard                = moduser32.NewProc("EmptyClipboard")
	procGetClipboardFormatName        = moduser32.NewProc("GetClipboardFormatNameW")
	procIsClipboardFormatAvailable    = moduser32.NewProc("IsClipboardFormatAvailable")
	procBeginPaint                    = moduser32.NewProc("BeginPaint")
	procEndPaint                      = moduser32.NewProc("EndPaint")
	procGetKeyboardState              = moduser32.NewProc("GetKeyboardState")
	procMapVirtualKey                 = moduser32.NewProc("MapVirtualKeyExW")
	procGetAsyncKeyState              = moduser32.NewProc("GetAsyncKeyState")
	procToAscii                       = moduser32.NewProc("ToAscii")
	procSwapMouseButton               = moduser32.NewProc("SwapMouseButton")
	procGetCursorPos                  = moduser32.NewProc("GetCursorPos")
	procSetCursorPos                  = moduser32.NewProc("SetCursorPos")
	procSetCursor                     = moduser32.NewProc("SetCursor")
	procCreateIcon                    = moduser32.NewProc("CreateIcon")
	procDestroyIcon                   = moduser32.NewProc("DestroyIcon")
	procMonitorFromPoint              = moduser32.NewProc("MonitorFromPoint")
	procMonitorFromRect               = moduser32.NewProc("MonitorFromRect")
	procMonitorFromWindow             = moduser32.NewProc("MonitorFromWindow")
	procGetMonitorInfo                = moduser32.NewProc("GetMonitorInfoW")
	procEnumDisplayMonitors           = moduser32.NewProc("EnumDisplayMonitors")
	procEnumDisplaySettingsEx         = moduser32.NewProc("EnumDisplaySettingsExW")
	procChangeDisplaySettingsEx       = moduser32.NewProc("ChangeDisplaySettingsExW")
	procSendInput                     = moduser32.NewProc("SendInput")
	procFindWindow                    = moduser32.NewProc("FindWindowW")
)

// Windows
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms632595
// TODO: AllowSetForegroundWindow
// TODO: AnimateWindow
// TODO: AnyPopup
// TODO: ArrangeIconicWindows
// TODO: BeginDeferWindowPos
// TODO: BringWindowToTop
// TODO: CalculatePopupWindowPosition
// TODO: CascadeWindows
// TODO: ChangeWindowMessageFilter
// TODO: ChangeWindowMessageFilterEx
// TODO: ChildWindowFromPoint
// TODO: ChildWindowFromPointEx
// TODO: CloseWindow
// TODO: CreateWindow
// TODO: DeferWindowPos
// TODO: DeregisterShellHookWindow
// TODO: EndDeferWindowPos
// TODO: EndTask
// TODO: EnumChildProc
// TODO: EnumChildWindows
// TODO: EnumThreadWindows
// TODO: EnumThreadWndProc
// TODO: EnumWindows
// TODO: EnumWindowsProc
// TODO: FindWindowEx
// TODO: GetAltTabInfo
// TODO: GetAncestor
// TODO: GetDesktopWindow
// TODO: GetForegroundWindow
// TODO: GetGUIThreadInfo
// TODO: GetLastActivePopup
// TODO: GetLayeredWindowAttributes
// TODO: GetNextWindow
// TODO: GetParent
// TODO: GetProcessDefaultLayout
// TODO: GetShellWindow
// TODO: GetSysColor
// TODO: GetTitleBarInfo
// TODO: GetTopWindow
// TODO: GetWindow
// TODO: GetWindowDisplayAffinity
// TODO: GetWindowInfo
// TODO: GetWindowModuleFileName
// TODO: GetWindowPlacement
// TODO: InternalGetWindowText
// TODO: IsChild
// TODO: IsGUIThread
// TODO: IsHungAppWindow
// TODO: IsIconic
// TODO: IsProcessDPIAware
// TODO: IsWindowUnicode
// TODO: IsZoomed
// TODO: LockSetForegroundWindow
// TODO: LogicalToPhysicalPoint
// TODO: OpenIcon
// TODO: PhysicalToLogicalPoint
// TODO: RealChildWindowFromPoint
// TODO: RealGetWindowClass
// TODO: RegisterShellHookWindow
// TODO: SetForegroundWindow
// TODO: SetLayeredWindowAttributes
// TODO: SetParent
// TODO: SetProcessDefaultLayout
// TODO: SetProcessDPIAware
// TODO: SetSysColors
// TODO: SetWindowDisplayAffinity
// TODO: SetWindowFeedbackSettings
// TODO: SetWindowPlacement
// TODO: ShowOwnedPopups
// TODO: ShowWindowAsync
// TODO: SoundSentry
// TODO: SwitchToThisWindow
// TODO: TileWindows
// TODO: UpdateLayeredWindow
// TODO: UpdateLayeredWindowIndirect
// TODO: WindowFromPhysicalPoint
// TODO: WindowFromPoint
// TODO: WinMain - Needed?

// AdjustWindowRect calculates the required size of the window rectangle, based on the desired
// client-rectangle size. The window rectangle can then be passed to the CreateWindow function to
// create a window whose client area is the desired size.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms632665
func AdjustWindowRect(rect *RECT, style uint, menu bool) bool {
	ret, _, _ := procAdjustWindowRect.Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(style),
		uintptr(BoolToBOOL(menu)))

	return ret != 0
}

// AdjustWindowRectEx calculates the required size of the window rectangle, based on the desired
// size of the client rectangle. The window rectangle can then be passed to the CreateWindowEx
// function to create a window whose client area is the desired size.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms632667
func AdjustWindowRectEx(rect *RECT, style uint, menu bool, exStyle uint) bool {
	ret, _, _ := procAdjustWindowRectEx.Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(style),
		uintptr(BoolToBOOL(menu)),
		uintptr(exStyle))

	return ret != 0
}

// CreateWindowEx creates an overlapped, pop-up, or child window with an extended window style;
// otherwise, this function is identical to the CreateWindow function. For more information about
// creating a window and for full descriptions of the other parameters of CreateWindowEx,
// see CreateWindow.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms632680
func CreateWindowEx(exStyle uint, className, windowName *uint16,
	style uint, x, y, width, height int, parent HWND, menu HMENU,
	instance HINSTANCE, param unsafe.Pointer) HWND {
	ret, _, _ := procCreateWindowEx.Call(
		uintptr(exStyle),
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(windowName)),
		uintptr(style),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(parent),
		uintptr(menu),
		uintptr(instance),
		uintptr(param))

	return HWND(ret)
}

// DestroyWindow destroys the specified window. The function sends WM_DESTROY and WM_NCDESTROY
// messages to the window to deactivate it and remove the keyboard focus from it. The function also
// destroys the window's menu, flushes the thread message queue, destroys timers, removes clipboard
// ownership, and breaks the clipboard viewer chain (if the window is at the top of the viewer
// chain).
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms632682
func DestroyWindow(hwnd HWND) bool {
	ret, _, _ := procDestroyWindow.Call(
		uintptr(hwnd))

	return ret != 0
}

// FindWindow retrieves a handle to the top-level window whose class name and window name match the
// specified strings. This function is not case-senstive and does not search child windows. If a
// window cannot be found, an error is returned.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms633499
func FindWindow(lpClassName string, lpWindowName string) (HWND, error) {
	var strHelper uintptr
	if lpClassName != "" {
		strHelper = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpClassName)))
	}
	ret, _, _ := procFindWindow.Call(
		strHelper,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpWindowName))))
	if ret == 0 {
		return HWND(ret), errors.New("Unable to Find Window")
	}
	return HWND(ret), nil
}

// GetClientRect retrieves the coordinates of a window's client area. The client coordinates specify
// the upper-left and lower-right corners of the client area. Because client coordinates are
// relative to the upper-left corner of a window's client area, the coordinates of the upper-left
// corner are (0,0).
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms633503
func GetClientRect(hwnd HWND) *RECT {
	var rect RECT
	ret, _, _ := procGetClientRect.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&rect)))

	if ret == 0 {
		panic(fmt.Sprintf("GetClientRect(%d) failed", hwnd))
	}

	return &rect
}

// GetWindowRect retrieves the dimensions of the bounding rectangle of the specified window. The
// dimensions are given in screen coordinates that are relative to the upper-left corner of the
// screen.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms633519
func GetWindowRect(hwnd HWND) *RECT {
	var rect RECT
	procGetWindowRect.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&rect)))

	return &rect
}

// GetWindowText returns the text of the specified window's title bar (if it has one).
// If the specified window is a control, the text of the control is copied. However, GetWindowText
// cannot retrieve the text of a control in another application.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms633520
func GetWindowText(hwnd HWND) string {
	textLen := GetWindowTextLength(hwnd) + 1

	buf := make([]uint16, textLen)
	procGetWindowText.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(textLen))

	return syscall.UTF16ToString(buf)
}

// GetWindowTextLength retrieves the length, in characters, of the specified window's title bar text
// (if the window has a title bar). If the specified window is a control, the function retrieves the
// length of the text within the control. However, GetWindowTextLength cannot retrieve the length of
// the text of an edit control in another application.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms633521
func GetWindowTextLength(hwnd HWND) int {
	ret, _, _ := procGetWindowTextLength.Call(
		uintptr(hwnd))

	return int(ret)
}

// GetWindowThreadProcessId retrieves the identifier of the thread that created the specified window
// and, optionally, the identifier of the process that created the window.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms633522
func GetWindowThreadProcessId(hwnd HWND) (HANDLE, int) {
	var processId int
	ret, _, _ := procGetWindowThreadProcessId.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&processId)))
	return HANDLE(ret), processId
}

// IsWindow determines whether the specified window handle identifies an existing window.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms633528
func IsWindow(hwnd HWND) bool {
	ret, _, _ := procIsWindow.Call(
		uintptr(hwnd))
	return ret != 0
}

// IsWindowVisible determines the visibility state of the specified window.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms633530
func IsWindowVisible(hwnd HWND) bool {
	ret, _, _ := procIsWindowVisible.Call(
		uintptr(hwnd))
	return ret != 0
}

// MoveWindow changes the position and dimensions of the specified window. For a top-level window,
// the position and dimensions are relative to the upper-left corner of the screen. For a child
// window, they are relative to the upper-left corner of the parent window's client area.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms633534
func MoveWindow(hwnd HWND, x, y, width, height int, repaint bool) bool {
	ret, _, _ := procMoveWindow.Call(
		uintptr(hwnd),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(BoolToBOOL(repaint)))
	return ret != 0
}

// SetWindowPos changes the size, position, and Z order of a child, pop-up, or top-level window.
// These windows are ordered according to their appearance on the screen. The topmost window
// receives the highest rank and is the first window in the Z order.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms633545
func SetWindowPos(hwnd, hWndInsertAfter HWND, x, y, cx, cy int, uFlags uint) bool {
	ret, _, _ := procSetWindowPos.Call(
		uintptr(hwnd),
		uintptr(hWndInsertAfter),
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(uFlags))

	return ret != 0
}

// SetWindowText changes the text of the specified window's title bar (if it has one). If the
// specified window is a control, the text of the control is changed. However, SetWindowText cannot
// change the text of a control in another application.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms633546
func SetWindowText(hwnd HWND, text string) {
	procSetWindowText.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))))
}

// ShowWindow sets the specified window's show state.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms633548
func ShowWindow(hwnd HWND, cmdshow int) bool {
	ret, _, _ := procShowWindow.Call(
		uintptr(hwnd),
		uintptr(cmdshow))

	return ret != 0

}

// Window Classes
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms632596
// TODO: GetClassInfo
// TODO: GetClassInfoEx
// TODO: GetClassLong
// TODO: GetClassLongPtr
// TODO: GetClassName
// TODO: GetClassWord
// TODO: SetClassLong
// TODO: SetClassLongPtr
// TODO: SetClassWord
// TODO: UnregisterClass

// GetWindowLong retrieves information about the specified window. The function also retrieves the
// 32-bit (DWORD) value at the specified offset into the extra window memory.
//
// Note: If you are retrieving a pointer or a handle, this function has been superseded by the
// GetWindowLongPtr function. (Pointers and handles are 32 bits on 32-bit Windows and 64 bits on
// 64-bit Windows.) To write code that is compatible with both 32-bit and 64-bit versions of
// Windows, use GetWindowLongPtr.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms633584
func GetWindowLong(hwnd HWND, index int) int32 {
	ret, _, _ := procGetWindowLong.Call(
		uintptr(hwnd),
		uintptr(index))

	return int32(ret)
}

// Retrieves information about the specified window. The function also retrieves the value at a
// specified offset into the extra window memory.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms633585
func GetWindowLongPtr(hwnd HWND, index int) uintptr {
	ret, _, _ := procGetWindowLongPtr.Call(
		uintptr(hwnd),
		uintptr(index))

	return ret
}

// RegisterClassEx registers a window class for subsequent use in calls to the CreateWindow or
// CreateWindowEx function.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms633587
func RegisterClassEx(wndClassEx *WNDCLASSEX) ATOM {
	ret, _, _ := procRegisterClassEx.Call(uintptr(unsafe.Pointer(wndClassEx)))
	return ATOM(ret)
}

// SetWindowLong changes an attribute of the specified window. The function also sets the 32-bit
// (long) value at the specified offset into the extra window memory.
//
// Note: This function has been superseded by the SetWindowLongPtr function. To write code that is
// compatible with both 32-bit and 64-bit versions of Windows, use the SetWindowLongPtr function.
//
//
func SetWindowLong(hwnd HWND, index int, value uint32) uint32 {
	ret, _, _ := procSetWindowLong.Call(
		uintptr(hwnd),
		uintptr(index),
		uintptr(value))

	return uint32(ret)
}

// SetWindowLongPtr changes an attribute of the specified window. The function also sets a value at
// the specified offset in the extra window memory.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms644898
func SetWindowLongPtr(hwnd HWND, index int, value uintptr) uintptr {
	ret, _, _ := procSetWindowLongPtr.Call(
		uintptr(hwnd),
		uintptr(index),
		value)

	return ret
}

// Window Procedures
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms632593

// CallWindowProc passes message information to the specified window procedure.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms633571
func CallWindowProc(preWndProc uintptr, hwnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := procCallWindowProc.Call(
		preWndProc,
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)

	return ret
}

// DefWindowProc calls the default window procedure to provide default processing for any window
// messages that an application does not process. This function ensures that every message is
// processed. DefWindowProc is called with the same parameters received by the window procedure.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms633572
func DefWindowProc(hwnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := procDefWindowProc.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)

	return ret
}

// Messages and Message Queues
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms632590
// TODO: BroadcastSystemMessage
// TODO: BroadcastSystemMessageEx
// TODO: GetInputState
// TODO: GetMessageExtraInfo
// TODO: GetMessagePos
// TODO: GetMessageTime
// TODO: GetQueueStatus
// TODO: InSendMessage
// TODO: InSendMessageEx
// TODO: PostThreadMessage
// TODO: RegisterWindowMessage
// TODO: ReplyMessage
// TODO: SendAsyncProc
// TODO: SendMessageCallback
// TODO: SendMessageTimeout
// TODO: SendNotifyMessage
// TODO: SetMessageExtraInfo

// DispatchMessage dispatches a message to a window procedure. It is typically used to dispatch a
// message retrieved by the GetMessage function.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms644934
func DispatchMessage(msg *MSG) uintptr {
	ret, _, _ := procDispatchMessage.Call(
		uintptr(unsafe.Pointer(msg)))

	return ret

}

// GetMessage retrieves a message from the calling thread's message queue. The function dispatches
// incoming sent messages until a posted message is available for retrieval.
//
// Unlike GetMessage, the PeekMessage function does not wait for a message to be posted before
// returning.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms644936
func GetMessage(msg *MSG, hwnd HWND, msgFilterMin, msgFilterMax uint32) int {
	ret, _, _ := procGetMessage.Call(
		uintptr(unsafe.Pointer(msg)),
		uintptr(hwnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax))

	return int(ret)
}

// PeekMessage dispatches incoming sent messages, checks the thread message queue for a posted
// message, and retrieves the message (if any exist).
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms644943
func PeekMessage(lpMsg *MSG, hwnd HWND, wMsgFilterMin, wMsgFilterMax, wRemoveMsg uint32) bool {
	ret, _, _ := procPeekMessage.Call(
		uintptr(unsafe.Pointer(lpMsg)),
		uintptr(hwnd),
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
		uintptr(wRemoveMsg))

	return ret != 0
}

// PostMessage places (posts) a message in the message queue associated with the thread that created
// the specified window and returns without waiting for the thread to process the message.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms644944
func PostMessage(hwnd HWND, msg uint32, wParam, lParam uintptr) bool {
	ret, _, _ := procPostMessage.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)

	return ret != 0
}

// PostQuitMessage indicates to the system that a thread has made a request to terminate (quit). It
// is typically used in response to a WM_DESTROY message.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms644945
func PostQuitMessage(exitCode int) {
	procPostQuitMessage.Call(
		uintptr(exitCode))
}

// SendMessage sends the specified message to a window or windows. The SendMessage function calls
// the window procedure for the specified window and does not return until the window procedure has
// processed the message.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms644950
func SendMessage(hwnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := procSendMessage.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)

	return ret
}

// TranslateMessage translates virtual-key messages into character messages. The character messages
// are posted to the calling thread's message queue, to be read the next time the thread calls the
// GetMessage or PeekMessage function.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms644955
func TranslateMessage(msg *MSG) bool {
	ret, _, _ := procTranslateMessage.Call(
		uintptr(unsafe.Pointer(msg)))

	return ret != 0

}

// WaitMessage yields control to other threads when a thread has no other messages in its message
// queue. The WaitMessage function suspends the thread and does not return until a new message is
// placed in the thread's message queue.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms644956
func WaitMessage() bool {
	ret, _, _ := procWaitMessage.Call()
	return ret != 0
}

// Timers
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms632592

// Window Properties
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms632594

// Configuration
// http://msdn.microsoft.com/en-us/library/windows/desktop/ff625301

// Multiple Document Interface
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms632591

// Dialog Boxes
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms632588

// Carets
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms646968

// Cursors
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms646970

// Icons
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms646973
// TODO: CopyIcon
// TODO: CreateIconFromResource
// TODO: CreateIconFromResourceEx
// TODO: CreateIconIndirect
// TODO: DrawIconEx
// TODO: DuplicateIcon
// TODO: ExtractAssociatedIcon
// TODO: ExtractIcon
// TODO: ExtractIconEx
// TODO: GetIconInfo
// TODO: GetIconInfoEx
// TODO: LookupIconIdFromDirectory
// TODO: LookupIconIdFromDirectoryEx
// TODO: PrivateExtractIcons

// CreateIcon creates an icon that has the specified size, colors, and bit patterns.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms648059
func CreateIcon(instance HINSTANCE, nWidth, nHeight int, cPlanes, cBitsPerPixel byte, ANDbits, XORbits *byte) HICON {
	ret, _, _ := procCreateIcon.Call(
		uintptr(instance),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(cPlanes),
		uintptr(cBitsPerPixel),
		uintptr(unsafe.Pointer(ANDbits)),
		uintptr(unsafe.Pointer(XORbits)),
	)
	return HICON(ret)
}

// DestroyIcon destroys an icon and frees any memory the icon occupied.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms648063
func DestroyIcon(icon HICON) bool {
	ret, _, _ := procDestroyIcon.Call(
		uintptr(icon),
	)
	return ret != 0
}

// DrawIcon draws an icon or cursor into the specified device context.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms648064
func DrawIcon(hDC HDC, x, y int, hIcon HICON) bool {
	ret, _, _ := procDrawIcon.Call(
		uintptr(unsafe.Pointer(hDC)),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(hIcon)))

	return ret != 0
}

// LoadIcon loads the specified icon resource from the executable (.exe) file associated with an
// application instance.
//
// Note: This function has been superseded by the LoadImage function.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms648072
func LoadIcon(instance HINSTANCE, iconName *uint16) HICON {
	ret, _, _ := procLoadIcon.Call(
		uintptr(instance),
		uintptr(unsafe.Pointer(iconName)))

	return HICON(ret)

}

// Keyboard Accelerators
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms645526

// Menus
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms646977

// Strings - Possibly Useless
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms646979

// Version Information
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms646981

// Data Exchange
// Clipboard
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms648709

func LoadCursor(instance HINSTANCE, cursorName *uint16) HCURSOR {
	ret, _, _ := procLoadCursor.Call(
		uintptr(instance),
		uintptr(unsafe.Pointer(cursorName)))

	return HCURSOR(ret)

}

func UpdateWindow(hwnd HWND) bool {
	ret, _, _ := procUpdateWindow.Call(
		uintptr(hwnd))
	return ret != 0
}

func DefDlgProc(hwnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := procDefDlgProc.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)

	return ret
}

func ScreenToClient(hwnd HWND, x, y int) (X, Y int, ok bool) {
	pt := POINT{X: int32(x), Y: int32(y)}
	ret, _, _ := procScreenToClient.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&pt)))

	return int(pt.X), int(pt.Y), ret != 0
}

func InvalidateRect(hwnd HWND, rect *RECT, erase bool) bool {
	ret, _, _ := procInvalidateRect.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(rect)),
		uintptr(BoolToBOOL(erase)))

	return ret != 0
}

func SetCapture(hwnd HWND) HWND {
	ret, _, _ := procSetCapture.Call(
		uintptr(hwnd))

	return HWND(ret)
}

func ReleaseCapture() bool {
	ret, _, _ := procReleaseCapture.Call()

	return ret != 0
}

func MessageBox(hwnd HWND, title, caption string, flags uint) int {
	ret, _, _ := procMessageBox.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		uintptr(flags))

	return int(ret)
}

func GetSystemMetrics(index int) int {
	ret, _, _ := procGetSystemMetrics.Call(
		uintptr(index))

	return int(ret)
}

func CopyRect(dst, src *RECT) bool {
	ret, _, _ := procCopyRect.Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)))

	return ret != 0
}

func EqualRect(rect1, rect2 *RECT) bool {
	ret, _, _ := procEqualRect.Call(
		uintptr(unsafe.Pointer(rect1)),
		uintptr(unsafe.Pointer(rect2)))

	return ret != 0
}

func InflateRect(rect *RECT, dx, dy int) bool {
	ret, _, _ := procInflateRect.Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(dx),
		uintptr(dy))

	return ret != 0
}

func IntersectRect(dst, src1, src2 *RECT) bool {
	ret, _, _ := procIntersectRect.Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src1)),
		uintptr(unsafe.Pointer(src2)))

	return ret != 0
}

func IsRectEmpty(rect *RECT) bool {
	ret, _, _ := procIsRectEmpty.Call(
		uintptr(unsafe.Pointer(rect)))

	return ret != 0
}

func OffsetRect(rect *RECT, dx, dy int) bool {
	ret, _, _ := procOffsetRect.Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(dx),
		uintptr(dy))

	return ret != 0
}

func PtInRect(rect *RECT, x, y int) bool {
	pt := POINT{X: int32(x), Y: int32(y)}
	ret, _, _ := procPtInRect.Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(&pt)))

	return ret != 0
}

func SetRect(rect *RECT, left, top, right, bottom int) bool {
	ret, _, _ := procSetRect.Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(left),
		uintptr(top),
		uintptr(right),
		uintptr(bottom))

	return ret != 0
}

func SetRectEmpty(rect *RECT) bool {
	ret, _, _ := procSetRectEmpty.Call(
		uintptr(unsafe.Pointer(rect)))

	return ret != 0
}

func SubtractRect(dst, src1, src2 *RECT) bool {
	ret, _, _ := procSubtractRect.Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src1)),
		uintptr(unsafe.Pointer(src2)))

	return ret != 0
}

func UnionRect(dst, src1, src2 *RECT) bool {
	ret, _, _ := procUnionRect.Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src1)),
		uintptr(unsafe.Pointer(src2)))

	return ret != 0
}

func CreateDialog(hInstance HINSTANCE, lpTemplate *uint16, hWndParent HWND, lpDialogProc uintptr) HWND {
	ret, _, _ := procCreateDialogParam.Call(
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpTemplate)),
		uintptr(hWndParent),
		lpDialogProc,
		0)

	return HWND(ret)
}

func DialogBox(hInstance HINSTANCE, lpTemplateName *uint16, hWndParent HWND, lpDialogProc uintptr) int {
	ret, _, _ := procDialogBoxParam.Call(
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpTemplateName)),
		uintptr(hWndParent),
		lpDialogProc,
		0)

	return int(ret)
}

func GetDlgItem(hDlg HWND, nIDDlgItem int) HWND {
	ret, _, _ := procGetDlgItem.Call(
		uintptr(unsafe.Pointer(hDlg)),
		uintptr(nIDDlgItem))

	return HWND(ret)
}

func ClientToScreen(hwnd HWND, x, y int) (int, int) {
	pt := POINT{X: int32(x), Y: int32(y)}

	procClientToScreen.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&pt)))

	return int(pt.X), int(pt.Y)
}

func IsDialogMessage(hwnd HWND, msg *MSG) bool {
	ret, _, _ := procIsDialogMessage.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(msg)))

	return ret != 0
}

func EndDialog(hwnd HWND, nResult uintptr) bool {
	ret, _, _ := procEndDialog.Call(
		uintptr(hwnd),
		nResult)

	return ret != 0
}

func TranslateAccelerator(hwnd HWND, hAccTable HACCEL, lpMsg *MSG) bool {
	ret, _, _ := procTranslateMessage.Call(
		uintptr(hwnd),
		uintptr(hAccTable),
		uintptr(unsafe.Pointer(lpMsg)))

	return ret != 0
}

func FillRect(hDC HDC, lprc *RECT, hbr HBRUSH) bool {
	ret, _, _ := procFillRect.Call(
		uintptr(hDC),
		uintptr(unsafe.Pointer(lprc)),
		uintptr(hbr))

	return ret != 0
}

func DrawText(hDC HDC, text string, uCount int, lpRect *RECT, uFormat uint) int {
	ret, _, _ := procDrawText.Call(
		uintptr(hDC),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
		uintptr(uCount),
		uintptr(unsafe.Pointer(lpRect)),
		uintptr(uFormat))

	return int(ret)
}

func AddClipboardFormatListener(hwnd HWND) bool {
	ret, _, _ := procAddClipboardFormatListener.Call(
		uintptr(hwnd))
	return ret != 0
}

func RemoveClipboardFormatListener(hwnd HWND) bool {
	ret, _, _ := procRemoveClipboardFormatListener.Call(
		uintptr(hwnd))
	return ret != 0
}

func OpenClipboard(hWndNewOwner HWND) bool {
	ret, _, _ := procOpenClipboard.Call(
		uintptr(hWndNewOwner))
	return ret != 0
}

func CloseClipboard() bool {
	ret, _, _ := procCloseClipboard.Call()
	return ret != 0
}

func EnumClipboardFormats(format uint) uint {
	ret, _, _ := procEnumClipboardFormats.Call(
		uintptr(format))
	return uint(ret)
}

func GetClipboardData(uFormat uint) HANDLE {
	ret, _, _ := procGetClipboardData.Call(
		uintptr(uFormat))
	return HANDLE(ret)
}

func SetClipboardData(uFormat uint, hMem HANDLE) HANDLE {
	ret, _, _ := procSetClipboardData.Call(
		uintptr(uFormat),
		uintptr(hMem))
	return HANDLE(ret)
}

func EmptyClipboard() bool {
	ret, _, _ := procEmptyClipboard.Call()
	return ret != 0
}

func GetClipboardFormatName(format uint) (string, bool) {
	cchMaxCount := 255
	buf := make([]uint16, cchMaxCount)
	ret, _, _ := procGetClipboardFormatName.Call(
		uintptr(format),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(cchMaxCount))

	if ret > 0 {
		return syscall.UTF16ToString(buf), true
	}

	return "Requested format does not exist or is predefined", false
}

func IsClipboardFormatAvailable(format uint) bool {
	ret, _, _ := procIsClipboardFormatAvailable.Call(uintptr(format))
	return ret != 0
}

func BeginPaint(hwnd HWND, paint *PAINTSTRUCT) HDC {
	ret, _, _ := procBeginPaint.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(paint)))
	return HDC(ret)
}

func EndPaint(hwnd HWND, paint *PAINTSTRUCT) {
	procBeginPaint.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(paint)))
}

func SwapMouseButton(fSwap bool) bool {
	ret, _, _ := procSwapMouseButton.Call(
		uintptr(BoolToBOOL(fSwap)))
	return ret != 0
}

func GetCursorPos() (x, y int, ok bool) {
	pt := POINT{}
	ret, _, _ := procGetCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	return int(pt.X), int(pt.Y), ret != 0
}

func SetCursorPos(x, y int) bool {
	ret, _, _ := procSetCursorPos.Call(
		uintptr(x),
		uintptr(y),
	)
	return ret != 0
}

func SetCursor(cursor HCURSOR) HCURSOR {
	ret, _, _ := procSetCursor.Call(
		uintptr(cursor),
	)
	return HCURSOR(ret)
}

// Multiple Display Monitors
// http://msdn.microsoft.com/en-us/library/windows/desktop/dd145072
// TODO: MonitorEnumProc

// EnumDisplayMonitors enumerates display monitors (including invisible pseudo-monitors associated
// with the mirroring drivers) that intersect a region formed by the intersection of a specified
// clipping rectangle and the visible region of a device context. EnumDisplayMonitors calls an
// application-defined MonitorEnumProc callback function once for each monitor that is enumerated.
// Note that GetSystemMetrics (SM_CMONITORS) counts only the display monitors.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/dd162610
func EnumDisplayMonitors(hdc HDC, clip *RECT, fnEnum, dwData uintptr) bool {
	ret, _, _ := procEnumDisplayMonitors.Call(
		uintptr(hdc),
		uintptr(unsafe.Pointer(clip)),
		fnEnum,
		dwData,
	)
	return ret != 0
}

// GetMonitorInfo retrieves information about a display monitor.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/dd144901
func GetMonitorInfo(hMonitor HMONITOR, lmpi *MONITORINFO) bool {
	ret, _, _ := procGetMonitorInfo.Call(
		uintptr(hMonitor),
		uintptr(unsafe.Pointer(lmpi)),
	)
	return ret != 0
}

// MonitorFromPoint retrieves a handle to the display monitor that contains a specified point.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/dd145062
func MonitorFromPoint(x, y int, dwFlags uint32) HMONITOR {
	ret, _, _ := procMonitorFromPoint.Call(
		uintptr(x),
		uintptr(y),
		uintptr(dwFlags),
	)
	return HMONITOR(ret)
}

// MonitorFromRect retrieves a handle to the display monitor that has the largest area of
// intersection with a specified rectangle.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/dd145063
func MonitorFromRect(rc *RECT, dwFlags uint32) HMONITOR {
	ret, _, _ := procMonitorFromRect.Call(
		uintptr(unsafe.Pointer(rc)),
		uintptr(dwFlags),
	)
	return HMONITOR(ret)
}

// MonitorFromWindow retrieves a handle to the display monitor that has the largest area of
// intersection with the bounding rectangle of a specified window.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/dd145064
func MonitorFromWindow(hwnd HWND, dwFlags uint32) HMONITOR {
	ret, _, _ := procMonitorFromWindow.Call(
		uintptr(hwnd),
		uintptr(dwFlags),
	)
	return HMONITOR(ret)
}

// Device Context Functions
// http://msdn.microsoft.com/en-us/library/windows/desktop/dd183554
// TODO: CancelDC
// TODO: CreateCompatibleDC
// TODO: CreateDC
// TODO: CreateIC
// TODO: DeleteDC
// TODO: DeleteObject
// TODO: DrawEscape
// TODO: EnumDisplayDevices
// TODO: EnumObjects
// TODO: EnumObjectsProc
// TODO: GetCurrentObject
// TODO: GetDCBrushColor
// TODO: GetDCEx
// TODO: GetDCOrgEx
// TODO: GetDCPenColor
// TODO: GetDeviceCaps
// TODO: GetLayout
// TODO: GetObject
// TODO: GetStockObject
// TODO: ResetDC
// TODO: RestoreDC
// TODO: SaveDC
// TODO: SelectObject
// TODO: SetDCBrushColor
// TODO: SetDCPenColor
// TODO: SetLayout
// TODO: WindowFromDC

// ChangeDisplaySettingsEx changes the settings of the specified display device to the specified
// graphics mode.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/dd183413
func ChangeDisplaySettingsEx(szDeviceName *uint16, devMode *DEVMODE, hwnd HWND, dwFlags uint32, lParam uintptr) int32 {
	ret, _, _ := procChangeDisplaySettingsEx.Call(
		uintptr(unsafe.Pointer(szDeviceName)),
		uintptr(unsafe.Pointer(devMode)),
		uintptr(hwnd),
		uintptr(dwFlags),
		lParam,
	)
	return int32(ret)
}

// EnumDisplaySettingsEx retrieves information about one of the graphics modes for a display device.
// To retrieve information for all the graphics modes for a display device, make a series of calls
// to this function.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/dd162612
func EnumDisplaySettingsEx(szDeviceName *uint16, iModeNum uint32, devMode *DEVMODE, dwFlags uint32) bool {
	ret, _, _ := procEnumDisplaySettingsEx.Call(
		uintptr(unsafe.Pointer(szDeviceName)),
		uintptr(iModeNum),
		uintptr(unsafe.Pointer(devMode)),
		uintptr(dwFlags),
	)
	return ret != 0
}

// The GetDC function retrieves a handle to a device context (DC) for the client area of a specified
// window or for the entire screen. You can use the returned handle in subsequent GDI functions to
// draw in the DC. The device context is an opaque data structure, whose values are used internally
// by GDI.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/dd144871
func GetDC(hwnd HWND) HDC {
	ret, _, _ := procGetDC.Call(
		uintptr(hwnd))

	return HDC(ret)
}

// ReleaseDC releases a device context (DC), freeing it for use by other applications. The effect of
// the ReleaseDC function depends on the type of DC. It frees only common and window DCs. It has no
// effect on class or private DCs.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/dd162920
func ReleaseDC(hwnd HWND, hDC HDC) bool {
	ret, _, _ := procReleaseDC.Call(
		uintptr(hwnd),
		uintptr(hDC))

	return ret != 0
}

// Keyboard Input
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms645530
// TODO: ActivateKeyboardLayout
// TODO: BlockInput
// TODO: GetActiveWindow
// TODO: GetFocus
// TODO: GetKBCodePage
// TODO: GetKeyboardLayout
// TODO: GetKeyboardLayoutList
// TODO: GetKeyboardLayoutName
// TODO: GetKeyboardType
// TODO: GetKeyNameText
// TODO: GetKeyState
// TODO: GetLastInputInfo
// TODO: LoadKeyboardLayout
// TODO: OemKeyScan
// TODO: RegisterHotKey
// TODO: SetActiveWindow
// TODO: SetKeyboardState
// TODO: ToAsciiEx
// TODO: ToUnicode
// TODO: ToUnicodeEx
// TODO: UnloadKeyboardLayout
// TODO: UnregisterHotKey
// TODO: VkKeyScan
// TODO: VkKeyScanEx

// EnableWindow enables or disables mouse and keyboard input to the specified window or control.
// When input is disabled, the window does not receive input such as mouse clicks and key presses.
// When input is enabled, the window receives all input.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms646291
func EnableWindow(hwnd HWND, b bool) bool {
	ret, _, _ := procEnableWindow.Call(
		uintptr(hwnd),
		uintptr(BoolToBOOL(b)))
	return ret != 0
}

// GetAsyncKeyState determines whether a key is up or down at the time the function is called, and
// whether the key was pressed after a previous call to GetAsyncKeyState.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms646293
func GetAsyncKeyState(vKey int) uint16 {
	ret, _, _ := procGetAsyncKeyState.Call(uintptr(vKey))
	return uint16(ret)
}

// GetKeyboardState copies the status of the 256 virtual keys to lpKeyState buffer.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms646299
func GetKeyboardState(lpKeyState *[]byte) bool {
	ret, _, _ := procGetKeyboardState.Call(
		uintptr(unsafe.Pointer(&(*lpKeyState)[0])))
	return ret != 0
}

// IsWindowEnabled determines whether the specified window is enabled for mouse and keyboard input.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms646303
func IsWindowEnabled(hwnd HWND) bool {
	ret, _, _ := procIsWindowEnabled.Call(
		uintptr(hwnd))

	return ret != 0
}

// MapVirtualKeyEx translates (maps) a virtual-key code into a scan code or character value, or
// translates a scan code into a virtual-key code. The function translates the codes using the input
// language and an input locale identifier.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms646307
func MapVirtualKeyEx(uCode, uMapType uint, dwhkl HKL) uint {
	ret, _, _ := procMapVirtualKey.Call(
		uintptr(uCode),
		uintptr(uMapType),
		uintptr(dwhkl))
	return uint(ret)
}

// SendInput synthesizes keystrokes, mouse motions, and button clicks. Inputs are sent serially
// starting from beginning of slice. Returns number of successful inputs and the last error number
// in the set (if any).
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms646310
func SendInput(inputs []INPUT) uint32 {
	var validInputs []C.INPUT

	for _, oneInput := range inputs {
		input := C.INPUT{_type: C.DWORD(oneInput.Type)}

		switch oneInput.Type {
		case INPUT_MOUSE:
			(*MouseInput)(unsafe.Pointer(&input)).mi = oneInput.Mi
		case INPUT_KEYBOARD:
			(*KbdInput)(unsafe.Pointer(&input)).ki = oneInput.Ki
		case INPUT_HARDWARE:
			(*HardwareInput)(unsafe.Pointer(&input)).hi = oneInput.Hi
		default:
			panic("unkown type")
		}

		validInputs = append(validInputs, input)
	}

	ret, _, _ := procSendInput.Call(
		uintptr(len(validInputs)),
		uintptr(unsafe.Pointer(&validInputs[0])),
		uintptr(unsafe.Sizeof(C.INPUT{})),
	)
	return uint32(ret)
}

// SetFocus sets the keyboard focus to the specified window. The window must be attached to the
// calling thread's message queue.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms646312
func SetFocus(hwnd HWND) HWND {
	ret, _, _ := procSetFocus.Call(
		uintptr(hwnd))

	return HWND(ret)
}

// Translates the specified virtual-key code and keyboard state to the corresponding character or
// characters. The function translates the code using the input language and physical keyboard
// layout identified by the keyboard layout handle.
//
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms646316
func ToAscii(uVirtKey, uScanCode uint, lpKeyState *byte, lpChar *uint16, uFlags uint) int {
	ret, _, _ := procToAscii.Call(
		uintptr(uVirtKey),
		uintptr(uScanCode),
		uintptr(unsafe.Pointer(lpKeyState)),
		uintptr(unsafe.Pointer(lpChar)),
		uintptr(uFlags))
	return int(ret)
}
