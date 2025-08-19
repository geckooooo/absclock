# The Absolute Clock
_A novel way to represent time_

### Overview
You may be asking yourself... what exactly is this Absolute Clock thing suppose to be? Is this strange, seemingly useless invention perhaps the product of a fevered imagination?

The Absolute Clock is a major departure from time as we experience it in the boringly familiar ways. It represents a single unique moment or event since the dawn of time--absolute time, as perceived from the observer's frame of reference--as a colon-delimted string of hierarchically contained time elements. Example:

`E:4:3:3:2:3:2025:07:23:09:28:15:042:654321` (9:28am UTC on July 23rd, 2025)

More on this below. 

The Absolute Clock is... my answer to a chronic nuisance. I do not like how we have enculturated time, and here I reject it. There is no elegance in having clocks and calendars as separate, basically incompatible tools. There is no value in trapping the clunky, split representation in its visual presentation, as if we’re stuck in some cognitive skeuomorphic fixation. There is no respect for the profound, almost extravagant timespans of the cosmos and all that has occurred. I see no reflection in it of the deep truths of relativity in spacetime. It is incomplete, incoherent, and hidebound. It doesn't connect us to the cosmos the way it can.

It is refreshingly unencumbered with practical considerations. For the most part! (the truncation, wildcards, and duration expressions--defined below--notwithstanding.)

Is there any real-world value to this? Who knows! I think it’s kinda cool, with some pretty interesting qualities, some of which I didn’t expect, and I find it satisfyingly elegant. That’s why I created it. Check it out!


### Problem Statement

The usual ways we express our experience of time is quite janky:

* We use three different modes of representation to make sense of a single universal natural phenomenon: clock time, calendar time, geological time.
* They each have their own tools, frameworks, and conceptions. A clock or watch for hours, minutes, and seconds, but a calendar for days, weeks, months, and years. The tools are completely different and unrelated, splitting out our experience of time irrationally.
* As there is no need for multiple representations, Occam's Razor prefers a single representation with a single tool.
* There is currently no universal time system in use that goes back to the Big Bang.
* There is currently no consistent geological time primitives across multiple planets or moons. Each celestial body has its own system.
* It's awkward at best to express durations or repeating/cyclical events. Calendars sometime use visualizations to give this (size of the box on the calendar) but repetition has a separate representation (such as a "repeats weekly" dropdown).

What's missing is a way of expressing time in a way that deeply connects to the cosmos--which ultimately is the true scope of time. The Absolute Clock rationalizes these different inconsistent, incomplete, and narrow perspectives into a unified coherent representation. Every single moment since the Big Bang has an "address" that is unique across all time. By doing so it creates a single namespace for absolute time that can be implemented as (say) a single device such as a digital clock.


### Absolute Time Representation

#### Structure

* Each time component is a subcomponent of the prior one (immediate left).
* Values for each time element are in chronological order.
* Time units greater than the year are based on International Commission on Stratigraphy (ICS) geochronology.

> [!NOTE]
> By default, geological time represents years by Before Present (BP, prior to 1 Jan 1950 CE). However, non-geological time in the past is expressed using Common Era/Before Common Era (e.g. using BCE prior to 1 CE).
> As there's a 1,950-year differential between the two, and the differential is of no consequence for geological timeframes, the Absolute Clock standardizes on CE/BCE.

#### Format

The general format of Absolute Clock time is as follows:

`E:O:R:P:C:A:YYYY:MM:DD:HH:MM:SS:MSS:NSSSSS`

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
    * Paleoproterozoic: `1` (first era of the Proterozoic eon. first to be divided into geological periods)
    * Mesoproterozoic: `2`
    * Neoproterozoic: `3`
    * Paleozoic: `1` (first era of the Phanerozoic eon)
    * Mesozoic: `2`
    * Cenozoic: `3` (the current era)
    * `1..n` (first era of a future eon)
* `P`: **Period.**
    * pre-Siderian: `0` 
    * Siderian: `1` (first period of the Paleoproterozoic era)
    * Rhyacian: `2`
    * Orosirian: `3`
    * Statherian: `4`
    * Calymmian: `1` (first period of the Mesoproterozoic era)
    * Ectasian: `2`
    * Stenian: `3`
    * Tonian: `1` (first period of the Neoproterozoic era)
    * Cryogenian: `2`
    * Ediacaran: `3`
    * Cambrian: `1` (first period of the Paleozoic era. first to be divided into geological epochs)
    * Ordovician: `2`
    * Silurian: `3`
    * Devonian: `4`
    * Carboniferous: `5` 
    * Permian:  `6`
    * Triassic: `1` (first period of the Mesozoic era)
    * Jurassic: `2`
    * Cretaceous: `3`
    * Paleogene: `1` (first period of the Cenozoic era)
    * Neogene: `2`
    * Quaternary: `3` (the current period)
    * `1..n` (first period of a future era)
