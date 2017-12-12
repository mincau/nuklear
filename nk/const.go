// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Tue, 12 Dec 2017 20:09:32 MSK.
// By https://git.io/c-for-go. DO NOT EDIT.

package nk

/*
#include "nuklear.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// IncludeFixedTypes as defined in nuklear/<predefine>:24
	IncludeFixedTypes = 1
	// IncludeStandardIo as defined in nuklear/<predefine>:25
	IncludeStandardIo = 1
	// IncludeDefaultAllocator as defined in nuklear/<predefine>:26
	IncludeDefaultAllocator = 1
	// IncludeFontBaking as defined in nuklear/<predefine>:27
	IncludeFontBaking = 1
	// IncludeDefaultFont as defined in nuklear/<predefine>:28
	IncludeDefaultFont = 1
	// IncludeVertexBufferOutput as defined in nuklear/<predefine>:29
	IncludeVertexBufferOutput = 1
	// Undefined as defined in nk/nuklear.h:253
	Undefined = (-1.0)
	// UtfInvalid as defined in nk/nuklear.h:254
	UtfInvalid = 0xFFFD
	// UtfSize as defined in nk/nuklear.h:255
	UtfSize = 4
	// InputMax as defined in nk/nuklear.h:257
	InputMax = 16
	// MaxNumberBuffer as defined in nk/nuklear.h:260
	MaxNumberBuffer = 64
	// ScrollbarHidingTimeout as defined in nk/nuklear.h:263
	ScrollbarHidingTimeout = 4.0
	// TexteditUndostatecount as defined in nk/nuklear.h:2812
	TexteditUndostatecount = 99
	// TexteditUndocharcount as defined in nk/nuklear.h:2816
	TexteditUndocharcount = 999
	// MaxLayoutRowTemplateColumns as defined in nk/nuklear.h:3817
	MaxLayoutRowTemplateColumns = 16
	// ChartMaxSlot as defined in nk/nuklear.h:3820
	ChartMaxSlot = 4
	// WindowMaxName as defined in nk/nuklear.h:3918
	WindowMaxName = 64
	// ButtonBehaviorStackSize as defined in nk/nuklear.h:4031
	ButtonBehaviorStackSize = 8
	// FontStackSize as defined in nk/nuklear.h:4035
	FontStackSize = 8
	// StyleItemStackSize as defined in nk/nuklear.h:4039
	StyleItemStackSize = 16
	// FloatStackSize as defined in nk/nuklear.h:4043
	FloatStackSize = 32
	// VectorStackSize as defined in nk/nuklear.h:4047
	VectorStackSize = 16
	// FlagsStackSize as defined in nk/nuklear.h:4051
	FlagsStackSize = 32
	// ColorStackSize as defined in nk/nuklear.h:4055
	ColorStackSize = 32
	// Float as defined in nk/nuklear.h:4069
	Float = 0
	// Pi as defined in nk/nuklear.h:4183
	Pi = 3.141592654
	// MaxFloatPrecision as defined in nk/nuklear.h:4185
	MaxFloatPrecision = 2
)

// Heading as declared in nk/nuklear.h:465
type Heading int32

// Heading enumeration from nk/nuklear.h:465
const (
	Up    = iota
	Right = 1
	Down  = 2
	Left  = 3
)

// ButtonBehavior as declared in nk/nuklear.h:466
type ButtonBehavior int32

// ButtonBehavior enumeration from nk/nuklear.h:466
const (
	ButtonDefault  = iota
	ButtonRepeater = 1
)

// Modify as declared in nk/nuklear.h:467
type Modify int32

// Modify enumeration from nk/nuklear.h:467
const (
	Fixed      = False
	Modifiable = True
)

// Orientation as declared in nk/nuklear.h:468
type Orientation int32

// Orientation enumeration from nk/nuklear.h:468
const (
	Vertical   = iota
	Horizontal = 1
)

// CollapseStates as declared in nk/nuklear.h:469
type CollapseStates int32

// CollapseStates enumeration from nk/nuklear.h:469
const (
	Minimized = False
	Maximized = True
)

// ShowStates as declared in nk/nuklear.h:470
type ShowStates int32

// ShowStates enumeration from nk/nuklear.h:470
const (
	Hidden = False
	Shown  = True
)

// ChartType as declared in nk/nuklear.h:471
type ChartType int32

// ChartType enumeration from nk/nuklear.h:471
const (
	ChartLines  = iota
	ChartColumn = 1
	ChartMax    = 2
)

// ChartEvent as declared in nk/nuklear.h:472
type ChartEvent int32

// ChartEvent enumeration from nk/nuklear.h:472
const (
	ChartHovering = 0x01
	ChartClicked  = 0x02
)

// ColorFormat as declared in nk/nuklear.h:473
type ColorFormat int32

// ColorFormat enumeration from nk/nuklear.h:473
const (
	ColorFormatRGB  = iota
	ColorFormatRGBA = 1
)

// PopupType as declared in nk/nuklear.h:474
type PopupType int32

// PopupType enumeration from nk/nuklear.h:474
const (
	PopupStatic  = iota
	PopupDynamic = 1
)

// LayoutFormat as declared in nk/nuklear.h:475
type LayoutFormat int32

// LayoutFormat enumeration from nk/nuklear.h:475
const (
	Dynamic = iota
	Static  = 1
)

// TreeType as declared in nk/nuklear.h:476
type TreeType int32

// TreeType enumeration from nk/nuklear.h:476
const (
	TreeNode = iota
	TreeTab  = 1
)

// SymbolType as declared in nk/nuklear.h:489
type SymbolType int32

// SymbolType enumeration from nk/nuklear.h:489
const (
	SymbolNone          = iota
	SymbolX             = 1
	SymbolUnderscore    = 2
	SymbolCircleSolid   = 3
	SymbolCircleOutline = 4
	SymbolRectSolid     = 5
	SymbolRectOutline   = 6
	SymbolTriangleUp    = 7
	SymbolTriangleDown  = 8
	SymbolTriangleLeft  = 9
	SymbolTriangleRight = 10
	SymbolPlus          = 11
	SymbolMinus         = 12
	SymbolMax           = 13
)

// Keys as declared in nk/nuklear.h:660
type Keys int32

// Keys enumeration from nk/nuklear.h:660
const (
	KeyNone            = iota
	KeyShift           = 1
	KeyCtrl            = 2
	KeyDel             = 3
	KeyEnter           = 4
	KeyTab             = 5
	KeyBackspace       = 6
	KeyCopy            = 7
	KeyCut             = 8
	KeyPaste           = 9
	KeyUp              = 10
	KeyDown            = 11
	KeyLeft            = 12
	KeyRight           = 13
	KeyTextInsertMode  = 14
	KeyTextReplaceMode = 15
	KeyTextResetMode   = 16
	KeyTextLineStart   = 17
	KeyTextLineEnd     = 18
	KeyTextStart       = 19
	KeyTextEnd         = 20
	KeyTextUndo        = 21
	KeyTextRedo        = 22
	KeyTextSelectAll   = 23
	KeyTextWordLeft    = 24
	KeyTextWordRight   = 25
	KeyScrollStart     = 26
	KeyScrollEnd       = 27
	KeyScrollDown      = 28
	KeyScrollUp        = 29
	KeyMax             = 30
)

// Buttons as declared in nk/nuklear.h:695
type Buttons int32

// Buttons enumeration from nk/nuklear.h:695
const (
	ButtonLeft   = iota
	ButtonMiddle = 1
	ButtonRight  = 2
	ButtonDouble = 3
	ButtonMax    = 4
)

// AntiAliasing as declared in nk/nuklear.h:969
type AntiAliasing int32

// AntiAliasing enumeration from nk/nuklear.h:969
const (
	AntiAliasingOff = iota
	AntiAliasingOn  = 1
)

// ConvertResult as declared in nk/nuklear.h:970
type ConvertResult int32

// ConvertResult enumeration from nk/nuklear.h:970
const (
	ConvertSuccess           = iota
	ConvertInvalidParam      = 1
	ConvertCommandBufferFull = (1 << (1))
	ConvertVertexBufferFull  = (1 << (2))
	ConvertElementBufferFull = (1 << (3))
)

// PanelFlags as declared in nk/nuklear.h:1173
type PanelFlags int32

// PanelFlags enumeration from nk/nuklear.h:1173
const (
	WindowBorder         = (1 << (0))
	WindowMovable        = (1 << (1))
	WindowScalable       = (1 << (2))
	WindowClosable       = (1 << (3))
	WindowMinimizable    = (1 << (4))
	WindowNoScrollbar    = (1 << (5))
	WindowTitle          = (1 << (6))
	WindowScrollAutoHide = (1 << (7))
	WindowBackground     = (1 << (8))
	WindowScaleLeft      = (1 << (9))
	WindowNoInput        = (1 << (10))
)

// WidgetLayoutStates as declared in nk/nuklear.h:1833
type WidgetLayoutStates int32

// WidgetLayoutStates enumeration from nk/nuklear.h:1833
const (
	WidgetInvalid = iota
	WidgetValid   = 1
	WidgetRom     = 2
)

// WidgetStates as declared in nk/nuklear.h:1838
type WidgetStates int32

// WidgetStates enumeration from nk/nuklear.h:1838
const (
	WidgetStateModified = (1 << (1))
	WidgetStateInactive = (1 << (2))
	WidgetStateEntered  = (1 << (3))
	WidgetStateHover    = (1 << (4))
	WidgetStateActived  = (1 << (5))
	WidgetStateLeft     = (1 << (6))
	WidgetStateHovered  = WidgetStateHover | WidgetStateModified
	WidgetStateActive   = WidgetStateActived | WidgetStateModified
)

// TextAlign as declared in nk/nuklear.h:1864
type TextAlign int32

// TextAlign enumeration from nk/nuklear.h:1864
const (
	TextAlignLeft     = 0x01
	TextAlignCentered = 0x02
	TextAlignRight    = 0x04
	TextAlignTop      = 0x08
	TextAlignMiddle   = 0x10
	TextAlignBottom   = 0x20
)

// TextAlignment as declared in nk/nuklear.h:1872
type TextAlignment int32

// TextAlignment enumeration from nk/nuklear.h:1872
const (
	TextLeft     = TextAlignMiddle | TextAlignLeft
	TextCentered = TextAlignMiddle | TextAlignCentered
	TextRight    = TextAlignMiddle | TextAlignRight
)

// EditFlags as declared in nk/nuklear.h:1999
type EditFlags int32

// EditFlags enumeration from nk/nuklear.h:1999
const (
	EditDefault            = iota
	EditReadOnly           = (1 << (0))
	EditAutoSelect         = (1 << (1))
	EditSigEnter           = (1 << (2))
	EditAllowTab           = (1 << (3))
	EditNoCursor           = (1 << (4))
	EditSelectable         = (1 << (5))
	EditClipboard          = (1 << (6))
	EditCtrlEnterNewline   = (1 << (7))
	EditNoHorizontalScroll = (1 << (8))
	EditAlwaysInsertMode   = (1 << (9))
	EditMultiline          = (1 << (10))
	EditGotoEndOnActivate  = (1 << (11))
)

// EditTypes as declared in nk/nuklear.h:2014
type EditTypes int32

// EditTypes enumeration from nk/nuklear.h:2014
const (
	EditSimple = EditAlwaysInsertMode
	EditField  = EditSimple | EditSelectable | EditClipboard
	EditBox    = EditAlwaysInsertMode | EditSelectable | EditMultiline | EditAllowTab | EditClipboard
	EditEditor = EditSelectable | EditMultiline | EditAllowTab | EditClipboard
)

// EditEvents as declared in nk/nuklear.h:2020
type EditEvents int32

// EditEvents enumeration from nk/nuklear.h:2020
const (
	EditActive      = (1 << (0))
	EditInactive    = (1 << (1))
	EditActivated   = (1 << (2))
	EditDeactivated = (1 << (3))
	EditCommited    = (1 << (4))
)

// StyleColors as declared in nk/nuklear.h:2142
type StyleColors int32

// StyleColors enumeration from nk/nuklear.h:2142
const (
	ColorText                  = iota
	ColorWindow                = 1
	ColorHeader                = 2
	ColorBorder                = 3
	ColorButton                = 4
	ColorButtonHover           = 5
	ColorButtonActive          = 6
	ColorToggle                = 7
	ColorToggleHover           = 8
	ColorToggleCursor          = 9
	ColorSelect                = 10
	ColorSelectActive          = 11
	ColorSlider                = 12
	ColorSliderCursor          = 13
	ColorSliderCursorHover     = 14
	ColorSliderCursorActive    = 15
	ColorProperty              = 16
	ColorEdit                  = 17
	ColorEditCursor            = 18
	ColorCombo                 = 19
	ColorChart                 = 20
	ColorChartColor            = 21
	ColorChartColorHighlight   = 22
	ColorScrollbar             = 23
	ColorScrollbarCursor       = 24
	ColorScrollbarCursorHover  = 25
	ColorScrollbarCursorActive = 26
	ColorTabHeader             = 27
	ColorCount                 = 28
)

// StyleCursor as declared in nk/nuklear.h:2173
type StyleCursor int32

// StyleCursor enumeration from nk/nuklear.h:2173
const (
	CursorArrow                  = iota
	CursorText                   = 1
	CursorMove                   = 2
	CursorResizeVertical         = 3
	CursorResizeHorizontal       = 4
	CursorResizeTopLeftDownRight = 5
	CursorResizeTopRightDownLeft = 6
	CursorCount                  = 7
)

// FontCoordType as declared in nk/nuklear.h:2503
type FontCoordType int32

// FontCoordType enumeration from nk/nuklear.h:2503
const (
	CoordUv    = iota
	CoordPixel = 1
)

// FontAtlasFormat as declared in nk/nuklear.h:2577
type FontAtlasFormat int32

// FontAtlasFormat enumeration from nk/nuklear.h:2577
const (
	FontAtlasAlpha8 = iota
	FontAtlasRgba32 = 1
)

// AllocationType as declared in nk/nuklear.h:2674
type AllocationType int32

// AllocationType enumeration from nk/nuklear.h:2674
const (
	BufferFixed   = iota
	BufferDynamic = 1
)

// BufferAllocationType as declared in nk/nuklear.h:2679
type BufferAllocationType int32

// BufferAllocationType enumeration from nk/nuklear.h:2679
const (
	BufferFront = iota
	BufferBack  = 1
	BufferMax   = 2
)

// TextEditType as declared in nk/nuklear.h:2842
type TextEditType int32

// TextEditType enumeration from nk/nuklear.h:2842
const (
	TextEditSingleLine = iota
	TextEditMultiLine  = 1
)

// TextEditMode as declared in nk/nuklear.h:2847
type TextEditMode int32

// TextEditMode enumeration from nk/nuklear.h:2847
const (
	TextEditModeView    = iota
	TextEditModeInsert  = 1
	TextEditModeReplace = 2
)

// CommandType as declared in nk/nuklear.h:2947
type CommandType int32

// CommandType enumeration from nk/nuklear.h:2947
const (
	CommandTypeNop            = iota
	CommandTypeScissor        = 1
	CommandTypeLine           = 2
	CommandTypeCurve          = 3
	CommandTypeRect           = 4
	CommandTypeRectFilled     = 5
	CommandTypeRectMultiColor = 6
	CommandTypeCircle         = 7
	CommandTypeCircleFilled   = 8
	CommandTypeArc            = 9
	CommandTypeArcFilled      = 10
	CommandTypeTriangle       = 11
	CommandTypeTriangleFilled = 12
	CommandTypePolygon        = 13
	CommandTypePolygonFilled  = 14
	CommandTypePolyline       = 15
	CommandTypeText           = 16
	CommandTypeImage          = 17
	CommandTypeCustom         = 18
)

// CommandClipping as declared in nk/nuklear.h:3130
type CommandClipping int32

// CommandClipping enumeration from nk/nuklear.h:3130
const (
	ClippingOff = False
	ClippingOn  = True
)

// DrawListStroke as declared in nk/nuklear.h:3239
type DrawListStroke int32

// DrawListStroke enumeration from nk/nuklear.h:3239
const (
	StrokeOpen   = False
	StrokeClosed = True
)

// DrawVertexLayoutAttribute as declared in nk/nuklear.h:3246
type DrawVertexLayoutAttribute int32

// DrawVertexLayoutAttribute enumeration from nk/nuklear.h:3246
const (
	VertexPosition       = iota
	VertexColor          = 1
	VertexTexcoord       = 2
	VertexAttributeCount = 3
)

// DrawVertexLayoutFormat as declared in nk/nuklear.h:3253
type DrawVertexLayoutFormat int32

// DrawVertexLayoutFormat enumeration from nk/nuklear.h:3253
const (
	FormatSchar              = iota
	FormatSshort             = 1
	FormatSint               = 2
	FormatUchar              = 3
	FormatUshort             = 4
	FormatUint               = 5
	FormatFloat              = 6
	FormatDouble             = 7
	FormatColorBegin         = 8
	FormatR8g8b8             = FormatColorBegin
	FormatR16g15b16          = 9
	FormatR32g32b32          = 10
	FormatR8g8b8a8           = 11
	FormatB8g8r8a8           = 12
	FormatR16g15b16a16       = 13
	FormatR32g32b32a32       = 14
	FormatR32g32b32a32Float  = 15
	FormatR32g32b32a32Double = 16
	FormatRgb32              = 17
	FormatRgba32             = 18
	FormatColorEnd           = FormatRgba32
	FormatCount              = 19
)

// StyleItemType as declared in nk/nuklear.h:3376
type StyleItemType int32

// StyleItemType enumeration from nk/nuklear.h:3376
const (
	StyleItemColor = iota
	StyleItemImage = 1
)

// StyleHeaderAlign as declared in nk/nuklear.h:3717
type StyleHeaderAlign int32

// StyleHeaderAlign enumeration from nk/nuklear.h:3717
const (
	HeaderLeft  = iota
	HeaderRight = 1
)

// PanelType as declared in nk/nuklear.h:3823
type PanelType int32

// PanelType enumeration from nk/nuklear.h:3823
const (
	PanelWindow     = (1 << (0))
	PanelGroup      = (1 << (1))
	PanelPopup      = (1 << (2))
	PanelContextual = (1 << (4))
	PanelCombo      = (1 << (5))
	PanelMenu       = (1 << (6))
	PanelTooltip    = (1 << (7))
)

// PanelSet as declared in nk/nuklear.h:3832
type PanelSet int32

// PanelSet enumeration from nk/nuklear.h:3832
const (
	PanelSetNonblock = PanelContextual | PanelCombo | PanelMenu | PanelTooltip
	PanelSetPopup    = PanelSetNonblock | PanelPopup
	PanelSetSub      = PanelSetPopup | PanelGroup
)

// PanelRowLayoutType as declared in nk/nuklear.h:3854
type PanelRowLayoutType int32

// PanelRowLayoutType enumeration from nk/nuklear.h:3854
const (
	LayoutDynamicFixed = iota
	LayoutDynamicRow   = 1
	LayoutDynamicFree  = 2
	LayoutDynamic      = 3
	LayoutStaticFixed  = 4
	LayoutStaticRow    = 5
	LayoutStaticFree   = 6
	LayoutStatic       = 7
	LayoutTemplate     = 8
	LayoutCount        = 9
)

// WindowFlags as declared in nk/nuklear.h:3922
type WindowFlags int32

// WindowFlags enumeration from nk/nuklear.h:3922
const (
	WindowPrivate        = (1 << (11))
	WindowDynamic        = WindowPrivate
	WindowRom            = (1 << (12))
	WindowNotInteractive = WindowRom | WindowNoInput
	WindowHidden         = (1 << (13))
	WindowClosed         = (1 << (14))
	WindowMinimized      = (1 << (15))
	WindowRemoveRom      = (1 << (16))
)

const (
	// False as declared in nk/nuklear.h:452
	False = iota
	// True as declared in nk/nuklear.h:452
	True = 1
)
