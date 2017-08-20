
package types


func ExampleSafeSet(){

	// Create new SafeSet
	safeSet := NewSafeSet()

	// Creating elements for SafeSet
	apple := new(StringElement)
	apple.element = "Apple" 
	
	cat := new(StringElement)
	cat.element = "Cat"

	// Adding Elements to the SafeSet
	safeSet.Add(apple)
	safeSet.Add(cat)

	// Print the entire SafeSet
	safeSet.PrintAll()


	// Output: [ Apple Cat ]	
}

