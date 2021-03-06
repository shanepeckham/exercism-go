package erratum

import (
	"io"
)

const testVersion = 2

type TransientError struct {
	err error
}

func (e TransientError) Error() string {
	return e.err.Error()
}

type FrobError struct {
	defrobTag string
	inner     error
}

func (e FrobError) Error() string {
	return e.inner.Error()
}

type Resource interface {
	io.Closer

	// Frob does something with the input string.
	// Because this is an incredibly badly designed system if there is an error
	// it panics.
	//
	// The paniced error may be a FrobError in which case Defrob should be called
	// with the defrobTag string.
	Frob(string)

	Defrob(string)
}

// ResourceOpener is a function that creates a resource.
//
// It may return a wrapped error of type TransientError. In this case the resource
// is temporarily unavailable and the caller should retry soon.
type ResourceOpener func() (Resource, error)
type err error

func Use(ro ResourceOpener, s string) (err error) {

	var res Resource

	for done := false; done != true; {
		res, err = ro()
		if _, ok := err.(TransientError); ok {
			continue
		}
		if err != nil {
			return err
		}
		break
	}

	defer func() {
		defer res.Close()
		if r := recover(); r != nil {
			if nerr, ok := r.(FrobError); ok {
				res.Defrob(nerr.defrobTag)
			}
			err = r.(error)
		}
	}()

	res.Frob(s)
	return err

}