* `C`: **Epoch.** 
    * pre-Terreneuvian: `0`
    * Terreneuvian: `1` (first epoch of the Cambrian period. first to be divided into geological ages)
    * Cambrian Series 2: `2` (epoch yet to be named)
    * Miaolingian: `3`
    * Furongian: `4`
    * Early Ordovician: `1` (first epoch of the Ordovician period)
    * Middle Ordovician: `2`
    * Late Ordovician: `3`
    * Llandovery: `1` (first epoch of the Silurian period)
    * Wenlock: `2`
    * Ludlow: `3`
    * Přídolí: `4` (has no ages) 
    * Early Devonian: `1` (first epoch of the Devonian period)
    * Middle Devonian: `2`
    * Late Devonian: `3`
    * Mississippian: `1` (first epoch of the Carboniferous period)
    * Pennsylvanian: `2`
    * Cisuralian: `1` (first epoch of the Permian period)
    * Guadalupian: `2`
    * Lopingian: `3`
    * Early Triassic: `1` (first epoch of the Triassic period)
    * Middle Triassic: `2`
    * Late Triassic: `3`
    * Early Jurassic: `1` (first epoch of the Jurassic period)
    * Middle Jurassic: `2`
    * Late Jurassic: `3`
    * Early Cretaceous: `1` (first epoch of the Cretaceous period)
    * Late Cretaceous: `2`
    * Paleocene: `1` (first epoch of the Paleogene period)
    * Eocene: `2`
    * Oligocene: `3`
    * Miocene: `1` (first epoch of the Neogene period)
    * Pliocene: `2`
    * Pleistocene: `1` (first epoch of the Quaternary period)
    * Holocene `2` (the current epoch)
    * `1..n` (first epoch of a future period)
* `A`: **Age.**
    * pre-Fortunian: `0`
    * Fortunian: `1` (first age of the Terreneuvian epoch)
    * Stage 2: `2` (age yet to be named)
    * Stage 3: `1` (first age of the Cambrian Series 2 epoch)
    * Stage 4: `2`
    * Wuliuan: `1` (first age of the Miaolingian epoch)
    * Drumian: `2`
    * Guzhangian: `3`
    * Paibian: `1` (first age of the Furongian epoch)
    * Jiangshanian: `2`
    * Stage 10: `3` (age yet to be named)
    * Tremadocian: `1` (first age of the Early Ordovician epoch)
    * Floian: `2`
    * Dapingian: `1` (first age of the Middle Ordovician epoch)
    * Darriwilian: `2`
    * Sandbian: `1` (first age of the Late Ordovician epoch)
    * Katian: `2`
    * Hirnantian: `3`
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
    * Visean: `2`
    * Serpukhovian: `3`
    * Bashkirian: `1` (first age of the Pennsylvanian epoch)
    * Moscovian: `2`
    * Kasimovian: `3`
    * Gzhelian: `4`
    * Asselian: `1` (first age of the Cisuralian epoch)
    * Sakmarian: `2`
    * Artinskian: `3`
    * Kungurian: `4`
    * Roadian: `1` (first age of the Guadalupian epoch)
    * Wordian: `2`
    * Capitanian: `3`
    * Wuchiapingian: `1` (first age of the Lopingian epoch)
    * Changhsingian: `2`
    * Induan: `1` (first age of the Early Triassic epoch)
    * Olenekian: `2`
    * Anisian: `1` (first age of the Middle Triassic epoch)
    * Ladinian: `2`
    * Carnian: `1` (first age of the Late Triassic epoch)
    * Norian: `2`
    * Rhaetian: `3`
    * Hettangian: `1` (first age of the Early Jurassic epoch)
    * Sinemurian: `2`
    * Pliensbachian: `3`
    * Toarcian: `4`
    * Aalenian: `1` (first age of the Middle Jurassic epoch)
    * Bajocian: `2`
    * Bathonian: `3`
    * Callovian: `4`
    * Oxfordian: `1` (first age of the Late Jurassic epoch)
    * Kimmeridgian: `2`
    * Tithonian: `3`
    * Berriasian: `1` (first age of the Early Cretaceous epoch)
    * Valanginian: `2`
    * Hauterivian: `3`
    * Barremian: `4`
    * Aptian: `5`
    * Albian: `6`
    * Cenomanian: `1` (first age of the Late Cretaceous epoch)
    * Turonian: `2`
    * Coniacian: `3`
    * Santonian: `4`
    * Campanian: `5`
    * Maastrichtian: `6`
    * Danian: `1` (first age of the Paleocene epoch)
    * Selandian: `2`
    * Thanetian: `3`
    * Ypresian: `1` (first age of the Eocene epoch)
    * Lutetian: `2`
    * Bartonian: `3`
    * Priabonian: `4`
    * Rupelian: `1` (first age of the Oligocene epoch)
    * Chattian: `2`
    * Aquitanian: `1` (first age of the Miocene epoch)
    * Burdigalian: `2`
    * Langhian: `3`
    * Serravallian: `4`
    * Tortonian: `5`
    * Messinian: `6`
    * Zanclean: `1` (first age of the Pliocene epoch)
    * Piacenzian: `2`
    * Gelasian: `1` (first age of the Pleistocene epoch)
    * Calabrian: `2`
    * Chibanian: `3`
    * Late Pleistocene: `4`
    * Greenlandian: `1` (first age of the Holocene epoch)
    * Northgrippian: `2`
    * Meghalayan `3` (the current age)
    * `1..n` (first age of a future epoch)
