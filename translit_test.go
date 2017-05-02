package translit

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"testing"

	. "pkg.re/check.v1"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func Test(t *testing.T) { TestingT(t) }

type TranslitSuite struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

var _ = Suite(&TranslitSuite{})

// ////////////////////////////////////////////////////////////////////////////////// //

func (ts *TranslitSuite) TestUniversal(c *C) {
	c.Assert(Universal(""), Equals, "")

	c.Assert(Universal(
		"Pack my box with five dozen liquor jugs."), Equals,
		"Pack my box with five dozen liquor jugs.")

	c.Assert(
		Universal("Эй, жлоб! Где туз? Прячь юных съёмщиц в шкаф."), Equals,
		"Jej, zhlob! Gde tuz? Prjach' junyh s\"jomshhic v shkaf.")
}

func (ts *TranslitSuite) TestGOST(c *C) {
	c.Assert(GOST779B(""), Equals, "")

	c.Assert(GOST779B(
		"Pack my box with five dozen liquor jugs."), Equals,
		"Pack my box with five dozen liquor jugs.")

	c.Assert(
		GOST779B("Эй, жлоб! Где туз? Прячь юных съёмщиц в шкаф."), Equals,
		"Ehj, zhlob! Gde tuz? Pryach' yunyx s\"yomshhic v shkaf.")
}

func (ts *TranslitSuite) TestISO(c *C) {
	c.Assert(ISO9(""), Equals, "")

	c.Assert(ISO9(
		"Pack my box with five dozen liquor jugs."), Equals,
		"Pack my box with five dozen liquor jugs.")

	c.Assert(
		ISO9("Эй, жлоб! Где туз? Прячь юных съёмщиц в шкаф."), Equals,
		"Èj, žlob! Gde tuz? Prâč' ûnyh s\"ëmŝic v škaf.")
}

func (ts *TranslitSuite) TestYo(c *C) {
	c.Assert(Universal("Ёжик"), Equals, "Jozhik")
}
