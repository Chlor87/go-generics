package main

// This is a demo. It seems that we can't do a concrete type in a method of
// a generic type. What's weird is that the compiler doesn't complain which
// results in a runtime stack overflow caused probably by infinite type

type RecursiveType[T any] func() RecursiveType[T]

func (d RecursiveType[T]) Method() {
	// change this type to a concrete one (eg. any) to get the error
	Function[T]()
}

func Function[O any]() RecursiveType[O] {
	return nil
}
