package branch

import "time"

type Branch struct {
	ID        int64
	Name      string
	Lat       float64
	Long      float64
	Radius    int
	Status    int
	CreatedAt int64
	UpdatedAt int64
}

func NewBranch(name string, lat, long float64, radius, status int) *Branch {
	now := time.Now().Unix()
	return &Branch{
		Name:      name,
		Lat:       lat,
		Long:      long,
		Radius:    radius,
		Status:    status,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (b *Branch) Update(name string, lat, long float64, radius, status int) {
	b.Name = name
	b.Lat = lat
	b.Long = long
	b.Radius = radius
	b.Status = status
	b.UpdatedAt = time.Now().Unix()
}
