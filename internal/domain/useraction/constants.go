package useraction

const (
	StatusOnTime     = 1
	StatusLate       = 2
	StatusBeforeTime = 3
)

var StatusLabels = map[int]string{
	StatusOnTime:     "O'z vaqtida",
	StatusLate:       "Kechikkan",
	StatusBeforeTime: "Vaqtidan oldin",
}
