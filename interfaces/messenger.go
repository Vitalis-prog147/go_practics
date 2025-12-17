package interfaces

type Messenger interface {
	SendMessage(chatID int64, text string) error
}
