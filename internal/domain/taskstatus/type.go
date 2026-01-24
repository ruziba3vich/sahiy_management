package taskstatus

type Type struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

const (
	Pending  = 0 // Not Started
	Started  = 1
	Finished = 10
	Frozen   = -1
	Simple   = 5
)

const (
	PendingName  = "Pending"
	StartedName  = "Started"
	FinishedName = "Finished"
	FrozenName   = "Frozen"
	SimpleName   = "Simple"
)

func GetAllTypes() []Type {
	return []Type{
		{Name: PendingName, Value: Pending},
		{Name: StartedName, Value: Started},
		{Name: FinishedName, Value: Finished},
		{Name: FrozenName, Value: Frozen},
		{Name: SimpleName, Value: Simple},
	}
}

func GetTypeByValue(value int) string {
	switch value {
	case Pending:
		return PendingName
	case Started:
		return StartedName
	case Finished:
		return FinishedName
	case Frozen:
		return FrozenName
	default:
		return SimpleName
	}
}

func IsValidType(value int) bool {
	return value == Pending || value == Started || value == Finished || value == Frozen || value == Simple
}
