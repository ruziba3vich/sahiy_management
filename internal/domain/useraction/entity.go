package useraction

type UserAction struct {
	ID            int64
	UserID        int64
	VisitBranchID int64
	LeaveBranchID *int64
	ComeStatus    int
	OutStatus     *int
	StartedAt     int64
	FinishedAt    *int64
}

func NewUserAction(userID, visitBranchID int64, leaveBranchID *int64, comeStatus int, outStatus *int, startedAt int64, finishedAt *int64) *UserAction {
	return &UserAction{
		UserID:        userID,
		VisitBranchID: visitBranchID,
		LeaveBranchID: leaveBranchID,
		ComeStatus:    comeStatus,
		OutStatus:     outStatus,
		StartedAt:     startedAt,
		FinishedAt:    finishedAt,
	}
}

func (u *UserAction) Update(userID, visitBranchID int64, leaveBranchID *int64, comeStatus int, outStatus *int, startedAt int64, finishedAt *int64) {
	u.UserID = userID
	u.VisitBranchID = visitBranchID
	u.LeaveBranchID = leaveBranchID
	u.ComeStatus = comeStatus
	u.OutStatus = outStatus
	u.StartedAt = startedAt
	u.FinishedAt = finishedAt
}
