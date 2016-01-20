# Dioder-API

This is an API for the IKEA dioder.

# Dependencies

This is what you need to run the API:
* a working Go environment ([Instructions](https://golang.org/doc/install))
* Pi-Blaster and the `/dev/pi-blaster`-device
* IKEA-Dioder assembled like in the blog post from @lakrizz (It got deleted, the WayBackMachine archived it: [The piDioder](https://web.archive.org/web/20140427154216/http://krizzblog.de/2013/12/the-pidioder/))

# How do i access the API?

By default, the API listens on port `8080`. There is no API-prefix.

## Methods

* `color/COLOR`
  * COLOR can be anything of the following:
    * red
	* green
	* blue


* `/hex/HEXCODE`
  * `HEXCODE` must be a valid hex-code without `#` (3 or 6 characters)


* `/rgb/REDVALUE/GREENVALUE/BLUEVALUE`
  * The values must be a number within the range `0-255`
