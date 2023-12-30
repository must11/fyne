package fyne

// Layout defines how CanvasObjects may be laid out in a specified Size.
type LayoutExtent interface {
	Layout
	Padding()
	Paddint(float32)
}
