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

func easyjson8674d8afDecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *TopologySpreadConstraint) {
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
		case "labelSelector":
			(out.LabelSelector).UnmarshalEasyJSON(in)
		case "maxSkew":
			if in.IsNull() {
				in.Skip()
				out.MaxSkew = nil
			} else {
				if out.MaxSkew == nil {
					out.MaxSkew = new(int32)
				}
				*out.MaxSkew = int32(in.Int32())
			}
		case "minDomains":
			out.MinDomains = int32(in.Int32())
		case "topologyKey":
			if in.IsNull() {
				in.Skip()
				out.TopologyKey = nil
			} else {
				if out.TopologyKey == nil {
					out.TopologyKey = new(string)
				}
				*out.TopologyKey = string(in.String())
			}
		case "whenUnsatisfiable":
			if in.IsNull() {
				in.Skip()
				out.WhenUnsatisfiable = nil
			} else {
				if out.WhenUnsatisfiable == nil {
					out.WhenUnsatisfiable = new(string)
				}
				*out.WhenUnsatisfiable = string(in.String())
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
func easyjson8674d8afEncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in TopologySpreadConstraint) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		const prefix string = ",\"labelSelector\":"
		first = false
		out.RawString(prefix[1:])
		(in.LabelSelector).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"maxSkew\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.MaxSkew == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.MaxSkew))
		}
	}
	if in.MinDomains != 0 {
		const prefix string = ",\"minDomains\":"
		out.RawString(prefix)
		out.Int32(int32(in.MinDomains))
	}
	{
		const prefix string = ",\"topologyKey\":"
		out.RawString(prefix)
		if in.TopologyKey == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.TopologyKey))
		}
	}
	{
		const prefix string = ",\"whenUnsatisfiable\":"
		out.RawString(prefix)
		if in.WhenUnsatisfiable == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.WhenUnsatisfiable))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TopologySpreadConstraint) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8674d8afEncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TopologySpreadConstraint) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8674d8afEncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TopologySpreadConstraint) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8674d8afDecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TopologySpreadConstraint) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8674d8afDecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
