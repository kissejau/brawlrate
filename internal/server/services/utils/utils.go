package utils

import "fmt"

type rankingsData struct {
	Braket string
	Region string
	Page   int
}

func NewRankingsData(braket string, region string, page int) (*rankingsData, error) {
	rd := &rankingsData{}

	if braket == "" || region == "" {
		return nil, fmt.Errorf("braket n region must be passed")
	}

	if page > 1000 {
		return nil, fmt.Errorf("page must be < 1000")
	}

	if page <= 0 {
		page = 1
	}

	rd.Braket, rd.Region, rd.Page = braket, region, page
	return rd, nil
}
