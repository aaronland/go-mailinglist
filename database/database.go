package database

type NoRecordError string

func (err NoRecordError) Error() string {
	return "Database record does not exist"
}

func IsNotExist(err error) bool {

	switch err.(type) {
	case *NoRecordError, NoRecordError:
		return true
	default:
		return false
	}
}


