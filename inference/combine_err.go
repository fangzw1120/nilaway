package inference

import "errors"

func joinErrors(errs ...error) error {
	var combinedErr string
	for _, err := range errs {
		if err != nil {
			if combinedErr == "" {
				combinedErr = err.Error()
			} else {
				combinedErr += "; " + err.Error()
			}
		}
	}

	if combinedErr == "" {
		return nil
	}

	return errors.New(combinedErr)
}
