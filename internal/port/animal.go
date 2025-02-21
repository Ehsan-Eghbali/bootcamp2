package port

type Animal interface {
	Walk() string
	Breed() string
	GetName() string
}

type Lion struct {
	Name string
	Age  int
}

func (l *Lion) Walk() string {
	return "lion can walk"
}
func (l *Lion) Breed() string {
	return "lion breed"
}
func (l *Lion) GetName() string {
	return l.Name
}

type Cat struct {
	Name string
}

func (c *Cat) Walk() string {
	return "cat can walk"
}
func (c *Cat) Breed() string {
	return "cat breed"
}
func (c *Cat) GetName() string {
	return c.Name
}
