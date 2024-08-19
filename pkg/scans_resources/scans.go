package scans_resources

type ScansResources struct{}

type Scan struct {
	ID string
}

func New() *ScansResources {
	return &ScansResources{}
}

func (sr *ScansResources) GetScans() []*Scan {
	var scans []*Scan

	scans = []*Scan{
		{ID: "1"},
		{ID: "2"},
		{ID: "3"},
	}
	return scans
}
