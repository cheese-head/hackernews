package models

const (
	TypeAsk      = "ask"
	TypeComment  = "comment"
	TypeJob      = "job"
	TypePoll     = "poll"
	TypepPollopt = "pollopt"
	TypeStory    = "story"
)

func IsAsk(s string) bool {
	return s == TypeAsk
}

func IsComment(s string) bool {
	return s == TypeComment
}

func IsJob(s string) bool {
	return s == TypeJob
}

func IsPoll(s string) bool {
	return s == TypePoll
}

func IsPollopt(s string) bool {
	return s == TypepPollopt
}

func IsStory(s string) bool {
	return s == TypeStory
}
