package TextEdit

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A multiline text editor. It also has limited facilities for editing code, such as syntax highlighting support. For more advanced facilities for editing code, see [CodeEdit].
[b]Note:[/b] Most viewport, caret, and edit methods contain a [code]caret_index[/code] argument for [member caret_multiple] support. The argument should be one of the following: [code]-1[/code] for all carets, [code]0[/code] for the main caret, or greater than [code]0[/code] for secondary carets in the order they were created.
[b]Note:[/b] When holding down [kbd]Alt[/kbd], the vertical scroll wheel will scroll 5 times as fast as it would normally do. This also works in the Godot script editor.
	// TextEdit methods that can be overridden by a [Class] that extends it.
	type TextEdit interface {
		//Override this method to define what happens when the user types in the provided key [param unicode_char].
		HandleUnicodeInput(unicode_char int, caret_index int) 
		//Override this method to define what happens when the user presses the backspace key.
		Backspace(caret_index int) 
		//Override this method to define what happens when the user performs a cut operation.
		Cut(caret_index int) 
		//Override this method to define what happens when the user performs a copy operation.
		Copy(caret_index int) 
		//Override this method to define what happens when the user performs a paste operation.
		Paste(caret_index int) 
		//Override this method to define what happens when the user performs a paste operation with middle mouse button.
		//[b]Note:[/b] This method is only implemented on Linux.
		PastePrimaryClipboard(caret_index int) 
	}

*/
type Go [1]classdb.TextEdit

/*
Override this method to define what happens when the user types in the provided key [param unicode_char].
*/
func (Go) _handle_unicode_input(impl func(ptr unsafe.Pointer, unicode_char int, caret_index int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var unicode_char = gd.UnsafeGet[gd.Int](p_args,0)
		var caret_index = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, int(unicode_char), int(caret_index))
		gc.End()
	}
}

/*
Override this method to define what happens when the user presses the backspace key.
*/
func (Go) _backspace(impl func(ptr unsafe.Pointer, caret_index int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var caret_index = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, int(caret_index))
		gc.End()
	}
}

/*
Override this method to define what happens when the user performs a cut operation.
*/
func (Go) _cut(impl func(ptr unsafe.Pointer, caret_index int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var caret_index = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, int(caret_index))
		gc.End()
	}
}

/*
Override this method to define what happens when the user performs a copy operation.
*/
func (Go) _copy(impl func(ptr unsafe.Pointer, caret_index int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var caret_index = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, int(caret_index))
		gc.End()
	}
}

/*
Override this method to define what happens when the user performs a paste operation.
*/
func (Go) _paste(impl func(ptr unsafe.Pointer, caret_index int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var caret_index = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, int(caret_index))
		gc.End()
	}
}

/*
Override this method to define what happens when the user performs a paste operation with middle mouse button.
[b]Note:[/b] This method is only implemented on Linux.
*/
func (Go) _paste_primary_clipboard(impl func(ptr unsafe.Pointer, caret_index int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var caret_index = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, int(caret_index))
		gc.End()
	}
}

/*
Returns [code]true[/code] if the user has text in the [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] (IME).
*/
func (self Go) HasImeText() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasImeText())
}

/*
Closes the [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] (IME) if it is open. Any text in the IME will be lost.
*/
func (self Go) CancelIme() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).CancelIme()
}

/*
Applies text from the [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] (IME) to each caret and closes the IME if it is open.
*/
func (self Go) ApplyIme() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ApplyIme()
}

/*
Sets the tab size for the [TextEdit] to use.
*/
func (self Go) SetTabSize(size int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTabSize(gd.Int(size))
}

/*
Returns the [TextEdit]'s' tab size.
*/
func (self Go) GetTabSize() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetTabSize()))
}

/*
If [code]true[/code], sets the user into overtype mode. When the user types in this mode, it will override existing text.
*/
func (self Go) SetOvertypeModeEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOvertypeModeEnabled(enabled)
}

/*
Returns whether the user is in overtype mode.
*/
func (self Go) IsOvertypeModeEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsOvertypeModeEnabled())
}

/*
Performs a full reset of [TextEdit], including undo history.
*/
func (self Go) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Clear()
}

/*
Returns the number of lines in the text.
*/
func (self Go) GetLineCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetLineCount()))
}

/*
Sets the text for a specific [param line].
Carets on the line will attempt to keep their visual x position.
*/
func (self Go) SetLine(line int, new_text string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLine(gd.Int(line), gc.String(new_text))
}

/*
Returns the text of a specific line.
*/
func (self Go) GetLine(line int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetLine(gc, gd.Int(line)).String())
}

/*
Returns the width in pixels of the [param wrap_index] on [param line].
*/
func (self Go) GetLineWidth(line int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetLineWidth(gd.Int(line), gd.Int(-1))))
}

/*
Returns the maximum value of the line height among all lines.
[b]Note:[/b] The return value is influenced by [theme_item line_spacing] and [theme_item font_size]. And it will not be less than [code]1[/code].
*/
func (self Go) GetLineHeight() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetLineHeight()))
}

/*
Returns the number of spaces and [code]tab * tab_size[/code] before the first char.
*/
func (self Go) GetIndentLevel(line int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetIndentLevel(gd.Int(line))))
}

/*
Returns the first column containing a non-whitespace character.
*/
func (self Go) GetFirstNonWhitespaceColumn(line int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetFirstNonWhitespaceColumn(gd.Int(line))))
}

/*
Swaps the two lines. Carets will be swapped with the lines.
*/
func (self Go) SwapLines(from_line int, to_line int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SwapLines(gd.Int(from_line), gd.Int(to_line))
}

/*
Inserts a new line with [param text] at [param line].
*/
func (self Go) InsertLineAt(line int, text string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).InsertLineAt(gd.Int(line), gc.String(text))
}

/*
Removes the line of text at [param line]. Carets on this line will attempt to match their previous visual x position.
If [param move_carets_down] is [code]true[/code] carets will move to the next line down, otherwise carets will move up.
*/
func (self Go) RemoveLineAt(line int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveLineAt(gd.Int(line), true)
}

/*
Insert the specified text at the caret position.
*/
func (self Go) InsertTextAtCaret(text string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).InsertTextAtCaret(gc.String(text), gd.Int(-1))
}

/*
Inserts the [param text] at [param line] and [param column].
If [param before_selection_begin] is [code]true[/code], carets and selections that begin at [param line] and [param column] will moved to the end of the inserted text, along with all carets after it.
If [param before_selection_end] is [code]true[/code], selections that end at [param line] and [param column] will be extended to the end of the inserted text. These parameters can be used to insert text inside of or outside of selections.
*/
func (self Go) InsertText(text string, line int, column int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).InsertText(gc.String(text), gd.Int(line), gd.Int(column), true, false)
}

/*
Removes text between the given positions.
*/
func (self Go) RemoveText(from_line int, from_column int, to_line int, to_column int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveText(gd.Int(from_line), gd.Int(from_column), gd.Int(to_line), gd.Int(to_column))
}

/*
Returns the last unhidden line in the entire [TextEdit].
*/
func (self Go) GetLastUnhiddenLine() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetLastUnhiddenLine()))
}

/*
Returns the count to the next visible line from [param line] to [code]line + visible_amount[/code]. Can also count backwards. For example if a [TextEdit] has 5 lines with lines 2 and 3 hidden, calling this with [code]line = 1, visible_amount = 1[/code] would return 3.
*/
func (self Go) GetNextVisibleLineOffsetFrom(line int, visible_amount int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetNextVisibleLineOffsetFrom(gd.Int(line), gd.Int(visible_amount))))
}

/*
Similar to [method get_next_visible_line_offset_from], but takes into account the line wrap indexes. In the returned vector, [code]x[/code] is the line, [code]y[/code] is the wrap index.
*/
func (self Go) GetNextVisibleLineIndexOffsetFrom(line int, wrap_index int, visible_amount int) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(class(self).GetNextVisibleLineIndexOffsetFrom(gd.Int(line), gd.Int(wrap_index), gd.Int(visible_amount)))
}

/*
Called when the user presses the backspace key. Can be overridden with [method _backspace].
*/
func (self Go) Backspace() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Backspace(gd.Int(-1))
}

/*
Cut's the current selection. Can be overridden with [method _cut].
*/
func (self Go) Cut() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Cut(gd.Int(-1))
}

/*
Copies the current text selection. Can be overridden with [method _copy].
*/
func (self Go) Copy() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Copy(gd.Int(-1))
}

/*
Paste at the current location. Can be overridden with [method _paste].
*/
func (self Go) Paste() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Paste(gd.Int(-1))
}

/*
Pastes the primary clipboard.
*/
func (self Go) PastePrimaryClipboard() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PastePrimaryClipboard(gd.Int(-1))
}

/*
Starts an action, will end the current action if [param action] is different.
An action will also end after a call to [method end_action], after [member ProjectSettings.gui/timers/text_edit_idle_detect_sec] is triggered or a new undoable step outside the [method start_action] and [method end_action] calls.
*/
func (self Go) StartAction(action classdb.TextEditEditAction) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).StartAction(action)
}

/*
Marks the end of steps in the current action started with [method start_action].
*/
func (self Go) EndAction() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).EndAction()
}

/*
Starts a multipart edit. All edits will be treated as one action until [method end_complex_operation] is called.
*/
func (self Go) BeginComplexOperation() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).BeginComplexOperation()
}

/*
Ends a multipart edit, started with [method begin_complex_operation]. If called outside a complex operation, the current operation is pushed onto the undo/redo stack.
*/
func (self Go) EndComplexOperation() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).EndComplexOperation()
}

/*
Returns [code]true[/code] if an "undo" action is available.
*/
func (self Go) HasUndo() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasUndo())
}

/*
Returns [code]true[/code] if a "redo" action is available.
*/
func (self Go) HasRedo() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasRedo())
}

/*
Perform undo operation.
*/
func (self Go) Undo() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Undo()
}

/*
Perform redo operation.
*/
func (self Go) Redo() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Redo()
}

/*
Clears the undo history.
*/
func (self Go) ClearUndoHistory() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearUndoHistory()
}

/*
Tag the current version as saved.
*/
func (self Go) TagSavedVersion() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).TagSavedVersion()
}

/*
Returns the current version of the [TextEdit]. The version is a count of recorded operations by the undo/redo history.
*/
func (self Go) GetVersion() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetVersion()))
}

/*
Returns the last tagged saved version from [method tag_saved_version].
*/
func (self Go) GetSavedVersion() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetSavedVersion()))
}

/*
Sets the search text. See [method set_search_flags].
*/
func (self Go) SetSearchText(search_text string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSearchText(gc.String(search_text))
}

/*
Sets the search [param flags]. This is used with [method set_search_text] to highlight occurrences of the searched text. Search flags can be specified from the [enum SearchFlags] enum.
*/
func (self Go) SetSearchFlags(flags int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSearchFlags(gd.Int(flags))
}

/*
Perform a search inside the text. Search flags can be specified in the [enum SearchFlags] enum.
In the returned vector, [code]x[/code] is the column, [code]y[/code] is the line. If no results are found, both are equal to [code]-1[/code].
[codeblocks]
[gdscript]
var result = search("print", SEARCH_WHOLE_WORDS, 0, 0)
if result.x != -1:
    # Result found.
    var line_number = result.y
    var column_number = result.x
[/gdscript]
[csharp]
Vector2I result = Search("print", (uint)TextEdit.SearchFlags.WholeWords, 0, 0);
if (result.X != -1)
{
    // Result found.
    int lineNumber = result.Y;
    int columnNumber = result.X;
}
[/csharp]
[/codeblocks]
*/
func (self Go) Search(text string, flags int, from_line int, from_column int) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(class(self).Search(gc.String(text), gd.Int(flags), gd.Int(from_line), gd.Int(from_column)))
}

