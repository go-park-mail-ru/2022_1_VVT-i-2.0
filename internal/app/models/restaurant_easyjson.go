// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
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

func easyjson16134a91DecodeGithubComGoParkMailRu20221VVTI20InternalAppModels(in *jlexer.Lexer, out *RestaurantResp) {
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
			out.Id = int(in.Int())
		case "restName":
			out.Name = string(in.String())
		case "imgPath":
			out.ImagePath = string(in.String())
		case "slug":
			out.Slug = string(in.String())
		case "price":
			out.MinPrice = int(in.Int())
		case "rating":
			out.Rating = float64(in.Float64())
		case "timeToDeliver":
			out.TimeToDelivery = string(in.String())
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
func easyjson16134a91EncodeGithubComGoParkMailRu20221VVTI20InternalAppModels(out *jwriter.Writer, in RestaurantResp) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"restName\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"imgPath\":"
		out.RawString(prefix)
		out.String(string(in.ImagePath))
	}
	{
		const prefix string = ",\"slug\":"
		out.RawString(prefix)
		out.String(string(in.Slug))
	}
	{
		const prefix string = ",\"price\":"
		out.RawString(prefix)
		out.Int(int(in.MinPrice))
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		out.Float64(float64(in.Rating))
	}
	{
		const prefix string = ",\"timeToDeliver\":"
		out.RawString(prefix)
		out.String(string(in.TimeToDelivery))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RestaurantResp) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson16134a91EncodeGithubComGoParkMailRu20221VVTI20InternalAppModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RestaurantResp) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson16134a91EncodeGithubComGoParkMailRu20221VVTI20InternalAppModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RestaurantResp) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson16134a91DecodeGithubComGoParkMailRu20221VVTI20InternalAppModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RestaurantResp) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson16134a91DecodeGithubComGoParkMailRu20221VVTI20InternalAppModels(l, v)
}
func easyjson16134a91DecodeGithubComGoParkMailRu20221VVTI20InternalAppModels1(in *jlexer.Lexer, out *AllRestaurantsResp) {
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
		case "restaurants":
			if in.IsNull() {
				in.Skip()
				out.Restaurants = nil
			} else {
				in.Delim('[')
				if out.Restaurants == nil {
					if !in.IsDelim(']') {
						out.Restaurants = make([]RestaurantResp, 0, 0)
					} else {
						out.Restaurants = []RestaurantResp{}
					}
				} else {
					out.Restaurants = (out.Restaurants)[:0]
				}
				for !in.IsDelim(']') {
					var v1 RestaurantResp
					(v1).UnmarshalEasyJSON(in)
					out.Restaurants = append(out.Restaurants, v1)
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
func easyjson16134a91EncodeGithubComGoParkMailRu20221VVTI20InternalAppModels1(out *jwriter.Writer, in AllRestaurantsResp) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"restaurants\":"
		out.RawString(prefix[1:])
		if in.Restaurants == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Restaurants {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AllRestaurantsResp) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson16134a91EncodeGithubComGoParkMailRu20221VVTI20InternalAppModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AllRestaurantsResp) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson16134a91EncodeGithubComGoParkMailRu20221VVTI20InternalAppModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AllRestaurantsResp) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson16134a91DecodeGithubComGoParkMailRu20221VVTI20InternalAppModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AllRestaurantsResp) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson16134a91DecodeGithubComGoParkMailRu20221VVTI20InternalAppModels1(l, v)
}