package betypes

type BotMessage struct {
	Message struct {
		Message_id int
		From struct {
			User
		}
	}
}