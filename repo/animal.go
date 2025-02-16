package repo

type Animal struct {
	Type string
}

func (a *Animal) SetType(typ string) {
	a.Type = typ
}
func (a *Animal) GetType() string {
	return a.Type
}

func (r *Person) GetName() string {
	return r.name
}
func (r *Person) SetName(name string) {
	r.name = name
}