/*
Provide custom tooltip text. The callback method must take the following args: [code]hovered_word: String[/code].
*/
func (self Go) SetTooltipRequestFunc(callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTooltipRequestFunc(callback)
}

/*
Returns the local mouse position adjusted for the text direction.
*/
func (self Go) GetLocalMousePos() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetLocalMousePos())
}

/*
Returns the word at [param position].
*/
func (self Go) GetWordAtPos(position gd.Vector2) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetWordAtPos(gc, position).String())
}

/*
Returns the line and column at the given position. In the returned vector, [code]x[/code] is the column, [code]y[/code] is the line. If [param allow_out_of_bounds] is [code]false[/code] and the position is not over the text, both vector values will be set to [code]-1[/code].
*/
func (self Go) GetLineColumnAtPos(position gd.Vector2i) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(class(self).GetLineColumnAtPos(position, true))
}

/*
Returns the local position for the given [param line] and [param column]. If [code]x[/code] or [code]y[/code] of the returned vector equal [code]-1[/code], the position is outside of the viewable area of the control.
[b]Note:[/b] The Y position corresponds to the bottom side of the line. Use [method get_rect_at_line_column] to get the top side position.
*/
func (self Go) GetPosAtLineColumn(line int, column int) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(class(self).GetPosAtLineColumn(gd.Int(line), gd.Int(column)))
}

/*
Returns the local position and size for the grapheme at the given [param line] and [param column]. If [code]x[/code] or [code]y[/code] position of the returned rect equal [code]-1[/code], the position is outside of the viewable area of the control.
[b]Note:[/b] The Y position of the returned rect corresponds to the top side of the line, unlike [method get_pos_at_line_column] which returns the bottom side.
*/
func (self Go) GetRectAtLineColumn(line int, column int) gd.Rect2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2i(class(self).GetRectAtLineColumn(gd.Int(line), gd.Int(column)))
}

/*
Returns the equivalent minimap line at [param position].
*/
func (self Go) GetMinimapLineAtPos(position gd.Vector2i) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetMinimapLineAtPos(position)))
}

/*
Returns [code]true[/code] if the user is dragging their mouse for scrolling, selecting, or text dragging.
*/
func (self Go) IsDraggingCursor() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsDraggingCursor())
}

/*
Returns whether the mouse is over selection. If [param edges] is [code]true[/code], the edges are considered part of the selection.
*/
func (self Go) IsMouseOverSelection(edges bool) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsMouseOverSelection(edges, gd.Int(-1)))
}

/*
Adds a new caret at the given location. Returns the index of the new caret, or [code]-1[/code] if the location is invalid.
*/
func (self Go) AddCaret(line int, column int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).AddCaret(gd.Int(line), gd.Int(column))))
}

/*
Removes the given caret index.
[b]Note:[/b] This can result in adjustment of all other caret indices.
*/
func (self Go) RemoveCaret(caret int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveCaret(gd.Int(caret))
}

/*
Removes all additional carets.
*/
func (self Go) RemoveSecondaryCarets() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveSecondaryCarets()
}

/*
Returns the number of carets in this [TextEdit].
*/
func (self Go) GetCaretCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCaretCount()))
}

/*
Adds an additional caret above or below every caret. If [param below] is [code]true[/code] the new caret will be added below and above otherwise.
*/
func (self Go) AddCaretAtCarets(below bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddCaretAtCarets(below)
}

/*
Returns the carets sorted by selection beginning from lowest line and column to highest (from top to bottom of text).
If [param include_ignored_carets] is [code]false[/code], carets from [method multicaret_edit_ignore_caret] will be ignored.
*/
func (self Go) GetSortedCarets() []int32 {
	gc := gd.GarbageCollector(); _ = gc
	return []int32(class(self).GetSortedCarets(gc, false).AsSlice())
}

/*
Collapse all carets in the given range to the [param from_line] and [param from_column] position.
[param inclusive] applies to both ends.
If [method is_in_mulitcaret_edit] is [code]true[/code], carets that are collapsed will be [code]true[/code] for [method multicaret_edit_ignore_caret].
[method merge_overlapping_carets] will be called if any carets were collapsed.
*/
func (self Go) CollapseCarets(from_line int, from_column int, to_line int, to_column int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).CollapseCarets(gd.Int(from_line), gd.Int(from_column), gd.Int(to_line), gd.Int(to_column), false)
}

/*
Merges any overlapping carets. Will favor the newest caret, or the caret with a selection.
If [method is_in_mulitcaret_edit] is [code]true[/code], the merge will be queued to happen at the end of the multicaret edit. See [method begin_multicaret_edit] and [method end_multicaret_edit].
[b]Note:[/b] This is not called when a caret changes position but after certain actions, so it is possible to get into a state where carets overlap.
*/
func (self Go) MergeOverlappingCarets() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).MergeOverlappingCarets()
}

/*
Starts an edit for multiple carets. The edit must be ended with [method end_multicaret_edit]. Multicaret edits can be used to edit text at multiple carets and delay merging the carets until the end, so the caret indexes aren't affected immediately. [method begin_multicaret_edit] and [method end_multicaret_edit] can be nested, and the merge will happen at the last [method end_multicaret_edit].
Example usage:
[codeblock]
begin_complex_operation()
begin_multicaret_edit()
for i in range(get_caret_count()):
    if multicaret_edit_ignore_caret(i):
        continue
    # Logic here.
end_multicaret_edit()
end_complex_operation()
[/codeblock]
*/
func (self Go) BeginMulticaretEdit() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).BeginMulticaretEdit()
}

/*
Ends an edit for multiple carets, that was started with [method begin_multicaret_edit]. If this was the last [method end_multicaret_edit] and [method merge_overlapping_carets] was called, carets will be merged.
*/
func (self Go) EndMulticaretEdit() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).EndMulticaretEdit()
}

/*
Returns [code]true[/code] if a [method begin_multicaret_edit] has been called and [method end_multicaret_edit] has not yet been called.
*/
func (self Go) IsInMulitcaretEdit() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsInMulitcaretEdit())
}

/*
Returns [code]true[/code] if the given [param caret_index] should be ignored as part of a multicaret edit. See [method begin_multicaret_edit] and [method end_multicaret_edit]. Carets that should be ignored are ones that were part of removed text and will likely be merged at the end of the edit, or carets that were added during the edit.
It is recommended to [code]continue[/code] within a loop iterating on multiple carets if a caret should be ignored.
*/
func (self Go) MulticaretEditIgnoreCaret(caret_index int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).MulticaretEditIgnoreCaret(gd.Int(caret_index)))
}

/*
Returns [code]true[/code] if the caret is visible on the screen.
*/
func (self Go) IsCaretVisible() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsCaretVisible(gd.Int(0)))
}

/*
Returns the caret pixel draw position.
*/
func (self Go) GetCaretDrawPos() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetCaretDrawPos(gd.Int(0)))
}

/*
Moves the caret to the specified [param line] index. The caret column will be moved to the same visual position it was at the last time [method set_caret_column] was called, or clamped to the end of the line.
If [param adjust_viewport] is [code]true[/code], the viewport will center at the caret position after the move occurs.
If [param can_be_hidden] is [code]true[/code], the specified [param line] can be hidden.
If [param wrap_index] is [code]-1[/code], the caret column will be clamped to the [param line]'s length. If [param wrap_index] is greater than [code]-1[/code], the column will be moved to attempt to match the visual x position on the line's [param wrap_index] to the position from the last time [method set_caret_column] was called.
[b]Note:[/b] If supporting multiple carets this will not check for any overlap. See [method merge_overlapping_carets].
*/
func (self Go) SetCaretLine(line int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCaretLine(gd.Int(line), true, true, gd.Int(0), gd.Int(0))
}

/*
Returns the line the editing caret is on.
*/
func (self Go) GetCaretLine() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCaretLine(gd.Int(0))))
}

/*
Moves the caret to the specified [param column] index.
If [param adjust_viewport] is [code]true[/code], the viewport will center at the caret position after the move occurs.
[b]Note:[/b] If supporting multiple carets this will not check for any overlap. See [method merge_overlapping_carets].
*/
func (self Go) SetCaretColumn(column int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCaretColumn(gd.Int(column), true, gd.Int(0))
}

/*
Returns the column the editing caret is at.
*/
func (self Go) GetCaretColumn() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCaretColumn(gd.Int(0))))
}

/*
Returns the wrap index the editing caret is on.
*/
func (self Go) GetCaretWrapIndex() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCaretWrapIndex(gd.Int(0))))
}

/*
Returns a [String] text with the word under the caret's location.
*/
func (self Go) GetWordUnderCaret() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetWordUnderCaret(gc, gd.Int(-1)).String())
}

/*
Sets the current selection mode.
*/
func (self Go) SetSelectionMode(mode classdb.TextEditSelectionMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSelectionMode(mode)
}

/*
Returns the current selection mode.
*/
func (self Go) GetSelectionMode() classdb.TextEditSelectionMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextEditSelectionMode(class(self).GetSelectionMode())
}

/*
Select all the text.
If [member selecting_enabled] is [code]false[/code], no selection will occur.
*/
func (self Go) SelectAll() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SelectAll()
}

/*
Selects the word under the caret.
*/
func (self Go) SelectWordUnderCaret() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SelectWordUnderCaret(gd.Int(-1))
}

/*
Adds a selection and a caret for the next occurrence of the current selection. If there is no active selection, selects word under caret.
*/
func (self Go) AddSelectionForNextOccurrence() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddSelectionForNextOccurrence()
}

/*
Moves a selection and a caret for the next occurrence of the current selection. If there is no active selection, moves to the next occurrence of the word under caret.
*/
func (self Go) SkipSelectionForNextOccurrence() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SkipSelectionForNextOccurrence()
}

/*
Selects text from [param origin_line] and [param origin_column] to [param caret_line] and [param caret_column] for the given [param caret_index]. This moves the selection origin and the caret. If the positions are the same, the selection will be deselected.
If [member selecting_enabled] is [code]false[/code], no selection will occur.
[b]Note:[/b] If supporting multiple carets this will not check for any overlap. See [method merge_overlapping_carets].
*/
func (self Go) Select(origin_line int, origin_column int, caret_line int, caret_column int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Select(gd.Int(origin_line), gd.Int(origin_column), gd.Int(caret_line), gd.Int(caret_column), gd.Int(0))
}

/*
Returns [code]true[/code] if the user has selected text.
*/
func (self Go) HasSelection() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasSelection(gd.Int(-1)))
}

/*
Returns the text inside the selection of a caret, or all the carets if [param caret_index] is its default value [code]-1[/code].
*/
func (self Go) GetSelectedText() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetSelectedText(gc, gd.Int(-1)).String())
}

/*
Returns the caret index of the selection at the given [param line] and [param column], or [code]-1[/code] if there is none.
If [param include_edges] is [code]false[/code], the position must be inside the selection and not at either end. If [param only_selections] is [code]false[/code], carets without a selection will also be considered.
*/
func (self Go) GetSelectionAtLineColumn(line int, column int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetSelectionAtLineColumn(gd.Int(line), gd.Int(column), true, true)))
}

/*
Returns an [Array] of line ranges where [code]x[/code] is the first line and [code]y[/code] is the last line. All lines within these ranges will have a caret on them or be part of a selection. Each line will only be part of one line range, even if it has multiple carets on it.
If a selection's end column ([method get_selection_to_column]) is at column [code]0[/code], that line will not be included. If a selection begins on the line after another selection ends and [param merge_adjacent] is [code]true[/code], or they begin and end on the same line, one line range will include both selections.
*/
func (self Go) GetLineRangesFromCarets() gd.ArrayOf[gd.Vector2i] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Vector2i](class(self).GetLineRangesFromCarets(gc, false, true))
}

/*
Returns the origin line of the selection. This is the opposite end from the caret.
*/
func (self Go) GetSelectionOriginLine() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetSelectionOriginLine(gd.Int(0))))
}

