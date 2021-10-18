package property

var AllLSNodeProperties = []string{
	Key,
	ID,
	RouterHash,
	DomainID,
	RouterIP,
	PeerHash,
	PeerIP,
	PeerASN,
	Timestamp,
	IGPRouterID,
	ASN,
	MTID,
	AreaID,
	Protocol,
	ProtocolID,
	Name,
	IsPrepolicy,
	IsAdjRIBIn,
}

var AllLSLinkProperties = []string{
	Key,
	ID,
	RouterHash,
	RouterIP,
	DomainID,
	PeerHash,
	PeerIP,
	PeerASN,
	Timestamp,
	IGPRouterID,
	Protocol,
	AreaID,
	Nexthop,
	MTID,
	LocalLinkIP,
	RemoteLinkIP,
	IGPMetric,
	RemoteNodeHash,
	LocalNodeHash,
	RemoteIGPRouterID,
}

var AllLSPrefixProperties = []string{
	Key,
	ID,
	RouterHash,
	RouterIP,
	DomainID,
	PeerHash,
	PeerIP,
	PeerASN,
	Timestamp,
	IGPRouterID,
	Protocol,
	AreaID,
	Nexthop,
	LocalNodeHash,
	MTID,
	Prefix,
	PrefixLen,
	PrefixMetric,
	IsPrepolicy,
	IsAdjRIBIn,
}

var AllLSSRv6SIDProperties = []string{
	Key,
	ID,
	RouterHash,
	RouterIP,
	DomainID,
	PeerHash,
	PeerIP,
	PeerASN,
	Timestamp,
	IGPRouterID,
	LocalNodeASN,
	Protocol,
	Nexthop,
	LocalNodeHash,
	MTID,
	IGPFlags,
	IsPrepolicy,
	IsAdjRIBIn,
	SRv6SID,
}

var AllPhysicalInterfaceProperties = []string{
	DataRate,
	PacketsSent,
	PacketsReceived,
}

var AllLoopbackInterfaceProperties = []string{
	State,
	LastStateTransitionTime,
}

var AllLSNodeEdgeProperties = []string{
	Key,
	ID,
	From,
	To,
	Link,
}