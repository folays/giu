package giu

type Widget interface {
	Build()
}

type Layout []Widget

func (l Layout) Build() {
	for _, w := range l {
		if w != nil {
			w.Build()
		}
	}
}
