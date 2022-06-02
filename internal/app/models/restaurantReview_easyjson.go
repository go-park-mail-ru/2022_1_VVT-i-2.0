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

func easyjson547d70d5DecodeGithubComGoParkMailRu20221VVTI20InternalAppModels(in *jlexer.Lexer, out *ReviewResp) {
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
		case "author":
			out.Author = string(in.String())
		case "text":
			out.Text = string(in.String())
		case "stars":
			out.Stars = int(in.Int())
		case "date":
			out.Date = string(in.String())
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
func easyjson547d70d5EncodeGithubComGoParkMailRu20221VVTI20InternalAppModels(out *jwriter.Writer, in ReviewResp) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"author\":"
		out.RawString(prefix[1:])
		out.String(string(in.Author))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"stars\":"
		out.RawString(prefix)
		out.Int(int(in.Stars))
	}
	{
		const prefix string = ",\"date\":"
		out.RawString(prefix)
		out.String(string(in.Date))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReviewResp) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson547d70d5EncodeGithubComGoParkMailRu20221VVTI20InternalAppModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReviewResp) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson547d70d5EncodeGithubComGoParkMailRu20221VVTI20InternalAppModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReviewResp) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson547d70d5DecodeGithubComGoParkMailRu20221VVTI20InternalAppModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReviewResp) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson547d70d5DecodeGithubComGoParkMailRu20221VVTI20InternalAppModels(l, v)
}
func easyjson547d70d5DecodeGithubComGoParkMailRu20221VVTI20InternalAppModels1(in *jlexer.Lexer, out *GetRestaurantReviews) {
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
		case "comment":
			if in.IsNull() {
				in.Skip()
				out.Reviews = nil
			} else {
				in.Delim('[')
				if out.Reviews == nil {
					if !in.IsDelim(']') {
						out.Reviews = make([]ReviewResp, 0, 1)
					} else {
						out.Reviews = []ReviewResp{}
					}
				} else {
					out.Reviews = (out.Reviews)[:0]
				}
				for !in.IsDelim(']') {
					var v1 ReviewResp
					(v1).UnmarshalEasyJSON(in)
					out.Reviews = append(out.Reviews, v1)
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
func easyjson547d70d5EncodeGithubComGoParkMailRu20221VVTI20InternalAppModels1(out *jwriter.Writer, in GetRestaurantReviews) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"comment\":"
		out.RawString(prefix[1:])
		if in.Reviews == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Reviews {
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
func (v GetRestaurantReviews) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson547d70d5EncodeGithubComGoParkMailRu20221VVTI20InternalAppModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetRestaurantReviews) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson547d70d5EncodeGithubComGoParkMailRu20221VVTI20InternalAppModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetRestaurantReviews) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson547d70d5DecodeGithubComGoParkMailRu20221VVTI20InternalAppModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetRestaurantReviews) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson547d70d5DecodeGithubComGoParkMailRu20221VVTI20InternalAppModels1(l, v)
}
func easyjson547d70d5DecodeGithubComGoParkMailRu20221VVTI20InternalAppModels2(in *jlexer.Lexer, out *AddRestaurantReviewResp) {
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
		case "author":
			out.Author = string(in.String())
		case "text":
			out.Text = string(in.String())
		case "stars":
			out.Stars = int(in.Int())
		case "date":
			out.Date = string(in.String())
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
func easyjson547d70d5EncodeGithubComGoParkMailRu20221VVTI20InternalAppModels2(out *jwriter.Writer, in AddRestaurantReviewResp) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"author\":"
		out.RawString(prefix[1:])
		out.String(string(in.Author))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"stars\":"
		out.RawString(prefix)
		out.Int(int(in.Stars))
	}
	{
		const prefix string = ",\"date\":"
		out.RawString(prefix)
		out.String(string(in.Date))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AddRestaurantReviewResp) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson547d70d5EncodeGithubComGoParkMailRu20221VVTI20InternalAppModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AddRestaurantReviewResp) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson547d70d5EncodeGithubComGoParkMailRu20221VVTI20InternalAppModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AddRestaurantReviewResp) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson547d70d5DecodeGithubComGoParkMailRu20221VVTI20InternalAppModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AddRestaurantReviewResp) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson547d70d5DecodeGithubComGoParkMailRu20221VVTI20InternalAppModels2(l, v)
}
func easyjson547d70d5DecodeGithubComGoParkMailRu20221VVTI20InternalAppModels3(in *jlexer.Lexer, out *AddRestaurantReviewReq) {
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
		case "slug":
			out.Slug = string(in.String())
		case "text":
			out.Text = string(in.String())
		case "stars":
			out.Rating = int(in.Int())
		case "orderId":
			out.OrderId = int64(in.Int64())
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
func easyjson547d70d5EncodeGithubComGoParkMailRu20221VVTI20InternalAppModels3(out *jwriter.Writer, in AddRestaurantReviewReq) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"slug\":"
		out.RawString(prefix[1:])
		out.String(string(in.Slug))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"stars\":"
		out.RawString(prefix)
		out.Int(int(in.Rating))
	}
	{
		const prefix string = ",\"orderId\":"
		out.RawString(prefix)
		out.Int64(int64(in.OrderId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AddRestaurantReviewReq) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson547d70d5EncodeGithubComGoParkMailRu20221VVTI20InternalAppModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AddRestaurantReviewReq) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson547d70d5EncodeGithubComGoParkMailRu20221VVTI20InternalAppModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AddRestaurantReviewReq) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson547d70d5DecodeGithubComGoParkMailRu20221VVTI20InternalAppModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AddRestaurantReviewReq) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson547d70d5DecodeGithubComGoParkMailRu20221VVTI20InternalAppModels3(l, v)
}
