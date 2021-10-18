package topology

import "encoding/json"

func (doc LSNode) MarshalBinary() ([]byte, error) {
	return json.Marshal(doc)
}

func (doc LSLink) MarshalBinary() ([]byte, error) {
	return json.Marshal(doc)
}

func (doc LSPrefix) MarshalBinary() ([]byte, error) {
	return json.Marshal(doc)
}

func (doc LSSRv6SID) MarshalBinary() ([]byte, error) {
	return json.Marshal(doc)
}

func (doc LSNodeEdge) MarshalBinary() ([]byte, error) {
	return json.Marshal(doc)
}