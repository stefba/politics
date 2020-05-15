package server

import (
	"politics/go/entry"
	"politics/go/entry/types/set"
	"politics/go/entry/types/tree"
	"flag"
	p "path/filepath"
)

type Server struct {
	Paths  *paths
	Tree   *tree.Tree
	Files  entry.Entries
	Recent entry.Entries
	Flags  *flags
}

type paths struct {
	Root, Data string
}

type flags struct {
	Reload bool
}

func NewServer() *Server {
	
	path := flag.String("path", ".", "root path of this app")
	reload := flag.Bool("reload", false, "reload files on every request")

	flag.Parse()

	s := &Server{}

	s.Paths = &paths{
		Root: *path,
		Data: p.Clean(*path + "/data"),
	}
	
	s.Flags = &flags{
		Reload: *reload,
	}

	return s
}

func (s *Server) Load() error {
	return s.LoadEntries()
}

func (s *Server) LoadEntries() error {
	t, err := tree.ReadTree(s.Paths.Data, nil)
	if err != nil {
		return err
	}

	s.Tree = t
	s.Recent = t.TraverseEntries().Desc()
	s.Files = BlobList(s.Recent)

	return nil
}

func BlobList(es entry.Entries) entry.Entries {
	blobs := entry.Entries{}
	for _, e := range es {
		if e.IsBlob() {
			blobs = append(blobs, e)
			continue
		}
		s, ok := e.(*set.Set)
		if ok {
			blobs = append(blobs, BlobList(s.Entries())...)
		}
	}
	return blobs
}
