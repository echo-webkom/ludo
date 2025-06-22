package git

type Branch struct {
	Name     string
	Remote   string
	Commits  int
	UpToDate bool
}

func CreateBranch(name string) error {
	return nil
}

func PushBranch(name string) error {
	return nil
}

func GetBranch(name string) Branch {
	return Branch{}
}
