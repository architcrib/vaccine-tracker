package main

type ValidCenter struct {
	CenterName        string `json:"name"`
	BlockName         string `json:"block_name"`
	Pincode           int    `json:"pincode"`
	Date              string `json:"date"`
	AvailableCapacity int    `json:"available_capacity"`
	MinAgeLimit       int    `json:"min_age_limit"`
	Vaccine           string `json:"vaccine"`
}
type Session struct {
	SessionID         string   `json:"session_id"`
	Date              string   `json:"date"`
	AvailableCapacity int      `json:"available_capacity"`
	MinAgeLimit       int      `json:"min_age_limit"`
	Vaccine           string   `json:"vaccine"`
	Slots             []string `json:"slots"`
}
type Center struct {
	CenterID     int       `json:"center_id"`
	Name         string    `json:"name"`
	StateName    string    `json:"state_name"`
	DistrictName string    `json:"district_name"`
	BlockName    string    `json:"block_name"`
	Pincode      int       `json:"pincode"`
	Lat          int       `json:"lat"`
	Long         int       `json:"long"`
	From         string    `json:"from"`
	To           string    `json:"to"`
	FeeType      string    `json:"fee_type"`
	Sessions     []Session `json:"sessions"`
}
type CowinResponse struct {
	Centers []Center `json:"centers"`
}

func (cowinResponse CowinResponse) FilterValidCenter(validCenters *[]ValidCenter) {
	for _, center := range cowinResponse.Centers {
		if len(center.Sessions) != 0 {
			for _, session := range center.Sessions {
				if session.MinAgeLimit == 18 && session.AvailableCapacity != 0 {
					validCenter := ValidCenter{
						CenterName:        center.Name,
						BlockName:         center.BlockName,
						Pincode:           center.Pincode,
						Date:              session.Date,
						AvailableCapacity: session.AvailableCapacity,
						Vaccine:           session.Vaccine,
					}
					*validCenters = append(*validCenters, validCenter)
				}
			}
		}
	}
}
