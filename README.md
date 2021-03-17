# GoStations
### console based radio station selector and player

This is a port of a Python script I wrote, called "radiostations". It is a simple console tool to grab a list of radio stations retrieved from `radio-browser.info`, put them into a menu, and then use your local installation of a console stream player, to play the station for you.


### Requirements
* OS: Linux or macOS
* the `mpv` player (or `mpg123` or `mplayer`, if you change the config)
* Go 1.14+ (if you build your own binary)

### Build
1. clone the repository
2. cd into the root of the project
3. run the following: 
```
> go mod vendor
> go mod tidy
> go build -o /wherever/you/want/gostations github.com/gmgauthier/gostations
```

### Install
1. copy `gostations` to a location that is on your `PATH`
2. copy `radiostations.ini` to a location that is on your `XDG_CONFIG_HOME`
    * Note: if you skip step 2, the app will build an ini file for you automatically.
    
### Execute
```
> gostations -h 
Usage:
 gostations  [-n "name"]  [-c "home country"]  [-s "home state"]  [-t "ordered,tag,list"] [-x]
  -c string
    	Home country.
  -n string
    	Station name (or identifier).
  -s string
    	Home state (if in the United States).
  -t string
    	Tag (or comma-separated tag list)
  -x	If toggled, will show stations that are down
  -h (or none)
	This help message
```
### Examples
```
greg.gauthier@C02DRPKUMD6M $ gostations -c "United Kingdom" -t "news" -n "LBC" 
...Radio Menu...
1) LBC London (London stream) AAC             AAC+  0     http://media-sov.musicradio.com:80/LBCLondon
2) LBC London (London stream) MP3             MP3   48    http://media-ice.musicradio.com/LBCLondonMP3Low
3) LBC London (National stream)               MP3   48    http://media-ice.musicradio.com/LBCUKMP3Low
4) LBC London News                            AAC   48    http://media-ice.musicradio.com/LBC1152.m3u
5) LBC UK                                     AAC   48    http://media-ice.musicradio.com/LBCUK
6) *Quit
What is your choice?
```
```
greg.gauthier@C02DRPKUMD6M $ gostations -c "Gambia"
...Radio Menu...
1) Choice FM                                  MP3   128   http://uk3-pn.webcast-server.net/8276/stream.mp3
2) *Quit
What is your choice?
```
```
greg.gauthier@C02DRPKUMD6M $ gostations -t "chicago,classical"
...Radio Menu...
1) WFMT 98.7 Chicago, IL (MP3)                MP3   0     http://stream.wfmt.com/main-mp3
2) *Quit
What is your choice?
```
When you want to play, you just choose an entry from the menu and hit enter. Your selected stream player will begin playing, and its output will be seen on the console:
```
greg.gauthier@C02DRPKUMD6M $ gostations -s "Illinois" -t "classical"
...Radio Menu...
1) Ancient Faith Radio                        AAC+  96    https://ancientfaith.streamguys1.com/music
2) Lutheran Public Radio - Collinsville, IL   AAC+  0     http://lpr.streamguys1.com/lpr-aac
3) Majesty Radio                              MP3   128   http://primary.moodyradiostream.org/majesty.mp3
4) WFMT 98.7 Chicago, IL (AAC)                AAC   256   http://wowza.wfmt.com/live/smil:wfmt.smil/playlist.m3u8
5) WFMT 98.7 Chicago, IL (MP3)                MP3   0     http://stream.wfmt.com/main-mp3
6) WNIU 90.5 Northern Public Radio Classica   MP3   128   http://peace.str3am.com:6840/live-128k.mp3
7) WUIS-HD2 NPR Illinois Classic - Springfi   MP3   96    http://war.str3am.com:7780/WUISRIS-2
8) WVIK                                       MP3   0     https://wvik.streamguys1.com//live.mp3
9) *Quit
What is your choice?
5
Streaming: WFMT 98.7 Chicago, IL (MP3)                MP3   0     http://stream.wfmt.com/main-mp3
 (+) Audio --aid=1 (mp3 2ch 44100Hz)
AO: [coreaudio] 44100Hz stereo 2ch floatp
A: 00:00:00 / 00:00:06 (14%) Cache: 5.5s/195KB
File tags:
 icy-title: Steiner - Treasure of the Sierra Madre (1948) -  -  - Centaur
A: 00:00:23 / 00:00:46 (49%) Cache: 23s/825KB
```
**NOTE:** If you are using `mpv` and you are on a Mac with a touchbar, and Catalina or better, you will see this error message:
```
2021-03-17 10:51:41.328 mpv[40610:6351061] This NSLayoutConstraint is being configured with a constant that exceeds 
internal limits.  A smaller value will be substituted, but this problem should be fixed. Break 
on BOOL _NSLayoutConstraintNumberExceedsLimit(void) to debug. This will be logged only once.  This may break in the future.
```
This is due to an issue between `mpv` and Apple at the moment, regarding the creation of `mpv` touchbar controls, and can be ignored.

