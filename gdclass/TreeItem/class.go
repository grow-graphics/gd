package TreeItem

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A single item of a [Tree] control. It can contain other [TreeItem]s as children, which allows it to create a hierarchy. It can also contain text and buttons. [TreeItem] is not a [Node], it is internal to the [Tree].
To create a [TreeItem], use [method Tree.create_item] or [method TreeItem.create_child]. To remove a [TreeItem], use [method Object.free].
[b]Note:[/b] The ID values used for buttons are 32-bit, unlike [int] which is always 64-bit. They go from [code]-2147483648[/code] to [code]2147483647[/code].

*/
type Go [1]classdb.TreeItem

/*
Sets the given column's cell mode to [param mode]. This determines how the cell is displayed and edited. See [enum TreeCellMode] constants for details.
*/
func (self Go) SetCellMode(column int, mode classdb.TreeItemTreeCellMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCellMode(gd.Int(column), mode)
}

/*
Returns the column's cell mode.
*/
func (self Go) GetCellMode(column int) classdb.TreeItemTreeCellMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TreeItemTreeCellMode(class(self).GetCellMode(gd.Int(column)))
}

/*
If [param multiline] is [code]true[/code], the given [param column] is multiline editable.
[b]Note:[/b] This option only affects the type of control ([LineEdit] or [TextEdit]) that appears when editing the column. You can set multiline values with [method set_text] even if the column is not multiline editable.
*/
func (self Go) SetEditMultiline(column int, multiline bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEditMultiline(gd.Int(column), multiline)
}

/*
Returns [code]true[/code] if the given [param column] is multiline editable.
*/
func (self Go) IsEditMultiline(column int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsEditMultiline(gd.Int(column)))
}

/*
If [param checked] is [code]true[/code], the given [param column] is checked. Clears column's indeterminate status.
*/
func (self Go) SetChecked(column int, checked bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetChecked(gd.Int(column), checked)
}

/*
If [param indeterminate] is [code]true[/code], the given [param column] is marked indeterminate.
[b]Note:[/b] If set [code]true[/code] from [code]false[/code], then column is cleared of checked status.
*/
func (self Go) SetIndeterminate(column int, indeterminate bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIndeterminate(gd.Int(column), indeterminate)
}

/*
Returns [code]true[/code] if the given [param column] is checked.
*/
func (self Go) IsChecked(column int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsChecked(gd.Int(column)))
}

/*
Returns [code]true[/code] if the given [param column] is indeterminate.
*/
func (self Go) IsIndeterminate(column int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsIndeterminate(gd.Int(column)))
}

/*
Propagates this item's checked status to its children and parents for the given [param column]. It is possible to process the items affected by this method call by connecting to [signal Tree.check_propagated_to_item]. The order that the items affected will be processed is as follows: the item invoking this method, children of that item, and finally parents of that item. If [param emit_signal] is [code]false[/code], then [signal Tree.check_propagated_to_item] will not be emitted.
*/
func (self Go) PropagateCheck(column int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PropagateCheck(gd.Int(column), true)
}

/*
Sets the given column's text value.
*/
func (self Go) SetText(column int, text string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetText(gd.Int(column), gc.String(text))
}

/*
Returns the given column's text.
*/
func (self Go) GetText(column int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetText(gc, gd.Int(column)).String())
}

/*
Sets item's text base writing direction.
*/
func (self Go) SetTextDirection(column int, direction classdb.ControlTextDirection) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextDirection(gd.Int(column), direction)
}

/*
Returns item's text base writing direction.
*/
func (self Go) GetTextDirection(column int) classdb.ControlTextDirection {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ControlTextDirection(class(self).GetTextDirection(gd.Int(column)))
}

/*
Sets the autowrap mode in the given [param column]. If set to something other than [constant TextServer.AUTOWRAP_OFF], the text gets wrapped inside the cell's bounding rectangle.
*/
func (self Go) SetAutowrapMode(column int, autowrap_mode classdb.TextServerAutowrapMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAutowrapMode(gd.Int(column), autowrap_mode)
}

/*
Returns the text autowrap mode in the given [param column]. By default it is [constant TextServer.AUTOWRAP_OFF].
*/
func (self Go) GetAutowrapMode(column int) classdb.TextServerAutowrapMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextServerAutowrapMode(class(self).GetAutowrapMode(gd.Int(column)))
}

/*
Sets the clipping behavior when the text exceeds the item's bounding rectangle in the given [param column].
*/
func (self Go) SetTextOverrunBehavior(column int, overrun_behavior classdb.TextServerOverrunBehavior) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextOverrunBehavior(gd.Int(column), overrun_behavior)
}

/*
Returns the clipping behavior when the text exceeds the item's bounding rectangle in the given [param column]. By default it is [constant TextServer.OVERRUN_TRIM_ELLIPSIS].
*/
func (self Go) GetTextOverrunBehavior(column int) classdb.TextServerOverrunBehavior {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextServerOverrunBehavior(class(self).GetTextOverrunBehavior(gd.Int(column)))
}

/*
Set BiDi algorithm override for the structured text. Has effect for cells that display text.
*/
func (self Go) SetStructuredTextBidiOverride(column int, parser classdb.TextServerStructuredTextParser) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStructuredTextBidiOverride(gd.Int(column), parser)
}

/*
Returns the BiDi algorithm override set for this cell.
*/
func (self Go) GetStructuredTextBidiOverride(column int) classdb.TextServerStructuredTextParser {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextServerStructuredTextParser(class(self).GetStructuredTextBidiOverride(gd.Int(column)))
}

/*
Set additional options for BiDi override. Has effect for cells that display text.
*/
func (self Go) SetStructuredTextBidiOverrideOptions(column int, args gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStructuredTextBidiOverrideOptions(gd.Int(column), args)
}

