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

``` gomanga -s "mangareader" -m "bleach" -c 100"```

### nhentai ###

``` gomanga -s "nhentai" -c 88848 ````

### mangatown ###

``` gomanga -s "mangatown" -m "Minamoto-kun Monogatari" -c 150```



## TODO ##

- fix bugs
- refactor
- add multiple manga sources 
- batch download 

