# The Absolute Clock

### TL;DR
* To get the current absolute time: `% go run absclock.go`

The Absolute Clock represents a single unique moment or event since the dawn of time (absolute time) as a colon-delimted string of time elements. Example:

`E:4:10:?:?:?:2025:07:23:09:28:15:042:654321`


### Absolute Time Representation

The general format of Absolute Clock time is as follows:

`E:O:ER:PE:EP:AG:YYYY:MM:DD:HH:MM:SS:MSS:NSSSSS`

* `E`: **Eternity.** The scope of all time. This is an invariant `E`.
* `O`: **Eon.**
    * pre-Hadean: `0`
    * Hadean: `1`
    * Archean: `2`
    * Proterozoic: `3`
    * Phanerozoic: `4` (the current eon)
    * `5..n` (subsequent eons)
* `ER`: **Era.**
    * pre-Eoarchean: `00` (covering the pre-Hadean and Hadean; neither have eras)
    * Eoarchean: `01`
    * Paleoarchean: `02`
    * Mesoarchean: `03`
    * Neoarchean: `04`
    * Paleoproterozoic: `05`
    * Mesoproterozoic: `06`
    * Neoproterozoic: `07`
    * Paleozoic: `08`
    * Mesozoic: `09`
    * Cenozoic: `10` (the current era)
    * `11..n` (subsequent eras)
* `PE`: **Period.**
    * pre-?: `00`
    * ...
    * Cambrian: `??`
    * ...
    * Quaternary: `??` (the current period)
    * `?..n` (subsequent periods)
* `EP`: **Epoch.** 
    * pre-?: `00`
    * ...
    * Holocene `??` (the current epoch)
    * `?..n` (subsequent epochs)
* `AG`: **Age.**
    * pre-?: `00`
    * ...
    * Meghalayan `??` (the current age)
    * `??..n` (subsequent ages)
* `YYYY`: **Year.**
    * Common Era notation. Negative numbers represent BCE (e.g. `-500` is 500 BCE). There is no `0` value for either BCE or CE.
* `MM`: **Month.** 01-12.
* `DD`: **Day.** 01-31.
* `HH`: **Hour.** 24-hour format. 00-23.
* `MI`: **Minute.** 00-59.
* `SS`: **Second.** 00-59.
* `MSS`: **Millisecond.** 000-999. (1s = 1000ms.)
* `NSSSSS`: **Nanosecond.** 000000-999999. (1ms = 1M ns.)

### More Examples

This is the absolute time of the Big Bang:

`E:0:00:00:00:00:-13800000000:01:01:00:00:00:000:000000`

(this assumes the Big Bang started 13.8 billion years ago on January 1st of that year.)

### Notes
* Time is represented in Zulu time.
* Each time component is a subcomponent of the prior one (immediate left).
* Values for each time element are in chronological order.
* Resolution is determined by what temporal elements, if any, are omitted from the far right.
* Only works on Earth pending implementation of timezones on other planets or moons. Till then UTC is good enough.
* Also, other planets and satellites (like Mars, Venus, and the moon) can have their own geological time systems. There is currently no universal time system in use that goes back to the Big Bang.

### To Do:
* subepoch, subperiod support? 
* picoseconds, microseconds, etc. support?
* geological time element subunits vary based on the element; determine this representation
* Possibly add a way to express durations for cyclical time elements (e.g. something starting at 2pm lasting for an hour could have an hour element of `14-15`)