/*
Returns the additional BiDi options set for this cell.
*/
func (self Go) GetStructuredTextBidiOverrideOptions(column int) gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(class(self).GetStructuredTextBidiOverrideOptions(gc, gd.Int(column)))
}

/*
Sets language code of item's text used for line-breaking and text shaping algorithms, if left empty current locale is used instead.
*/
func (self Go) SetLanguage(column int, language string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLanguage(gd.Int(column), gc.String(language))
}

/*
Returns item's text language code.
*/
func (self Go) GetLanguage(column int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetLanguage(gc, gd.Int(column)).String())
}

/*
Sets a string to be shown after a column's value (for example, a unit abbreviation).
*/
func (self Go) SetSuffix(column int, text string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSuffix(gd.Int(column), gc.String(text))
}

/*
Gets the suffix string shown after the column value.
*/
func (self Go) GetSuffix(column int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetSuffix(gc, gd.Int(column)).String())
}

/*
Sets the given cell's icon [Texture2D]. The cell has to be in [constant CELL_MODE_ICON] mode.
*/
func (self Go) SetIcon(column int, texture gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIcon(gd.Int(column), texture)
}

/*
Returns the given column's icon [Texture2D]. Error if no icon is set.
*/
func (self Go) GetIcon(column int) gdclass.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Texture2D(class(self).GetIcon(gc, gd.Int(column)))
}

/*
Sets the given column's icon's texture region.
*/
func (self Go) SetIconRegion(column int, region gd.Rect2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIconRegion(gd.Int(column), region)
}

/*
Returns the icon [Texture2D] region as [Rect2].
*/
func (self Go) GetIconRegion(column int) gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(class(self).GetIconRegion(gd.Int(column)))
}

/*
Sets the maximum allowed width of the icon in the given [param column]. This limit is applied on top of the default size of the icon and on top of [theme_item Tree.icon_max_width]. The height is adjusted according to the icon's ratio.
*/
func (self Go) SetIconMaxWidth(column int, width int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIconMaxWidth(gd.Int(column), gd.Int(width))
}

/*
Returns the maximum allowed width of the icon in the given [param column].
*/
func (self Go) GetIconMaxWidth(column int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetIconMaxWidth(gd.Int(column))))
}

/*
Modulates the given column's icon with [param modulate].
*/
func (self Go) SetIconModulate(column int, modulate gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIconModulate(gd.Int(column), modulate)
}

/*
Returns the [Color] modulating the column's icon.
*/
func (self Go) GetIconModulate(column int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(class(self).GetIconModulate(gd.Int(column)))
}

/*
Sets the value of a [constant CELL_MODE_RANGE] column.
*/
func (self Go) SetRange(column int, value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRange(gd.Int(column), gd.Float(value))
}

/*
Returns the value of a [constant CELL_MODE_RANGE] column.
*/
func (self Go) GetRange(column int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetRange(gd.Int(column))))
}

/*
Sets the range of accepted values for a column. The column must be in the [constant CELL_MODE_RANGE] mode.
If [param expr] is [code]true[/code], the edit mode slider will use an exponential scale as with [member Range.exp_edit].
*/
func (self Go) SetRangeConfig(column int, min float64, max float64, step float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRangeConfig(gd.Int(column), gd.Float(min), gd.Float(max), gd.Float(step), false)
}

/*
Returns a dictionary containing the range parameters for a given column. The keys are "min", "max", "step", and "expr".
*/
func (self Go) GetRangeConfig(column int) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(class(self).GetRangeConfig(gc, gd.Int(column)))
}

/*
Sets the metadata value for the given column, which can be retrieved later using [method get_metadata]. This can be used, for example, to store a reference to the original data.
*/
func (self Go) SetMetadata(column int, meta gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMetadata(gd.Int(column), meta)
}

/*
Returns the metadata value that was set for the given column using [method set_metadata].
*/
func (self Go) GetMetadata(column int) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(class(self).GetMetadata(gc, gd.Int(column)))
}

/*
Sets the given column's custom draw callback to the [param callback] method on [param object].
The method named [param callback] should accept two arguments: the [TreeItem] that is drawn and its position and size as a [Rect2].
*/
func (self Go) SetCustomDraw(column int, obj gd.Object, callback string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCustomDraw(gd.Int(column), obj, gc.StringName(callback))
}

/*
Sets the given column's custom draw callback. Use an empty [Callable] ([code skip-lint]Callable()[/code]) to clear the custom callback. The cell has to be in [constant CELL_MODE_CUSTOM] to use this feature.
The [param callback] should accept two arguments: the [TreeItem] that is drawn and its position and size as a [Rect2].
*/
func (self Go) SetCustomDrawCallback(column int, callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCustomDrawCallback(gd.Int(column), callback)
}

/*
Returns the custom callback of column [param column].
*/
func (self Go) GetCustomDrawCallback(column int) gd.Callable {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Callable(class(self).GetCustomDrawCallback(gc, gd.Int(column)))
}

/*
Collapses or uncollapses this [TreeItem] and all the descendants of this item.
*/
func (self Go) SetCollapsedRecursive(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollapsedRecursive(enable)
}

/*
Returns [code]true[/code] if this [TreeItem], or any of its descendants, is collapsed.
If [param only_visible] is [code]true[/code] it ignores non-visible [TreeItem]s.
*/
func (self Go) IsAnyCollapsed() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsAnyCollapsed(false))
}

/*
Returns [code]true[/code] if [member visible] is [code]true[/code] and all its ancestors are also visible.
*/
func (self Go) IsVisibleInTree() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsVisibleInTree())
}

