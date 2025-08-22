package ierrors

import "fmt"

type CityNotFoundError struct {
	CityName string
}

func (c *CityNotFoundError) Error() string {
	return fmt.Sprintf("City \"%s\" not found.", c.CityName)
}

func NewCityNotFoundError(cityName string) error {
	return &CityNotFoundError{CityName: cityName}
}
