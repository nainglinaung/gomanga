# GOMANGA #


Manga Crawler app written in golang


## Purpose ##

Learning Golang by doing familiar things 


## Installation ##

So far, you can build with  

``` go build main.go -o . ```

and if you want to install globally 

``` go get -u github.com/nainglinaung/gomanga```


## Usage ##

Depending on the provider, you can run as follows:

### mangareader,mangatown,mangazuki,mangapanda,isekaiscan ###

```gomanga -s "mangareader" -m "World Trigger" -c 15```

### nhentai, hentaicafe, hentainexus ###

```gomanga -s "nhentai" -c 88848 ```


## Definition ##



### -s, --site ###

Selecting the website you wish to download. (eg. nhentai,mangareader,mangatown...)

### -m, --manga ##

Selecting the title of the manga you wish to download (eg. bleach)

### -c, --chapter ###

Selecting the chapter you want to download.

### -o, --output (optional) ###

Selecting the destinated location of the downloaded images. If it's not selected, it'll automatically choose `.` as default.

``` ./gomanga -s "mangareader"  -m "bleach" -c 120 -o "output/" ```



## TODO ##

- fix bugs
- refactor
- add multiple manga sources 
- batch download 
- need to zip and change to cbz 
- add mangasy, topmanhua,mangakakalot, mangafreak,manga-here
- otakusmash (comic)
- goroutine to all
