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
	c.Assert(EncodeToScientific(""), Equals, "")

	c.Assert(EncodeToScientific(
		"Pack my box with five dozen liquor jugs."), Equals,
		"Pack my box with five dozen liquor jugs.")

	c.Assert(
		EncodeToScientific("Эй, жлоб! Где туз? Прячь юных съёмщиц в шкаф."), Equals,
		"Èj, žlob! Gde tuz? Prjač′ junych s″ëmščic v škaf.")
}

func (ts *TranslitSuite) TestISO9A(c *C) {
	c.Assert(EncodeToISO9A(""), Equals, "")

	c.Assert(EncodeToISO9A(
		"Pack my box with five dozen liquor jugs."), Equals,
		"Pack my box with five dozen liquor jugs.")

	c.Assert(
		EncodeToISO9A("Эй, жлоб! Где туз? Прячь юных съёмщиц в шкаф."), Equals,
		"Èj, žlob! Gde tuz? Prâč′ ûnyh s″ëmŝic v škaf.")
}

func (ts *TranslitSuite) TestISO9B(c *C) {
	c.Assert(EncodeToISO9B(""), Equals, "")

	c.Assert(EncodeToISO9B(
		"Pack my box with five dozen liquor jugs."), Equals,
		"Pack my box with five dozen liquor jugs.")

	c.Assert(
		EncodeToISO9B("Эй, жлоб! Где туз? Прячь юных съёмщиц в шкаф."), Equals,
		"E`j, zhlob! Gde tuz? Pryach` yuny`x s``yomshhicz v shkaf.")

	c.Assert(EncodeToISO9B("Цепленок"), Equals, "Сeplenok")
}

func (ts *TranslitSuite) TestBGN(c *C) {
	c.Assert(EncodeToBGN(""), Equals, "")

	c.Assert(EncodeToBGN(
		"Pack my box with five dozen liquor jugs."), Equals,
		"Pack my box with five dozen liquor jugs.")

	c.Assert(
		EncodeToPCGN("Эй, жлоб! Где туз? Прячь юных съёмщиц в шкаф."), Equals,
		"Ey, zhlob! Gde tuz? Pryach′ yunykh s″ëmshchits v shkaf.")

	c.Assert(EncodeToPCGN("Ёжик"), Equals, "Yëzhik")
}

func (ts *TranslitSuite) TestALALC(c *C) {
	c.Assert(EncodeToALALC(""), Equals, "")

	c.Assert(EncodeToALALC(
		"Pack my box with five dozen liquor jugs."), Equals,
		"Pack my box with five dozen liquor jugs.")

	c.Assert(
		EncodeToALALC("Эй, жлоб! Где туз? Прячь юных съёмщиц в шкаф."), Equals,
		"Ėĭ, zhlob! Gde tuz? Pri͡ach′ i͡unykh s″ëmshchit͡s v shkaf.")
}

func (ts *TranslitSuite) TestBS(c *C) {
	c.Assert(EncodeToBS(""), Equals, "")

	c.Assert(EncodeToBS(
		"Pack my box with five dozen liquor jugs."), Equals,
		"Pack my box with five dozen liquor jugs.")

	c.Assert(
		EncodeToBS("Эй, жлоб! Где туз? Прячь юных съёмщиц в шкаф."), Equals,
		"Éĭ, zhlob! Gde tuz? Pryach′ yunȳkh s″ëmshchits v shkaf.")
}

func (ts *TranslitSuite) TestICAO(c *C) {
	c.Assert(EncodeToICAO(""), Equals, "")

	c.Assert(EncodeToICAO(
		"Pack my box with five dozen liquor jugs."), Equals,
		"Pack my box with five dozen liquor jugs.")

	c.Assert(
		EncodeToICAO("Эй, жлоб! Где туз? Прячь юных съёмщиц в шкаф."), Equals,
		"Ei, zhlob! Gde tuz? Priach iunykh sieemshchits v shkaf.")
}
