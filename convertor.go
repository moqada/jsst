package main

import (
	"fmt"
	"go/format"
	"io"
	"sort"

	"github.com/achiku/varfmt"
	"github.com/lestrrat/go-jshschema"
	"github.com/lestrrat/go-jsschema"
)

const (
	// DefaultPackage is default package name
	DefaultPackage = "main"
)

// Convertor convert JSON Schema to Struct
type Convertor struct {
	schema   *schema.Schema
	Package  string
	Resolved StructMap
}

// StructMap is map of Struct
type StructMap map[string]*Struct

// SortedKeys returns keys
func (sm StructMap) SortedKeys() []string {
	var keys []string
	for key := range sm {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// New Convertor
func New(s *schema.Schema) *Convertor {
	return &Convertor{
		schema:   s,
		Package:  DefaultPackage,
		Resolved: make(StructMap),
	}
}

// Read []byte, returns New Convertor
func Read(in io.Reader) (*Convertor, error) {
	s, err := schema.Read(in)
	if err != nil {
		return nil, err
	}
	return New(s), nil
}

// ReadFile from filepath, returns New Convertor
func ReadFile(file string) (*Convertor, error) {
	s, err := schema.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return New(s), nil
}

func (con *Convertor) Write(out io.Writer) error {
	head := "package " + con.Package + "\n"
	imports := make(map[string]string)
	source := ""
	for _, key := range con.Resolved.SortedKeys() {
		s := con.Resolved[key]
		for _, p := range s.Packages {
			imports[p] = p
		}
		source += structToString(s, &con.Resolved, true)
	}
	if len(imports) > 0 {
		head += "import (\n"
		for _, p := range imports {
			head += fmt.Sprintf("\"%s\"", p)
		}
		head += ")\n"
	}
	source = head + source
	b, err := format.Source([]byte(source))
	if err != nil {
		return err
	}
	if _, err := fmt.Fprint(out, string(b)); err != nil {
		return err
	}
	return nil
}

// Extract Schema to StructMap
func (con *Convertor) Extract() error {
	for name, sc := range con.schema.Properties {
		_, err := extractProps(name, sc, &con.Resolved)
		if err != nil {
			return err
		}
		links := sc.Extras["links"]
		if links == nil {
			continue
		}
		hsc := hschema.New()
		hsc.Extract(sc.Extras)
		for _, link := range hsc.Links {
			if link.TargetSchema != nil {
				// TODO: Extract TargetSchema
				continue
			}
		}
	}
	return nil
}

// SetPackage set package naem
func (con *Convertor) SetPackage(name string) {
	con.Package = name
}

// extractProps extract Schema.Properties to Struct
func extractProps(name string, sc *schema.Schema, resolved *StructMap) (*Struct, error) {
	var err error
	var ref string
	if !sc.IsResolved() {
		ref = sc.Reference
		sc, err = sc.Resolve(nil)
		if err != nil {
			return nil, err
		}
	}
	t, pkg, err := getPropertyType(sc)
	if err != nil {
		return nil, err
	}
	st := Struct{Name: name, Type: t, Ref: ref}
	if pkg != "" {
		st.AddPkg(pkg)
	}
	switch t {
	case "object":
		for k, v := range sc.Properties {
			s, err := extractProps(k, v, resolved)
			if err != nil {
				return nil, err
			}
			st.Properties = append(st.Properties, *s)
			for _, p := range s.Packages {
				st.AddPkg(p)
			}
		}
		// TODO: ref == "" -> uuid?
		if ref != "" {
			(*resolved)[ref] = &st
		}
	case "array":
		if len(sc.Items.Schemas) != 1 {
			// TODO: Support multiple types
			return nil, fmt.Errorf("Multiple Items doesnot support.")
		}
		s, err := extractProps(name, sc.Items.Schemas[0], resolved)
		if err != nil {
			return nil, err
		}
		st.Properties = append(st.Properties, *s)
		for _, p := range s.Packages {
			st.AddPkg(p)
		}
		// TODO: ref == "" -> uuid?
		if ref != "" {
			(*resolved)[ref] = &st
		}
	}
	return &st, nil
}

// getPropertyType convert Schema into type of Go
func getPropertyType(s *schema.Schema) (string, string, error) {
	pkg := ""
	if len(s.Type) != 1 {
		// TODO: Support multiple types
		// TODO: Support Nullable
		return "", pkg, fmt.Errorf("Multiple Types doesnot Support.")
	}
	t := s.Type[0].String()
	switch t {
	case "number":
		t = "float64"
	case "boolean":
		t = "bool"
	case "string":
		if s.Format == "date-time" {
			return "time.Time", "time", nil
		}
	}
	return t, pkg, nil
}

// Convert Struct.Property into string
func propToString(name, goType string) string {
	return fmt.Sprintf("%s %s `json:\"%s\"`\n", varfmt.PublicVarName(name), goType, name)
}

// Convert Struct into string
func structToString(st *Struct, resolved *StructMap, root bool) string {
	typePre := ""
	typeDef := fmt.Sprintf("type %s ", varfmt.PublicVarName(st.Key()))
	if st.Type == "array" {
		st = &st.Properties[0]
		typePre = "[]"
	}
	if !root && st.Ref != "" {
		if res, ok := (*resolved)[st.Ref]; ok {
			return propToString(st.Name, typePre+varfmt.PublicVarName(res.Key()))
		}
	}
	t := st.Type
	if st.Type == "object" {
		t = "struct {\n"
		sort.Sort(st.Properties)
		for _, prop := range st.Properties {
			t += structToString(&prop, resolved, false)
		}
		t += "}"
	}
	t = typePre + t
	if root {
		return typeDef + t + "\n"
	}
	return propToString(st.Name, t)
}
