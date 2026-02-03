package entities

type TrackRequest struct {
	Power bool `json:"power"`
}

type VehicleSpeedRequest struct {
	Speed   int  `json:"speed"`
	Forward bool `json:"forward"`
}

type VehicleFunctionRequest struct {
	Function int `json:"function"`
	State    int `json:"state"`
}