/*
Uncollapses all [TreeItem]s necessary to reveal this [TreeItem], i.e. all ancestor [TreeItem]s.
*/
func (self Go) UncollapseTree() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).UncollapseTree()
}

/*
If [param selectable] is [code]true[/code], the given [param column] is selectable.
*/
func (self Go) SetSelectable(column int, selectable bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSelectable(gd.Int(column), selectable)
}

/*
Returns [code]true[/code] if the given [param column] is selectable.
*/
func (self Go) IsSelectable(column int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsSelectable(gd.Int(column)))
}

/*
Returns [code]true[/code] if the given [param column] is selected.
*/
func (self Go) IsSelected(column int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsSelected(gd.Int(column)))
}

/*
Selects the given [param column].
*/
func (self Go) Select(column int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Select(gd.Int(column))
}

/*
Deselects the given column.
*/
func (self Go) Deselect(column int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Deselect(gd.Int(column))
}

/*
If [param enabled] is [code]true[/code], the given [param column] is editable.
*/
func (self Go) SetEditable(column int, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEditable(gd.Int(column), enabled)
}

/*
Returns [code]true[/code] if the given [param column] is editable.
*/
func (self Go) IsEditable(column int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsEditable(gd.Int(column)))
}

/*
Sets the given column's custom color.
*/
func (self Go) SetCustomColor(column int, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCustomColor(gd.Int(column), color)
}

/*
Returns the custom color of column [param column].
*/
func (self Go) GetCustomColor(column int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(class(self).GetCustomColor(gd.Int(column)))
}

/*
Resets the color for the given column to default.
*/
func (self Go) ClearCustomColor(column int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearCustomColor(gd.Int(column))
}

/*
Sets custom font used to draw text in the given [param column].
*/
func (self Go) SetCustomFont(column int, font gdclass.Font) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCustomFont(gd.Int(column), font)
}

/*
Returns custom font used to draw text in the column [param column].
*/
func (self Go) GetCustomFont(column int) gdclass.Font {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Font(class(self).GetCustomFont(gc, gd.Int(column)))
}

/*
Sets custom font size used to draw text in the given [param column].
*/
func (self Go) SetCustomFontSize(column int, font_size int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCustomFontSize(gd.Int(column), gd.Int(font_size))
}

/*
Returns custom font size used to draw text in the column [param column].
*/
func (self Go) GetCustomFontSize(column int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCustomFontSize(gd.Int(column))))
}

/*
Sets the given column's custom background color and whether to just use it as an outline.
*/
func (self Go) SetCustomBgColor(column int, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCustomBgColor(gd.Int(column), color, false)
}

/*
Resets the background color for the given column to default.
*/
func (self Go) ClearCustomBgColor(column int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearCustomBgColor(gd.Int(column))
}

/*
Returns the custom background color of column [param column].
*/
func (self Go) GetCustomBgColor(column int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(class(self).GetCustomBgColor(gd.Int(column)))
}

/*
Makes a cell with [constant CELL_MODE_CUSTOM] display as a non-flat button with a [StyleBox].
*/
func (self Go) SetCustomAsButton(column int, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCustomAsButton(gd.Int(column), enable)
}

/*
Returns [code]true[/code] if the cell was made into a button with [method set_custom_as_button].
*/
func (self Go) IsCustomSetAsButton(column int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsCustomSetAsButton(gd.Int(column)))
}

/*
Adds a button with [Texture2D] [param button] at column [param column]. The [param id] is used to identify the button in the according [signal Tree.button_clicked] signal and can be different from the buttons index. If not specified, the next available index is used, which may be retrieved by calling [method get_button_count] immediately before this method. Optionally, the button can be [param disabled] and have a [param tooltip_text].
*/
func (self Go) AddButton(column int, button gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddButton(gd.Int(column), button, gd.Int(-1), false, gc.String(""))
}

/*
Returns the number of buttons in column [param column].
*/
func (self Go) GetButtonCount(column int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetButtonCount(gd.Int(column))))
}

/*
Returns the tooltip text for the button at index [param button_index] in column [param column].
*/
func (self Go) GetButtonTooltipText(column int, button_index int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetButtonTooltipText(gc, gd.Int(column), gd.Int(button_index)).String())
}

/*
Returns the ID for the button at index [param button_index] in column [param column].
*/
func (self Go) GetButtonId(column int, button_index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetButtonId(gd.Int(column), gd.Int(button_index))))
}

/*
Returns the button index if there is a button with ID [param id] in column [param column], otherwise returns -1.
*/
func (self Go) GetButtonById(column int, id int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetButtonById(gd.Int(column), gd.Int(id))))
}

/*
Returns the color of the button with ID [param id] in column [param column]. If the specified button does not exist, returns [constant Color.BLACK].
*/
func (self Go) GetButtonColor(column int, id int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(class(self).GetButtonColor(gd.Int(column), gd.Int(id)))
}

/*
Returns the [Texture2D] of the button at index [param button_index] in column [param column].
*/
func (self Go) GetButton(column int, button_index int) gdclass.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Texture2D(class(self).GetButton(gc, gd.Int(column), gd.Int(button_index)))
}

/*
Sets the tooltip text for the button at index [param button_index] in the given [param column].
*/
func (self Go) SetButtonTooltipText(column int, button_index int, tooltip string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetButtonTooltipText(gd.Int(column), gd.Int(button_index), gc.String(tooltip))
}

/*
Sets the given column's button [Texture2D] at index [param button_index] to [param button].
*/
func (self Go) SetButton(column int, button_index int, button gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetButton(gd.Int(column), gd.Int(button_index), button)
}

/*
Removes the button at index [param button_index] in column [param column].
*/
func (self Go) EraseButton(column int, button_index int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).EraseButton(gd.Int(column), gd.Int(button_index))
}

