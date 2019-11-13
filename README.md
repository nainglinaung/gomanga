# GOMANGA #


Manga Crawler app written in golang


## Purpose ##

Learning Golang by doing familiar things 


## Installation ##

So far, you can build with  

``` go build -o gomanga ./ ```

and if you want to install globally 

``` go get -u github.com/nainglinaung/gomanga```


## Usage ##

Depending on the provider, you can run as follows:

### mangareader,mangatown,mangazuki,mangapanda,isekaiscan,mngdoom,topmanhua ###


If you're downloading manga you need to provide the name of manga and the chapter. 

```gomanga -s "mangareader" -m "World Trigger" -c 15```

### nhentai, hentaicafe, hentainexus ###


You only need to provide code to download hentai. (P.S it's more efficient to download from registered nhentai account than this).

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


## How to Read ##

After Downloading the designated pictures, You need to zip the folder and change the extension to either .cbz or .cbr and feed it to your comic reader of desired platform. In my case, I use [Panel](https://apps.apple.com/us/app/panels-comic-reader/id1236567663)



## TODO ##

- refactor mangatown
- goroutine to hentainexus 
- drop support for nhentai
- add multiple manga sources 
- batch download 
- need to zip and change to cbz 
- add mangasy,mangakakalot, mangafreak
- otakusmash (comic)

