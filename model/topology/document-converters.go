package topology

import "github.com/jalapeno-api-gateway/jagw-core/arango"

func ConvertLSNode(doc arango.LSNode) LSNode {
	return LSNode{
		ID: doc.ID,
		Key: doc.Key,
		RouterHash: doc.RouterHash,
		DomainID: doc.DomainID,
		RouterIP: doc.RouterIP,
		PeerHash: doc.PeerHash,
		PeerIP: doc.PeerIP,
		PeerASN: doc.PeerASN,
		Timestamp: doc.Timestamp,
		IGPRouterID: doc.IGPRouterID,
		ASN: doc.ASN,
		MTID: convertMTIDSlice(doc.MTID),
		AreaID: doc.AreaID,
		Protocol: doc.Protocol,
		ProtocolID: doc.ProtocolID,
		Name: doc.Name,
		IsPrepolicy: doc.IsPrepolicy,
		IsAdjRIBIn: doc.IsAdjRIBIn,
	}
}

func ConvertLSLink(doc arango.LSLink) LSLink {
	return LSLink{
		ID: doc.ID,
		Key: doc.Key,
		RouterHash: doc.RouterHash,
		RouterIP: doc.RouterIP,
		DomainID: doc.DomainID,
		PeerHash: doc.PeerHash,
		PeerIP: doc.PeerIP,
		PeerASN: doc.PeerASN,
		Timestamp: doc.Timestamp,
		IGPRouterID: doc.IGPRouterID,
		Protocol: doc.Protocol,
		AreaID: doc.AreaID,
		Nexthop: doc.Nexthop,
		MTID: convertMTID(doc.MTID),
		LocalLinkIP: doc.LocalLinkIP,
		RemoteLinkIP: doc.RemoteLinkIP,
		IGPMetric: doc.IGPMetric,
		RemoteNodeHash: doc.RemoteNodeHash,
		LocalNodeHash: doc.LocalNodeHash,
		RemoteIGPRouterID: doc.RemoteIGPRouterID,
	}
}

func ConvertLSPrefix(doc arango.LSPrefix) LSPrefix {
	return LSPrefix{
		Key: doc.Key,
		ID: doc.ID,
		RouterHash: doc.RouterHash,
		RouterIP: doc.RouterIP,
		DomainID: doc.DomainID,
		PeerHash: doc.PeerHash,
		PeerIP: doc.PeerIP,
		PeerASN: doc.PeerASN,
		Timestamp: doc.Timestamp,
		IGPRouterID: doc.IGPRouterID,
		Protocol: doc.Protocol,
		AreaID: doc.AreaID,
		Nexthop: doc.Nexthop,
		LocalNodeHash: doc.LocalNodeHash,
		MTID: convertMTID(doc.MTID),
		Prefix: doc.Prefix,
		PrefixLen: doc.PrefixLen,
		PrefixMetric: doc.PrefixMetric,
		IsPrepolicy: doc.IsPrepolicy,
		IsAdjRIBIn: doc.IsAdjRIBIn,
	}
}

func ConvertLSSRv6SID(doc arango.LSSRv6SID) LSSRv6SID {
	return LSSRv6SID{
		Key: doc.Key,
		ID: doc.ID,
		RouterHash: doc.RouterHash,
		RouterIP: doc.RouterIP,
		DomainID: doc.DomainID,
		PeerHash: doc.PeerHash,
		PeerIP: doc.PeerIP,
		PeerASN: doc.PeerASN,
		Timestamp: doc.Timestamp,
		IGPRouterID: doc.IGPRouterID,
		LocalNodeASN: doc.LocalNodeASN,
		Protocol: doc.Protocol,
		Nexthop: doc.Nexthop,
		LocalNodeHash: doc.LocalNodeHash,
		MTID: convertMTID(doc.MTID),
		IGPFlags: doc.IGPFlags,
		IsPrepolicy: doc.IsPrepolicy,
		IsAdjRIBIn: doc.IsAdjRIBIn,
		SRv6SID: doc.SRv6SID,
	}
}