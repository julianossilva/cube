package cube

type Test interface {
	Fail(message string)
}

type Cube interface {
	AddProvider(provider any) error
	AddRunner(runner any) error
	AddTest(test any)
}
