package section

import "context"

type Repository interface {
	CreateSection(ctx context.Context, sec *Section) (*Section, error)
	UpdateSection(ctx context.Context, sec *Section) (*Section, error)
	DeleteSection(ctx context.Context, id int64) error
	GetSectionByID(ctx context.Context, id int64) (*Section, error)
	GetAllSections(ctx context.Context) ([]*Section, error)
}
