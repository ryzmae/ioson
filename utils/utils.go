package handlers

func HandlerConfig(ttl int) error {
	setTTLConfig(ttl)
	return nil
}

func HandlerPing(website string) {
	ping(website)
}