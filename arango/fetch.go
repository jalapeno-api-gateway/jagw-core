package arango

import (
	"context"
)

//
// ---> FETCH SINGLE <---
//

func FetchLsNode(ctx context.Context, key string) LSNode {
	cursor := queryArangoDbDatabase(ctx, "FOR d IN LSNode FILTER d._key == \"" + key + "\" RETURN d");
	var document LSNode
	readDocument(cursor.ReadDocument(ctx, &document))
	return document
}

func FetchLsLink(ctx context.Context, key string) LSLink {
	cursor := queryArangoDbDatabase(ctx, "FOR d IN LSLink FILTER d._key == \"" + key + "\" RETURN d");
	var document LSLink
	readDocument(cursor.ReadDocument(ctx, &document))
	return document
}

func FetchLsPrefix(ctx context.Context, key string) LSPrefix {
	cursor := queryArangoDbDatabase(ctx, "FOR d IN LSPrefix FILTER d._key == \"" + key + "\" RETURN d");
	var document LSPrefix
	readDocument(cursor.ReadDocument(ctx, &document))
	return document
}

func FetchLsSrv6Sid(ctx context.Context, key string) LSSRv6SID {
	cursor := queryArangoDbDatabase(ctx, "FOR d IN LSSRv6SID FILTER d._key == \"" + key + "\" RETURN d");
	var document LSSRv6SID
	readDocument(cursor.ReadDocument(ctx, &document))
	return document
}

func FetchLsNodeEdge(ctx context.Context, key string) LSNode_Edge {
	cursor := queryArangoDbDatabase(ctx, "FOR d IN LSNodeEdge FILTER d._key == \"" + key + "\" RETURN d");
	var document LSNode_Edge
	readDocument(cursor.ReadDocument(ctx, &document))
	return document
}

//
// ---> FETCH ALL <---
//

func FetchAllLsNodes(ctx context.Context) []LSNode {
	cursor := queryArangoDbDatabase(ctx, "FOR d IN LSNode RETURN d");
	var documents []LSNode
	for {
		var document LSNode
		if (!readDocument(cursor.ReadDocument(ctx, &document))) {
			break
		}
		documents = append(documents, document)
	}
	return documents
}

func FetchAllLsLinks(ctx context.Context) []LSLink {
	cursor := queryArangoDbDatabase(ctx, "FOR d IN LSLink RETURN d");
	var documents []LSLink
	for {
		var document LSLink
		if (!readDocument(cursor.ReadDocument(ctx, &document))) {
			break
		}
		documents = append(documents, document)
	}
	return documents
}

func FetchAllLsPrefixes(ctx context.Context) []LSPrefix {
	cursor := queryArangoDbDatabase(ctx, "FOR d IN LSPrefix RETURN d");
	var documents []LSPrefix
	for {
		var document LSPrefix
		if (!readDocument(cursor.ReadDocument(ctx, &document))) {
			break
		}
		documents = append(documents, document)
	}
	return documents
}

func FetchAllLsSrv6Sids(ctx context.Context) []LSSRv6SID {
	cursor := queryArangoDbDatabase(ctx, "FOR d IN LSSRv6SID RETURN d");
	var documents []LSSRv6SID
	for {
		var document LSSRv6SID
		if (!readDocument(cursor.ReadDocument(ctx, &document))) {
			break
		}
		documents = append(documents, document)
	}
	return documents
}

func FetchAllLsNodeEdges(ctx context.Context) []LSNode_Edge {
	cursor := queryArangoDbDatabase(ctx, "FOR d IN LSNode_Edge RETURN d");
	var documents []LSNode_Edge
	for {
		var document LSNode_Edge
		if (!readDocument(cursor.ReadDocument(ctx, &document))) {
			break
		}
		documents = append(documents, document)
	}
	return documents
}