/*
Returns the origin column of the selection. This is the opposite end from the caret.
*/
func (self Go) GetSelectionOriginColumn() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetSelectionOriginColumn(gd.Int(0))))
}

/*
Sets the selection origin line to the [param line] for the given [param caret_index]. If the selection origin is moved to the caret position, the selection will deselect.
If [param can_be_hidden] is [code]false[/code], The line will be set to the nearest unhidden line below or above.
If [param wrap_index] is [code]-1[/code], the selection origin column will be clamped to the [param line]'s length. If [param wrap_index] is greater than [code]-1[/code], the column will be moved to attempt to match the visual x position on the line's [param wrap_index] to the position from the last time [method set_selection_origin_column] or [method select] was called.
*/
func (self Go) SetSelectionOriginLine(line int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSelectionOriginLine(gd.Int(line), true, gd.Int(-1), gd.Int(0))
}

/*
Sets the selection origin column to the [param column] for the given [param caret_index]. If the selection origin is moved to the caret position, the selection will deselect.
*/
func (self Go) SetSelectionOriginColumn(column int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSelectionOriginColumn(gd.Int(column), gd.Int(0))
}

/*
Returns the selection begin line. Returns the caret line if there is no selection.
*/
func (self Go) GetSelectionFromLine() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetSelectionFromLine(gd.Int(0))))
}

/*
Returns the selection begin column. Returns the caret column if there is no selection.
*/
func (self Go) GetSelectionFromColumn() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetSelectionFromColumn(gd.Int(0))))
}

/*
Returns the selection end line. Returns the caret line if there is no selection.
*/
func (self Go) GetSelectionToLine() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetSelectionToLine(gd.Int(0))))
}

/*
Returns the selection end column. Returns the caret column if there is no selection.
*/
func (self Go) GetSelectionToColumn() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetSelectionToColumn(gd.Int(0))))
}

/*
Returns [code]true[/code] if the caret of the selection is after the selection origin. This can be used to determine the direction of the selection.
*/
func (self Go) IsCaretAfterSelectionOrigin() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsCaretAfterSelectionOrigin(gd.Int(0)))
}

/*
Deselects the current selection.
*/
func (self Go) Deselect() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Deselect(gd.Int(-1))
}

/*
Deletes the selected text.
*/
func (self Go) DeleteSelection() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DeleteSelection(gd.Int(-1))
}

/*
Returns if the given line is wrapped.
*/
func (self Go) IsLineWrapped(line int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsLineWrapped(gd.Int(line)))
}

/*
Returns the number of times the given line is wrapped.
*/
func (self Go) GetLineWrapCount(line int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetLineWrapCount(gd.Int(line))))
}

/*
Returns the wrap index of the given line column.
*/
func (self Go) GetLineWrapIndexAtColumn(line int, column int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetLineWrapIndexAtColumn(gd.Int(line), gd.Int(column))))
}

/*
Returns an array of [String]s representing each wrapped index.
*/
func (self Go) GetLineWrappedText(line int) []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetLineWrappedText(gc, gd.Int(line)).Strings(gc))
}

/*
Returns the [VScrollBar] of the [TextEdit].
*/
func (self Go) GetVScrollBar() gdclass.VScrollBar {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.VScrollBar(class(self).GetVScrollBar(gc))
}

/*
Returns the [HScrollBar] used by [TextEdit].
*/
func (self Go) GetHScrollBar() gdclass.HScrollBar {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.HScrollBar(class(self).GetHScrollBar(gc))
}

/*
Returns the scroll position for [param wrap_index] of [param line].
*/
func (self Go) GetScrollPosForLine(line int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetScrollPosForLine(gd.Int(line), gd.Int(0))))
}

/*
Positions the [param wrap_index] of [param line] at the top of the viewport.
*/
func (self Go) SetLineAsFirstVisible(line int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLineAsFirstVisible(gd.Int(line), gd.Int(0))
}

/*
Returns the first visible line.
*/
func (self Go) GetFirstVisibleLine() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetFirstVisibleLine()))
}

/*
Positions the [param wrap_index] of [param line] at the center of the viewport.
*/
func (self Go) SetLineAsCenterVisible(line int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLineAsCenterVisible(gd.Int(line), gd.Int(0))
}

/*
Positions the [param wrap_index] of [param line] at the bottom of the viewport.
*/
func (self Go) SetLineAsLastVisible(line int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLineAsLastVisible(gd.Int(line), gd.Int(0))
}

/*
Returns the last visible line. Use [method get_last_full_visible_line_wrap_index] for the wrap index.
*/
func (self Go) GetLastFullVisibleLine() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetLastFullVisibleLine()))
}

/*
Returns the last visible wrap index of the last visible line.
*/
func (self Go) GetLastFullVisibleLineWrapIndex() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetLastFullVisibleLineWrapIndex()))
}

/*
Returns the number of visible lines, including wrapped text.
*/
func (self Go) GetVisibleLineCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetVisibleLineCount()))
}

/*
Returns the total number of visible + wrapped lines between the two lines.
*/
func (self Go) GetVisibleLineCountInRange(from_line int, to_line int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetVisibleLineCountInRange(gd.Int(from_line), gd.Int(to_line))))
}

/*
Returns the number of lines that may be drawn.
*/
func (self Go) GetTotalVisibleLineCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetTotalVisibleLineCount()))
}

/*
Adjust the viewport so the caret is visible.
*/
func (self Go) AdjustViewportToCaret() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AdjustViewportToCaret(gd.Int(0))
}

/*
Centers the viewport on the line the editing caret is at. This also resets the [member scroll_horizontal] value to [code]0[/code].
*/
func (self Go) CenterViewportToCaret() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).CenterViewportToCaret(gd.Int(0))
}

/*
Returns the number of lines that may be drawn on the minimap.
*/
func (self Go) GetMinimapVisibleLines() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetMinimapVisibleLines()))
}

/*
Register a new gutter to this [TextEdit]. Use [param at] to have a specific gutter order. A value of [code]-1[/code] appends the gutter to the right.
*/
func (self Go) AddGutter() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddGutter(gd.Int(-1))
}

/*
Removes the gutter from this [TextEdit].
*/
func (self Go) RemoveGutter(gutter int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveGutter(gd.Int(gutter))
}

/*
Returns the number of gutters registered.
*/
func (self Go) GetGutterCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetGutterCount()))
}

/*
Sets the name of the gutter.
*/
func (self Go) SetGutterName(gutter int, name string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGutterName(gd.Int(gutter), gc.String(name))
}

/*
Returns the name of the gutter at the given index.
*/
func (self Go) GetGutterName(gutter int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetGutterName(gc, gd.Int(gutter)).String())
}

/*
Sets the type of gutter. Gutters can contain icons, text, or custom visuals. See [enum TextEdit.GutterType] for options.
*/
func (self Go) SetGutterType(gutter int, atype classdb.TextEditGutterType) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGutterType(gd.Int(gutter), atype)
}

/*
Returns the type of the gutter at the given index. Gutters can contain icons, text, or custom visuals. See [enum TextEdit.GutterType] for options.
*/
func (self Go) GetGutterType(gutter int) classdb.TextEditGutterType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextEditGutterType(class(self).GetGutterType(gd.Int(gutter)))
}

/*
Set the width of the gutter.
*/
func (self Go) SetGutterWidth(gutter int, width int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGutterWidth(gd.Int(gutter), gd.Int(width))
}

/*
Returns the width of the gutter at the given index.
*/
func (self Go) GetGutterWidth(gutter int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetGutterWidth(gd.Int(gutter))))
}

/*
Sets whether the gutter should be drawn.
*/
func (self Go) SetGutterDraw(gutter int, draw bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGutterDraw(gd.Int(gutter), draw)
}

/*
Returns whether the gutter is currently drawn.
*/
func (self Go) IsGutterDrawn(gutter int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsGutterDrawn(gd.Int(gutter)))
}

/*
Sets the gutter as clickable. This will change the mouse cursor to a pointing hand when hovering over the gutter.
*/
func (self Go) SetGutterClickable(gutter int, clickable bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGutterClickable(gd.Int(gutter), clickable)
}

/*
Returns whether the gutter is clickable.
*/
func (self Go) IsGutterClickable(gutter int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsGutterClickable(gd.Int(gutter)))
}

/*
Sets the gutter to overwritable. See [method merge_gutters].
*/
func (self Go) SetGutterOverwritable(gutter int, overwritable bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGutterOverwritable(gd.Int(gutter), overwritable)
}

/*
Returns whether the gutter is overwritable.
*/
func (self Go) IsGutterOverwritable(gutter int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsGutterOverwritable(gd.Int(gutter)))
}

/*
Merge the gutters from [param from_line] into [param to_line]. Only overwritable gutters will be copied.
*/
func (self Go) MergeGutters(from_line int, to_line int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).MergeGutters(gd.Int(from_line), gd.Int(to_line))
}

/*
Set a custom draw method for the gutter. The callback method must take the following args: [code]line: int, gutter: int, Area: Rect2[/code]. This only works when the gutter type is [constant GUTTER_TYPE_CUSTOM] (see [method set_gutter_type]).
*/
func (self Go) SetGutterCustomDraw(column int, draw_callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGutterCustomDraw(gd.Int(column), draw_callback)
}

/*
Returns the total width of all gutters and internal padding.
*/
func (self Go) GetTotalGutterWidth() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetTotalGutterWidth()))
}

/*
Sets the metadata for [param gutter] on [param line] to [param metadata].
*/
func (self Go) SetLineGutterMetadata(line int, gutter int, metadata gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLineGutterMetadata(gd.Int(line), gd.Int(gutter), metadata)
}

/*
Returns the metadata currently in [param gutter] at [param line].
*/
func (self Go) GetLineGutterMetadata(line int, gutter int) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(class(self).GetLineGutterMetadata(gc, gd.Int(line), gd.Int(gutter)))
}

/*
Sets the text for [param gutter] on [param line] to [param text]. This only works when the gutter type is [constant GUTTER_TYPE_STRING] (see [method set_gutter_type]).
*/
func (self Go) SetLineGutterText(line int, gutter int, text string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLineGutterText(gd.Int(line), gd.Int(gutter), gc.String(text))
}

/*
Returns the text currently in [param gutter] at [param line]. This only works when the gutter type is [constant GUTTER_TYPE_STRING] (see [method set_gutter_type]).
*/
func (self Go) GetLineGutterText(line int, gutter int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetLineGutterText(gc, gd.Int(line), gd.Int(gutter)).String())
}

/*
Sets the icon for [param gutter] on [param line] to [param icon]. This only works when the gutter type is [constant GUTTER_TYPE_ICON] (see [method set_gutter_type]).
*/
func (self Go) SetLineGutterIcon(line int, gutter int, icon gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLineGutterIcon(gd.Int(line), gd.Int(gutter), icon)
}

/*
Returns the icon currently in [param gutter] at [param line]. This only works when the gutter type is [constant GUTTER_TYPE_ICON] (see [method set_gutter_type]).
*/
func (self Go) GetLineGutterIcon(line int, gutter int) gdclass.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Texture2D(class(self).GetLineGutterIcon(gc, gd.Int(line), gd.Int(gutter)))
}

/*
Sets the color for [param gutter] on [param line] to [param color].
*/
func (self Go) SetLineGutterItemColor(line int, gutter int, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLineGutterItemColor(gd.Int(line), gd.Int(gutter), color)
}

/*
Returns the color currently in [param gutter] at [param line].
*/
func (self Go) GetLineGutterItemColor(line int, gutter int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(class(self).GetLineGutterItemColor(gd.Int(line), gd.Int(gutter)))
}

/*
If [param clickable] is [code]true[/code], makes the [param gutter] on [param line] clickable. See [signal gutter_clicked].
*/
func (self Go) SetLineGutterClickable(line int, gutter int, clickable bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLineGutterClickable(gd.Int(line), gd.Int(gutter), clickable)
}

