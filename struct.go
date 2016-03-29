package main

import "strings"

// Struct is Go struct
type Struct struct {
	Name       string
	Nullable   bool
	Type       string
	Ref        string
	Properties PropertyList
}

// PropertyList is array of Struct
type PropertyList []Struct

// Key is name of Struct
func (st *Struct) Key() string {
	name := st.Name
	if st.Ref != "" {
		slice := strings.Split(st.Ref, "/")
		name = slice[len(slice)-1]
	}
	return name
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
