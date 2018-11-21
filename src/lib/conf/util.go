package conf

func confToBrokerMap(bs []Broker) map[string]Broker {
	res := make(map[string]Broker, len(bs))
	for _, v := range bs {
		res[v.Name] = v
	}
	return res
}
