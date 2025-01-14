package sbs1

import (
	"encoding/csv"
	"errors"
	"io"
	"strconv"
	"time"
)

const (
	timeFormat = "2006/01/02T15:04:05.999"
)

var (
	ErrUnkownMessageType       = errors.New("unknown message type")
	ErrUnknownTransmissionType = errors.New("unknown transmission type")
)

func NewReader(r io.Reader) *Reader {
	var csvr = csv.NewReader(r)

	return &Reader{csvr: csvr}
}

type Reader struct {
	csvr *csv.Reader
}

func (r *Reader) Read() (*Message, error) {
	fields, err := r.csvr.Read()

	if err != nil {
		return nil, err
	}

	return r.parse(fields)
}

func (r *Reader) parse(fields []string) (*Message, error) {
	message := &Message{}

	switch fields[0] {
	case "SEL":
		message.MessageType = MessageTypeSelectionChange
	case "ID":
		message.MessageType = MessageTypeNewId
	case "AIR":
		message.MessageType = MessageTypeNewAircraft
	case "STA":
		message.MessageType = MessageTypeStatusAircraft
	case "CLK":
		message.MessageType = MessageTypeClick
	case "MSG":
		message.MessageType = MessageTypeTransmission
	default:
		return nil, ErrUnkownMessageType
	}

	switch fields[1] {
	case "1":
		message.TransmissionType = TransmissionTypeESIdentAndCategory
	case "2":
		message.TransmissionType = TransmissionTypeESSurfacePos
	case "3":
		message.TransmissionType = TransmissionTypeESAirbornePos
	case "4":
		message.TransmissionType = TransmissionTypeESAirborneVel
	case "5":
		message.TransmissionType = TransmissionTypeSurveillanceAlt
	case "6":
		message.TransmissionType = TransmissionTypeSurveillanceId
	case "7":
		message.TransmissionType = TransmissionTypeAirToAir
	case "8":
		message.TransmissionType = TransmissionTypeAllCallReply
	default:
		return nil, errors.New("unknown transmission type: " + fields[1])
	}

	message.SessionId = fields[2]
	message.AircraftId = fields[3]
	message.HexId = fields[4]
	message.FlightId = fields[5]

	if len(fields[6]) > 0 {
		generated, err := time.Parse(timeFormat, fields[6]+"T"+fields[7])

		if err != nil {
			return nil, err
		}

	message.Generated = generated
	}

	if len(fields[8]) > 0 {
		logged, err := time.Parse(timeFormat, fields[8]+"T"+fields[9])

		if err != nil {
			return nil, err
		}

		message.Logged = logged
	}

	message.Callsign = fields[10]

	if len(fields[11]) > 0 {
		altitude, err := strconv.ParseInt(fields[11], 10, 32)

		if err != nil {
			return nil, err
		}

		message.Altitude = int32(altitude)
	}

	if len(fields[12]) > 0 {
		groundSpeed, err := strconv.ParseFloat(fields[12], 16)

		if err != nil {
			return nil, err
		}

		message.GroundSpeed = int32(groundSpeed)
	}

	if len(fields[13]) > 0 {
		track, err := strconv.ParseFloat(fields[13], 64)

		if err != nil {
			return nil, err
		}

		message.Track = track
	}

	if len(fields[14]) > 0 {
		latitude, err := strconv.ParseFloat(fields[14], 64)

		if err != nil {
			return nil, err
		}

		message.Latitude = latitude

		longitude, err := strconv.ParseFloat(fields[15], 64)

		if err != nil {
			return nil, err
		}

		message.Longitude = longitude
	}

	if len(fields[16]) > 0 {
		verticalRate, err := strconv.ParseFloat(fields[16], 16)

		if err != nil {
			return nil, err
		}

		message.VerticalRate = int16(verticalRate)
	}

	return message, nil
}