/*
If [code]true[/code], disables the button at index [param button_index] in the given [param column].
*/
func (self Go) SetButtonDisabled(column int, button_index int, disabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetButtonDisabled(gd.Int(column), gd.Int(button_index), disabled)
}

/*
Sets the given column's button color at index [param button_index] to [param color].
*/
func (self Go) SetButtonColor(column int, button_index int, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetButtonColor(gd.Int(column), gd.Int(button_index), color)
}

/*
Returns [code]true[/code] if the button at index [param button_index] for the given [param column] is disabled.
*/
func (self Go) IsButtonDisabled(column int, button_index int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsButtonDisabled(gd.Int(column), gd.Int(button_index)))
}

/*
Sets the given column's tooltip text.
*/
func (self Go) SetTooltipText(column int, tooltip string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTooltipText(gd.Int(column), gc.String(tooltip))
}

/*
Returns the given column's tooltip text.
*/
func (self Go) GetTooltipText(column int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetTooltipText(gc, gd.Int(column)).String())
}

/*
Sets the given column's text alignment. See [enum HorizontalAlignment] for possible values.
*/
func (self Go) SetTextAlignment(column int, text_alignment gd.HorizontalAlignment) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextAlignment(gd.Int(column), text_alignment)
}

/*
Returns the given column's text alignment.
*/
func (self Go) GetTextAlignment(column int) gd.HorizontalAlignment {
	gc := gd.GarbageCollector(); _ = gc
	return gd.HorizontalAlignment(class(self).GetTextAlignment(gd.Int(column)))
}

/*
If [param enable] is [code]true[/code], the given [param column] is expanded to the right.
*/
func (self Go) SetExpandRight(column int, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetExpandRight(gd.Int(column), enable)
}

/*
Returns [code]true[/code] if [code]expand_right[/code] is set.
*/
func (self Go) GetExpandRight(column int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetExpandRight(gd.Int(column)))
}

/*
Creates an item and adds it as a child.
The new item will be inserted as position [param index] (the default value [code]-1[/code] means the last position), or it will be the last child if [param index] is higher than the child count.
*/
func (self Go) CreateChild() gdclass.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.TreeItem(class(self).CreateChild(gc, gd.Int(-1)))
}

/*
Adds a previously unparented [TreeItem] as a direct child of this one. The [param child] item must not be a part of any [Tree] or parented to any [TreeItem]. See also [method remove_child].
*/
func (self Go) AddChild(child gdclass.TreeItem) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddChild(child)
}

/*
Removes the given child [TreeItem] and all its children from the [Tree]. Note that it doesn't free the item from memory, so it can be reused later (see [method add_child]). To completely remove a [TreeItem] use [method Object.free].
[b]Note:[/b] If you want to move a child from one [Tree] to another, then instead of removing and adding it manually you can use [method move_before] or [method move_after].
*/
func (self Go) RemoveChild(child gdclass.TreeItem) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveChild(child)
}

/*
Returns the [Tree] that owns this TreeItem.
*/
func (self Go) GetTree() gdclass.Tree {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Tree(class(self).GetTree(gc))
}

/*
Returns the next sibling TreeItem in the tree or a null object if there is none.
*/
func (self Go) GetNext() gdclass.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.TreeItem(class(self).GetNext(gc))
}

/*
Returns the previous sibling TreeItem in the tree or a null object if there is none.
*/
func (self Go) GetPrev() gdclass.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.TreeItem(class(self).GetPrev(gc))
}

/*
Returns the parent TreeItem or a null object if there is none.
*/
func (self Go) GetParent() gdclass.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.TreeItem(class(self).GetParent(gc))
}

/*
Returns the TreeItem's first child.
*/
func (self Go) GetFirstChild() gdclass.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.TreeItem(class(self).GetFirstChild(gc))
}

/*
Returns the next TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the first element in the tree when called on the last element, otherwise it returns [code]null[/code].
*/
func (self Go) GetNextInTree() gdclass.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.TreeItem(class(self).GetNextInTree(gc, false))
}

/*
Returns the previous TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the last element in the tree when called on the first visible element, otherwise it returns [code]null[/code].
*/
func (self Go) GetPrevInTree() gdclass.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.TreeItem(class(self).GetPrevInTree(gc, false))
}

/*
Returns the next visible TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the first visible element in the tree when called on the last visible element, otherwise it returns [code]null[/code].
*/
func (self Go) GetNextVisible() gdclass.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.TreeItem(class(self).GetNextVisible(gc, false))
}

/*
Returns the previous visible sibling TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the last visible element in the tree when called on the first visible element, otherwise it returns [code]null[/code].
*/
func (self Go) GetPrevVisible() gdclass.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.TreeItem(class(self).GetPrevVisible(gc, false))
}

/*
Returns a child item by its [param index] (see [method get_child_count]). This method is often used for iterating all children of an item.
Negative indices access the children from the last one.
*/
func (self Go) GetChild(index int) gdclass.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.TreeItem(class(self).GetChild(gc, gd.Int(index)))
}

/*
Returns the number of child items.
*/
func (self Go) GetChildCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetChildCount()))
}

/*
Returns an array of references to the item's children.
*/
func (self Go) GetChildren() gd.ArrayOf[gdclass.TreeItem] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gdclass.TreeItem](class(self).GetChildren(gc))
}

/*
Returns the node's order in the tree. For example, if called on the first child item the position is [code]0[/code].
*/
func (self Go) GetIndex() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetIndex()))
}

/*
Moves this TreeItem right before the given [param item].
[b]Note:[/b] You can't move to the root or move the root.
*/
func (self Go) MoveBefore(item gdclass.TreeItem) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).MoveBefore(item)
}

