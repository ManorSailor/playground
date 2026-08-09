package main

import "fmt"

// No optionality. Instead we have undefined flavors of respective types
// Use pointer to define a true optional field because a pointer by default is always nil
// Fabulous design decisions right here. Learned something from C, I see.
// It would have killed to write ? or even a union type - Oh wait. They don't support that either.
type Pen struct {
	quantity  uint8
	color     string
	ownerInfo Owner
	model     string
}

type Owner struct {
	name string
}

type Pencil struct {
	length    uint8
	color     string
	ownerInfo Owner
}

type EmptyStruct struct {}

// Interfaces only allow defining method contracts.
// Struct properties aren't allowed.
type WritingInstrument interface {
	replenish()
	write()
}

// Go has made everything so confusing?? It's all implicit?
// How tf do I know if a type implements an interface?
// I CAN'T because this language excluded the `implements` keyword just to maintain "simplicity"
// Official Hack to ensure a type implements an interface
var _ WritingInstrument = &Pen{}
var _ WritingInstrument = &Pencil{}

// This is a Pointer Receiver method. We are defining a method on type Pen which receives a pointer to Pen instance.
// Value receivers are also methods or functions which accept `p Pen` instead of pointer to Pen instance.
func (p *Pen) replenish() {
	// Go automatically dereferences p.
	// (*p).quantity == p.quantity
	// Why not? p->quantity?? I guess this syntax could be confused with Channels that's why.

	if p.quantity < 100 {
		p.quantity = 100
		fmt.Println("Your pen has been refilled. Happy writing.")
	}
}

func (p *Pencil) replenish() {
	if p.length <= 0 {
		p.length = 10
		fmt.Println("You took out a new pencil. Happy writing.")
	}
}

func (p *Pen) write() {
	if p.quantity <= 0 {
		fmt.Println("Pen is empty! Please refill")
		return
	}

	fmt.Println("Wrote something...")
	p.quantity -= 50
}

func (p *Pencil) write() {
	if p.length <= 1 {
		fmt.Println("Pencil is too short to write anything. Please replenish.")
		return
	}

	fmt.Println("Wrote something...")
	p.length -= 5
}

func main() {
	pen := Pen{
		quantity: 100,
		color:    "Black",
		model:    "Cello CK100",
		ownerInfo: Owner{
			name: "Me",
		},
	}

	pencil := Pencil{
		length: 10,
		color:  "Black",
		ownerInfo: Owner{
			name: "Me",
		},
	}

	empty := EmptyStruct{}

	fmt.Printf("%v %T\n", empty, empty)

	fmt.Printf("Your pen is %v. It's capacity is at %v\n", pen.model, pen.quantity)

	// Pass the address of the struct to write
	// BUT! write itself doesn't accept a pointer??!! It expects an Interface!
	// and a pointer to interface is invalid.
	// Oh! The reason for address here is that the contract which this interface expects
	// only accepts a pointer, i.e., a Pointer Receiver
	// You pass a pointer to a function, expect a pointer, but if the function's argument is an interface...
	// Suddenly, the pointer argument is not required???!!!
	write(&pen)
	fmt.Println("")
	write(&pencil)
}

func write(p WritingInstrument) {
	(p).write()
	(p).write()
	(p).write()
	(p).replenish()
	(p).write()
}
