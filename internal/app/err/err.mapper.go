package err

import "errors"

const (
	Auth   ErrDomain = "auth"
	Assets ErrDomain = "assets"
)

func (e ErrDomain) Error() string {
	return string(e)
}

type ErrApp struct {
	MapErr Mapper
}

func NewErrApp() *ErrApp {
	return &ErrApp{
		MapErr: make(Mapper),
	}
}

func (e *ErrApp) Register(domain ErrDomain, errs []ErrEntry) {
	for _, err := range errs {
		e.MapErr[domain][err.Err] = err
	}
}

func (e *ErrApp) Translate(err error, domain ErrDomain) *ErrEntry {
	if err == nil {
		return nil
	}
	var Err *ErrEntry
	for _, errs := range e.MapErr[domain] {
		if errs.Err == err {
			Err = &ErrEntry{
				Err:        err,
				Code:       errs.Code,
				StatusCode: errs.StatusCode,
				Message:    errs.Message,
			}
			return Err
		}
	}

	return &ErrEntry{
		Err:        errors.New("internal server error"),
		Code:       "INTERNAL_SERVER_ERROR",
		StatusCode: 500,
		Message:    "Internal Server Error",
	}
}
