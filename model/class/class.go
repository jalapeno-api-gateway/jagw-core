package class

type Class string

const (
	LsNode Class = "LSNode"
	LsLink Class = "LSLink"
	LsPrefix Class = "LSPrefix"
	LsSrv6Sid Class = "LSSRv6SID"
	LsNodeEdge Class = "LSNode_Edge"
	LsNodeCoordinates Class = "LSNode_Coordinates"
	PhysicalInterface Class = "PhysicalInterface"
	LoopbackInterface Class = "LoopbackInterface"
)