package domain

const (
	USER = "user"
)

const (
	GENERIC_START_BUTTON  = "/start"
	GENERIC_BACK_BUTTON   = "Назад"
	GENERIC_CANCEL_BUTTON = "Отмена"
	GENERIC_YES_BUTTON    = "Да"
	GENERIC_NO_BUTTON     = "Нет"
	GENERIC_DONE_BUTTON   = "Готово"
)

const (
	START_STATE = "START_STATE"
)

const (
	START_MESSAGE = "Добро пожаловать в <b>Geo bot!</b>\n\n" +
		"🌐 Для получения рейтинга приложения из Google Play Market введите пакет в формате:\n" +
		"<b>･ org.telegram.messenger</b>\n\n" +
		"📊 Рейтинг из App Store:\n <b>･ id686449807</b>\n\n" +
		"👆Если комбинация выше не работает, добавляем тайтл:\n <b>･ telegram-messenger/id686449807</b>"
	CANCEL_MESSAGE = "Действие отменено"
)

var GenericCancelMarkupButtons = [][]string{
	{GENERIC_BACK_BUTTON, GENERIC_CANCEL_BUTTON},
}

var GenericConfirmMarkupButtons = [][]string{
	{GENERIC_YES_BUTTON},
	{GENERIC_NO_BUTTON},
}

var GenericDoneButtons = [][]string{
	{GENERIC_DONE_BUTTON, GENERIC_CANCEL_BUTTON},
}
