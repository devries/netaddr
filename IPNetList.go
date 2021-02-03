package netaddr

type IPNetList struct {
	Net4List IPv4NetList
	Net6List IPv6NetList
}

func NewIPNetList(networks []string) (IPNetList, error) {
	net4 := make(IPv4NetList, 0)
	net6 := make(IPv6NetList, 0)

	for _, addr := range networks {
		net, err := ParseIPNet(addr)
		if err != nil {
			return IPNetList{net4, net6}, err
		}

		switch nets := net.(type) {
		case *IPv4Net:
			net4 = append(net4, nets)
		case *IPv6Net:
			net6 = append(net6, nets)
		}
	}

	return IPNetList{net4, net6}, nil
}

func (list IPNetList) Summ() IPNetList {
	return IPNetList{list.Net4List.Summ(), list.Net6List.Summ()}
}
