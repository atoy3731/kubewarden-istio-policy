// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

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

func easyjson44744437DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *WeightedPodAffinityTerm) {
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
		case "podAffinityTerm":
			if in.IsNull() {
				in.Skip()
				out.PodAffinityTerm = nil
			} else {
				if out.PodAffinityTerm == nil {
					out.PodAffinityTerm = new(PodAffinityTerm)
				}
				(*out.PodAffinityTerm).UnmarshalEasyJSON(in)
			}
		case "weight":
			if in.IsNull() {
				in.Skip()
				out.Weight = nil
			} else {
				if out.Weight == nil {
					out.Weight = new(int32)
				}
				*out.Weight = int32(in.Int32())
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
func easyjson44744437EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in WeightedPodAffinityTerm) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"podAffinityTerm\":"
		out.RawString(prefix[1:])
		if in.PodAffinityTerm == nil {
			out.RawString("null")
		} else {
			(*in.PodAffinityTerm).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"weight\":"
		out.RawString(prefix)
		if in.Weight == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.Weight))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v WeightedPodAffinityTerm) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson44744437EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v WeightedPodAffinityTerm) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson44744437EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *WeightedPodAffinityTerm) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson44744437DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *WeightedPodAffinityTerm) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson44744437DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
