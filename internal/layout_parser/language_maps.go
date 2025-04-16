package layout_parser

var RuneMaps = []map[rune]rune{
	bengalNumbers,
	nepalNumbers,
	persianNumbers,
	symbolsToChange,
}

var bengalNumbers = map[rune]rune{
	'০': '0',
	'১': '1',
	'২': '2',
	'৩': '3',
	'৪': '4',
	'৫': '5',
	'৬': '6',
	'৭': '7',
	'৮': '8',
	'৯': '9',
}

var nepalNumbers = map[rune]rune{
	'०': '0',
	'१': '1',
	'२': '2',
	'३': '3',
	'४': '4',
	'५': '5',
	'६': '6',
	'७': '7',
	'८': '8',
	'९': '9',
}

var persianNumbers = map[rune]rune{
	'۰': '0',
	'۱': '1',
	'۲': '2',
	'۳': '3',
	'۴': '4',
	'۵': '5',
	'۶': '6',
	'۷': '7',
	'۸': '8',
	'۹': '9',
}

var symbolsToChange = map[rune]rune{
	',': '.',
	'٫': '.',
}
