<p align="center"><a href="#readme"><img src=".github/images/card.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/g/translit.v2"><img src=".github/images/godoc.svg"/></a>
  <a href="https://kaos.sh/r/translit"><img src="https://kaos.sh/r/translit.svg" alt="GoReportCard" /></a>
  <a href="https://kaos.sh/y/translit"><img src="https://kaos.sh/y/222ebbb777bf4867b05d302c23c3f77e.svg" alt="Codacy badge" /></a>
  <br/>
  <a href="https://kaos.sh/w/translit/ci"><img src="https://kaos.sh/w/translit/ci.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/w/translit/codeql"><img src="https://kaos.sh/w/translit/codeql.svg" alt="GitHub Actions CodeQL Status" /></a>
  <a href="https://kaos.sh/c/translit"><img src="https://kaos.sh/c/translit.svg" alt="Coverage Status" /></a>
  <a href="#license"><img src=".github/images/license.svg"/></a>
</p>

<p align="center"><a href="#ci-status">CI Status</a> • <a href="#usage-example">Usage Example</a> • <a href="#license">License</a></p>

<br/>

`translit` is a package for [Go 1.23+](https://github.com/essentialkaos/.github/blob/master/GO-VERSION-SUPPORT.md) for transliteration of Russian text.

Supported output formats:

* Scientific
* ISO 9:1995/A ГОСТ 7.79-2000/A
* ISO 9:1995/B ГОСТ 7.79-2000/Б
* BGN/PCGN
* ALA-LC
* BS 2979:1958
* ICAO (_ИКАО_)

### [Usage example](https://go.dev/play/p/lrrxNRKu8rm)

```go
package main

import (
  "fmt"
  "strings"

  "github.com/essentialkaos/translit/v3"
)

func main() {
  firstName := "Владислав"
  lastName := "Чернявенький"

  transliterator := translit.ICAO

  tFistName := transliterator(strings.ToLower(firstName))[:1]
  tLastName := transliterator(strings.ToLower(lastName))

  fmt.Printf("%s %s → %s.%s\n", firstName, lastName, tFistName, tLastName)
}
```

### CI Status

| Branch | Status |
|------------|--------|
| `master` | [![CI](https://kaos.sh/w/translit/ci.svg?branch=master)](https://kaos.sh/w/translit/ci?query=branch:master) |
| `develop` | [![CI](https://kaos.sh/w/translit/ci.svg?branch=develop)](https://kaos.sh/w/translit/ci?query=branch:develop) |

### License

[Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
