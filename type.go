package mem

import (
	"encoding/json"
)

type dbErr struct {
	Number int
	Message string
}

func (e *dbErr) Error() string {
	return e.Message
}

func (e *dbErr) mapping(err error) error {
	errBytes, errJ := json.Marshal(err)
	if errJ != nil {
		return errJ
	}

	errJ = json.Unmarshal(errBytes, e)
	if errJ != nil {
		return errJ
	}

	if e.Number == 0 {
		e.Message = err.Error()
	}

	return nil
}