/*
Moves this TreeItem right after the given [param item].
[b]Note:[/b] You can't move to the root or move the root.
*/
func (self Go) MoveAfter(item gdclass.TreeItem) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).MoveAfter(item)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.TreeItem
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("TreeItem"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Collapsed() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsCollapsed())
}

func (self Go) SetCollapsed(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCollapsed(value)
}

func (self Go) Visible() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsVisible())
}

func (self Go) SetVisible(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVisible(value)
}

func (self Go) DisableFolding() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsFoldingDisabled())
}

func (self Go) SetDisableFolding(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDisableFolding(value)
}

func (self Go) CustomMinimumHeight() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCustomMinimumHeight()))
}

func (self Go) SetCustomMinimumHeight(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCustomMinimumHeight(gd.Int(value))
}

/*
Sets the given column's cell mode to [param mode]. This determines how the cell is displayed and edited. See [enum TreeCellMode] constants for details.
*/
//go:nosplit
func (self class) SetCellMode(column gd.Int, mode classdb.TreeItemTreeCellMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_cell_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the column's cell mode.
*/
//go:nosplit
func (self class) GetCellMode(column gd.Int) classdb.TreeItemTreeCellMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[classdb.TreeItemTreeCellMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_cell_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param multiline] is [code]true[/code], the given [param column] is multiline editable.
[b]Note:[/b] This option only affects the type of control ([LineEdit] or [TextEdit]) that appears when editing the column. You can set multiline values with [method set_text] even if the column is not multiline editable.
*/
//go:nosplit
func (self class) SetEditMultiline(column gd.Int, multiline bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, multiline)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_edit_multiline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the given [param column] is multiline editable.
*/
//go:nosplit
func (self class) IsEditMultiline(column gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_edit_multiline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param checked] is [code]true[/code], the given [param column] is checked. Clears column's indeterminate status.
*/
//go:nosplit
func (self class) SetChecked(column gd.Int, checked bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, checked)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_checked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [param indeterminate] is [code]true[/code], the given [param column] is marked indeterminate.
[b]Note:[/b] If set [code]true[/code] from [code]false[/code], then column is cleared of checked status.
*/
//go:nosplit
func (self class) SetIndeterminate(column gd.Int, indeterminate bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, indeterminate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_indeterminate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the given [param column] is checked.
*/
//go:nosplit
func (self class) IsChecked(column gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_checked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the given [param column] is indeterminate.
*/
//go:nosplit
func (self class) IsIndeterminate(column gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_indeterminate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Propagates this item's checked status to its children and parents for the given [param column]. It is possible to process the items affected by this method call by connecting to [signal Tree.check_propagated_to_item]. The order that the items affected will be processed is as follows: the item invoking this method, children of that item, and finally parents of that item. If [param emit_signal] is [code]false[/code], then [signal Tree.check_propagated_to_item] will not be emitted.
*/
//go:nosplit
func (self class) PropagateCheck(column gd.Int, emit_signal bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, emit_signal)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_propagate_check, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the given column's text value.
*/
//go:nosplit
func (self class) SetText(column gd.Int, text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the given column's text.
*/
//go:nosplit
func (self class) GetText(ctx gd.Lifetime, column gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets item's text base writing direction.
*/
//go:nosplit
func (self class) SetTextDirection(column gd.Int, direction classdb.ControlTextDirection)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns item's text base writing direction.
*/
//go:nosplit
func (self class) GetTextDirection(column gd.Int) classdb.ControlTextDirection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[classdb.ControlTextDirection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the autowrap mode in the given [param column]. If set to something other than [constant TextServer.AUTOWRAP_OFF], the text gets wrapped inside the cell's bounding rectangle.
*/
//go:nosplit
func (self class) SetAutowrapMode(column gd.Int, autowrap_mode classdb.TextServerAutowrapMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, autowrap_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the text autowrap mode in the given [param column]. By default it is [constant TextServer.AUTOWRAP_OFF].
*/
//go:nosplit
func (self class) GetAutowrapMode(column gd.Int) classdb.TextServerAutowrapMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[classdb.TextServerAutowrapMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the clipping behavior when the text exceeds the item's bounding rectangle in the given [param column].
*/
//go:nosplit
func (self class) SetTextOverrunBehavior(column gd.Int, overrun_behavior classdb.TextServerOverrunBehavior)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, overrun_behavior)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the clipping behavior when the text exceeds the item's bounding rectangle in the given [param column]. By default it is [constant TextServer.OVERRUN_TRIM_ELLIPSIS].
*/
//go:nosplit
func (self class) GetTextOverrunBehavior(column gd.Int) classdb.TextServerOverrunBehavior {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[classdb.TextServerOverrunBehavior](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Set BiDi algorithm override for the structured text. Has effect for cells that display text.
*/
//go:nosplit
func (self class) SetStructuredTextBidiOverride(column gd.Int, parser classdb.TextServerStructuredTextParser)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, parser)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the BiDi algorithm override set for this cell.
*/
//go:nosplit
func (self class) GetStructuredTextBidiOverride(column gd.Int) classdb.TextServerStructuredTextParser {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[classdb.TextServerStructuredTextParser](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Set additional options for BiDi override. Has effect for cells that display text.
*/
//go:nosplit
func (self class) SetStructuredTextBidiOverrideOptions(column gd.Int, args gd.Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mmm.Get(args))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the additional BiDi options set for this cell.
*/
//go:nosplit
func (self class) GetStructuredTextBidiOverrideOptions(ctx gd.Lifetime, column gd.Int) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets language code of item's text used for line-breaking and text shaping algorithms, if left empty current locale is used instead.
*/
//go:nosplit
func (self class) SetLanguage(column gd.Int, language gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mmm.Get(language))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns item's text language code.
*/
//go:nosplit
func (self class) GetLanguage(ctx gd.Lifetime, column gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets a string to be shown after a column's value (for example, a unit abbreviation).
*/
//go:nosplit
func (self class) SetSuffix(column gd.Int, text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_suffix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Gets the suffix string shown after the column value.
*/
//go:nosplit
func (self class) GetSuffix(ctx gd.Lifetime, column gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_suffix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the given cell's icon [Texture2D]. The cell has to be in [constant CELL_MODE_ICON] mode.
*/
//go:nosplit
func (self class) SetIcon(column gd.Int, texture gdclass.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the given column's icon [Texture2D]. Error if no icon is set.
*/
//go:nosplit
func (self class) GetIcon(ctx gd.Lifetime, column gd.Int) gdclass.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets the given column's icon's texture region.
*/
//go:nosplit
func (self class) SetIconRegion(column gd.Int, region gd.Rect2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, region)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_icon_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the icon [Texture2D] region as [Rect2].
*/
//go:nosplit
func (self class) GetIconRegion(column gd.Int) gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_icon_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the maximum allowed width of the icon in the given [param column]. This limit is applied on top of the default size of the icon and on top of [theme_item Tree.icon_max_width]. The height is adjusted according to the icon's ratio.
*/
//go:nosplit
func (self class) SetIconMaxWidth(column gd.Int, width gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_icon_max_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the maximum allowed width of the icon in the given [param column].
*/
//go:nosplit
func (self class) GetIconMaxWidth(column gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_icon_max_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Modulates the given column's icon with [param modulate].
*/
//go:nosplit
func (self class) SetIconModulate(column gd.Int, modulate gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, modulate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_icon_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Color] modulating the column's icon.
*/
//go:nosplit
func (self class) GetIconModulate(column gd.Int) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_icon_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the value of a [constant CELL_MODE_RANGE] column.
*/
//go:nosplit
func (self class) SetRange(column gd.Int, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of a [constant CELL_MODE_RANGE] column.
*/
//go:nosplit
func (self class) GetRange(column gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the range of accepted values for a column. The column must be in the [constant CELL_MODE_RANGE] mode.
If [param expr] is [code]true[/code], the edit mode slider will use an exponential scale as with [member Range.exp_edit].
*/
//go:nosplit
func (self class) SetRangeConfig(column gd.Int, min gd.Float, max gd.Float, step gd.Float, expr bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, min)
	callframe.Arg(frame, max)
	callframe.Arg(frame, step)
	callframe.Arg(frame, expr)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_range_config, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a dictionary containing the range parameters for a given column. The keys are "min", "max", "step", and "expr".
*/
//go:nosplit
func (self class) GetRangeConfig(ctx gd.Lifetime, column gd.Int) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_range_config, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the metadata value for the given column, which can be retrieved later using [method get_metadata]. This can be used, for example, to store a reference to the original data.
*/
//go:nosplit
func (self class) SetMetadata(column gd.Int, meta gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mmm.Get(meta))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the metadata value that was set for the given column using [method set_metadata].
*/
//go:nosplit
func (self class) GetMetadata(ctx gd.Lifetime, column gd.Int) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the given column's custom draw callback to the [param callback] method on [param object].
The method named [param callback] should accept two arguments: the [TreeItem] that is drawn and its position and size as a [Rect2].
*/
//go:nosplit
func (self class) SetCustomDraw(column gd.Int, obj gd.Object, callback gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mmm.End(obj.AsPointer())[0])
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_custom_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the given column's custom draw callback. Use an empty [Callable] ([code skip-lint]Callable()[/code]) to clear the custom callback. The cell has to be in [constant CELL_MODE_CUSTOM] to use this feature.
The [param callback] should accept two arguments: the [TreeItem] that is drawn and its position and size as a [Rect2].
*/
//go:nosplit
func (self class) SetCustomDrawCallback(column gd.Int, callback gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_custom_draw_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the custom callback of column [param column].
*/
//go:nosplit
func (self class) GetCustomDrawCallback(ctx gd.Lifetime, column gd.Int) gd.Callable {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_custom_draw_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Callable](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollapsed(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_collapsed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCollapsed() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_collapsed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Collapses or uncollapses this [TreeItem] and all the descendants of this item.
*/
//go:nosplit
func (self class) SetCollapsedRecursive(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_collapsed_recursive, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if this [TreeItem], or any of its descendants, is collapsed.
If [param only_visible] is [code]true[/code] it ignores non-visible [TreeItem]s.
*/
//go:nosplit
func (self class) IsAnyCollapsed(only_visible bool) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, only_visible)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_any_collapsed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVisible(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsVisible() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if [member visible] is [code]true[/code] and all its ancestors are also visible.
*/
//go:nosplit
func (self class) IsVisibleInTree() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_visible_in_tree, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Uncollapses all [TreeItem]s necessary to reveal this [TreeItem], i.e. all ancestor [TreeItem]s.
*/
//go:nosplit
func (self class) UncollapseTree()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_uncollapse_tree, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetCustomMinimumHeight(height gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_custom_minimum_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCustomMinimumHeight() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_custom_minimum_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param selectable] is [code]true[/code], the given [param column] is selectable.
*/
//go:nosplit
func (self class) SetSelectable(column gd.Int, selectable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, selectable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_selectable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the given [param column] is selectable.
*/
//go:nosplit
func (self class) IsSelectable(column gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_selectable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the given [param column] is selected.
*/
//go:nosplit
func (self class) IsSelected(column gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Selects the given [param column].
*/
//go:nosplit
func (self class) Select(column gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_select_, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Deselects the given column.
*/
//go:nosplit
func (self class) Deselect(column gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_deselect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [param enabled] is [code]true[/code], the given [param column] is editable.
*/
//go:nosplit
func (self class) SetEditable(column gd.Int, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the given [param column] is editable.
*/
//go:nosplit
func (self class) IsEditable(column gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the given column's custom color.
*/
//go:nosplit
func (self class) SetCustomColor(column gd.Int, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_custom_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the custom color of column [param column].
*/
//go:nosplit
func (self class) GetCustomColor(column gd.Int) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_custom_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Resets the color for the given column to default.
*/
//go:nosplit
func (self class) ClearCustomColor(column gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_clear_custom_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets custom font used to draw text in the given [param column].
*/
//go:nosplit
func (self class) SetCustomFont(column gd.Int, font gdclass.Font)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mmm.Get(font[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_custom_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns custom font used to draw text in the column [param column].
*/
//go:nosplit
func (self class) GetCustomFont(ctx gd.Lifetime, column gd.Int) gdclass.Font {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_custom_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Font
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets custom font size used to draw text in the given [param column].
*/
//go:nosplit
func (self class) SetCustomFontSize(column gd.Int, font_size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, font_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_custom_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns custom font size used to draw text in the column [param column].
*/
//go:nosplit
func (self class) GetCustomFontSize(column gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_custom_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the given column's custom background color and whether to just use it as an outline.
*/
//go:nosplit
func (self class) SetCustomBgColor(column gd.Int, color gd.Color, just_outline bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, color)
	callframe.Arg(frame, just_outline)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_custom_bg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Resets the background color for the given column to default.
*/
//go:nosplit
func (self class) ClearCustomBgColor(column gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_clear_custom_bg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the custom background color of column [param column].
*/
//go:nosplit
func (self class) GetCustomBgColor(column gd.Int) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_custom_bg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Makes a cell with [constant CELL_MODE_CUSTOM] display as a non-flat button with a [StyleBox].
*/
//go:nosplit
func (self class) SetCustomAsButton(column gd.Int, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_custom_as_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the cell was made into a button with [method set_custom_as_button].
*/
//go:nosplit
func (self class) IsCustomSetAsButton(column gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_custom_set_as_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a button with [Texture2D] [param button] at column [param column]. The [param id] is used to identify the button in the according [signal Tree.button_clicked] signal and can be different from the buttons index. If not specified, the next available index is used, which may be retrieved by calling [method get_button_count] immediately before this method. Optionally, the button can be [param disabled] and have a [param tooltip_text].
*/
//go:nosplit
func (self class) AddButton(column gd.Int, button gdclass.Texture2D, id gd.Int, disabled bool, tooltip_text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mmm.Get(button[0].AsPointer())[0])
	callframe.Arg(frame, id)
	callframe.Arg(frame, disabled)
	callframe.Arg(frame, mmm.Get(tooltip_text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_add_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of buttons in column [param column].
*/
//go:nosplit
func (self class) GetButtonCount(column gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_button_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the tooltip text for the button at index [param button_index] in column [param column].
*/
//go:nosplit
func (self class) GetButtonTooltipText(ctx gd.Lifetime, column gd.Int, button_index gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_button_tooltip_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the ID for the button at index [param button_index] in column [param column].
*/
//go:nosplit
func (self class) GetButtonId(column gd.Int, button_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_button_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the button index if there is a button with ID [param id] in column [param column], otherwise returns -1.
*/
//go:nosplit
func (self class) GetButtonById(column gd.Int, id gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_button_by_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the color of the button with ID [param id] in column [param column]. If the specified button does not exist, returns [constant Color.BLACK].
*/
//go:nosplit
func (self class) GetButtonColor(column gd.Int, id gd.Int) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_button_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [Texture2D] of the button at index [param button_index] in column [param column].
*/
//go:nosplit
func (self class) GetButton(ctx gd.Lifetime, column gd.Int, button_index gd.Int) gdclass.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets the tooltip text for the button at index [param button_index] in the given [param column].
*/
//go:nosplit
func (self class) SetButtonTooltipText(column gd.Int, button_index gd.Int, tooltip gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	callframe.Arg(frame, mmm.Get(tooltip))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_button_tooltip_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the given column's button [Texture2D] at index [param button_index] to [param button].
*/
//go:nosplit
func (self class) SetButton(column gd.Int, button_index gd.Int, button gdclass.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	callframe.Arg(frame, mmm.Get(button[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the button at index [param button_index] in column [param column].
*/
//go:nosplit
func (self class) EraseButton(column gd.Int, button_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_erase_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], disables the button at index [param button_index] in the given [param column].
*/
//go:nosplit
func (self class) SetButtonDisabled(column gd.Int, button_index gd.Int, disabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_button_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the given column's button color at index [param button_index] to [param color].
*/
//go:nosplit
func (self class) SetButtonColor(column gd.Int, button_index gd.Int, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_button_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the button at index [param button_index] for the given [param column] is disabled.
*/
//go:nosplit
func (self class) IsButtonDisabled(column gd.Int, button_index gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_button_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the given column's tooltip text.
*/
//go:nosplit
func (self class) SetTooltipText(column gd.Int, tooltip gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mmm.Get(tooltip))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_tooltip_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the given column's tooltip text.
*/
//go:nosplit
func (self class) GetTooltipText(ctx gd.Lifetime, column gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_tooltip_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the given column's text alignment. See [enum HorizontalAlignment] for possible values.
*/
//go:nosplit
func (self class) SetTextAlignment(column gd.Int, text_alignment gd.HorizontalAlignment)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, text_alignment)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_text_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the given column's text alignment.
*/
//go:nosplit
func (self class) GetTextAlignment(column gd.Int) gd.HorizontalAlignment {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.HorizontalAlignment](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_text_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param enable] is [code]true[/code], the given [param column] is expanded to the right.
*/
//go:nosplit
func (self class) SetExpandRight(column gd.Int, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_expand_right, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if [code]expand_right[/code] is set.
*/
//go:nosplit
func (self class) GetExpandRight(column gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_expand_right, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDisableFolding(disable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, disable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_disable_folding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFoldingDisabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_folding_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates an item and adds it as a child.
The new item will be inserted as position [param index] (the default value [code]-1[/code] means the last position), or it will be the last child if [param index] is higher than the child count.
*/
//go:nosplit
func (self class) CreateChild(ctx gd.Lifetime, index gd.Int) gdclass.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_create_child, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Adds a previously unparented [TreeItem] as a direct child of this one. The [param child] item must not be a part of any [Tree] or parented to any [TreeItem]. See also [method remove_child].
*/
//go:nosplit
func (self class) AddChild(child gdclass.TreeItem)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(child[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_add_child, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the given child [TreeItem] and all its children from the [Tree]. Note that it doesn't free the item from memory, so it can be reused later (see [method add_child]). To completely remove a [TreeItem] use [method Object.free].
[b]Note:[/b] If you want to move a child from one [Tree] to another, then instead of removing and adding it manually you can use [method move_before] or [method move_after].
*/
//go:nosplit
func (self class) RemoveChild(child gdclass.TreeItem)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(child[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_remove_child, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Tree] that owns this TreeItem.
*/
//go:nosplit
func (self class) GetTree(ctx gd.Lifetime) gdclass.Tree {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_tree, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Tree
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the next sibling TreeItem in the tree or a null object if there is none.
*/
//go:nosplit
func (self class) GetNext(ctx gd.Lifetime) gdclass.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_next, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the previous sibling TreeItem in the tree or a null object if there is none.
*/
//go:nosplit
func (self class) GetPrev(ctx gd.Lifetime) gdclass.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_prev, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the parent TreeItem or a null object if there is none.
*/
//go:nosplit
func (self class) GetParent(ctx gd.Lifetime) gdclass.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the TreeItem's first child.
*/
//go:nosplit
func (self class) GetFirstChild(ctx gd.Lifetime) gdclass.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_first_child, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the next TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the first element in the tree when called on the last element, otherwise it returns [code]null[/code].
*/
//go:nosplit
func (self class) GetNextInTree(ctx gd.Lifetime, wrap bool) gdclass.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, wrap)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_next_in_tree, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the previous TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the last element in the tree when called on the first visible element, otherwise it returns [code]null[/code].
*/
//go:nosplit
func (self class) GetPrevInTree(ctx gd.Lifetime, wrap bool) gdclass.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, wrap)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_prev_in_tree, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the next visible TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the first visible element in the tree when called on the last visible element, otherwise it returns [code]null[/code].
*/
//go:nosplit
func (self class) GetNextVisible(ctx gd.Lifetime, wrap bool) gdclass.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, wrap)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_next_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the previous visible sibling TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the last visible element in the tree when called on the first visible element, otherwise it returns [code]null[/code].
*/
//go:nosplit
func (self class) GetPrevVisible(ctx gd.Lifetime, wrap bool) gdclass.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, wrap)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_prev_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns a child item by its [param index] (see [method get_child_count]). This method is often used for iterating all children of an item.
Negative indices access the children from the last one.
*/
//go:nosplit
func (self class) GetChild(ctx gd.Lifetime, index gd.Int) gdclass.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_child, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the number of child items.
*/
//go:nosplit
func (self class) GetChildCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_child_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an array of references to the item's children.
*/
//go:nosplit
func (self class) GetChildren(ctx gd.Lifetime) gd.ArrayOf[gdclass.TreeItem] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_children, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.TreeItem](ret)
}
/*
Returns the node's order in the tree. For example, if called on the first child item the position is [code]0[/code].
*/
//go:nosplit
func (self class) GetIndex() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Moves this TreeItem right before the given [param item].
[b]Note:[/b] You can't move to the root or move the root.
*/
//go:nosplit
func (self class) MoveBefore(item gdclass.TreeItem)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(item[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_move_before, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Moves this TreeItem right after the given [param item].
[b]Note:[/b] You can't move to the root or move the root.
*/
//go:nosplit
func (self class) MoveAfter(item gdclass.TreeItem)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(item[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_move_after, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsTreeItem() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTreeItem() Go { return *((*Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("TreeItem", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type TreeCellMode = classdb.TreeItemTreeCellMode

const (
/*Cell shows a string label. When editable, the text can be edited using a [LineEdit], or a [TextEdit] popup if [method set_edit_multiline] is used.*/
	CellModeString TreeCellMode = 0
/*Cell shows a checkbox, optionally with text. The checkbox can be pressed, released, or indeterminate (via [method set_indeterminate]). The checkbox can't be clicked unless the cell is editable.*/
	CellModeCheck TreeCellMode = 1
/*Cell shows a numeric range. When editable, it can be edited using a range slider. Use [method set_range] to set the value and [method set_range_config] to configure the range.
This cell can also be used in a text dropdown mode when you assign a text with [method set_text]. Separate options with a comma, e.g. [code]"Option1,Option2,Option3"[/code].*/
	CellModeRange TreeCellMode = 2
/*Cell shows an icon. It can't be edited nor display text.*/
	CellModeIcon TreeCellMode = 3
/*Cell shows as a clickable button. It will display an arrow similar to [OptionButton], but doesn't feature a dropdown (for that you can use [constant CELL_MODE_RANGE]). Clicking the button emits the [signal Tree.item_edited] signal. The button is flat by default, you can use [method set_custom_as_button] to display it with a [StyleBox].
This mode also supports custom drawing using [method set_custom_draw_callback].*/
	CellModeCustom TreeCellMode = 4
)