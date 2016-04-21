package main

import "strings"

// Struct is Go struct
type Struct struct {
	Name       string
	Nullable   bool
	Type       string
	Ref        string
	Packages   map[string]string
	Properties PropertyList
	Required   bool
	// FIXME: too dirty
	Link bool
}

// PropertyList is array of Struct
type PropertyList []Struct

// Key is name of Struct
func (st *Struct) Key() string {
	name := st.Name
	if !st.Link && st.Ref != "" {
		slice := strings.Split(st.Ref, "/")
		name = slice[len(slice)-1]
	}
	return name
}

// AddPkg is adding package
func (st *Struct) AddPkg(name string) {
	if st.Packages == nil {
		st.Packages = make(map[string]string)
	}
	st.Packages[name] = name
}

func (pl PropertyList) Len() int {
	return len(pl)
}

func (pl PropertyList) Less(i, j int) bool {
	return pl[i].Name < pl[j].Name
}

func (pl PropertyList) Swap(i, j int) {
	pl[i], pl[j] = pl[j], pl[i]
}
