package dto

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	domain "github.com/Friend-zva/game-server/internal/domain"
)

type Event struct {
	Time     time.Time
	IdPlayer int
	IdEvent  domain.EventIncomingID
	Param    string
}

type DecoderEvent struct {
	formatTime string
}

func NewDecoderEvent(formatTime string) *DecoderEvent {
	return &DecoderEvent{formatTime: formatTime}
}

func (d *DecoderEvent) Execute(line string) (Event, error) {
	var event Event

	parts := strings.Split(line, "] ")
	if len(parts) < 2 {
		return event, fmt.Errorf("invalid format")
	}

	timeRaw := strings.TrimPrefix(parts[0], "[")
	time, err := time.Parse(d.formatTime, timeRaw)
	if err != nil {
		return event, fmt.Errorf("failed to parse time: %w", err)
	}
	event.Time = time

	tokens := strings.Fields(parts[1])
	if len(tokens) < 2 {
		return event, fmt.Errorf("failed to parse event args")
	}

	idPlayer, err := strconv.Atoi(tokens[0])
	if err != nil {
		return event, fmt.Errorf("failed to parse player id: %w", err)
	}
	event.IdPlayer = idPlayer

	eventRaw, err := strconv.Atoi(tokens[1])
	if err != nil {
		return event, fmt.Errorf("failed to parse event id: %w", err)
	}
	event.IdEvent = domain.EventIncomingID(eventRaw)

	if len(tokens) >= 3 {
		event.Param = strings.Join(tokens[2:], " ")
	}

	return event, nil
}
