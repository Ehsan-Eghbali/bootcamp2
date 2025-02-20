package port

type Animal interface {
	Breed() string
	Breath() string
	Walk() string
	GetName() string
}

type Lion struct {
	Name string
	Age  int
}

func (l *Lion) Breed() string {
	l.Walk()
	return "animal can breed"

}

func (l *Lion) Breath() string {
	return "animal can breath"

}

func (l *Lion) Walk() string {
	return "animal can walk"
}

func (l *Lion) GetName() string {
	return l.Name
}
