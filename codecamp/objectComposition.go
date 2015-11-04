package main

import "fmt" //OMIT

// Type Definition
type Gopher struct { // HL
	Name  string
	Owner Owner
}
type Owner struct { // HL
	Name    string
	Address string
}

func (o *Owner) Talk() { // HL
	fmt.Println("Hi, my name is", o.Name)
}

func (o *Gopher) Tag() { // HL
	fmt.Printf("Name: %s Owner: %s, Address: %s", o.Name, o.Owner.Name, o.Owner.Address)

}

// ENDTYPE OMIT
func main() {
	// Gopher Example
	g := Gopher{
		Name: "Steve",
		Owner: Owner{
			Name:    "Cosmin",
			Address: "Some Address",
		},
	}
	g.Owner.Talk()
	g.Tag()
	// ENDEXAMPLE OMIT
}