While you are listening, all the normal stdin controls for `mpv` should work properly. So, `9` will lover the volume, `0` will raise the volume, and `m` will mute. To quit, type `q`. When you do, you'll automatically be delivered back to your original menu:
```
 (+) Audio --aid=1 (mp3 2ch 44100Hz)
AO: [coreaudio] 44100Hz stereo 2ch floatp
A: 00:00:00 / 00:00:06 (14%) Cache: 5.5s/195KB
File tags:
 icy-title: Steiner - Treasure of the Sierra Madre (1948) -  -  - Centaur

A: 00:05:52 / 00:06:16 (94%) Cache: 23s/824KB
File tags:
 icy-title: Korngold, Erich Wolfgang - The Sea Wolf film music -  -  - Koch
A: 00:07:42 / 00:08:06 (95%) Cache: 23s/827KB

Exiting... (Quit)
[]
1) Ancient Faith Radio                        AAC+  96    https://ancientfaith.streamguys1.com/music
2) Lutheran Public Radio - Collinsville, IL   AAC+  0     http://lpr.streamguys1.com/lpr-aac
3) Majesty Radio                              MP3   128   http://primary.moodyradiostream.org/majesty.mp3
4) WFMT 98.7 Chicago, IL (AAC)                AAC   256   http://wowza.wfmt.com/live/smil:wfmt.smil/playlist.m3u8
5) WFMT 98.7 Chicago, IL (MP3)                MP3   0     http://stream.wfmt.com/main-mp3
6) WNIU 90.5 Northern Public Radio Classica   MP3   128   http://peace.str3am.com:6840/live-128k.mp3
7) WUIS-HD2 NPR Illinois Classic - Springfi   MP3   96    http://war.str3am.com:7780/WUISRIS-2
8) WVIK                                       MP3   0     https://wvik.streamguys1.com//live.mp3
9) *Quit
What is your choice?
```
To exit the program entirely, choose the __*Quit__ option, or just hit `[ENTER]`

### Additional Notes:

1. The ini file sets a bunch of environmental defaults necessary for the program to work. Gostations looks for it on your `XDG_CONFIG_HOME` path. If that path is not set, or the file is not found on the path, a default version of the ini will be generated automatically, called `radiostations.ini`. The default location for the `XDG_CONFIG_HOME` is `$HOME/.config/gostations`, if you need to edit it. 
2. One of the defaults set in that ini file, is a limit on the number of entries that the program will retrieve from the radio-browser.info api, and the number of entries that will be displayed in the menu. `gostations` isn't really designed to dump the entire radio-browser database to a menu.  I've set the limit to 9999 by default, but it could probably be less. Tuning your searches should help improve the experience.
3. Speaking of that, there are a number of things to keep in mind when searching: 
    * The country search is by country NAME, not CODE. So, "United States" will work, but "US" will not (likewise for "UK") 
    * if you supply multiple tags in a comma-separated list, you may unintentionally filter out results. Unfortunately, the api at radio-info is such that the tag list you search for, has to be in precisely the order it is returned from the host. So, for example, if you search for "classical,chicago", your search will filter out WFMT, because their tags are "chicago,classical".
    * It seems many of the stations put their city in the tag list. So, you can reduce the size of your results by doing something like this: `-c "United States" -t "atlanta"`, which makes more sense for radio stations anyway, for radio stations.
    
### TODO

* Change the precheck, to do ini file validation once, rather than every time a config value is called for.
* Add a way to capture favorite selections. Perhaps a preloaded search or something.
* Add color or a dashboard to the menu and player.
