package translit

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "github.com/essentialkaos/check"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type TranslitSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&TranslitSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (ts *TranslitSuite) TestScientific(c *C) {
	c.Assert(Scientific(""), Equals, "")

	c.Assert(Scientific(
		"Pack my box with five dozen liquor jugs."), Equals,
		"Pack my box with five dozen liquor jugs.")

	c.Assert(
		Scientific(
			"Эй, жлоб! Где туз? Прячь юных съёмщиц в шкаф."), Equals,
		"Èj, žlob! Gde tuz? Prjač′ junych s″ëmščic v škaf.")
}

func (ts *TranslitSuite) TestISO9A(c *C) {
	c.Assert(ISO9A(""), Equals, "")

	c.Assert(ISO9A(
		"Pack my box with five dozen liquor jugs."), Equals,
		"Pack my box with five dozen liquor jugs.")

	c.Assert(
		ISO9A("Эй, жлоб! Где туз? Прячь юных съёмщиц в шкаф."), Equals,
		"Èj, žlob! Gde tuz? Prâč′ ûnyh s″ëmŝic v škaf.")
}

func (ts *TranslitSuite) TestISO9B(c *C) {
	c.Assert(ISO9B(""), Equals, "")

	c.Assert(ISO9B(
		"Pack my box with five dozen liquor jugs."), Equals,
		"Pack my box with five dozen liquor jugs.")

	c.Assert(
		ISO9B("Эй, жлоб! Где туз? Прячь юных съёмщиц в шкаф."), Equals,
		"E`j, zhlob! Gde tuz? Pryach` yuny`x s``yomshhicz v shkaf.")

	c.Assert(ISO9B("Цепочка"), Equals, "Cepochka")
}

func (ts *TranslitSuite) TestBGN(c *C) {
	c.Assert(BGN(""), Equals, "")

	c.Assert(BGN(
		"Pack my box with five dozen liquor jugs."), Equals,
		"Pack my box with five dozen liquor jugs.")

	c.Assert(
		PCGN("Эй, жлоб! Где туз? Прячь юных съёмщиц в шкаф."), Equals,
		"Ey, zhlob! Gde tuz? Pryach′ yunykh s″ëmshchits v shkaf.")

	c.Assert(PCGN("Ёжик"), Equals, "Yëzhik")
}

func (ts *TranslitSuite) TestALALC(c *C) {
	c.Assert(ALALC(""), Equals, "")

	c.Assert(ALALC(
		"Pack my box with five dozen liquor jugs."), Equals,
		"Pack my box with five dozen liquor jugs.")

	c.Assert(
		ALALC("Эй, жлоб! Где туз? Прячь юных съёмщиц в шкаф."), Equals,
		"Ėĭ, zhlob! Gde tuz? Pri͡ach′ i͡unykh s″ëmshchit͡s v shkaf.")
}

func (ts *TranslitSuite) TestBS(c *C) {
	c.Assert(BS(""), Equals, "")

	c.Assert(BS(
		"Pack my box with five dozen liquor jugs."), Equals,
		"Pack my box with five dozen liquor jugs.")

	c.Assert(
		BS("Эй, жлоб! Где туз? Прячь юных съёмщиц в шкаф."), Equals,
		"Éĭ, zhlob! Gde tuz? Pryach′ yunȳkh s″ëmshchits v shkaf.")
}

func (ts *TranslitSuite) TestICAO(c *C) {
	c.Assert(ICAO(""), Equals, "")

	c.Assert(ICAO(
		"Pack my box with five dozen liquor jugs."), Equals,
		"Pack my box with five dozen liquor jugs.")

	c.Assert(
		ICAO("Эй, жлоб! Где туз? Прячь юных съёмщиц в шкаф."), Equals,
		"Ei, zhlob! Gde tuz? Priach iunykh sieemshchits v shkaf.")
}

func (ts *TranslitSuite) TestCustom(c *C) {
	c.Assert(Custom("", nil), Equals, "")

	c.Assert(Custom(
		"Pack my box with five dozen liquor jugs.", nil), Equals,
		"Pack my box with five dozen liquor jugs.")

	c.Assert(
		Custom("Эй, жлоб! Где туз? Прячь юных съёмщиц в шкаф.", nil), Equals,
		"Eiy, jlob! Gde tuz? Pryach ynyh sieemshchic v shkaf.")
}

// ////////////////////////////////////////////////////////////////////////////////// //

func (ts *TranslitSuite) BenchmarkBasic(c *C) {
	for i := 0; i < c.N; i++ {
		ICAO("Эй, жлоб! Где туз? Прячь юных съёмщиц в шкаф.")
	}
}

func (ts *TranslitSuite) BenchmarkExt(c *C) {
	for i := 0; i < c.N; i++ {
		ISO9B("Эй, жлоб! Где туз? Прячь юных съёмщиц в шкаф.")
	}
}
