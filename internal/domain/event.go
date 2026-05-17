package domain

type EventIncomingID int

const (
	EventIncomingUnspecified    EventIncomingID = 0
	EventIncomingRegisterPlayer EventIncomingID = 1
	EventIncomingEnterDungeon   EventIncomingID = 2
	EventIncomingKillMonster    EventIncomingID = 3
	EventIncomingNextFloor      EventIncomingID = 4
	EventIncomingPreviousFloor  EventIncomingID = 5
	EventIncomingEnterBossFloor EventIncomingID = 6
	EventIncomingKillBoss       EventIncomingID = 7
	EventIncomingLeaveDungeon   EventIncomingID = 8
	EventIncomingCannotContinue EventIncomingID = 9
	EventIncomingRestoreHealth  EventIncomingID = 10
	EventIncomingReceiveDamage  EventIncomingID = 11
)

type EventOutgoingID int

const (
	EventOutgoingUnspecified    EventOutgoingID = 0
	EventOutgoingDisqualified   EventOutgoingID = 31
	EventOutgoingDead           EventOutgoingID = 32
	EventOutgoingImpossibleMove EventOutgoingID = 33
)
