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
