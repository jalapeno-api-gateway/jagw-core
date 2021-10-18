package arango

import (
	"context"
)

//
// ---> FETCH SINGLE <---
//

func FetchLSNode(ctx context.Context, key string) LSNode {
	cursor := queryArangoDbDatabase(ctx, "FOR d IN LSNode FILTER d._key == \"" + key + "\" RETURN d");
	var document LSNode
	readDocument(cursor.ReadDocument(ctx, &document))
	return document
}

func FetchLSLink(ctx context.Context, key string) LSLink {
	cursor := queryArangoDbDatabase(ctx, "FOR d IN LSLink FILTER d._key == \"" + key + "\" RETURN d");
	var document LSLink
	readDocument(cursor.ReadDocument(ctx, &document))
	return document
}

func FetchLSPrefix(ctx context.Context, key string) LSPrefix {
	cursor := queryArangoDbDatabase(ctx, "FOR d IN LSPrefix FILTER d._key == \"" + key + "\" RETURN d");
	var document LSPrefix
	readDocument(cursor.ReadDocument(ctx, &document))
	return document
}

func FetchLSSRv6SID(ctx context.Context, key string) LSSRv6SID {
	cursor := queryArangoDbDatabase(ctx, "FOR d IN LSSRv6SID FILTER d._key == \"" + key + "\" RETURN d");
	var document LSSRv6SID
	readDocument(cursor.ReadDocument(ctx, &document))
	return document
}

func FetchLSNodeEdge(ctx context.Context, key string) LSNodeEdge {
	cursor := queryArangoDbDatabase(ctx, "FOR d IN LSNodeEdge FILTER d._key == \"" + key + "\" RETURN d");
	var document LSNodeEdge
	readDocument(cursor.ReadDocument(ctx, &document))
	return document
}

//
// ---> FETCH ALL <---
//

func FetchAllLSNodes(ctx context.Context) []LSNode {
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

func FetchAllLSLinks(ctx context.Context) []LSLink {
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

func FetchAllLSPrefixes(ctx context.Context) []LSPrefix {
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

func FetchAllLSSRv6SIDs(ctx context.Context) []LSSRv6SID {
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

func FetchAllLSNodeEdges(ctx context.Context) []LSNodeEdge {
	cursor := queryArangoDbDatabase(ctx, "FOR d IN LSNodeEdge RETURN d");
	var documents []LSNodeEdge
	for {
		var document LSNodeEdge
		if (!readDocument(cursor.ReadDocument(ctx, &document))) {
			break
		}
		documents = append(documents, document)
	}
	return documents
}