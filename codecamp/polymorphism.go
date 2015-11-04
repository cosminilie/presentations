package main //OMIT

import "fmt" //OMIT

// Type Definition
type Pet interface { // HL
	Tag()
}

type Gopher struct { // HL
	Name  string
	Owner Owner
}
type Owner struct { // OMIT
	Name    string // OMIT
	Address string // OMIT
} // OMIT

func (o *Owner) Talk() { // OMIT
	fmt.Println("Hi, my name is", o.Name) // OMIT
} // OMIT
// OMIT
func (o *Gopher) Tag() { // HL
	//...
	fmt.Printf("Tag: Pet Gopher, Name:%s \n Owner: %s, Address: %s\n", o.Name, o.Owner.Name, o.Owner.Address) // OMIT
} // OMIT

type Dog struct { // HL
	Name  string
	Owner Owner
}

func (d *Dog) Tag() { // HL
	//...
	fmt.Printf("Tag: Pet Dog, Name: %s \n Owner: %s, Address: %s \n", d.Name, d.Owner.Name, d.Owner.Address) // OMIT
} // OMIT

func PetTag(p Pet) { // HL
	p.Tag() // HL
}

// ENDTYPE OMIT
func main() {
	// Tag Example
	p := Gopher{
		Name: "Steve",
		Owner: Owner{
			Name:    "Cosmin",
			Address: "Some Address",
		},
	}
	p2 := Dog{
		Name: "Spoc",
		Owner: Owner{
			Name:    "Cosmin",
			Address: "Some Address",
		},
	}
	PetTag(&p)  // HL
	PetTag(&p2) // HL
	// ENDEXAMPLE OMIT
}
