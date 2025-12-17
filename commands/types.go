package commands

type CommandType string

const (
	CommandStart CommandType = "/start"
	CommandScore CommandType = "/score"
	CommandAdd   CommandType = "/add"
	CommandHelp  CommandType = "/help"
	CommandNone  CommandType = ""
)

func GetAllCommands() []CommandType {
	return []CommandType{
		CommandStart,
		CommandScore,
		CommandAdd,
		CommandHelp,
	}
}

func (c CommandType) GetDescription() string {
	switch c {
	case CommandStart:
		return "начать работу с ботом"
	case CommandScore:
		return "показать таблицу очков"
	case CommandAdd:
		return "добавить очко игроку (использование: /add [имя])"
	case CommandHelp:
		return "показать справку по командам"
	default:
		return "неизвестная команда - обратитесь в техническую поддержку"
	}
}

func (c CommandType) IsValid() bool {
	switch c {
	case CommandStart, CommandScore, CommandAdd, CommandHelp:
		return true
	default:
		return false
	}
}

func (c CommandType) String() string {
	return string(c)
}
