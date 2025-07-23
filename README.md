# absclock
## The Absolute Clock

_TL;DR_
* To get the current absolute time: `go run absclock.go`

Represents a single unique moment or event since the dawn of time (absolute time) as a colon-delimted string of time elements. Example:

`E:4:3:?:?:?:2025:07:23:09:28:15:042:?`

The general format:

`E:O:R:P:C:A:YYYY:MM:DD:HH:MM:SS:MS:NS`


* `E`: Eternity. The scope of all time. This is an invariant `E`.
* `O`: Eon. Values: 
    * pre-Hadean `0`
    * Hadean `1`
    * Archean `2`
    * Proterozoic `3`
    * Phanerozoic `4` (the current eon)
    * `5..n` (subsequent eons)
* `R`: Era. Values:
    * pre-Eoarchean `0`
    * Neoproterozoic `?`
    * Paleozoic `1`
    * Mesozoic `2`
    * Cenozoic `3` (the current era)
    * `?..n` (subsequent eras)
* `P`: Period. Quaternary `?` (the current period), , or `?..n` (subsequent periods).
* `C`: Epoch. Holocene `?` (the current epoch), , or `?..n` (subsequent epochs).
* `A`: Age. Meghalayan `?` (the current age), , or `?..n` (subsequent ages).
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
* subepoch, subperiod support? 
* periods are different for different eras:
** 
* Possibly add a way to express durations for cyclical time elements (e.g. something starting at 2pm lasting for an hour could have an hour element of `14-15`)
