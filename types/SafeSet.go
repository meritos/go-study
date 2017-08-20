
package types

import (

	"fmt"
	"sync"
)

// Interface for elements 
// to add to the SafeSet
type Element interface {

	toString() string  
}

// Thread safe structure
type SafeSet struct {

	elements map[string]Element 
	sync.RWMutex
} 

// CTor
func NewSafeSet() *SafeSet {

	return &SafeSet {

		elements : make(map[string]Element),
	}
}


func (this *SafeSet) Add(e Element){

	defer this.Unlock()

	this.Lock()
	this.elements[e.toString()] = e
}

// 
// Remove function removes given element 
// from the SafeSet
//
//  @param e Element - to remove from the SafeSet
//
func (this *SafeSet) Remove(e Element){

	defer this.Unlock()

	this.Lock()
	delete(this.elements, e.toString())
}

//
// The function returns the number of the elements 
// in the SafeSet.
//
//  @return int - number of the elements in the SafeSet
//
func (this *SafeSet) Size() int {
	
	defer this.RUnlock()

	this.RLock()
	return len(this.elements)
}

// 
// Clean all the elements from the SafeSet
//
func (this *SafeSet) Clean() {
	
	defer this.Unlock()

	this.Lock()
	this.elements = make(map[string]Element)
}

//
// The function returns true if the given 
// element exist in that SafeSet false otherwise.
//
//  @param e Element - element to check for existence 
//  @return bool - true for existence, otherwise false
//
func (this *SafeSet) Has(e Element) bool {

	defer this.RUnlock()

	this.RLock()
 	if this.elements[e.toString()] == nil {

 		return false
 	} else {

 		return true
 	}
}

//
// The function returns true if the SafeSet
// is empty, otherwise false.
//
//  @return bool - true if empty, otherwise 
//                 false
//
func (this *SafeSet) IsEmpty() bool {

	defer this.RUnlock()

	this.RLock()
	if len(this.elements) > 0 {

		return false
	} else {

		return true
	}
}

//
// For any given SafeSet the function will create
// and return list of elements
//
func (this *SafeSet) ToList() []Element {

	defer this.Unlock()

	this.Lock()
	list := make([]Element, len(this.elements))
	i    := 0

	for _, element := range(this.elements) {

		list[i] = element
		i++
	}

	return list
}

//
// Print all elements in the given SafeSet
//
func (this *SafeSet) PrintAll() {

	defer this.RUnlock()

	this.RLock()

	fmt.Print("[ ")
	for _, e := range this.elements {

		fmt.Printf("%s ", e.toString())
	}
	fmt.Print("]")

	fmt.Printf("\n")
}


