# absclock
The Absolute Clock

Represents a single unique moment or event since the dawn of time (absolute time) as a colon-delimted string of time elements. Example:

`E:P:?:?:?:?:2025:07:23:09:28:15`

The Format:

* `E`: Eternity. The scope of all time. This is an invariant `E`.
* `O`: Eon. Can be Hadean `H`, Archean `A`, Proterozoic `R`, or Phanerozoic `P` (the current eon).
* `R`: Era. Can be Paleozoic `P`, Mesozoic `M`, or Cenozoic `C` (the current era).
* `P`: Period.
* `C`: Epoch.
* `A`: Age.
* `YYYY`: The current year. Unbounded.
* `MM`: The current month. 01-12.
* `DD`: The current day. 01-31.
* `HH`: The current hour. Military format. 00-23.
* `MI`: The current minute. 00-59.
* `SS`: The current second. 00-59.

This is the absolute time of the Big Bang:
`(add this once the last representation challenges are resolved)`

_Notes:_
* Time is represented in Zulu time.
* Resolution is determined by what temporal elements, if any, are omitted from the far right.

_To do:_
* Figure out how to represent the time between the Big Bang and the Hadean. (complication: all geologic elements don't exist before the Hadean)
* Figure out how to represent B.C.E. dates
