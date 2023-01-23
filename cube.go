package cube

type Test interface {
	Fail(message string)
}

type Cube interface {
	AddProvider(provider any) error
	ReplaceProvider(provider any) error
	Run(runner any) (error, error)
	AddRunner(runner any) error
	AddTest(test any)
}
