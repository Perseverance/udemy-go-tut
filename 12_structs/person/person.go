package person

import "fmt"

//Person provides interface to work with person
type Person struct {
	firstName string
	LastName  string
	sayHi     func() string
	gender    rune
}

//Woman is a type of Person
type Woman struct {
	Person
}

func (w Woman) String() string {
	return fmt.Sprintf("%v %v", w.firstName, w.LastName)
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

//NewWoman Creates a new woman
func NewWoman() Woman {
	person := NewPerson()
	person.firstName = "Ivanka"
	person.LastName = "Georgieva"
	person.gender = 'f'
	return Woman{
		Person: person,
	}
}

//SayHi provides ability for the Person to say Hi
func (p *Person) SayHi() string {
	return p.sayHi()
}

//SayYo provides ability for the Person to say Hi
func (p *Person) SayYo() string {
	return p.sayHi()
}
