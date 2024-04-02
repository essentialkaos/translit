package translit

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var source = "Широкая электрификация южных губерний даст мощный толчок подъёму сельского хозяйства."

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleScientific() {
	fmt.Println(Scientific(source))
	// output: Širokaja èlektrifikacija južnych gubernij dast moščnyj tolčok pod″ëmu sel′skogo chozjajstva.
}

func ExampleISO9A() {
	fmt.Println(ISO9A(source))
	// output: Širokaâ èlektrifikaciâ ûžnyh gubernij dast moŝnyj tolčok pod″ëmu sel′skogo hozâjstva.
}

func ExampleISO9B() {
	fmt.Println(ISO9B(source))
	// output: Shirokaya e`lektrifikaсiya yuzhny`x gubernij dast moshhny`j tolchok pod``yomu sel`skogo xozyajstva.
}

func ExampleBGN() {
	fmt.Println(BGN(source))
	// output: Shirokaya elektrifikatsiya yuzhnykh guberniy dast moshchnyy tolchok pod″ëmu sel′skogo khozyaystva.
}

func ExamplePCGN() {
	fmt.Println(PCGN(source))
	// output: Shirokaya elektrifikatsiya yuzhnykh guberniy dast moshchnyy tolchok pod″ëmu sel′skogo khozyaystva.
}

func ExampleALALC() {
	fmt.Println(ALALC(source))
	// output: Shirokai͡a ėlektrifikat͡sii͡a i͡uzhnykh guberniĭ dast moshchnyĭ tolchok pod″ëmu sel′skogo khozi͡aĭstva.
}

func ExampleBS() {
	fmt.Println(BS(source))
	// output: Shirokaya élektrifikatsiya yuzhnȳkh guberniĭ dast moshchnȳĭ tolchok pod″ëmu sel′skogo khozyaĭstva.
}

func ExampleICAO() {
	fmt.Println(ICAO(source))
	// output: Shirokaia elektrifikatsiia iuzhnykh gubernii dast moshchnyi tolchok podieemu selskogo khoziaistva.
}
