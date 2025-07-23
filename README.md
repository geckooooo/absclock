# absclock
The Absolute Clock

_TL;DR_
To get the current absolute time: `go run absclock.go`

Represents a single unique moment or event since the dawn of time (absolute time) as a colon-delimted string of time elements. Example:

`E:4:3:?:?:?:2025:07:23:09:28:15:042:?`

The general format:

`E:O:R:P:C:A:YYYY:MM:DD:HH:MM:SS:MS:NS`


* `E`: Eternity. The scope of all time. This is an invariant `E`.
* `O`: Eon. Can be pre-Hadean `0`, Hadean `1`, Archean `2`, Proterozoic `3`, Phanerozoic `4` (the current eon), or `5..n` (subsequent eons).
* `R`: Era. Can be pre-Paleozoic `0`, Paleozoic `1`, Mesozoic `2`, Cenozoic `3` (the current era), or `4..n` (subsequent eras).
* `P`: Period.
* `C`: Epoch.
* `A`: Age.
* `YYYY`: The current year. Unbounded.
* `MM`: The current month. 01-12.
* `DD`: The current day. 01-31.
* `HH`: The current hour. Military format. 00-23.
* `MI`: The current minute. 00-59.
* `SS`: The current second. 00-59.
* `MS`: The current millisecond. 00-999.
* `NS`: The current nanosecond. ?.

This is the absolute time of the Big Bang:

`(add this once the last representation challenges are resolved)`

_Notes:_
* Time is represented in Zulu time.
* Resolution is determined by what temporal elements, if any, are omitted from the far right.

_To do:_
* Figure out how to represent B.C.E. dates (negative values?)
* Possibly add a way to express durations for cyclical time elements (e.g. something starting at 2pm lasting for an hour could have an hour element of `14-15`)
