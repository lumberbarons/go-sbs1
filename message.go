package sbs1

import (
	"time"
)

type MessageType uint8

const (
	MessageTypeSelectionChange MessageType = iota
	MessageTypeNewId
	MessageTypeNewAircraft
	MessageTypeStatusAircraft
	MessageTypeClick
	MessageTypeTransmission
)

type TransmissionType uint8

const (
	TransmissionTypeESIdentAndCategory TransmissionType = iota
	TransmissionTypeESSurfacePos
	TransmissionTypeESAirbornePos
	TransmissionTypeESAirborneVel
	TransmissionTypeSurveillanceAlt
	TransmissionTypeSurveillanceId
	TransmissionTypeAirToAir
	TransmissionTypeAllCallReply
)

type Message struct {
	MessageType      MessageType `json:"messageType"`
	TransmissionType TransmissionType `json:"transmissionType"`
	SessionId        string `json:"sessionId"`
	AircraftId       string `json:"aircraftId"`
	HexId            string `json:"hexId"`
	FlightId         string `json:"FlightId"`
	Generated        time.Time `json:"generated"`
	Logged           time.Time `json:"logged"`
	Callsign         string `json:"callsign"`
	Altitude         int32 `json:"altitude"`
	GroundSpeed      int32 `json:"groundSpeed"`
	Track            float64 `json:"track"`
	Latitude      	 float64 `json:"latitude"`
	Longitude		 float64 `json:"longitude"`
	VerticalRate     int16 `json:"verticalRate"`
	Squawk           string `json:"squawk"`
}
