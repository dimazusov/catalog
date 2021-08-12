// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package building

import (
	organization "catalog/internal/domain/organization"
	organization_phone "catalog/internal/domain/organization_phone"
	coords "catalog/internal/pkg/coords"
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson1cf24d94DecodeCatalogInternalDomainBuilding(in *jlexer.Lexer, out *Building) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint(in.Uint())
		case "address":
			out.Address = string(in.String())
		case "coords":
			easyjson1cf24d94DecodeCatalogInternalPkgCoords(in, &out.Coords)
		case "Organizations":
			if in.IsNull() {
				in.Skip()
				out.Organizations = nil
			} else {
				in.Delim('[')
				if out.Organizations == nil {
					if !in.IsDelim(']') {
						out.Organizations = make([]organization.Organization, 0, 1)
					} else {
						out.Organizations = []organization.Organization{}
					}
				} else {
					out.Organizations = (out.Organizations)[:0]
				}
				for !in.IsDelim(']') {
					var v1 organization.Organization
					easyjson1cf24d94DecodeCatalogInternalDomainOrganization(in, &v1)
					out.Organizations = append(out.Organizations, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1cf24d94EncodeCatalogInternalDomainBuilding(out *jwriter.Writer, in Building) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.ID))
	}
	{
		const prefix string = ",\"address\":"
		out.RawString(prefix)
		out.String(string(in.Address))
	}
	{
		const prefix string = ",\"coords\":"
		out.RawString(prefix)
		easyjson1cf24d94EncodeCatalogInternalPkgCoords(out, in.Coords)
	}
	{
		const prefix string = ",\"Organizations\":"
		out.RawString(prefix)
		if in.Organizations == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Organizations {
				if v2 > 0 {
					out.RawByte(',')
				}
				easyjson1cf24d94EncodeCatalogInternalDomainOrganization(out, v3)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Building) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1cf24d94EncodeCatalogInternalDomainBuilding(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Building) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1cf24d94EncodeCatalogInternalDomainBuilding(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Building) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1cf24d94DecodeCatalogInternalDomainBuilding(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Building) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1cf24d94DecodeCatalogInternalDomainBuilding(l, v)
}
func easyjson1cf24d94DecodeCatalogInternalDomainOrganization(in *jlexer.Lexer, out *organization.Organization) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint(in.Uint())
		case "name":
			out.Name = string(in.String())
		case "phones":
			if in.IsNull() {
				in.Skip()
				out.Phones = nil
			} else {
				in.Delim('[')
				if out.Phones == nil {
					if !in.IsDelim(']') {
						out.Phones = make([]organization_phone.OrganizationPhone, 0, 2)
					} else {
						out.Phones = []organization_phone.OrganizationPhone{}
					}
				} else {
					out.Phones = (out.Phones)[:0]
				}
				for !in.IsDelim(']') {
					var v4 organization_phone.OrganizationPhone
					easyjson1cf24d94DecodeCatalogInternalDomainOrganizationPhone(in, &v4)
					out.Phones = append(out.Phones, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1cf24d94EncodeCatalogInternalDomainOrganization(out *jwriter.Writer, in organization.Organization) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"phones\":"
		out.RawString(prefix)
		if in.Phones == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Phones {
				if v5 > 0 {
					out.RawByte(',')
				}
				easyjson1cf24d94EncodeCatalogInternalDomainOrganizationPhone(out, v6)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson1cf24d94DecodeCatalogInternalDomainOrganizationPhone(in *jlexer.Lexer, out *organization_phone.OrganizationPhone) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = uint(in.Uint())
		case "organizationId":
			out.OrganizationID = uint(in.Uint())
		case "number":
			out.Number = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1cf24d94EncodeCatalogInternalDomainOrganizationPhone(out *jwriter.Writer, in organization_phone.OrganizationPhone) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.ID))
	}
	{
		const prefix string = ",\"organizationId\":"
		out.RawString(prefix)
		out.Uint(uint(in.OrganizationID))
	}
	{
		const prefix string = ",\"number\":"
		out.RawString(prefix)
		out.String(string(in.Number))
	}
	out.RawByte('}')
}
func easyjson1cf24d94DecodeCatalogInternalPkgCoords(in *jlexer.Lexer, out *coords.Coords) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "lat":
			out.Lat = float64(in.Float64())
		case "lng":
			out.Lng = float64(in.Float64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson1cf24d94EncodeCatalogInternalPkgCoords(out *jwriter.Writer, in coords.Coords) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"lat\":"
		out.RawString(prefix[1:])
		out.Float64(float64(in.Lat))
	}
	{
		const prefix string = ",\"lng\":"
		out.RawString(prefix)
		out.Float64(float64(in.Lng))
	}
	out.RawByte('}')
}
