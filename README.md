# The Absolute Clock

### Overview
The Absolute Clock represents a single unique moment or event since the dawn of time (absolute time) as a colon-delimted string of time elements. Example:

`E:4:10:11:7:?:2025:07:23:09:28:15:042:654321` (9:28am UTC on July 23rd, 2025)


### Problem Statement

Our default representation of time is goofy:
* We use three different representations: clock time, calendar time, geological time.
* They each have their own tools, frameworks, and conceptions. A clock or watch for hours, minutes, and seconds, but a calendar for days, weeks, months, and years. The tools are completely different and unrelated, splitting out our experience of time irrationally.
* As there is no need for multiple representations, Occam's Razor prefers a single representation with a single tool.

The Absolute Clock rationalizes these different perspectives into a single, coherent representation. Every single moment since the Big Bang has an "address" that is unique across all time. By doing so it creates a single namespace for absolute time that can be implemented as (say) a signle device such as a digital clock.


### Absolute Time Representation

The general format of Absolute Clock time is as follows:

`E:O:ER:PE:EP:AG:YYYY:MM:DD:HH:MM:SS:MSS:NSSSSS`

* `E`: **Eternity.** The scope of all time. This is an invariant `E`.
* `O`: **Eon.**
    * pre-Hadean: `0` (Big Bang to the formation of Earth)
    * Hadean: `1`
    * Archean: `2`
    * Proterozoic: `3`
    * Phanerozoic: `4` (the current eon)
    * `5..n` (future eons)
* `ER`: **Era.**
    * pre-Eoarchean: `00` (covering the pre-Hadean and Hadean; neither have eras)
    * Eoarchean: `01`
    * Paleoarchean: `02`
    * Mesoarchean: `03`
    * Neoarchean: `04`
    * Paleoproterozoic: `05`
    * Mesoproterozoic: `06`
    * Neoproterozoic: `07`
    * Paleozoic: `08` (first era to be divided into geological periods)
    * Mesozoic: `09`
    * Cenozoic: `10` (the current era)
    * `11..n` (future eras)
* `PE`: **Period.**
    * Precambrian: `00` 
    * Cambrian: `01` (first period of the Paleozoic)
    * Ordovician: `02`
    * Silurian: `03`
    * Devonian: `04`
    * Carboniferous: `05` 
    * Permian:  `06`
    * Triassic: `07`
    * Jurassic: `08`
    * Paleogene: `09` (first period to be divided into geological epochs)
    * Neogene: `10`
    * Quaternary: `11` (the current period)
    * `12..n` (future periods)
* `EP`: **Epoch.** 
    * pre-Paleocene: `0`
    * Paleocene: `1` (first epoch of the Paleogene)
    * Eocene: `2`
    * Oligocene: `3`
    * Miocene: `4`
    * Pliocene: `5`
    * Pleistocene: `6`
    * Holocene `7` (the current epoch)
    * `8..n` (future epochs)
* `AG`: **Age.**
    * pre-?: `00`
    * ...
    * Meghalayan `??` (the current age)
    * `??..n` (future ages)
* `YYYY`: **Year.**
    * Common Era notation. Negative numbers represent BCE (e.g. `-500` is 500 BCE). There is no `0` value for either BCE or CE. Conventionally 4 digits, but can be any length.
* `MM`: **Month.** 01-12.
* `DD`: **Day.** 01-31.
* `HH`: **Hour.** 24-hour format. 00-23.
* `MI`: **Minute.** 00-59.
* `SS`: **Second.** 00-59.
* `MSS`: **Millisecond.** 000-999. (1s = 1000ms.)
* `NSSSSS`: **Nanosecond.** 000000-999999. (1ms = 1M ns.)

### More Examples
The Absolute Clock 

* Big Bang: `E:0:00:00:0:00:-13800000000:01:01:00:00:00:000:000000`
    * (this assumes the Big Bang started 13.8 billion years ago on January 1st of that year.)


### TL;DR
To get the current absolute time: `% go run absclock.go`


### Notes
* Time is represented in Zulu time.
* Each time component is a subcomponent of the prior one (immediate left).
* Values for each time element are in chronological order.
* Resolution is determined by what temporal elements, if any, are omitted from the far right.
* Only works on Earth pending implementation of timezones on other planets or moons. Till then UTC is good enough.
* Also, other planets and satellites (like Mars, Venus, and the moon) can have their own geological time systems. There is currently no universal time system in use that goes back to the Big Bang.
    * That said, there could be Mars, Venus, and moon versions of this clock!

### To Do:
* subepoch, subperiod support? 
* picoseconds, microseconds, etc. support?
* geological time element subunits vary based on the element; determine this representation
* Possibly add a way to express durations for cyclical time elements (e.g. something starting at 2pm lasting for an hour could have an hour element of `14-15`)
* Rationalize Before Present (BP, prior to 1 Jan 1950 CE) and Before Common Era (BCE, prior to 1 CE)
