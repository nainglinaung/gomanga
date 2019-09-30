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

### mangareader ###

```gomanga -s "mangareader" -m "bleach" -c 100```

### nhentai ###

```gomanga -s "nhentai" -c 88848 ```

### mangatown ###

```gomanga -s "mangatown" -m "Minamoto-kun Monogatari" -c 150```

## Definition ##


### -c, --chapter ###

Selecting the chapter you want to download 

### -o, --output (optional) ###

Selecting the destinated location of the downloaded images. If it's not selected, it'll automatically choose `.` as default.

``` ./gomanga -s "mangareader"  -m "bleach" -c 120 -o "output/" ```





## TODO ##

- fix bugs
- refactor
- add multiple manga sources 
- batch download 