* `YYYY`: **Year.**
    * Common Era notation. Negative numbers represent BCE (e.g. `-500` is 500 BCE). There is no `0` value for either BCE or CE. This field is conventionally 4 digits for current dates but can be any length.
* `MM`: **Month.** 01-12.
* `DD`: **Day.** 01-31.
* `HH`: **Hour.** 24-hour format. 00-23.
* `MI`: **Minute.** 00-59.
* `SS`: **Second.** 00-59.
* `MSS`: **Millisecond.** 000-999. (1s = 1000ms.)
* `NSSSSS`: **Nanosecond.** 000000-999999. (1ms = 1M ns.)

> [!NOTE]
> Geological time is used here because (as far as I can tell) there is no better expression for timespans leading back to the Big Bang. (And even this only goes back to the formation of Earth, which leaves a multi-billion year gap.) However, other planets and satellites (like Mars, Venus, and the moon) have their own specific (incompatible) geological time systems. So this implementation of the Absolute Clock only works on Earth absent the implementation of consistent geological time primitives across multiple planets or moons.

### Notation Variations

#### Truncation

the first portion of the string prior to the year--from eternity to age, `E:4:3:3:2:3:`--isn't likely to be relevant on a day to day basis for most people and can be omitted for practicality. The notation for ignoring unwanted larger time units on the left (for example, leaving only year to nanoseconds) is:

`YYYY:MM:DD:HH:MM:SS:MSS:NSSSSS`

Resolution of the expression is determined by what temporal elements, if any, are omitted from the right. To ignore smaller units, say for rounding to the minute, simply truncate them:

`E:O:R:P:C:A:YYYY:MM:DD:HH:MM`

And of couse you can truncate from both ends at once:

`YYYY:MM:DD:HH:MM`

#### Wildcards

Dots (`.`) can be used as wildcards for individual fields. If you want to indicate units to leave explicitly unspecified or that can be ignored, you can use the dot notation like this:

`E:O:R:P:C:A:YYYY:.:.:HH:.:.:.:.`

Setting a field to `.` (as above) explicitly states "can be any legal value" In this example, this use of dot notation could express an absolute time for an event that occurs at a specific hour every day for a specific year.

> [!NOTE]
> Dot notation is used in the special case of Přídolí epoch, which has no ages. When referring to that epoch, use `.` for the age.


#### Durations

Durations can be expressed with dashes (`-`) between two values in a field. For example:

`E:4:3:3:2:3:2025:12:25:16-21`

This expresses an event happening on Chrismas 2025 at 4-9pm (maybe Christmas dinner).


### Absolute Time and Relativity

The Absolute Clock is compatible with relativistic spacetime. A unique literal string identifies an event *t*. That may seem to imply an absolute frame of reference, which (per Einstein) doesn't exist. However, if two observers perceived a the chronology of an event differently, absolute time can have two separate values in the same schema--in other words, two different strings for the two different observers can reference the same event.

(give a simple example.)

> [!NOTE]
> As absolute time is scoped within a frame of reference, spacetime *(t, x, y, z)* is not represented formally here.

Practically speaking, as geologic time uses Earth as the frame of reference, so does this implementation of the Absolute Clock. 


### More Absolute Time Examples

* Big Bang: `E:0:0:0:0:0:-13800000000` (13.8 billion years ago)
* Declaration of Independence: `E:4:3:3:2:3:1776:07:04` (July 4th, 1776)
* The Hundred Years' War: `E:4:3:3:2:3:1337-1453` (1337–1453 CE)


### Reference Implementation

This project implements this time notation with a reference implementation written in Go. This simple program:

* spits out the current time in absolute time format
* illustrates how ISO 8601 dates are converted to absolute time

To get the current absolute time: `% go run absclock.go`


### Notes

* Time at the level of days, hours, etc. is represented in UTC.
* Also, other planets and satellites (like Mars, Venus, and the moon) can have their own geological time systems. There is currently no universal time system in use that goes back to the Big Bang.
    * That said, there could be Mars, Venus, and moon versions of this clock!
* Geologic time is still being defined and further changes are expected in the coming years.


### To Do:

* consider subepoch, subperiod support
* Consider support for specifying timezone (maybe as an arbitrary string appended to the hour or minute?)
* Consider changing the nanosecond field to "subsecond" with arbitrary precision. or another way to represent picoseconds, microseconds, etc. support
* Consider ways of representing relative time in absolute format (e.g. "the day before yesterday" computed as absolute time from the current day)
* consider expanded notation for expressing spacetime
* Give examples of date math. Provide function that computes dates.
