package configparser

type Provider interface {
	Get() (*Parser, error)
}

func Default() Provider {
	return NewFile()
}
