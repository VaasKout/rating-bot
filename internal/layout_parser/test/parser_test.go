package google_parser_test

import (
	"rating-bot/internal/layout_parser"
	"rating-bot/pkg/logger"
	"testing"
)

func TestGetApplicationInfo(t *testing.T) {
	zapLogger := logger.New()
	layoutParser := layout_parser.New(zapLogger, "../../../configs")

	result := layoutParser.GetAppInfo("telegram-messenger/id686449807")
	t.Log(result)

	//net := network.New(zapLogger)
	//response := net.MakeGetRequest("https://apps.apple.com/es/app/telegram-messenger/id686449807")
	//t.Log(response.StatusCode)
	//t.Log(response.Error)
	//t.Log(string(response.Body))
	//result = layoutParser.GetGoogleAppInfo("com.microsoft.math")
	//t.Log(result)
}
