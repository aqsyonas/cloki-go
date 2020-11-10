package service

import (
	"github.com/Jeffail/gabs/v2"
)

type SeriesService struct {
	ServiceData
}

// this method create new user in the database
// it doesn't check internally whether all the validation are applied or not
func (ss *SeriesService) GetSeries() string {
	reply := gabs.New()
	reply.Set("success", "status")
	reply.Set([]string{}, "data")
	return reply.String()
}
