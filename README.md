<p align="center"><a href="#readme"><img src="https://gh.kaos.st/go-translit.svg"/></a></p>

<p align="center"><a href="#installation">Installation</a> • <a href="#build-status">Build Status</a> • <a href="#license">License</a></p>

<p align="center">
  <a href="https://godoc.org/pkg.re/essentialkaos/translit.v2"><img src="https://godoc.org/pkg.re/essentialkaos/translit.v2?status.svg"></a>
  <a href="https://goreportcard.com/report/github.com/essentialkaos/translit"><img src="https://goreportcard.com/badge/github.com/essentialkaos/translit"></a>
  <a href="https://codebeat.co/projects/github-com-essentialkaos-translit-master"><img alt="codebeat badge" src="https://codebeat.co/badges/15851ebb-6715-44b9-be66-0d13dee8b1ee" /></a>
  <a href='https://coveralls.io/github/essentialkaos/translit?branch=master'><img src='https://coveralls.io/repos/github/essentialkaos/translit/badge.svg?branch=master' alt='Coverage Status' /></a>
  <a href="https://essentialkaos.com/ekol"><img src="https://gh.kaos.st/ekol.svg"></a>
</p>

`translit` is a Go package for Russian text transliteration.

Supported output formats:

* Scientific
* ISO 9:1995/A ГОСТ 7.79-2000/A
* ISO 9:1995/B ГОСТ 7.79-2000/Б
* BGN/PCGN
* ALA-LC
* BS 2979:1958
* ICAO

### Installation

Before the initial install allows git to use redirects for [pkg.re](https://github.com/essentialkaos/pkgre) service (_reason why you should do this described [here](https://github.com/essentialkaos/pkgre#git-support)_):

```
git config --global http.https://pkg.re.followRedirects true
```

Make sure you have a working Go 1.8+ workspace (_[instructions](https://golang.org/doc/install)_), then:

```
go get pkg.re/essentialkaos/translit.v2
```

For update to latest stable release, do:

```
go get -u pkg.re/essentialkaos/translit.v2
```

### Build Status

| Branch | Status |
|------------|--------|
| `master` | [![Build Status](https://travis-ci.org/essentialkaos/translit.svg?branch=master)](https://travis-ci.org/essentialkaos/translit) |
| `develop` | [![Build Status](https://travis-ci.org/essentialkaos/translit.svg?branch=develop)](https://travis-ci.org/essentialkaos/translit) |

### License

[EKOL](https://essentialkaos.com/ekol)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
