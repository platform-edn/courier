package registry

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

type Node struct {
	Id                  string
	Address             string
	Port                string
	SubscribedSubjects  []string
	BroadcastedSubjects []string
}

func NewNode(id string, address string, port string, subscribed []string, broadcasted []string) *Node {
	n := Node{
		Id:                  id,
		Address:             address,
		Port:                port,
		SubscribedSubjects:  subscribed,
		BroadcastedSubjects: broadcasted,
	}

	return &n
}

func RemovePointers(nodes []*Node) []Node {
	updated := []Node{}

	for _, n := range nodes {
		updated = append(updated, *n)
	}

	return updated
}

// TestNodeOptions that can be passed into CreateTestNodes
type TestNodeOptions struct {
	SubscribedSubjects  []string
	BroadcastedSubjects []string
	Id                  string
	Port                string
}

// CreateTestNodes creates a quantity of randomized nodes based on the options passed in
// TODO: clean this up with functional parameters
func CreateTestNodes(count int, options *TestNodeOptions) []*Node {
	nodes := []*Node{}
	var broadSubjects []string
	var subSubjects []string

	if len(options.SubscribedSubjects) == 0 {
		subSubjects = []string{"sub", "sub1", "sub2"}
	} else {
		subSubjects = options.SubscribedSubjects
	}
	if len(options.BroadcastedSubjects) == 0 {
		broadSubjects = []string{"broad", "broad1"}
	} else {
		broadSubjects = options.BroadcastedSubjects
	}

	for i := 0; i < count; i++ {
		ip := fmt.Sprintf("%v.%v.%v.%v", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
		subcount := (rand.Intn(len(subSubjects)) + 1)
		broadcount := rand.Intn(len(broadSubjects) + 1)
		var subs []string
		var broads []string

		for i := 0; i < subcount; i++ {
			subs = append(subs, subSubjects[i])
		}

		for i := 0; i < broadcount; i++ {
			broads = append(broads, broadSubjects[i])
		}

		var id string
		if options.Id != "" {
			id = options.Id
		} else {
			id = uuid.NewString()
		}
		var port string
		if options.Port != "" {
			port = options.Port
		} else {
			port = fmt.Sprint(rand.Intn(9999-1000) + 1000)
		}
		n := NewNode(id, ip, port, subs, broads)
		nodes = append(nodes, n)
	}

	return nodes
}
