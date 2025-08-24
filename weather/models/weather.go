package models

import (
	"encoding/json"
	"time"
)

type Weather struct {
	DescriptionList []Description `json:"weather"`
	CityName        string        `json:"name"`
	TimeZone        int           `json:"timezone"`
	Temperature     Temperature   `json:"main"`
	Wind            Wind          `json:"wind"`
	Sys             Sys           `json:"sys"`
}

type Description struct {
	Description string `json:"description"`
	IconID      string `json:"icon"`
}

type Temperature struct {
	Temp      float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
}

type Wind struct {
	Speed float32 `json:"speed"`
}

type Sys struct {
	Country string `json:"country"`
	Sunrise time.Time
	Sunset  time.Time
}

type internalSys struct {
	Country     string `json:"country"`
	SunriseUnix int64  `json:"sunrise"`
	SunsetUnix  int64  `json:"sunset"`
}

func (s *Sys) UnmarshalJSON(b []byte) error {

	var tempSys internalSys
	if err := json.Unmarshal(b, &tempSys); err != nil {
		return err
	}

	s.Country = tempSys.Country

	//offset := time.FixedZone("local")
	s.Sunrise = time.Unix(tempSys.SunriseUnix, 0).UTC()
	s.Sunset = time.Unix(tempSys.SunsetUnix, 0).UTC()

	return nil
}
