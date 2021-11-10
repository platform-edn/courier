package courier

import (
	"fmt"

	"github.com/platform-edn/courier/lock"
	"github.com/platform-edn/courier/node"
)

type ResponseMapper interface {
	Push(node.ResponseInfo)
	Pop(string) (string, error)
}

type responseMap struct {
	responses map[string]string
	lock      *lock.TicketLock
}

func newResponseMap() *responseMap {
	r := responseMap{
		responses: make(map[string]string),
		lock:      lock.NewTicketLock(),
	}

	return &r
}

func (r *responseMap) Push(info node.ResponseInfo) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.responses[info.MessageId] = info.NodeId
}

func (r *responseMap) Pop(messageId string) (string, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	nodeId, ok := r.responses[messageId]
	if !ok {
		return "", fmt.Errorf("response does not exist with %s as an id", messageId)
	}

	delete(r.responses, messageId)

	return nodeId, nil
}

func (r *responseMap) Length() int {
	r.lock.Lock()
	defer r.lock.Unlock()

	return len(r.responses)
}