## `translit` [![GoDoc](https://godoc.org/pkg.re/essentialkaos/translit.v2?status.svg)](https://godoc.org/pkg.re/essentialkaos/translit.v2) [![Go Report Card](https://goreportcard.com/badge/github.com/essentialkaos/translit)](https://goreportcard.com/report/github.com/essentialkaos/translit) [![codebeat badge](https://codebeat.co/badges/15851ebb-6715-44b9-be66-0d13dee8b1ee)](https://codebeat.co/projects/github-com-essentialkaos-translit-master) [![Coverage Status](https://coveralls.io/repos/github/essentialkaos/translit/badge.svg)](https://coveralls.io/github/essentialkaos/translit) [![License](https://gh.kaos.io/ekol.svg)](https://essentialkaos.com/ekol)

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

Before the initial install allows git to use redirects for [pkg.re](https://github.com/essentialkaos/pkgre) service (reason why you should do this described [here](https://github.com/essentialkaos/pkgre#git-support)):

```
git config --global http.https://pkg.re.followRedirects true
```

Make sure you have a working Go 1.5+ workspace ([instructions](https://golang.org/doc/install)), then:

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