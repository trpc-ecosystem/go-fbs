package fbs

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

//go:generate goyacc -o fbs.y.go -p fbs fbs.y

// Parser parses .fbs source into descriptors.
type Parser struct {
	accessor  FileAccessor
	filenames []string
	results   *parseResults
	handler   *errorHandler
	// IncludePaths stores paths used to search for dependencies
	// which is referred to by the include statement in .fbs
	// source files. If no include paths are provided then the
	// current directory (.) is assumed to be the only include
	// path.
	IncludePaths []string
	// Recursive decides whether to parse the file recursively (parse includes).
	// default is true.
	Recursive bool
}

// NewParser creates a parser.
func NewParser(includes ...string) *Parser {
	return &Parser{
		IncludePaths: includes,
		Recursive:    true, // Default is true.
	}
}

// SetRecursive configures whether parser parse the file recursively following the includes.
func (p *Parser) SetRecursive(recursive bool) {
	p.Recursive = recursive
}

// ParseFiles parse a list of .fbs files into descriptors.
func (p *Parser) ParseFiles(filenames ...string) ([]*SchemaDesc, error) {
	paths := extendPaths(p.IncludePaths, filenames)
	p.accessor = getAccessor(paths)
	p.filenames = filenames
	fbs := map[string]*parseResult{}
	p.results = &parseResults{
		resultsByFilename:   fbs,
		recursive:           p.Recursive,
		createDescriptorFbs: true,
	}
	p.handler = newErrorHandler()
	// Step1: source files => descriptors.
	// including lexing and parsing.
	if err := p.parseFiles(); err != nil {
		return nil, err
	}
	// Note: if recursive is set, here results will not only contain the ones specified in
	// `filenames`, but also the files that are included by them.
	// Step2: link all parsed descriptors.
	l := newLinker(p.results, p.handler)
	linkedFbs, err := l.linkFiles()
	if err != nil {
		return nil, err
	}
	fds := make([]*SchemaDesc, len(filenames))
	for i, name := range filenames {
		fd := linkedFbs[name]
		fds[i] = fd
	}
	return fds, nil
}

// parseFiles iterates all the given files to generate parse results.
func (p *Parser) parseFiles() error {
	for _, name := range p.filenames {
		p.parseFile(name)
		if p.handler.getError() != nil {
			return p.handler.err
		}
	}
	return nil
}

// parseFile parses a single file. If Parser.Recursive is set, it will iterate all
// its includes to parse recursively.
func (p *Parser) parseFile(filename string) {
	result, ok := p.parse(filename)
	if !ok {
		return
	}
	if p.Recursive {
		fd := result.fd
		schema := result.getSchemaNode(fd)
		// Even if the file is empty, schema would be nil.
		for _, incl := range schema.Includes {
			p.parseFile(incl.Name.Val)
			if p.handler.getError() != nil {
				return
			}
			result := p.results.resultsByFilename[incl.Name.Val]
			fd.Dependencies = append(fd.Dependencies, result.fd)
		}
	}
}

// parse does the actual parsing.
func (p *Parser) parse(filename string) (*parseResult, bool) {
	if p.results.has(filename) {
		return nil, false
	}
	in, err := p.accessor(filename)
	if err != nil {
		_ = p.handler.handleError(err)
		return nil, false
	}
	l := newLexer(in, filename, p.handler)
	fbsParse(l)
	result := newParseResult(filename, l.res, l.handler, p.results.createDescriptorFbs)
	_ = in.Close()
	p.results.add(filename, result)
	if p.handler.getError() != nil {
		return nil, false
	}
	return result, true
}

// extendPaths add necessary paths to current include paths, for example:
// empty string (representing current path), directories containing the .fbs files.
func extendPaths(paths []string, filenames []string) []string {
	// add empty string which represents current directory.
	paths = append(paths, "")
	existed := map[string]struct{}{}
	for _, path := range paths {
		existed[path] = struct{}{}
	}
	for _, filename := range filenames {
		separatorIdx := strings.LastIndexByte(filename, '/')
		if separatorIdx == -1 {
			continue
		}
		prefix := filename[:separatorIdx+1] // includes path separator
		if _, ok := existed[prefix]; ok {
			continue
		}
		paths = append(paths, prefix)
		existed[prefix] = struct{}{}
	}
	return paths
}

// getAccessor will create an file accessor that will try every include path as prefix
// to correctly open a file.
func getAccessor(paths []string) func(name string) (io.ReadCloser, error) {
	accessor := func(name string) (io.ReadCloser, error) {
		return os.Open(name)
	}
	if len(paths) > 0 {
		acc := accessor
		accessor = func(name string) (io.ReadCloser, error) {
			for _, path := range paths {
				f, err := acc(filepath.Join(path, name))
				if err != nil {
					continue
				}
				return f, nil
			}
			return nil, fmt.Errorf("cannot find file %v", name)
		}
	}
	return accessor
}

// FileAccessor abstracts how a file is opened.
type FileAccessor func(filename string) (io.ReadCloser, error)
