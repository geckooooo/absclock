# The Absolute Clock

### Overview
The Absolute Clock represents a single unique moment or event since the dawn of time (absolute time) as a colon-delimted string of time elements. Example:

`E:4:3:3:2:3:2025:07:23:09:28:15:042:654321` (9:28am UTC on July 23rd, 2025)


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

`E:O:R:PE:EP:AG:YYYY:MM:DD:HH:MM:SS:MSS:NSSSSS`

* `E`: **Eternity.** The scope of all time. This is an invariant `E`.
* `O`: **Eon.**
    * pre-Hadean: `0` (Big Bang to the formation of Earth)
    * Hadean: `1`  (first eon to be divided into geological eras)
    * Archean: `2`
    * Proterozoic: `3`
    * Phanerozoic: `4` (the current eon)
    * `5..n` (future eons)
* `R`: **Era.**
    * pre-Eoarchean: `0` (covering the pre-Hadean and Hadean; neither have eras)
    * Eoarchean: `1` (first era of the Archean eon)
    * Paleoarchean: `2`
    * Mesoarchean: `3`
    * Neoarchean: `4`
    * Paleoproterozoic: `1` (first era of the Proterozoic eon)
    * Mesoproterozoic: `2`
    * Neoproterozoic: `3`
    * Paleozoic: `1` (first Phanerozoic era, first to be divided into geological periods)
    * Mesozoic: `2`
    * Cenozoic: `3` (the current era)
    * `1..n` (future eras)
* `PE`: **Period.**
    * Precambrian: `0` 
    * Cambrian: `1` (first period of the Paleozoic)
    * Ordovician: `2`
    * Silurian: `3` (first period to be divided into geological epochs)
    * Devonian: `4`
    * Carboniferous: `5` 
    * Permian:  `6`
    * Triassic: `1` (first period of the Mesozoic)
    * Jurassic: `2`
    * Cretaceous: `3`
    * Paleogene: `1` (first period of the Cenozoic)
    * Neogene: `2`
    * Quaternary: `3` (the current period)
    * `13..n` (future periods)
* `EP`: **Epoch.** 
    * pre-Llandovery: `0`
    * Llandovery: `1` (first epoch of the Silurian period)
    * Wenlock: `2`
    * Ludlow: `3`
    * Přídolí: `1` (has no ages) 
    * Early Devonian: `1` (first epoch of the Devonian period)
    * Middle Devonian: `2`
    * Late Devonian: `3`
    * Mississippian: `1` (first epoch of the Carboniferous period)
    * Pennsylvanian: `2`
    * ...
    * Cisuralian: `n`
    * Guadalupian: `n`
    * Lopingian: `n`
    * Early Triassic: `1` (first epoch of the Triassic period)
    * Middle Triassic: `1`
    * Late Triassic: `3`
    * Early Jurassic: `1` (first epoch of the Jurassic period)
    * Middle Jurassic: `2`
    * Late Jurassic: `3`
    * Early Cretaceous: `1` (first epoch of the Cretaceous period)
    * Late Cretaceous: `2`
    * Paleocene: `n`
    * Eocene: `n`
    * Oligocene: `n`
    * Miocene: `n`
    * Pliocene: `n`
    * Pleistocene: `1` (first epoch of the Carboniferous period)
    * Holocene `2` (the current epoch)
    * `3..n` (future epochs)
* `AG`: **Age.**
    * pre-Rhuddanian: `0`
    * Rhuddanian: `1` (first age of the Llandovery epoch)
    * Aeronian: `2`
    * Telychian: `3`
    * Sheinwoodian: `1` (first age of the Wenlock epoch)
    * Homerian: `2`
    * Gorstian: `1` (first age of the Ludlow epoch)
    * Ludfordian: `2`
    * Lochkovian: `1` (first age of the Early Devonian epoch)
    * Pragian: `2`
    * Emsian: `3`
    * Eifelian: `1` (first age of the Middle Devonian epoch)
    * Givetian: `2`
    * Frasnian: `1` (first age of the Late Devonian epoch)
    * Famennian: `2`
    * Tournaisian: `1` (first age of the Mississippian epoch)
    * ...
    * Lower Lochkovian: `1` (first age of the Lochkovian epoch)
    * Middle Lochkovian: `2`
    * Upper Lochkovian: `3`
    * ...
    * Hettangian: `n`
    * Danian: `n` (first age of the Paleocene epoch)
    * Selandian: `n`
    * Thanetian: `n`
    * ...
    * Gelasian: `1` (first age of the Pleistocene)
    * Calabrian: `2`
    * Chibanian: `3`
    * Greenlandian: `1` (first age of the Holocene)
    * Northgrippian: `2`
    * Meghalayan `3` (the current age)
    * `1..n` (future ages)
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

* Big Bang: `E:0:0:00:00:00:-13800000000:01:01:00:00:00:000:000000`
    * (this assumes the Big Bang started exactly 13.8 billion years ago.)
* Declaration of Independence: 


### Reference Implementation

This project implements this time notation with a reference implementation written in Go. This simple program spits out the current time in absolute time format.

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
* Geologic time is still being defined and further changes are expected in the coming years.


### To Do:
* geological time enumerations should be scoped to the parent time unit. Fix epoch and age.
* should the Meghalayan age be softcoded/calculated in the reference implementation?
* subepoch, subperiod support? 
* picoseconds, microseconds, etc. support?
* geological time element subunits vary based on the element; determine this representation
* Possibly add a way to express durations for cyclical time elements (e.g. something starting at 2pm lasting for an hour could have an hour element of `14-15`)
* Rationalize Before Present (BP, prior to 1 Jan 1950 CE) and Before Common Era (BCE, prior to 1 CE). 1,950-year differential. (possible solution: make it "BP - 1950" formally, which is effectively identical to unmodified BP)
* Relate to the ISO 8601 time format
* how to represent the absence of any ages for the Přídolí epoch. An age of `0` indicates the time before the first age.
