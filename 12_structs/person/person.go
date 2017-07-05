package person

//Person provides interface to work with person
type Person struct {
	firstName string
	LastName  string
	sayHi     func() string
	gender    int
}

//Woman is a type of Person
type Woman struct {
	Person
}

//NewPerson Generates a new person
func NewPerson() Person {
	return Person{
		firstName: "Ivan",
		LastName:  "Ivanov",
		sayHi: func() string {
			return "hi!"
		}}
}

//SayHi provides ability for the Person to say Hi
func (p *Person) SayHi() string {
	return p.sayHi()
}

//SayYo provides ability for the Person to say Hi
func (p *Person) SayYo() string {
	return p.sayHi()
}
