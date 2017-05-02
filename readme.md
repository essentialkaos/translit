## `translit` [![GoDoc](https://godoc.org/pkg.re/essentialkaos/translit.v1?status.svg)](https://godoc.org/pkg.re/essentialkaos/translit.v1) [![Go Report Card](https://goreportcard.com/badge/github.com/essentialkaos/translit)](https://goreportcard.com/report/github.com/essentialkaos/translit) [![codebeat badge](https://codebeat.co/badges/ffffff)](https://codebeat.co/projects/github-com-essentialkaos-translit) [![Coverage Status](https://coveralls.io/repos/github/essentialkaos/translit/badge.svg?branch=develop)](https://coveralls.io/github/essentialkaos/translit?branch=develop) [![License](https://gh.kaos.io/ekol.svg)](https://essentialkaos.com/ekol)

`translit` is a Go package for Russian text transliteration.

### Installation

Before the initial install allows git to use redirects for [pkg.re](https://github.com/essentialkaos/pkgre) service (reason why you should do this described [here](https://github.com/essentialkaos/pkgre#git-support)):

```
git config --global http.https://pkg.re.followRedirects true
```

Make sure you have a working Go 1.5+ workspace ([instructions](https://golang.org/doc/install)), then:

```
go get pkg.re/essentialkaos/translit.v1
```

For update to latest stable release, do:

```
go get -u pkg.re/essentialkaos/translit.v1
```

### Build Status

| Branch | Status |
|------------|--------|
| `master` | [![Build Status](https://travis-ci.org/essentialkaos/translit.svg?branch=master)](https://travis-ci.org/essentialkaos/translit) |
| `develop` | [![Build Status](https://travis-ci.org/essentialkaos/translit.svg?branch=develop)](https://travis-ci.org/essentialkaos/translit) |

### License

[EKOL](https://essentialkaos.com/ekol)