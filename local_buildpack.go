package freezer

import "fmt"

type LocalBuildpack struct {
	Path        string
	Name        string
	UncachedKey string
	CachedKey   string
	Offline     bool
	Version     string
}

func NewLocalBuildpack(path, name string) LocalBuildpack {
	return LocalBuildpack{
		Path:        path,
		Name:        name,
		UncachedKey: name,
		CachedKey:   fmt.Sprintf("%s:cached", name),
	}
}

func (l LocalBuildpack) WithOffline(offline bool) LocalBuildpack {
	l.Offline = offline
	return l
}

func (l LocalBuildpack) WithVersion(version string) LocalBuildpack {
	l.Version = version
	return l
}
