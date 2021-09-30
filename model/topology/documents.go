package topology

type LSNode struct {
	ID			string
	Key			string
	RouterHash	string
	DomainID	int64
	RouterIP	string
	PeerHash	string
	PeerIP		string
	PeerASN		int32
	Timestamp	string
	IGPRouterID	string
	ASN			uint32
	MTID		[]*MultiTopologyIdentifier
	AreaID		string
	Protocol	string
	ProtocolID	uint8
	Name		string
	IsPrepolicy	bool
	IsAdjRIBIn	bool
}

type LSLink struct {
	ID					string
	Key					string
	RouterHash			string
	RouterIP			string
	DomainID			int64
	PeerHash			string
	PeerIP				string
	PeerASN				int32
	Timestamp			string
	IGPRouterID			string
	Protocol			string
	AreaID				string
	Nexthop				string
	MTID				*MultiTopologyIdentifier
	LocalLinkIP			string
	RemoteLinkIP		string
	IGPMetric			uint32
	RemoteNodeHash		string
	LocalNodeHash		string
	RemoteIGPRouterID	string
}

type LSPrefix struct {
	Key                  string
	ID                   string
	RouterHash           string
	RouterIP             string
	DomainID             int64
	PeerHash             string
	PeerIP               string
	PeerASN              int32
	Timestamp            string
	IGPRouterID          string
	Protocol             string
	AreaID               string
	Nexthop              string
	LocalNodeHash        string
	MTID                 *MultiTopologyIdentifier
	Prefix               string
	PrefixLen            int32
	PrefixMetric         uint32
	IsPrepolicy          bool
	IsAdjRIBIn           bool
}

type LSSRv6SID struct {
	Key                  string
	ID                   string
	RouterHash           string
	RouterIP             string
	DomainID             int64
	PeerHash             string
	PeerIP               string
	PeerASN              int32
	Timestamp            string
	IGPRouterID          string
	LocalNodeASN         uint32
	Protocol             string
	Nexthop              string
	LocalNodeHash        string
	MTID                 *MultiTopologyIdentifier
	IGPFlags             uint8
	IsPrepolicy          bool
	IsAdjRIBIn           bool
	SRv6SID              string
}

