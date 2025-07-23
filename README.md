# absclock
The Absolute Clock

Represents a single unique moment since the dawn of time as a colon-delimted string. Example:

`E:P:?:?:?:?:2025:07:23:09:28:15`

The Format:

* `E`: Eternity. The scope of all time. This is an invariant `E`.
* `O`: Eon. Can be Hadean `H`, Archean `A`, Proterozoic `R`, or Phanerozoic `P` (the current eon).
* `R`: Era.
* `P`: Period.
* `C`: Epoch.
* `A`: Age.
* `YYYY`: The current year. Unbounded.
* `MM`: The current month. 01-12.
* `DD`: The current day. 01-31.
* `HH`: The current hour. Military format. 00-23.
* `MI`: The current minute. 00-59.
* `SS`: The current second. 00-59.

Time is represented in Zulu time.

To do:
* Figure out how to represent the time between the Big Bang and the Hadean.