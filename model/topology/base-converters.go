package topology

import "github.com/jalapeno-api-gateway/jagw-core/arango"

func convertMTIDSlice(docs []*arango.MultiTopologyIdentifier) []*MultiTopologyIdentifier {
	mtids := []*MultiTopologyIdentifier{}
	for _, doc := range docs {
	   mtids = append(mtids, convertMTID(doc))
	}
	return mtids
}

func convertMTID(doc *arango.MultiTopologyIdentifier) *MultiTopologyIdentifier {
	return &MultiTopologyIdentifier{
		OFlag: doc.OFlag,
		AFlag: doc.AFlag,
		MTID: doc.MTID,
	}
}