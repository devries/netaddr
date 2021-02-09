package netaddr

import (
	"sort"
)

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

// Be sure to Summ first, or it will not work
func (list IPNetList) Contains(other IPNet) bool {
	switch nets := other.(type) {
	case *IPv4Net:
		return list.containsIPv4(nets)
	case *IPv6Net:
		return list.containsIPv6(nets)
	}
	return false
}

// The list must be sorted using Summ
func (list IPNetList) containsIPv4(other *IPv4Net) bool {
	listLength := list.Net4List.Len()
	m := sort.Search(listLength, func(i int) bool {
		cmp, _ := list.Net4List[i].Cmp(other)
		if cmp == 1 {
			return true
		}
		related, how := list.Net4List[i].Rel(other)
		if related && how >= 0 {
			return true
		}
		return false
	})

	if m == listLength {
		return false
	}

	related, how := list.Net4List[m].Rel(other)
	if related && how >= 0 {
		return true
	}

	return false
}

// The list must be sorted using Summ
func (list IPNetList) containsIPv6(other *IPv6Net) bool {
	listLength := list.Net6List.Len()
	m := sort.Search(listLength, func(i int) bool {
		cmp, _ := list.Net6List[i].Cmp(other)
		if cmp == 1 {
			return true
		}
		related, how := list.Net6List[i].Rel(other)
		if related && how >= 0 {
			return true
		}
		return false
	})

	if m == listLength {
		return false
	}

	related, how := list.Net6List[m].Rel(other)
	if related && how >= 0 {
		return true
	}

	return false
}
