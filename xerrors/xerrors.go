// package xerrors represents the proposed standard library errors package
// (only the relevant part for this demo purposes)
package xerrors

type Wrapper interface {
	error
	Wrapped() error
}

func Last(err error, f func(error) bool) error {
	var ok bool
	var we Wrapper

	for err != nil {
		if f(err) {
			return err
		}
		we, ok = err.(Wrapper)
		if !ok {
			return nil
		}
		err = we.Wrapped()
	}

	return nil
}