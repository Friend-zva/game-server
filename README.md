# System prototype

The prototype must be able to work with a configuration file and a set of external events of a certain format.
Solution should contain golang (1.22 or newer) source file/files and unit tests (optional)

## Description

A player is participating in a challenge. The goal is to completely clear a dungeon. The player navigates through floors and fights monsters. We need to process the events and compile the information into a final report

## Rules

1.  Only registered players are allowed to participate in the challenge
2.  The challenge ends if:
    1.  The player leaves the dungeon
    2.  The player cannot continue the challenge
    3.  The dungeon opening time has expired
    4.  Player is dead (health drops to 0)
3.  When entering the boss's floor, the player receives a notification
4.  The boss floor does not contain any monsters
5.  The dungeon is considered complete if:
    1.  All floors are cleared of monsters
    2.  The boss is defeated
6.  A floor is considered complete when all monsters or the boss have been killed; **_any time spent in that floor is no longer counted_**
7.  The player's health cannot exceed 100

## Usage

```sh
go build -v -o ./bin/app ./cmd/app.go
./bin/app
```

## Events

- All events occur sequentially in time. (**_Time of event N+1_**) >= (**_Time of event N_**)
- Time format **_[HH:MM:SS]_**. Trailing zeros are required in input and output
- The **_ExtraParam_** parameter can be a string containing multiple words.

### Incoming events

| EventID | ExtraParam | Comment                                         |
| ------- | :--------: | ----------------------------------------------- |
| 1       |            | Player [`id`] registered                        |
| 2       |            | Player [`id`] entered the dungeon               |
| 3       |            | Player [`id`] killed the monster                |
| 4       |            | Player [`id`] went to the next floor            |
| 5       |            | Player [`id`] went to the previous floor        |
| 6       |            | Player [`id`] entered the boss's floor          |
| 7       |            | Player [`id`] killed the boss                   |
| 8       |            | Player [`id`] left the dungeon                  |
| 9       |  `reason`  | Player [`id`] cannot continue due to [`reason`] |
| 10      |  `health`  | Player [`id`] has restored [`health`] of health |
| 11      |  `damage`  | Player [`id`] recieved [`damage`] of damage     |

> For example see [events](./events).

### Outgoing events

| EventID | ExtraParam | Comment                                        |
| ------- | :--------: | ---------------------------------------------- |
| 31      |            | Player [`id`] disqualified                     |
| 32      |            | Player [`id`] is dead                          |
| 33      |            | Player [`id`] makes imposible move [`eventID`] |

> For example, see [events_expected](./events_expected).

## States

| State   |                           Comment                            |
| ------- | :----------------------------------------------------------: |
| SUCCESS |                    All floors are cleared                    |
| FAIL    |  The player died or the dungeon is not considered completed  |
| DISQUAL | The player cannot continue or has not completed registration |

## Final report

1.  State `SUCCESS`/`FAIL`/`DISQUAL`
2.  Player ID
3.  Time spent in the dungeon (all the time until the player left the dungeon or the dungeon closed)
4.  Average time to clear a floor of monsters (the boss's floor is not included in the calculation)
5.  Time to kill the boss
6.  Player health at the end of the trial

> For example, see [events_expected](./events_expected).

## Configuration (json)

- **Floors** - Number of floors in the dungeon
- **Monsters** - Number of monsters on each floor of the dungeon
- **OpenAt** - Dungeon opening time
- **Duration** - Time until the dungeon closes in hours

> For example, see [config](./config/config.json).

## LICENSE

Distributed under the [MIT License](https://choosealicense.com/licenses/mit/). See [`LICENSE`](LICENSE) for more information.
