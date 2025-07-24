# The Absolute Clock

### Overview
The Absolute Clock represents a single unique moment or event since the dawn of time (absolute time) as a colon-delimted string of time elements. Example:

`E:4:10:12:31:?:2025:07:23:09:28:15:042:654321` (9:28am UTC on July 23rd, 2025)


### Problem Statement

Our default representation of time is goofy:
* We use three different representations: clock time, calendar time, geological time.
* They each have their own tools, frameworks, and conceptions. A clock or watch for hours, minutes, and seconds, but a calendar for days, weeks, months, and years. The tools are completely different and unrelated, splitting out our experience of time irrationally.
* As there is no need for multiple representations, Occam's Razor prefers a single representation with a single tool.
* There is currently no universal time system in use that goes back to the Big Bang.
* There is currently no consistent geological time primitives across multiple planets or moons. Each celestial body has its own system.

The Absolute Clock rationalizes these different inconsistent and incomplete perspectives into a single, coherent representation. Every single moment since the Big Bang has an "address" that is unique across all time. By doing so it creates a single namespace for absolute time that can be implemented as (say) a signle device such as a digital clock.


### Absolute Time Representation

The general format of Absolute Clock time is as follows:

`E:O:ER:PE:EP:AG:YYYY:MM:DD:HH:MM:SS:MSS:NSSSSS`

* `E`: **Eternity.** The scope of all time. This is an invariant `E`.
* `O`: **Eon.**
    * pre-Hadean: `0` (Big Bang to the formation of Earth)
    * Hadean: `1`  (first eon to be divided into geological eras)
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
    * Silurian: `03` (first period to be divided into geological epochs)
    * Devonian: `04`
    * Carboniferous: `05` 
    * Permian:  `06`
    * Triassic: `07`
    * Jurassic: `08`
    * Cretaceous: `09`
    * Paleogene: `10`
    * Neogene: `11`
    * Quaternary: `12` (the current period)
    * `13..n` (future periods)
* `EP`: **Epoch.** 
    * pre-Llandovery: `00`
    * Llandovery: `01` (first epoch of the Silurian period)
    * Wenlock:  `02`
    * Ludlow: `03`
    * Přídolí: `04` 
    * Lochkovian: `05`
    * Pragian: `06`
    * Emsian: `07`
    * Eifelian: `08`
    * Givetian: `09`
    * Frasnian: `10`
    * Famennian: `11`
    * Mississippian: `12`
    * Pennsylvanian: `13`
    * Cisuralian: `14`
    * Guadalupian: `15`
    * Lopingian: `16`
    * Early Triassic: `17`
    * Middle Triassic: `18`
    * Late Triassic: `19`
    * Early Jurassic: `20`
    * Middle Jurassic: `21`
    * Late Jurassic: `22`
    * Early Cretaceous: `23`
    * Late Cretaceous: `24`
    * Paleocene: `25`
    * Eocene: `26`
    * Oligocene: `27`
    * Miocene: `28`
    * Pliocene: `29`
    * Pleistocene: `30`
    * Holocene `31` (the current epoch)
    * `32..n` (future epochs)
* `AG`: **Age.**
    * pre-Rhuddanian: `00`
    * Rhuddanian: `01` (first age of the Llandovery epoch)
    * Aeronian
    * Telychian
    * Sheinwoodian
    * Homerian
    * Gorstian
    * Ludfordian
    * ...
    * Hettangian: `01`
    * Danian: `01` (first age of the Paleocene epoch)
    * Selandian: `02`
    * Thanetian: `03`
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

### More Absolute Time Examples

* Big Bang: `E:0:00:00:00:00:-13800000000:01:01:00:00:00:000:000000`
    * (this assumes the Big Bang started exactly 13.8 billion years ago.)
* Declaration of Independence: 


### TL;DR
To get the current absolute time: `% go run absclock.go`


### Notes
* Time is represented in UTC.
* Only works on Earth pending implementation of consistent geological time primitives across multiple planets or moons. Currently other planets and satellites (like Mars, Venus, and the moon) have their own specific geological time.
* Each time component is a subcomponent of the prior one (immediate left).
* Values for each time element are in chronological order.
* Resolution is determined by what temporal elements, if any, are omitted from the far right.
* Time units greater than the year are based on International Commission on Stratigraphy (ICS) geochronology.
* Also, other planets and satellites (like Mars, Venus, and the moon) can have their own geological time systems. There is currently no universal time system in use that goes back to the Big Bang.
    * That said, there could be Mars, Venus, and moon versions of this clock!

### To Do:
* subepoch, subperiod support? 
* picoseconds, microseconds, etc. support?
* geological time element subunits vary based on the element; determine this representation
* Possibly add a way to express durations for cyclical time elements (e.g. something starting at 2pm lasting for an hour could have an hour element of `14-15`)
* Rationalize Before Present (BP, prior to 1 Jan 1950 CE) and Before Common Era (BCE, prior to 1 CE). 1,950-year differential. (possible solution: make it "BP - 1950" formally, which is effectively identical to unmodified BP)
* Relate to the ISO 8601 time format
