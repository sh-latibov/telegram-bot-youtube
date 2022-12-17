package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	errInvalidURL   = errors.New("url is invalid")
	errUnauthorized = errors.New("user is not authorized")
	errUnableToSave = errors.New("unable to save")
)

func (b *Bot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, "Unknown error")

	switch err {
	case errInvalidURL:
		msg.Text = "Link not validated!"
		b.bot.Send(msg)
	case errUnauthorized:
		msg.Text = "You not authorized! Try command /start"
		b.bot.Send(msg)
	case errUnableToSave:
		msg.Text = "Oops! Try send again later."
		b.bot.Send(msg)
	default:
		b.bot.Send(msg)
	}
}
