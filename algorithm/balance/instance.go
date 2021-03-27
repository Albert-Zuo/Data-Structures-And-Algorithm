package balance

type Instance struct {
	host string
	port int

	extra * instextra
}

type instextra struct {
	zone string
	// ...
}

// new default instance
func NewInstance(host string, port int) *Instance {
	return &Instance{
		host: host,
		port: port,
	}
}

func (i *Instance) Host() string {
	return i.host
}

func (i *Instance) Port() int {
	return i.port
}