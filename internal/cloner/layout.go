package cloner

type Layouts struct {
	List []*Layout `json:"list" validate:"required,dive,required"`
}

func NewLayouts(layouts ...*Layout) *Layouts {
	var list []*Layout
	for _, l := range layouts {
		list = append(list, l)
	}
	return &Layouts{List: list}
}

func (l *Layouts) Get(name string) (*Layout, bool) {
	for _, v := range l.List {
		if name == v.Name {
			return v, true
		}
	}
	return nil, false
}

func (l *Layouts) GetNames() []string {
	var names []string
	for _, v := range l.List {
		names = append(names, v.Name)
	}
	return names
}

type Layout struct {
	Name        string `json:"name" validate:"required"`
	Namespace   string `json:"namespace" validate:"required"`
	URL         string `json:"url" validate:"required"`
	Description string `json:"description" validate:"required"`
}
