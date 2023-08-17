package fbs

// scope represents a lexical scope in a flatbuffers file in which
// tables/structs/enums/unions can be declared. Currently only
// schemaScope is used.
type scope func(firstName, fullName string) (fqn string, desc Desc)

// schemaScope creates a scope for symbol resolving.
func schemaScope(fd *SchemaDesc, l *linker) scope {
	// Example: { "rpc.app.server", "rpc.app", "rpc", "namespace2" "" }
	prefixes := createPrefixes(fd.Namespaces)
	querySymbol := func(n string) (d Desc) {
		return l.findSymbol(fd, n)
	}
	return func(firstName, fullName string) (string, Desc) {
		for _, prefix := range prefixes {
			var n1, n string
			if prefix == "" {
				n1, n = fullName, fullName
			} else {
				// example: "rpc.app.server.namespace2" "rpc.app.namespace2"
				n1 = prefix + "." + firstName
				// example: "rpc.app.server.namespace2.MyFieldTypeName" "rpc.app.namespace2.MyFieldTypeName"
				// ... "namespace2.MyFieldTypeName"
				// to resolve fullName such as "server2.MyFieldTypeName2" whose fully qualified name is
				// "rpc.app.server2.MyFieldTypeName2", try until prefix is "rpc.app".
				n = prefix + "." + fullName
			}
			d := findSymbolRelative(n1, n, querySymbol)
			if d != nil {
				return n, d
			}
		}
		return "", nil
	}
}

// findSymbolRelative find symbol relatively.
// Example:
//
// symbol to be resolved is in the same namespace:
//   firstName: "rpc.app.server.MyTable.MyFieldTypeName"
//   fullName: "rpc.app.server.MyTable.MyFieldTypeName"
//
// symbol to be resolved is in another namespace:
//   firstName: "rpc.app.server.MyTable.namespace2"
//   fullName: "rpc.app.server.MyTable.namespace2.MyFieldTypeName"
func findSymbolRelative(firstName, fullName string, query func(name string) (d Desc)) Desc {
	d := query(firstName)
	if d == nil {
		return nil
	}
	if firstName == fullName {
		return d
	}
	d = query(fullName)
	return d
}

// createPrefixes creates a list of prefixes out of
// a list of namespaces in reverse order. Example:
//
// Input: ["", "namespace2", "rpc.app.server"]
// Output:
// []string { "rpc.app.server", "rpc.app", "rpc", "namespace2", ""}
func createPrefixes(namespaces []string) []string {
	if namespaces == nil {
		return []string{""}
	}
	existed := map[string]struct{}{}
	var prefixes []string
	// fill up the slice backwards.
	for i := len(namespaces) - 1; i >= 0; i-- {
		if _, ok := existed[namespaces[i]]; !ok {
			existed[namespaces[i]] = struct{}{}
			prefixes = append(prefixes, namespaces[i])
		}
		prefixes = appendUniquePrefixes(prefixes, namespaces[i], existed)
	}
	// fd.Namespaces contain empty namespace string,
	// prefixes slice will include that empty string in the above steps.
	return prefixes
}

// appendUniquePrefixes creates unique prefixes from a single namespace.
func appendUniquePrefixes(prefixes []string, namespace string, existed map[string]struct{}) []string {
	for j := len(namespace) - 1; j >= 0; j-- {
		if namespace[j] != '.' {
			continue
		}
		if _, ok := existed[namespace[:j]]; !ok {
			existed[namespace[:j]] = struct{}{}
			prefixes = append(prefixes, namespace[:j])
		}
	}
	return prefixes
}
