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

func ExampleEncodeToScientific() {
	fmt.Println(EncodeToScientific(source))
	// output: Širokaja èlektrifikacija južnych gubernij dast moščnyj tolčok pod″ëmu sel′skogo chozjajstva.
}

func ExampleEncodeToISO9A() {
	fmt.Println(EncodeToISO9A(source))
	// output: Širokaâ èlektrifikaciâ ûžnyh gubernij dast moŝnyj tolčok pod″ëmu sel′skogo hozâjstva.
}

func ExampleEncodeToISO9B() {
	fmt.Println(EncodeToISO9B(source))
	// output: Shirokaya e`lektrifikaсiya yuzhny`x gubernij dast moshhny`j tolchok pod``yomu sel`skogo xozyajstva.
}

func ExampleEncodeToBGN() {
	fmt.Println(EncodeToBGN(source))
	// output: Shirokaya elektrifikatsiya yuzhnykh guberniy dast moshchnyy tolchok pod″ëmu sel′skogo khozyaystva.
}

func ExampleEncodeToPCGN() {
	fmt.Println(EncodeToPCGN(source))
	// output: Shirokaya elektrifikatsiya yuzhnykh guberniy dast moshchnyy tolchok pod″ëmu sel′skogo khozyaystva.
}

func ExampleEncodeToALALC() {
	fmt.Println(EncodeToALALC(source))
	// output: Shirokai͡a ėlektrifikat͡sii͡a i͡uzhnykh guberniĭ dast moshchnyĭ tolchok pod″ëmu sel′skogo khozi͡aĭstva.
}

func ExampleEncodeToBS() {
	fmt.Println(EncodeToBS(source))
	// output: Shirokaya élektrifikatsiya yuzhnȳkh guberniĭ dast moshchnȳĭ tolchok pod″ëmu sel′skogo khozyaĭstva.
}

func ExampleEncodeToICAO() {
	fmt.Println(EncodeToICAO(source))
	// output: Shirokaia elektrifikatsiia iuzhnykh gubernii dast moshchnyi tolchok podieemu selskogo khoziaistva.
}
