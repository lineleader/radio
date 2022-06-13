# LineLeader Radio

Streams Disney music from known internet radio stations from the command line.
Close those browser tabs and easily choose a stream to hear the music
you want.

## Current Streams

- [Sorcer Radio Atmospheres](http://srsounds.com/popperSRloops.php)
- [DPark Radio background music](https://www.dparkradio.com/dparkradioplayerbm.html)
- [Radio Wonderland Theme Park](https://www.radiowonderland.co.uk/listen-live-theme-park)
- [Sorcer Radio Seasons](https://srsounds.com/popperSRseasons.php)
- [DPark Radio Main Street/Christmas/Halloween](https://www.dparkradio.com/dparkradioplayer1ch3mainstreet.html)
- [Sorcer Radio Mocha](https://srsounds.com/popperSRmocha.php)
- [WDWNTunes](https://live365.com/station/WDWNTunes-a31769)
- [Sorcer Radio Main Stream](https://srsounds.com/popperSRmocha.php)
- [Radio Wonderland Main Stream](https://www.radiowonderland.co.uk/listen-live-main-station)
- [Sorcer Radio Spa Day](https://srsounds.com/popperSRspaday.php)
- [Radio Wonderlan Mellow](https://www.radiowonderland.co.uk/listen-live-mellow)
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

  [x] Atmospheres (Sorcer Radio)        You're Welcome (Moana) - Hollywood Studios Meet and Greet [Hollywood Studio     (01:17 / 02:49)
  [ ] Background (DPark Radio)          Disneyland Main Street Usa 2012 Loop - Disneyland [Disney Parks]        ⠸
  [ ] Theme Park (Radio Wonderland)     Pirates Of The Carribean (Yo Ho...) - Theme park music  (01:00 / 05:42)
  [ ] Seasons (Sorcer Radio)            Fiesta Latina - Oscar Lopez [Heat]      (00:47 / 03:53)
  [ ] Christmas (DPark Radio)           Disney Parks - Jingle Cruise Queue PT1  ⠸
  [ ] Mocha (Sorcer Radio)              Memory Hole - Sorcey [Lands]    ⠸
  [ ] WDWNTunes                         Area Music (Part 2) - Expedition Everest        ⠸
  [ ] Main Stream (Sorcer Radio)        Trashin' the Camp - DCappella   ⠸
  [ ] Main (Radio Wonderland)           Love is an Open Door - Frozen   (00:22 / 01:43)
  [ ] Spa Day (Sorcer Radio)            Inside Out Main Theme - Inside Out [Disney Movie Music] (01:56 / 04:17)
  [ ] Mellow (Radio Wonderland)         A whole new world (Mellow) - Aladdin    (42:43 / 42:53)
> [ ] Resort TV (DPark Radio)           Disney Parks - Disneyland Resort TV 2001        ⠸

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
