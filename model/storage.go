package model

var (
	storagies = map[string]Storage{}
)

func Register(name string, s Storage) {

	_, ok := storagies[name]

	if ok {
		panic(`storage already exists: ` + name)
	}

	storagies[name] = s
}

type Storage interface {
}
