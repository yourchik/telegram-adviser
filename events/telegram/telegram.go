package telegram

import "telegram-adviser/clients/telegram"

type Dispatcher struct {
	tf     *telegram.Client
	offset int
}
