package main

import "fmt"

// structes are data structure similar but not same dictinaries in python
// when we define a new struct for example person struct, we can make a new people from that struct.

type contactInfo struct { // we defined person struct like we define string slice in cards exercise
	email   string
	zipCode int
}

type person struct { // we defined person struct  like we define string slice in cards exercise
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// first way of create a struct alex := person{firstName: "Alex", lastName: "Anderson"}
	//  fmt.Println(alex)

	// second way: did not assign a value automatically assigned "" for default string for int 0 float 0 bool false
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Printf("%+v", alex) // prints all fiels ---property name: value--- syntax

	// Part 47. embedding structs
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	//  jimPointer := &jim // &jim means give me the memory address of the value this variable is pointing at. So jimPointer receives the address at the ram like: 0001.
	// when we use star operator *jimPointer it gives again the value this memory address pointing at.
	// Even though we did not define jimPointer and directly use jim, we can use it like jimPointer if we defined function argument
	// as *person. Although jim's type is person and argument required is pointerPerson. Go can deal with it.
	// jimPointer.updateNamewithPointer("Jimmy")
	jim.updateNamewithPointer("Jimmy") // Onemli olan fonksiyonun argumaninin tipinin pointer person olmasi.
	// Go kendisi otomatik olarak onu pointer person tipine cevirecek.
	jim.print()
	// jim.updateName("Jimmy") // person will not be updated with firstName Jimmy.
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// Part 48. Pass by Value. Go is a Pass by value language. creates new area calles p in the memory and adds same struct with jim to the new memory address for p when
// it is called by update name. So, struct in the old memory address (jim) did not change but we only change copy --> the struct (value of the firstName property) in the new memory address.
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

// Part 49. Struct with Pointers
func (pointerToPerson *person) updateNamewithPointer(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName // Memory adresinden ulastigimizda
	// gercek degeri degistirmis oluyoruz. Direkt value olarak verdigimizde yeni bir adreste yeni bir deger olusturmus oluyoruz.
}

// Part 53. Reference vs Value Types
// function call degil arguman verdigimizde degerde degisiklik oluyor pointer kullanmasak bile.
// Any time we make a slice, go creates both slice and array. Slice holds length, capacity and pointer to head values of the array.
// Bu yuzden slice bir reference type cunku memory deki baska bir data structure'a referans veriyor(arraye).
// Reference typelarda pass by value olsa da bir kopya olusturdugumuzda memeorydeki dogru adresi (ayni arrayi) referans verdigi icin
// olusturulmasinda bir sikinti yoktur. Ayni adreste degisiklik yapmis oluruz.

// value types: int float string bool structs --> use pointers to change these things in a function
// reference types slices maps channels pointers functions --> dont worry about pointers with these
