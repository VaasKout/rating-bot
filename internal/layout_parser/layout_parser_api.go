package layout_parser

type LayoutParserApi interface {
	GetAppInfo(packageName string) *MarketApp
}
