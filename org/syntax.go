package org

// Node is any element in the document tree.
type Node interface {
    Kind() string
    Children() []Node
}

// Document is the root.
type Document struct{ children []Node }
func (d *Document) Kind() string       { return "document" }
func (d *Document) Children() []Node   { return d.children }

// Headline, Paragraph, ListItem … keep it minimal.
type Headline struct {
    Level int
    Text  string
    children []Node
}
func (h *Headline) Kind() string     { return "headline" }
func (h *Headline) Children() []Node { return h.children }

// — Inline helpers (bold/italic/links) can go in a small
//   separate file; keep them string-based to start.

type Paragraph struct {
	text string
}
func (p *Paragraph) Kind() string		{ return "paragraph" }
func (p *Paragraph) Children() []Node	{ return nil }
