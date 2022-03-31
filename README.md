# LineLeader Radio

Streams Disney music from known internet radio stations from the command line.
Close those browser tabs and easily choose a stream to hear the music
you want.

## Current Streams

- [Sorcer Radio Atmospheres](http://srsounds.com/popperSRloops.php)
- [DPark Radio background music](https://www.dparkradio.com/dparkradioplayerbm.html)
- [Sorcer Radio Seasons](https://srsounds.com/popperSRseasons.php)
- [DPark Radio Main Street/Christmas](https://www.dparkradio.com/dparkradioplayer1ch3mainstreet.html)
- [Sorcer Radio Mocha](https://srsounds.com/popperSRmocha.php)
- [WDWNTunes](https://live365.com/station/WDWNTunes-a31769)
- [Sorcer Radio Main Stream](https://srsounds.com/popperSRmocha.php)
- [Sorcer Radio Spa Day](https://srsounds.com/popperSRspaday.php)
- [Sorcer Radio Resorrt TV](https://www.dparkradio.com/dparkradioplayer1ch4.html)

## Prerequisites

### Linux

```
sudo apt install libvlc-dev vlc libnotify-dev
```

## Installation

[Download the latest
release](https://github.com/lineleader/radio/releases) and add
to your path.

```
$ mv radio ~/bin/
```

From source, `go run main.go` is sufficient to try things out. Or you can `go
build` and run the created executable. 

## Usage

No args are required. When the program is running, the streams are displayed.
Use arrows or `j/k` to move the cursor between streams. Press `Enter` to select
a stream to start playing it.

```
$ radio
Station list

  [ ] Atmospheres (Sorcer Radio)        3 Cheeky Chicks Wax Company - magicallyscented.com [SRad2]      (Loading...)
  [ ] Background (DPark Radio)          Walt Disney Studios Park Daytime Entrance Loop pt2 - Walt Disney Studios [D     (Loading...)
  [ ] Seasons (Sorcer Radio)            Lava - Polynesian Resort [WDW Resort]   (01:35 / 05:44)
  [ ] Main Street (DPark Radio)         Disney Parks - Hollywood Boulevard Main Street pt 1 Hollywood Studios   (Loading...)
  [ ] Mocha (Sorcer Radio)              Festival of Food (Android & IOS) app - Festival of Food [SRad3] (00:27 / 00:45)
> [x] WDWNTunes                         Music Loop 9 - Plaza Inn        (Loading...)
  [ ] Main Stream (Sorcer Radio)        It's Ray-Ray Raining - The Dapper Dans  (Loading...)
  [ ] Spa Day (Sorcer Radio)            She's My Reason - Contemporary Resort [WDW Resort]      (02:03 / 05:40)
  [ ] Resort TV (DPark Radio)           Disney Parks - WDW Downtown Disney 2015 (Loading...)

Press q to quit.
```

## Contributing

The current status of this project is `hm, ok`. Many band-aids and duct
tape were used. Music started playing and information I wanted to see was shown
then I paused. I'm very much open to accepting contributions. Check out the 
current issues and comment on any that you'd like to work on. Please, create a
new issue before working on something new; just to discuss before you sink time
into something that might have other implications.

Thanks in advance for your work and help!
