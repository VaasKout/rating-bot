package layout_parser

import (
	"fmt"
	"net/http"
	"rating-bot/pkg/core/file"
	"rating-bot/pkg/core/network"
	"rating-bot/pkg/logger"
	"strings"
)

type LayoutParserImpl struct {
	configsPath string
	filesApi    file.FilesApi[[]CountryDB]
	network     network.NetworkApi
	logger      *logger.Logger
}

func New(
	logger *logger.Logger,
	configPath string,
) LayoutParserApi {
	filesApi := file.New[[]CountryDB]()
	networkApi := network.New(logger)
	return &LayoutParserImpl{
		configsPath: configPath,
		filesApi:    filesApi,
		network:     networkApi,
		logger:      logger,
	}
}

func (parser *LayoutParserImpl) GetAppInfo(packageName string) *MarketApp {
	countryDbArray := &[]CountryDB{}
	if strings.Contains(packageName, ".") {
		countryDbArray = parser.mapCountryDbJson("country_database_gp.json")
	}
	if !strings.Contains(packageName, ".") && strings.Contains(packageName, "id") {
		countryDbArray = parser.mapCountryDbJson("country_database_ios.json")
	}

	if len(*countryDbArray) == 0 {
		parser.logger.Log.Error("COUNTRY DB IS EMPTY")
		return &MarketApp{}
	}

	var app = parser.initMarketApp(packageName)
	if app == nil || app.Title == "" {
		return app
	}

	app.Rating = *parser.getStarRating(countryDbArray, packageName)
	return app
}

func (parser *LayoutParserImpl) initMarketApp(packageName string) *MarketApp {
	countryDbArray := parser.mapCountryDbJson("country_database_info.json")
	if len(*countryDbArray) == 0 {
		return &MarketApp{}
	}

	var marketApp MarketApp
	for _, item := range *countryDbArray {
		url := getUrlByPackage(packageName, &item)
		fmt.Println(url)
		var htmlResponse = parser.network.MakeGetRequest(url)
		if htmlResponse.Error != nil || htmlResponse.StatusCode != http.StatusOK || len(htmlResponse.Body) == 0 {
			fmt.Println(fmt.Sprintf("Error: %s", htmlResponse.Error))
			continue
		}
		marketApp.Title = getTitle(string(htmlResponse.Body))
		marketApp.Developer = getDeveloper(string(htmlResponse.Body))
		marketApp.Package = packageName
		marketApp.FoundUrl = url

		parser.logger.Log.Info(
			fmt.Sprintf(
				`Title : %s, Developer: %s, Package: %s, URL: %s`,
				marketApp.Title,
				marketApp.Developer,
				marketApp.Package,
				marketApp.FoundUrl,
			),
		)
		break
	}

	return &marketApp
}

func (parser *LayoutParserImpl) mapCountryDbJson(configName string) *[]CountryDB {
	var countryDbArray = make([]CountryDB, 0)
	err := parser.filesApi.ReadJsonFile(fmt.Sprintf("%s/%s", parser.configsPath, configName), &countryDbArray)
	if err != nil {
		fmt.Println(err)
	}
	return &countryDbArray
}
