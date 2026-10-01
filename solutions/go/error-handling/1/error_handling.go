package erratum

import "errors"

func Use(opener ResourceOpener, input string) (err error) {
	var res Resource
	for {
		res, err = opener()
		if err == nil {
			break
		}
		var transient TransientError
		if !errors.As(err, &transient) {
			return err
		}
	}

	defer res.Close()

	defer func() {
		if r := recover(); r != nil {
			if fe, ok := r.(FrobError); ok {
				res.Defrob(fe.defrobTag)
				err = fe
				return
			}
			if e, ok := r.(error); ok {
				err = e
				return
			}
			panic(r)
		}
	}()

	res.Frob(input)
	return err
}