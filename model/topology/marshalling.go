package topology

import "encoding/json"

func (node LSNode) MarshalBinary() ([]byte, error) {
	return json.Marshal(node)
}

func (link LSLink) MarshalBinary() ([]byte, error) {
	return json.Marshal(link)
}

func (node LSPrefix) MarshalBinary() ([]byte, error) {
	return json.Marshal(node)
}

func (link LSSRv6SID) MarshalBinary() ([]byte, error) {
	return json.Marshal(link)
}