/*
Returns whether the gutter on the given line is clickable.
*/
func (self Go) IsLineGutterClickable(line int, gutter int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsLineGutterClickable(gd.Int(line), gd.Int(gutter)))
}

/*
Sets the current background color of the line. Set to [code]Color(0, 0, 0, 0)[/code] for no color.
*/
func (self Go) SetLineBackgroundColor(line int, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLineBackgroundColor(gd.Int(line), color)
}

/*
Returns the current background color of the line. [code]Color(0, 0, 0, 0)[/code] is returned if no color is set.
*/
func (self Go) GetLineBackgroundColor(line int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(class(self).GetLineBackgroundColor(gd.Int(line)))
}

/*
Returns the [PopupMenu] of this [TextEdit]. By default, this menu is displayed when right-clicking on the [TextEdit].
You can add custom menu items or remove standard ones. Make sure your IDs don't conflict with the standard ones (see [enum MenuItems]). For example:
[codeblocks]
[gdscript]
func _ready():
    var menu = get_menu()
    # Remove all items after "Redo".
    menu.item_count = menu.get_item_index(MENU_REDO) + 1
    # Add custom items.
    menu.add_separator()
    menu.add_item("Insert Date", MENU_MAX + 1)
    # Connect callback.
    menu.id_pressed.connect(_on_item_pressed)

func _on_item_pressed(id):
    if id == MENU_MAX + 1:
        insert_text_at_caret(Time.get_date_string_from_system())
[/gdscript]
[csharp]
public override void _Ready()
{
    var menu = GetMenu();
    // Remove all items after "Redo".
    menu.ItemCount = menu.GetItemIndex(TextEdit.MenuItems.Redo) + 1;
    // Add custom items.
    menu.AddSeparator();
    menu.AddItem("Insert Date", TextEdit.MenuItems.Max + 1);
    // Add event handler.
    menu.IdPressed += OnItemPressed;
}

public void OnItemPressed(int id)
{
    if (id == TextEdit.MenuItems.Max + 1)
    {
        InsertTextAtCaret(Time.GetDateStringFromSystem());
    }
}
[/csharp]
[/codeblocks]
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
func (self Go) GetMenu() gdclass.PopupMenu {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.PopupMenu(class(self).GetMenu(gc))
}

/*
Returns whether the menu is visible. Use this instead of [code]get_menu().visible[/code] to improve performance (so the creation of the menu is avoided).
*/
func (self Go) IsMenuVisible() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsMenuVisible())
}

/*
Executes a given action as defined in the [enum MenuItems] enum.
*/
func (self Go) MenuOption(option int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).MenuOption(gd.Int(option))
}

/*
This method does nothing.
*/
func (self Go) AdjustCaretsAfterEdit(caret int, from_line int, from_col int, to_line int, to_col int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AdjustCaretsAfterEdit(gd.Int(caret), gd.Int(from_line), gd.Int(from_col), gd.Int(to_line), gd.Int(to_col))
}

/*
Returns a list of caret indexes in their edit order, this done from bottom to top. Edit order refers to the way actions such as [method insert_text_at_caret] are applied.
*/
func (self Go) GetCaretIndexEditOrder() []int32 {
	gc := gd.GarbageCollector(); _ = gc
	return []int32(class(self).GetCaretIndexEditOrder(gc).AsSlice())
}

/*
Returns the original start line of the selection.
*/
func (self Go) GetSelectionLine() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetSelectionLine(gd.Int(0))))
}

/*
Returns the original start column of the selection.
*/
func (self Go) GetSelectionColumn() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetSelectionColumn(gd.Int(0))))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.TextEdit
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("TextEdit"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Text() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetText(gc).String())
}

func (self Go) SetText(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetText(gc.String(value))
}

func (self Go) PlaceholderText() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetPlaceholder(gc).String())
}

func (self Go) SetPlaceholderText(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPlaceholder(gc.String(value))
}

func (self Go) Editable() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsEditable())
}

func (self Go) SetEditable(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEditable(value)
}

func (self Go) ContextMenuEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsContextMenuEnabled())
}

func (self Go) SetContextMenuEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetContextMenuEnabled(value)
}

func (self Go) ShortcutKeysEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsShortcutKeysEnabled())
}

func (self Go) SetShortcutKeysEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShortcutKeysEnabled(value)
}

func (self Go) SelectingEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsSelectingEnabled())
}

func (self Go) SetSelectingEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSelectingEnabled(value)
}

func (self Go) DeselectOnFocusLossEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDeselectOnFocusLossEnabled())
}

func (self Go) SetDeselectOnFocusLossEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDeselectOnFocusLossEnabled(value)
}

func (self Go) DragAndDropSelectionEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDragAndDropSelectionEnabled())
}

func (self Go) SetDragAndDropSelectionEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDragAndDropSelectionEnabled(value)
}

func (self Go) VirtualKeyboardEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsVirtualKeyboardEnabled())
}

func (self Go) SetVirtualKeyboardEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVirtualKeyboardEnabled(value)
}

func (self Go) MiddleMousePasteEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsMiddleMousePasteEnabled())
}

func (self Go) SetMiddleMousePasteEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMiddleMousePasteEnabled(value)
}

func (self Go) WrapMode() classdb.TextEditLineWrappingMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TextEditLineWrappingMode(class(self).GetLineWrappingMode())
}

func (self Go) SetWrapMode(value classdb.TextEditLineWrappingMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLineWrappingMode(value)
}

func (self Go) AutowrapMode() classdb.TextServerAutowrapMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TextServerAutowrapMode(class(self).GetAutowrapMode())
}

func (self Go) SetAutowrapMode(value classdb.TextServerAutowrapMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAutowrapMode(value)
}

func (self Go) IndentWrappedLines() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsIndentWrappedLines())
}

func (self Go) SetIndentWrappedLines(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIndentWrappedLines(value)
}

func (self Go) ScrollSmooth() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsSmoothScrollEnabled())
}

func (self Go) SetScrollSmooth(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSmoothScrollEnabled(value)
}

func (self Go) ScrollVScrollSpeed() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVScrollSpeed()))
}

func (self Go) SetScrollVScrollSpeed(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVScrollSpeed(gd.Float(value))
}

func (self Go) ScrollPastEndOfFile() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsScrollPastEndOfFileEnabled())
}

func (self Go) SetScrollPastEndOfFile(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetScrollPastEndOfFileEnabled(value)
}

func (self Go) ScrollVertical() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVScroll()))
}

func (self Go) SetScrollVertical(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVScroll(gd.Float(value))
}

func (self Go) ScrollHorizontal() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetHScroll()))
}

func (self Go) SetScrollHorizontal(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHScroll(gd.Int(value))
}

func (self Go) ScrollFitContentHeight() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsFitContentHeightEnabled())
}

func (self Go) SetScrollFitContentHeight(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFitContentHeightEnabled(value)
}

func (self Go) MinimapDraw() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDrawingMinimap())
}

func (self Go) SetMinimapDraw(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawMinimap(value)
}

func (self Go) MinimapWidth() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMinimapWidth()))
}

func (self Go) SetMinimapWidth(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMinimapWidth(gd.Int(value))
}

func (self Go) CaretType() classdb.TextEditCaretType {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TextEditCaretType(class(self).GetCaretType())
}

func (self Go) SetCaretType(value classdb.TextEditCaretType) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCaretType(value)
}

func (self Go) CaretBlink() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsCaretBlinkEnabled())
}

func (self Go) SetCaretBlink(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCaretBlinkEnabled(value)
}

func (self Go) CaretBlinkInterval() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetCaretBlinkInterval()))
}

func (self Go) SetCaretBlinkInterval(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCaretBlinkInterval(gd.Float(value))
}

func (self Go) CaretDrawWhenEditableDisabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDrawingCaretWhenEditableDisabled())
}

func (self Go) SetCaretDrawWhenEditableDisabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawCaretWhenEditableDisabled(value)
}

func (self Go) CaretMoveOnRightClick() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsMoveCaretOnRightClickEnabled())
}

func (self Go) SetCaretMoveOnRightClick(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMoveCaretOnRightClickEnabled(value)
}

func (self Go) CaretMidGrapheme() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsCaretMidGraphemeEnabled())
}

func (self Go) SetCaretMidGrapheme(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCaretMidGraphemeEnabled(value)
}

func (self Go) CaretMultiple() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsMultipleCaretsEnabled())
}

func (self Go) SetCaretMultiple(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMultipleCaretsEnabled(value)
}

func (self Go) UseDefaultWordSeparators() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDefaultWordSeparatorsEnabled())
}

func (self Go) SetUseDefaultWordSeparators(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUseDefaultWordSeparators(value)
}

func (self Go) UseCustomWordSeparators() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsCustomWordSeparatorsEnabled())
}

func (self Go) SetUseCustomWordSeparators(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUseCustomWordSeparators(value)
}

func (self Go) CustomWordSeparators() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetCustomWordSeparators(gc).String())
}

func (self Go) SetCustomWordSeparators(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCustomWordSeparators(gc.String(value))
}

func (self Go) SyntaxHighlighter() gdclass.SyntaxHighlighter {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.SyntaxHighlighter(class(self).GetSyntaxHighlighter(gc))
}

func (self Go) SetSyntaxHighlighter(value gdclass.SyntaxHighlighter) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSyntaxHighlighter(value)
}

func (self Go) HighlightAllOccurrences() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsHighlightAllOccurrencesEnabled())
}

func (self Go) SetHighlightAllOccurrences(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHighlightAllOccurrences(value)
}

func (self Go) HighlightCurrentLine() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsHighlightCurrentLineEnabled())
}

func (self Go) SetHighlightCurrentLine(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHighlightCurrentLine(value)
}

func (self Go) DrawControlChars() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetDrawControlChars())
}

func (self Go) SetDrawControlChars(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawControlChars(value)
}

func (self Go) DrawTabs() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDrawingTabs())
}

func (self Go) SetDrawTabs(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawTabs(value)
}

func (self Go) DrawSpaces() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDrawingSpaces())
}

func (self Go) SetDrawSpaces(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawSpaces(value)
}

func (self Go) TextDirection() classdb.ControlTextDirection {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.ControlTextDirection(class(self).GetTextDirection())
}

func (self Go) SetTextDirection(value classdb.ControlTextDirection) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextDirection(value)
}

func (self Go) Language() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetLanguage(gc).String())
}

func (self Go) SetLanguage(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLanguage(gc.String(value))
}

func (self Go) StructuredTextBidiOverride() classdb.TextServerStructuredTextParser {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TextServerStructuredTextParser(class(self).GetStructuredTextBidiOverride())
}

func (self Go) SetStructuredTextBidiOverride(value classdb.TextServerStructuredTextParser) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStructuredTextBidiOverride(value)
}

func (self Go) StructuredTextBidiOverrideOptions() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Array(class(self).GetStructuredTextBidiOverrideOptions(gc))
}

func (self Go) SetStructuredTextBidiOverrideOptions(value gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStructuredTextBidiOverrideOptions(value)
}

/*
Override this method to define what happens when the user types in the provided key [param unicode_char].
*/
func (class) _handle_unicode_input(impl func(ptr unsafe.Pointer, unicode_char gd.Int, caret_index gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var unicode_char = gd.UnsafeGet[gd.Int](p_args,0)
		var caret_index = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, unicode_char, caret_index)
		ctx.End()
	}
}

/*
Override this method to define what happens when the user presses the backspace key.
*/
func (class) _backspace(impl func(ptr unsafe.Pointer, caret_index gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var caret_index = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, caret_index)
		ctx.End()
	}
}

/*
Override this method to define what happens when the user performs a cut operation.
*/
func (class) _cut(impl func(ptr unsafe.Pointer, caret_index gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var caret_index = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, caret_index)
		ctx.End()
	}
}

/*
Override this method to define what happens when the user performs a copy operation.
*/
func (class) _copy(impl func(ptr unsafe.Pointer, caret_index gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var caret_index = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, caret_index)
		ctx.End()
	}
}

/*
Override this method to define what happens when the user performs a paste operation.
*/
func (class) _paste(impl func(ptr unsafe.Pointer, caret_index gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var caret_index = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, caret_index)
		ctx.End()
	}
}

/*
Override this method to define what happens when the user performs a paste operation with middle mouse button.
[b]Note:[/b] This method is only implemented on Linux.
*/
func (class) _paste_primary_clipboard(impl func(ptr unsafe.Pointer, caret_index gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var caret_index = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, caret_index)
		ctx.End()
	}
}

/*
Returns [code]true[/code] if the user has text in the [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] (IME).
*/
//go:nosplit
func (self class) HasImeText() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_has_ime_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Closes the [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] (IME) if it is open. Any text in the IME will be lost.
*/
//go:nosplit
func (self class) CancelIme()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_cancel_ime, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies text from the [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] (IME) to each caret and closes the IME if it is open.
*/
//go:nosplit
func (self class) ApplyIme()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_apply_ime, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetEditable(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsEditable() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTextDirection(direction classdb.ControlTextDirection)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextDirection() classdb.ControlTextDirection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ControlTextDirection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLanguage(language gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(language))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLanguage(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStructuredTextBidiOverride(parser classdb.TextServerStructuredTextParser)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, parser)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStructuredTextBidiOverride() classdb.TextServerStructuredTextParser {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerStructuredTextParser](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStructuredTextBidiOverrideOptions(args gd.Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(args))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStructuredTextBidiOverrideOptions(ctx gd.Lifetime) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the tab size for the [TextEdit] to use.
*/
//go:nosplit
func (self class) SetTabSize(size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_tab_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [TextEdit]'s' tab size.
*/
//go:nosplit
func (self class) GetTabSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_tab_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetIndentWrappedLines(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_indent_wrapped_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsIndentWrappedLines() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_indent_wrapped_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [code]true[/code], sets the user into overtype mode. When the user types in this mode, it will override existing text.
*/
//go:nosplit
func (self class) SetOvertypeModeEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_overtype_mode_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the user is in overtype mode.
*/
//go:nosplit
func (self class) IsOvertypeModeEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_overtype_mode_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetContextMenuEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_context_menu_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsContextMenuEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_context_menu_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShortcutKeysEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_shortcut_keys_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShortcutKeysEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_shortcut_keys_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVirtualKeyboardEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_virtual_keyboard_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsVirtualKeyboardEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_virtual_keyboard_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMiddleMousePasteEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_middle_mouse_paste_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsMiddleMousePasteEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_middle_mouse_paste_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Performs a full reset of [TextEdit], including undo history.
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetText(text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetText(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the number of lines in the text.
*/
//go:nosplit
func (self class) GetLineCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_line_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPlaceholder(text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_placeholder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPlaceholder(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_placeholder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the text for a specific [param line].
Carets on the line will attempt to keep their visual x position.
*/
//go:nosplit
func (self class) SetLine(line gd.Int, new_text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, mmm.Get(new_text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the text of a specific line.
*/
//go:nosplit
func (self class) GetLine(ctx gd.Lifetime, line gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the width in pixels of the [param wrap_index] on [param line].
*/
//go:nosplit
func (self class) GetLineWidth(line gd.Int, wrap_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, wrap_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_line_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the maximum value of the line height among all lines.
[b]Note:[/b] The return value is influenced by [theme_item line_spacing] and [theme_item font_size]. And it will not be less than [code]1[/code].
*/
//go:nosplit
func (self class) GetLineHeight() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_line_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of spaces and [code]tab * tab_size[/code] before the first char.
*/
//go:nosplit
func (self class) GetIndentLevel(line gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_indent_level, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the first column containing a non-whitespace character.
*/
//go:nosplit
func (self class) GetFirstNonWhitespaceColumn(line gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_first_non_whitespace_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Swaps the two lines. Carets will be swapped with the lines.
*/
//go:nosplit
func (self class) SwapLines(from_line gd.Int, to_line gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from_line)
	callframe.Arg(frame, to_line)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_swap_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Inserts a new line with [param text] at [param line].
*/
//go:nosplit
func (self class) InsertLineAt(line gd.Int, text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_insert_line_at, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the line of text at [param line]. Carets on this line will attempt to match their previous visual x position.
If [param move_carets_down] is [code]true[/code] carets will move to the next line down, otherwise carets will move up.
*/
//go:nosplit
func (self class) RemoveLineAt(line gd.Int, move_carets_down bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, move_carets_down)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_remove_line_at, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Insert the specified text at the caret position.
*/
//go:nosplit
func (self class) InsertTextAtCaret(text gd.String, caret_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_insert_text_at_caret, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Inserts the [param text] at [param line] and [param column].
If [param before_selection_begin] is [code]true[/code], carets and selections that begin at [param line] and [param column] will moved to the end of the inserted text, along with all carets after it.
If [param before_selection_end] is [code]true[/code], selections that end at [param line] and [param column] will be extended to the end of the inserted text. These parameters can be used to insert text inside of or outside of selections.
*/
//go:nosplit
func (self class) InsertText(text gd.String, line gd.Int, column gd.Int, before_selection_begin bool, before_selection_end bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	callframe.Arg(frame, line)
	callframe.Arg(frame, column)
	callframe.Arg(frame, before_selection_begin)
	callframe.Arg(frame, before_selection_end)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_insert_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes text between the given positions.
*/
//go:nosplit
func (self class) RemoveText(from_line gd.Int, from_column gd.Int, to_line gd.Int, to_column gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from_line)
	callframe.Arg(frame, from_column)
	callframe.Arg(frame, to_line)
	callframe.Arg(frame, to_column)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_remove_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the last unhidden line in the entire [TextEdit].
*/
//go:nosplit
func (self class) GetLastUnhiddenLine() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_last_unhidden_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the count to the next visible line from [param line] to [code]line + visible_amount[/code]. Can also count backwards. For example if a [TextEdit] has 5 lines with lines 2 and 3 hidden, calling this with [code]line = 1, visible_amount = 1[/code] would return 3.
*/
//go:nosplit
func (self class) GetNextVisibleLineOffsetFrom(line gd.Int, visible_amount gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, visible_amount)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_next_visible_line_offset_from, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Similar to [method get_next_visible_line_offset_from], but takes into account the line wrap indexes. In the returned vector, [code]x[/code] is the line, [code]y[/code] is the wrap index.
*/
//go:nosplit
func (self class) GetNextVisibleLineIndexOffsetFrom(line gd.Int, wrap_index gd.Int, visible_amount gd.Int) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, wrap_index)
	callframe.Arg(frame, visible_amount)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_next_visible_line_index_offset_from, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Called when the user presses the backspace key. Can be overridden with [method _backspace].
*/
//go:nosplit
func (self class) Backspace(caret_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_backspace, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Cut's the current selection. Can be overridden with [method _cut].
*/
//go:nosplit
func (self class) Cut(caret_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_cut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Copies the current text selection. Can be overridden with [method _copy].
*/
//go:nosplit
func (self class) Copy(caret_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_copy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Paste at the current location. Can be overridden with [method _paste].
*/
//go:nosplit
func (self class) Paste(caret_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_paste, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Pastes the primary clipboard.
*/
//go:nosplit
func (self class) PastePrimaryClipboard(caret_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_paste_primary_clipboard, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Starts an action, will end the current action if [param action] is different.
An action will also end after a call to [method end_action], after [member ProjectSettings.gui/timers/text_edit_idle_detect_sec] is triggered or a new undoable step outside the [method start_action] and [method end_action] calls.
*/
//go:nosplit
func (self class) StartAction(action classdb.TextEditEditAction)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, action)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_start_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Marks the end of steps in the current action started with [method start_action].
*/
//go:nosplit
func (self class) EndAction()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_end_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Starts a multipart edit. All edits will be treated as one action until [method end_complex_operation] is called.
*/
//go:nosplit
func (self class) BeginComplexOperation()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_begin_complex_operation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Ends a multipart edit, started with [method begin_complex_operation]. If called outside a complex operation, the current operation is pushed onto the undo/redo stack.
*/
//go:nosplit
func (self class) EndComplexOperation()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_end_complex_operation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if an "undo" action is available.
*/
//go:nosplit
func (self class) HasUndo() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_has_undo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if a "redo" action is available.
*/
//go:nosplit
func (self class) HasRedo() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_has_redo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Perform undo operation.
*/
//go:nosplit
func (self class) Undo()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_undo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Perform redo operation.
*/
//go:nosplit
func (self class) Redo()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_redo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clears the undo history.
*/
//go:nosplit
func (self class) ClearUndoHistory()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_clear_undo_history, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Tag the current version as saved.
*/
//go:nosplit
func (self class) TagSavedVersion()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_tag_saved_version, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the current version of the [TextEdit]. The version is a count of recorded operations by the undo/redo history.
*/
//go:nosplit
func (self class) GetVersion() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_version, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the last tagged saved version from [method tag_saved_version].
*/
//go:nosplit
func (self class) GetSavedVersion() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_saved_version, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the search text. See [method set_search_flags].
*/
//go:nosplit
func (self class) SetSearchText(search_text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(search_text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_search_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the search [param flags]. This is used with [method set_search_text] to highlight occurrences of the searched text. Search flags can be specified from the [enum SearchFlags] enum.
*/
//go:nosplit
func (self class) SetSearchFlags(flags gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_search_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Perform a search inside the text. Search flags can be specified in the [enum SearchFlags] enum.
In the returned vector, [code]x[/code] is the column, [code]y[/code] is the line. If no results are found, both are equal to [code]-1[/code].
[codeblocks]
[gdscript]
var result = search("print", SEARCH_WHOLE_WORDS, 0, 0)
if result.x != -1:
    # Result found.
    var line_number = result.y
    var column_number = result.x
[/gdscript]
[csharp]
Vector2I result = Search("print", (uint)TextEdit.SearchFlags.WholeWords, 0, 0);
if (result.X != -1)
{
    // Result found.
    int lineNumber = result.Y;
    int columnNumber = result.X;
}
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) Search(text gd.String, flags gd.Int, from_line gd.Int, from_column gd.Int) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	callframe.Arg(frame, flags)
	callframe.Arg(frame, from_line)
	callframe.Arg(frame, from_column)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_search, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Provide custom tooltip text. The callback method must take the following args: [code]hovered_word: String[/code].
*/
//go:nosplit
func (self class) SetTooltipRequestFunc(callback gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_tooltip_request_func, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the local mouse position adjusted for the text direction.
*/
//go:nosplit
func (self class) GetLocalMousePos() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_local_mouse_pos, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the word at [param position].
*/
//go:nosplit
func (self class) GetWordAtPos(ctx gd.Lifetime, position gd.Vector2) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_word_at_pos, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the line and column at the given position. In the returned vector, [code]x[/code] is the column, [code]y[/code] is the line. If [param allow_out_of_bounds] is [code]false[/code] and the position is not over the text, both vector values will be set to [code]-1[/code].
*/
//go:nosplit
func (self class) GetLineColumnAtPos(position gd.Vector2i, allow_out_of_bounds bool) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, allow_out_of_bounds)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_line_column_at_pos, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the local position for the given [param line] and [param column]. If [code]x[/code] or [code]y[/code] of the returned vector equal [code]-1[/code], the position is outside of the viewable area of the control.
[b]Note:[/b] The Y position corresponds to the bottom side of the line. Use [method get_rect_at_line_column] to get the top side position.
*/
//go:nosplit
func (self class) GetPosAtLineColumn(line gd.Int, column gd.Int) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_pos_at_line_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the local position and size for the grapheme at the given [param line] and [param column]. If [code]x[/code] or [code]y[/code] position of the returned rect equal [code]-1[/code], the position is outside of the viewable area of the control.
[b]Note:[/b] The Y position of the returned rect corresponds to the top side of the line, unlike [method get_pos_at_line_column] which returns the bottom side.
*/
//go:nosplit
func (self class) GetRectAtLineColumn(line gd.Int, column gd.Int) gd.Rect2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Rect2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_rect_at_line_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the equivalent minimap line at [param position].
*/
//go:nosplit
func (self class) GetMinimapLineAtPos(position gd.Vector2i) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_minimap_line_at_pos, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the user is dragging their mouse for scrolling, selecting, or text dragging.
*/
//go:nosplit
func (self class) IsDraggingCursor() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_dragging_cursor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns whether the mouse is over selection. If [param edges] is [code]true[/code], the edges are considered part of the selection.
*/
//go:nosplit
func (self class) IsMouseOverSelection(edges bool, caret_index gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, edges)
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_mouse_over_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCaretType(atype classdb.TextEditCaretType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_caret_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCaretType() classdb.TextEditCaretType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextEditCaretType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_caret_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCaretBlinkEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_caret_blink_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCaretBlinkEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_caret_blink_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCaretBlinkInterval(interval gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, interval)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_caret_blink_interval, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCaretBlinkInterval() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_caret_blink_interval, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDrawCaretWhenEditableDisabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_draw_caret_when_editable_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDrawingCaretWhenEditableDisabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_drawing_caret_when_editable_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMoveCaretOnRightClickEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_move_caret_on_right_click_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsMoveCaretOnRightClickEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_move_caret_on_right_click_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCaretMidGraphemeEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_caret_mid_grapheme_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCaretMidGraphemeEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_caret_mid_grapheme_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMultipleCaretsEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_multiple_carets_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsMultipleCaretsEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_multiple_carets_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a new caret at the given location. Returns the index of the new caret, or [code]-1[/code] if the location is invalid.
*/
//go:nosplit
func (self class) AddCaret(line gd.Int, column gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_add_caret, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes the given caret index.
[b]Note:[/b] This can result in adjustment of all other caret indices.
*/
//go:nosplit
func (self class) RemoveCaret(caret gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_remove_caret, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all additional carets.
*/
//go:nosplit
func (self class) RemoveSecondaryCarets()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_remove_secondary_carets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of carets in this [TextEdit].
*/
//go:nosplit
func (self class) GetCaretCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_caret_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds an additional caret above or below every caret. If [param below] is [code]true[/code] the new caret will be added below and above otherwise.
*/
//go:nosplit
func (self class) AddCaretAtCarets(below bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, below)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_add_caret_at_carets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the carets sorted by selection beginning from lowest line and column to highest (from top to bottom of text).
If [param include_ignored_carets] is [code]false[/code], carets from [method multicaret_edit_ignore_caret] will be ignored.
*/
//go:nosplit
func (self class) GetSortedCarets(ctx gd.Lifetime, include_ignored_carets bool) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, include_ignored_carets)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_sorted_carets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Collapse all carets in the given range to the [param from_line] and [param from_column] position.
[param inclusive] applies to both ends.
If [method is_in_mulitcaret_edit] is [code]true[/code], carets that are collapsed will be [code]true[/code] for [method multicaret_edit_ignore_caret].
[method merge_overlapping_carets] will be called if any carets were collapsed.
*/
//go:nosplit
func (self class) CollapseCarets(from_line gd.Int, from_column gd.Int, to_line gd.Int, to_column gd.Int, inclusive bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from_line)
	callframe.Arg(frame, from_column)
	callframe.Arg(frame, to_line)
	callframe.Arg(frame, to_column)
	callframe.Arg(frame, inclusive)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_collapse_carets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Merges any overlapping carets. Will favor the newest caret, or the caret with a selection.
If [method is_in_mulitcaret_edit] is [code]true[/code], the merge will be queued to happen at the end of the multicaret edit. See [method begin_multicaret_edit] and [method end_multicaret_edit].
[b]Note:[/b] This is not called when a caret changes position but after certain actions, so it is possible to get into a state where carets overlap.
*/
//go:nosplit
func (self class) MergeOverlappingCarets()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_merge_overlapping_carets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Starts an edit for multiple carets. The edit must be ended with [method end_multicaret_edit]. Multicaret edits can be used to edit text at multiple carets and delay merging the carets until the end, so the caret indexes aren't affected immediately. [method begin_multicaret_edit] and [method end_multicaret_edit] can be nested, and the merge will happen at the last [method end_multicaret_edit].
Example usage:
[codeblock]
begin_complex_operation()
begin_multicaret_edit()
for i in range(get_caret_count()):
    if multicaret_edit_ignore_caret(i):
        continue
    # Logic here.
end_multicaret_edit()
end_complex_operation()
[/codeblock]
*/
//go:nosplit
func (self class) BeginMulticaretEdit()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_begin_multicaret_edit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Ends an edit for multiple carets, that was started with [method begin_multicaret_edit]. If this was the last [method end_multicaret_edit] and [method merge_overlapping_carets] was called, carets will be merged.
*/
//go:nosplit
func (self class) EndMulticaretEdit()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_end_multicaret_edit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if a [method begin_multicaret_edit] has been called and [method end_multicaret_edit] has not yet been called.
*/
//go:nosplit
func (self class) IsInMulitcaretEdit() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_in_mulitcaret_edit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the given [param caret_index] should be ignored as part of a multicaret edit. See [method begin_multicaret_edit] and [method end_multicaret_edit]. Carets that should be ignored are ones that were part of removed text and will likely be merged at the end of the edit, or carets that were added during the edit.
It is recommended to [code]continue[/code] within a loop iterating on multiple carets if a caret should be ignored.
*/
//go:nosplit
func (self class) MulticaretEditIgnoreCaret(caret_index gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_multicaret_edit_ignore_caret, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the caret is visible on the screen.
*/
//go:nosplit
func (self class) IsCaretVisible(caret_index gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_caret_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the caret pixel draw position.
*/
//go:nosplit
func (self class) GetCaretDrawPos(caret_index gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_caret_draw_pos, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Moves the caret to the specified [param line] index. The caret column will be moved to the same visual position it was at the last time [method set_caret_column] was called, or clamped to the end of the line.
If [param adjust_viewport] is [code]true[/code], the viewport will center at the caret position after the move occurs.
If [param can_be_hidden] is [code]true[/code], the specified [param line] can be hidden.
If [param wrap_index] is [code]-1[/code], the caret column will be clamped to the [param line]'s length. If [param wrap_index] is greater than [code]-1[/code], the column will be moved to attempt to match the visual x position on the line's [param wrap_index] to the position from the last time [method set_caret_column] was called.
[b]Note:[/b] If supporting multiple carets this will not check for any overlap. See [method merge_overlapping_carets].
*/
//go:nosplit
func (self class) SetCaretLine(line gd.Int, adjust_viewport bool, can_be_hidden bool, wrap_index gd.Int, caret_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, adjust_viewport)
	callframe.Arg(frame, can_be_hidden)
	callframe.Arg(frame, wrap_index)
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_caret_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the line the editing caret is on.
*/
//go:nosplit
func (self class) GetCaretLine(caret_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_caret_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Moves the caret to the specified [param column] index.
If [param adjust_viewport] is [code]true[/code], the viewport will center at the caret position after the move occurs.
[b]Note:[/b] If supporting multiple carets this will not check for any overlap. See [method merge_overlapping_carets].
*/
//go:nosplit
func (self class) SetCaretColumn(column gd.Int, adjust_viewport bool, caret_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, adjust_viewport)
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_caret_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the column the editing caret is at.
*/
//go:nosplit
func (self class) GetCaretColumn(caret_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_caret_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the wrap index the editing caret is on.
*/
//go:nosplit
func (self class) GetCaretWrapIndex(caret_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_caret_wrap_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a [String] text with the word under the caret's location.
*/
//go:nosplit
func (self class) GetWordUnderCaret(ctx gd.Lifetime, caret_index gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_word_under_caret, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseDefaultWordSeparators(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_use_default_word_separators, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDefaultWordSeparatorsEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_default_word_separators_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseCustomWordSeparators(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_use_custom_word_separators, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCustomWordSeparatorsEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_custom_word_separators_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCustomWordSeparators(custom_word_separators gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(custom_word_separators))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_custom_word_separators, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCustomWordSeparators(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_custom_word_separators, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSelectingEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_selecting_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSelectingEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_selecting_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDeselectOnFocusLossEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_deselect_on_focus_loss_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDeselectOnFocusLossEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_deselect_on_focus_loss_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDragAndDropSelectionEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_drag_and_drop_selection_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDragAndDropSelectionEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_drag_and_drop_selection_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the current selection mode.
*/
//go:nosplit
func (self class) SetSelectionMode(mode classdb.TextEditSelectionMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_selection_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the current selection mode.
*/
//go:nosplit
func (self class) GetSelectionMode() classdb.TextEditSelectionMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextEditSelectionMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_selection_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Select all the text.
If [member selecting_enabled] is [code]false[/code], no selection will occur.
*/
//go:nosplit
func (self class) SelectAll()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_select_all, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Selects the word under the caret.
*/
//go:nosplit
func (self class) SelectWordUnderCaret(caret_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_select_word_under_caret, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a selection and a caret for the next occurrence of the current selection. If there is no active selection, selects word under caret.
*/
//go:nosplit
func (self class) AddSelectionForNextOccurrence()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_add_selection_for_next_occurrence, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Moves a selection and a caret for the next occurrence of the current selection. If there is no active selection, moves to the next occurrence of the word under caret.
*/
//go:nosplit
func (self class) SkipSelectionForNextOccurrence()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_skip_selection_for_next_occurrence, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Selects text from [param origin_line] and [param origin_column] to [param caret_line] and [param caret_column] for the given [param caret_index]. This moves the selection origin and the caret. If the positions are the same, the selection will be deselected.
If [member selecting_enabled] is [code]false[/code], no selection will occur.
[b]Note:[/b] If supporting multiple carets this will not check for any overlap. See [method merge_overlapping_carets].
*/
//go:nosplit
func (self class) Select(origin_line gd.Int, origin_column gd.Int, caret_line gd.Int, caret_column gd.Int, caret_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, origin_line)
	callframe.Arg(frame, origin_column)
	callframe.Arg(frame, caret_line)
	callframe.Arg(frame, caret_column)
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_select_, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the user has selected text.
*/
//go:nosplit
func (self class) HasSelection(caret_index gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_has_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the text inside the selection of a caret, or all the carets if [param caret_index] is its default value [code]-1[/code].
*/
//go:nosplit
func (self class) GetSelectedText(ctx gd.Lifetime, caret_index gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_selected_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the caret index of the selection at the given [param line] and [param column], or [code]-1[/code] if there is none.
If [param include_edges] is [code]false[/code], the position must be inside the selection and not at either end. If [param only_selections] is [code]false[/code], carets without a selection will also be considered.
*/
//go:nosplit
func (self class) GetSelectionAtLineColumn(line gd.Int, column gd.Int, include_edges bool, only_selections bool) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, column)
	callframe.Arg(frame, include_edges)
	callframe.Arg(frame, only_selections)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_selection_at_line_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an [Array] of line ranges where [code]x[/code] is the first line and [code]y[/code] is the last line. All lines within these ranges will have a caret on them or be part of a selection. Each line will only be part of one line range, even if it has multiple carets on it.
If a selection's end column ([method get_selection_to_column]) is at column [code]0[/code], that line will not be included. If a selection begins on the line after another selection ends and [param merge_adjacent] is [code]true[/code], or they begin and end on the same line, one line range will include both selections.
*/
//go:nosplit
func (self class) GetLineRangesFromCarets(ctx gd.Lifetime, only_selections bool, merge_adjacent bool) gd.ArrayOf[gd.Vector2i] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, only_selections)
	callframe.Arg(frame, merge_adjacent)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_line_ranges_from_carets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Vector2i](ret)
}
/*
Returns the origin line of the selection. This is the opposite end from the caret.
*/
//go:nosplit
func (self class) GetSelectionOriginLine(caret_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_selection_origin_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the origin column of the selection. This is the opposite end from the caret.
*/
//go:nosplit
func (self class) GetSelectionOriginColumn(caret_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_selection_origin_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the selection origin line to the [param line] for the given [param caret_index]. If the selection origin is moved to the caret position, the selection will deselect.
If [param can_be_hidden] is [code]false[/code], The line will be set to the nearest unhidden line below or above.
If [param wrap_index] is [code]-1[/code], the selection origin column will be clamped to the [param line]'s length. If [param wrap_index] is greater than [code]-1[/code], the column will be moved to attempt to match the visual x position on the line's [param wrap_index] to the position from the last time [method set_selection_origin_column] or [method select] was called.
*/
//go:nosplit
func (self class) SetSelectionOriginLine(line gd.Int, can_be_hidden bool, wrap_index gd.Int, caret_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, can_be_hidden)
	callframe.Arg(frame, wrap_index)
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_selection_origin_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the selection origin column to the [param column] for the given [param caret_index]. If the selection origin is moved to the caret position, the selection will deselect.
*/
//go:nosplit
func (self class) SetSelectionOriginColumn(column gd.Int, caret_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_selection_origin_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the selection begin line. Returns the caret line if there is no selection.
*/
//go:nosplit
func (self class) GetSelectionFromLine(caret_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_selection_from_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the selection begin column. Returns the caret column if there is no selection.
*/
//go:nosplit
func (self class) GetSelectionFromColumn(caret_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_selection_from_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the selection end line. Returns the caret line if there is no selection.
*/
//go:nosplit
func (self class) GetSelectionToLine(caret_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_selection_to_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the selection end column. Returns the caret column if there is no selection.
*/
//go:nosplit
func (self class) GetSelectionToColumn(caret_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_selection_to_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the caret of the selection is after the selection origin. This can be used to determine the direction of the selection.
*/
//go:nosplit
func (self class) IsCaretAfterSelectionOrigin(caret_index gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_caret_after_selection_origin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Deselects the current selection.
*/
//go:nosplit
func (self class) Deselect(caret_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_deselect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Deletes the selected text.
*/
//go:nosplit
func (self class) DeleteSelection(caret_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_delete_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetLineWrappingMode(mode classdb.TextEditLineWrappingMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_line_wrapping_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLineWrappingMode() classdb.TextEditLineWrappingMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextEditLineWrappingMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_line_wrapping_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutowrapMode(autowrap_mode classdb.TextServerAutowrapMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, autowrap_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAutowrapMode() classdb.TextServerAutowrapMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerAutowrapMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns if the given line is wrapped.
*/
//go:nosplit
func (self class) IsLineWrapped(line gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_line_wrapped, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of times the given line is wrapped.
*/
//go:nosplit
func (self class) GetLineWrapCount(line gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_line_wrap_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the wrap index of the given line column.
*/
//go:nosplit
func (self class) GetLineWrapIndexAtColumn(line gd.Int, column gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_line_wrap_index_at_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an array of [String]s representing each wrapped index.
*/
//go:nosplit
func (self class) GetLineWrappedText(ctx gd.Lifetime, line gd.Int) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_line_wrapped_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSmoothScrollEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_smooth_scroll_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSmoothScrollEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_smooth_scroll_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [VScrollBar] of the [TextEdit].
*/
//go:nosplit
func (self class) GetVScrollBar(ctx gd.Lifetime) gdclass.VScrollBar {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_v_scroll_bar, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.VScrollBar
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the [HScrollBar] used by [TextEdit].
*/
//go:nosplit
func (self class) GetHScrollBar(ctx gd.Lifetime) gdclass.HScrollBar {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_h_scroll_bar, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.HScrollBar
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVScroll(value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_v_scroll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVScroll() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_v_scroll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHScroll(value gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_h_scroll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHScroll() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_h_scroll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetScrollPastEndOfFileEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_scroll_past_end_of_file_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsScrollPastEndOfFileEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_scroll_past_end_of_file_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVScrollSpeed(speed gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, speed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_v_scroll_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVScrollSpeed() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_v_scroll_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFitContentHeightEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_fit_content_height_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFitContentHeightEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_fit_content_height_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the scroll position for [param wrap_index] of [param line].
*/
//go:nosplit
func (self class) GetScrollPosForLine(line gd.Int, wrap_index gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, wrap_index)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_scroll_pos_for_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Positions the [param wrap_index] of [param line] at the top of the viewport.
*/
//go:nosplit
func (self class) SetLineAsFirstVisible(line gd.Int, wrap_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, wrap_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_line_as_first_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the first visible line.
*/
//go:nosplit
func (self class) GetFirstVisibleLine() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_first_visible_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Positions the [param wrap_index] of [param line] at the center of the viewport.
*/
//go:nosplit
func (self class) SetLineAsCenterVisible(line gd.Int, wrap_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, wrap_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_line_as_center_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Positions the [param wrap_index] of [param line] at the bottom of the viewport.
*/
//go:nosplit
func (self class) SetLineAsLastVisible(line gd.Int, wrap_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, wrap_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_line_as_last_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the last visible line. Use [method get_last_full_visible_line_wrap_index] for the wrap index.
*/
//go:nosplit
func (self class) GetLastFullVisibleLine() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_last_full_visible_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the last visible wrap index of the last visible line.
*/
//go:nosplit
func (self class) GetLastFullVisibleLineWrapIndex() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_last_full_visible_line_wrap_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of visible lines, including wrapped text.
*/
//go:nosplit
func (self class) GetVisibleLineCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_visible_line_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the total number of visible + wrapped lines between the two lines.
*/
//go:nosplit
func (self class) GetVisibleLineCountInRange(from_line gd.Int, to_line gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from_line)
	callframe.Arg(frame, to_line)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_visible_line_count_in_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of lines that may be drawn.
*/
//go:nosplit
func (self class) GetTotalVisibleLineCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_total_visible_line_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adjust the viewport so the caret is visible.
*/
//go:nosplit
func (self class) AdjustViewportToCaret(caret_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_adjust_viewport_to_caret, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Centers the viewport on the line the editing caret is at. This also resets the [member scroll_horizontal] value to [code]0[/code].
*/
//go:nosplit
func (self class) CenterViewportToCaret(caret_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_center_viewport_to_caret, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetDrawMinimap(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_draw_minimap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDrawingMinimap() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_drawing_minimap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMinimapWidth(width gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_minimap_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMinimapWidth() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_minimap_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of lines that may be drawn on the minimap.
*/
//go:nosplit
func (self class) GetMinimapVisibleLines() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_minimap_visible_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Register a new gutter to this [TextEdit]. Use [param at] to have a specific gutter order. A value of [code]-1[/code] appends the gutter to the right.
*/
//go:nosplit
func (self class) AddGutter(at gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, at)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_add_gutter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the gutter from this [TextEdit].
*/
//go:nosplit
func (self class) RemoveGutter(gutter gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_remove_gutter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of gutters registered.
*/
//go:nosplit
func (self class) GetGutterCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_gutter_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the name of the gutter.
*/
//go:nosplit
func (self class) SetGutterName(gutter gd.Int, name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_gutter_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the name of the gutter at the given index.
*/
//go:nosplit
func (self class) GetGutterName(ctx gd.Lifetime, gutter gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_gutter_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the type of gutter. Gutters can contain icons, text, or custom visuals. See [enum TextEdit.GutterType] for options.
*/
//go:nosplit
func (self class) SetGutterType(gutter gd.Int, atype classdb.TextEditGutterType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_gutter_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the type of the gutter at the given index. Gutters can contain icons, text, or custom visuals. See [enum TextEdit.GutterType] for options.
*/
//go:nosplit
func (self class) GetGutterType(gutter gd.Int) classdb.TextEditGutterType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	var r_ret = callframe.Ret[classdb.TextEditGutterType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_gutter_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Set the width of the gutter.
*/
//go:nosplit
func (self class) SetGutterWidth(gutter gd.Int, width gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_gutter_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the width of the gutter at the given index.
*/
//go:nosplit
func (self class) GetGutterWidth(gutter gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_gutter_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether the gutter should be drawn.
*/
//go:nosplit
func (self class) SetGutterDraw(gutter gd.Int, draw bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	callframe.Arg(frame, draw)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_gutter_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the gutter is currently drawn.
*/
//go:nosplit
func (self class) IsGutterDrawn(gutter gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_gutter_drawn, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the gutter as clickable. This will change the mouse cursor to a pointing hand when hovering over the gutter.
*/
//go:nosplit
func (self class) SetGutterClickable(gutter gd.Int, clickable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	callframe.Arg(frame, clickable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_gutter_clickable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the gutter is clickable.
*/
//go:nosplit
func (self class) IsGutterClickable(gutter gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_gutter_clickable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the gutter to overwritable. See [method merge_gutters].
*/
//go:nosplit
func (self class) SetGutterOverwritable(gutter gd.Int, overwritable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	callframe.Arg(frame, overwritable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_gutter_overwritable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the gutter is overwritable.
*/
//go:nosplit
func (self class) IsGutterOverwritable(gutter gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_gutter_overwritable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Merge the gutters from [param from_line] into [param to_line]. Only overwritable gutters will be copied.
*/
//go:nosplit
func (self class) MergeGutters(from_line gd.Int, to_line gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from_line)
	callframe.Arg(frame, to_line)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_merge_gutters, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Set a custom draw method for the gutter. The callback method must take the following args: [code]line: int, gutter: int, Area: Rect2[/code]. This only works when the gutter type is [constant GUTTER_TYPE_CUSTOM] (see [method set_gutter_type]).
*/
//go:nosplit
func (self class) SetGutterCustomDraw(column gd.Int, draw_callback gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mmm.Get(draw_callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_gutter_custom_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the total width of all gutters and internal padding.
*/
//go:nosplit
func (self class) GetTotalGutterWidth() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_total_gutter_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the metadata for [param gutter] on [param line] to [param metadata].
*/
//go:nosplit
func (self class) SetLineGutterMetadata(line gd.Int, gutter gd.Int, metadata gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, gutter)
	callframe.Arg(frame, mmm.Get(metadata))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_line_gutter_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the metadata currently in [param gutter] at [param line].
*/
//go:nosplit
func (self class) GetLineGutterMetadata(ctx gd.Lifetime, line gd.Int, gutter gd.Int) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, gutter)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_line_gutter_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the text for [param gutter] on [param line] to [param text]. This only works when the gutter type is [constant GUTTER_TYPE_STRING] (see [method set_gutter_type]).
*/
//go:nosplit
func (self class) SetLineGutterText(line gd.Int, gutter gd.Int, text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, gutter)
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_line_gutter_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the text currently in [param gutter] at [param line]. This only works when the gutter type is [constant GUTTER_TYPE_STRING] (see [method set_gutter_type]).
*/
//go:nosplit
func (self class) GetLineGutterText(ctx gd.Lifetime, line gd.Int, gutter gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, gutter)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_line_gutter_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the icon for [param gutter] on [param line] to [param icon]. This only works when the gutter type is [constant GUTTER_TYPE_ICON] (see [method set_gutter_type]).
*/
//go:nosplit
func (self class) SetLineGutterIcon(line gd.Int, gutter gd.Int, icon gdclass.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, gutter)
	callframe.Arg(frame, mmm.Get(icon[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_line_gutter_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the icon currently in [param gutter] at [param line]. This only works when the gutter type is [constant GUTTER_TYPE_ICON] (see [method set_gutter_type]).
*/
//go:nosplit
func (self class) GetLineGutterIcon(ctx gd.Lifetime, line gd.Int, gutter gd.Int) gdclass.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, gutter)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_line_gutter_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets the color for [param gutter] on [param line] to [param color].
*/
//go:nosplit
func (self class) SetLineGutterItemColor(line gd.Int, gutter gd.Int, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, gutter)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_line_gutter_item_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the color currently in [param gutter] at [param line].
*/
//go:nosplit
func (self class) GetLineGutterItemColor(line gd.Int, gutter gd.Int) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, gutter)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_line_gutter_item_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param clickable] is [code]true[/code], makes the [param gutter] on [param line] clickable. See [signal gutter_clicked].
*/
//go:nosplit
func (self class) SetLineGutterClickable(line gd.Int, gutter gd.Int, clickable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, gutter)
	callframe.Arg(frame, clickable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_line_gutter_clickable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the gutter on the given line is clickable.
*/
//go:nosplit
func (self class) IsLineGutterClickable(line gd.Int, gutter gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, gutter)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_line_gutter_clickable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the current background color of the line. Set to [code]Color(0, 0, 0, 0)[/code] for no color.
*/
//go:nosplit
func (self class) SetLineBackgroundColor(line gd.Int, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_line_background_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the current background color of the line. [code]Color(0, 0, 0, 0)[/code] is returned if no color is set.
*/
//go:nosplit
func (self class) GetLineBackgroundColor(line gd.Int) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_line_background_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSyntaxHighlighter(syntax_highlighter gdclass.SyntaxHighlighter)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(syntax_highlighter[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_syntax_highlighter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSyntaxHighlighter(ctx gd.Lifetime) gdclass.SyntaxHighlighter {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_syntax_highlighter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.SyntaxHighlighter
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHighlightCurrentLine(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_highlight_current_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsHighlightCurrentLineEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_highlight_current_line_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHighlightAllOccurrences(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_highlight_all_occurrences, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsHighlightAllOccurrencesEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_highlight_all_occurrences_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetDrawControlChars() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_draw_control_chars, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDrawControlChars(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_draw_control_chars, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetDrawTabs(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_draw_tabs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDrawingTabs() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_drawing_tabs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDrawSpaces(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_set_draw_spaces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDrawingSpaces() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_drawing_spaces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [PopupMenu] of this [TextEdit]. By default, this menu is displayed when right-clicking on the [TextEdit].
You can add custom menu items or remove standard ones. Make sure your IDs don't conflict with the standard ones (see [enum MenuItems]). For example:
[codeblocks]
[gdscript]
func _ready():
    var menu = get_menu()
    # Remove all items after "Redo".
    menu.item_count = menu.get_item_index(MENU_REDO) + 1
    # Add custom items.
    menu.add_separator()
    menu.add_item("Insert Date", MENU_MAX + 1)
    # Connect callback.
    menu.id_pressed.connect(_on_item_pressed)

func _on_item_pressed(id):
    if id == MENU_MAX + 1:
        insert_text_at_caret(Time.get_date_string_from_system())
[/gdscript]
[csharp]
public override void _Ready()
{
    var menu = GetMenu();
    // Remove all items after "Redo".
    menu.ItemCount = menu.GetItemIndex(TextEdit.MenuItems.Redo) + 1;
    // Add custom items.
    menu.AddSeparator();
    menu.AddItem("Insert Date", TextEdit.MenuItems.Max + 1);
    // Add event handler.
    menu.IdPressed += OnItemPressed;
}

public void OnItemPressed(int id)
{
    if (id == TextEdit.MenuItems.Max + 1)
    {
        InsertTextAtCaret(Time.GetDateStringFromSystem());
    }
}
[/csharp]
[/codeblocks]
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
//go:nosplit
func (self class) GetMenu(ctx gd.Lifetime) gdclass.PopupMenu {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.PopupMenu
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns whether the menu is visible. Use this instead of [code]get_menu().visible[/code] to improve performance (so the creation of the menu is avoided).
*/
//go:nosplit
func (self class) IsMenuVisible() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_is_menu_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Executes a given action as defined in the [enum MenuItems] enum.
*/
//go:nosplit
func (self class) MenuOption(option gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, option)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_menu_option, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
This method does nothing.
*/
//go:nosplit
func (self class) AdjustCaretsAfterEdit(caret gd.Int, from_line gd.Int, from_col gd.Int, to_line gd.Int, to_col gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret)
	callframe.Arg(frame, from_line)
	callframe.Arg(frame, from_col)
	callframe.Arg(frame, to_line)
	callframe.Arg(frame, to_col)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_adjust_carets_after_edit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a list of caret indexes in their edit order, this done from bottom to top. Edit order refers to the way actions such as [method insert_text_at_caret] are applied.
*/
//go:nosplit
func (self class) GetCaretIndexEditOrder(ctx gd.Lifetime) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_caret_index_edit_order, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the original start line of the selection.
*/
//go:nosplit
func (self class) GetSelectionLine(caret_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_selection_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the original start column of the selection.
*/
//go:nosplit
func (self class) GetSelectionColumn(caret_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TextEdit.Bind_get_selection_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnTextSet(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("text_set"), gc.Callable(cb), 0)
}


func (self Go) OnTextChanged(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("text_changed"), gc.Callable(cb), 0)
}


func (self Go) OnLinesEditedFrom(cb func(from_line int, to_line int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("lines_edited_from"), gc.Callable(cb), 0)
}


func (self Go) OnCaretChanged(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("caret_changed"), gc.Callable(cb), 0)
}


func (self Go) OnGutterClicked(cb func(line int, gutter int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("gutter_clicked"), gc.Callable(cb), 0)
}


func (self Go) OnGutterAdded(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("gutter_added"), gc.Callable(cb), 0)
}


func (self Go) OnGutterRemoved(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("gutter_removed"), gc.Callable(cb), 0)
}


func (self class) AsTextEdit() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTextEdit() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_handle_unicode_input": return reflect.ValueOf(self._handle_unicode_input);
	case "_backspace": return reflect.ValueOf(self._backspace);
	case "_cut": return reflect.ValueOf(self._cut);
	case "_copy": return reflect.ValueOf(self._copy);
	case "_paste": return reflect.ValueOf(self._paste);
	case "_paste_primary_clipboard": return reflect.ValueOf(self._paste_primary_clipboard);
	default: return gd.VirtualByName(self.AsControl(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_handle_unicode_input": return reflect.ValueOf(self._handle_unicode_input);
	case "_backspace": return reflect.ValueOf(self._backspace);
	case "_cut": return reflect.ValueOf(self._cut);
	case "_copy": return reflect.ValueOf(self._copy);
	case "_paste": return reflect.ValueOf(self._paste);
	case "_paste_primary_clipboard": return reflect.ValueOf(self._paste_primary_clipboard);
	default: return gd.VirtualByName(self.AsControl(), name)
	}
}
func init() {classdb.Register("TextEdit", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type MenuItems = classdb.TextEditMenuItems

const (
/*Cuts (copies and clears) the selected text.*/
	MenuCut MenuItems = 0
/*Copies the selected text.*/
	MenuCopy MenuItems = 1
/*Pastes the clipboard text over the selected text (or at the cursor's position).*/
	MenuPaste MenuItems = 2
/*Erases the whole [TextEdit] text.*/
	MenuClear MenuItems = 3
/*Selects the whole [TextEdit] text.*/
	MenuSelectAll MenuItems = 4
/*Undoes the previous action.*/
	MenuUndo MenuItems = 5
/*Redoes the previous action.*/
	MenuRedo MenuItems = 6
/*ID of "Text Writing Direction" submenu.*/
	MenuSubmenuTextDir MenuItems = 7
/*Sets text direction to inherited.*/
	MenuDirInherited MenuItems = 8
/*Sets text direction to automatic.*/
	MenuDirAuto MenuItems = 9
/*Sets text direction to left-to-right.*/
	MenuDirLtr MenuItems = 10
/*Sets text direction to right-to-left.*/
	MenuDirRtl MenuItems = 11
/*Toggles control character display.*/
	MenuDisplayUcc MenuItems = 12
/*ID of "Insert Control Character" submenu.*/
	MenuSubmenuInsertUcc MenuItems = 13
/*Inserts left-to-right mark (LRM) character.*/
	MenuInsertLrm MenuItems = 14
/*Inserts right-to-left mark (RLM) character.*/
	MenuInsertRlm MenuItems = 15
/*Inserts start of left-to-right embedding (LRE) character.*/
	MenuInsertLre MenuItems = 16
/*Inserts start of right-to-left embedding (RLE) character.*/
	MenuInsertRle MenuItems = 17
/*Inserts start of left-to-right override (LRO) character.*/
	MenuInsertLro MenuItems = 18
/*Inserts start of right-to-left override (RLO) character.*/
	MenuInsertRlo MenuItems = 19
/*Inserts pop direction formatting (PDF) character.*/
	MenuInsertPdf MenuItems = 20
/*Inserts Arabic letter mark (ALM) character.*/
	MenuInsertAlm MenuItems = 21
/*Inserts left-to-right isolate (LRI) character.*/
	MenuInsertLri MenuItems = 22
/*Inserts right-to-left isolate (RLI) character.*/
	MenuInsertRli MenuItems = 23
/*Inserts first strong isolate (FSI) character.*/
	MenuInsertFsi MenuItems = 24
/*Inserts pop direction isolate (PDI) character.*/
	MenuInsertPdi MenuItems = 25
/*Inserts zero width joiner (ZWJ) character.*/
	MenuInsertZwj MenuItems = 26
/*Inserts zero width non-joiner (ZWNJ) character.*/
	MenuInsertZwnj MenuItems = 27
/*Inserts word joiner (WJ) character.*/
	MenuInsertWj MenuItems = 28
/*Inserts soft hyphen (SHY) character.*/
	MenuInsertShy MenuItems = 29
/*Represents the size of the [enum MenuItems] enum.*/
	MenuMax MenuItems = 30
)
type EditAction = classdb.TextEditEditAction

const (
/*No current action.*/
	ActionNone EditAction = 0
/*A typing action.*/
	ActionTyping EditAction = 1
/*A backwards delete action.*/
	ActionBackspace EditAction = 2
/*A forward delete action.*/
	ActionDelete EditAction = 3
)
type SearchFlags = classdb.TextEditSearchFlags

const (
/*Match case when searching.*/
	SearchMatchCase SearchFlags = 1
/*Match whole words when searching.*/
	SearchWholeWords SearchFlags = 2
/*Search from end to beginning.*/
	SearchBackwards SearchFlags = 4
)
type CaretType = classdb.TextEditCaretType

const (
/*Vertical line caret.*/
	CaretTypeLine CaretType = 0
/*Block caret.*/
	CaretTypeBlock CaretType = 1
)
type SelectionMode = classdb.TextEditSelectionMode

const (
/*Not selecting.*/
	SelectionModeNone SelectionMode = 0
/*Select as if [code]shift[/code] is pressed.*/
	SelectionModeShift SelectionMode = 1
/*Select single characters as if the user single clicked.*/
	SelectionModePointer SelectionMode = 2
/*Select whole words as if the user double clicked.*/
	SelectionModeWord SelectionMode = 3
/*Select whole lines as if the user triple clicked.*/
	SelectionModeLine SelectionMode = 4
)
type LineWrappingMode = classdb.TextEditLineWrappingMode

const (
/*Line wrapping is disabled.*/
	LineWrappingNone LineWrappingMode = 0
/*Line wrapping occurs at the control boundary, beyond what would normally be visible.*/
	LineWrappingBoundary LineWrappingMode = 1
)
type GutterType = classdb.TextEditGutterType

const (
/*When a gutter is set to string using [method set_gutter_type], it is used to contain text set via the [method set_line_gutter_text] method.*/
	GutterTypeString GutterType = 0
/*When a gutter is set to icon using [method set_gutter_type], it is used to contain an icon set via the [method set_line_gutter_icon] method.*/
	GutterTypeIcon GutterType = 1
/*When a gutter is set to custom using [method set_gutter_type], it is used to contain custom visuals controlled by a callback method set via the [method set_gutter_custom_draw] method.*/
	GutterTypeCustom GutterType = 2
)