// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
	_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
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

func easyjson4b13d882DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *ServiceStatus) {
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
		case "conditions":
			if in.IsNull() {
				in.Skip()
				out.Conditions = nil
			} else {
				in.Delim('[')
				if out.Conditions == nil {
					if !in.IsDelim(']') {
						out.Conditions = make([]_v1.Condition, 0, 1)
					} else {
						out.Conditions = []_v1.Condition{}
					}
				} else {
					out.Conditions = (out.Conditions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 _v1.Condition
					(v1).UnmarshalEasyJSON(in)
					out.Conditions = append(out.Conditions, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "loadBalancer":
			if in.IsNull() {
				in.Skip()
				out.LoadBalancer = nil
			} else {
				if out.LoadBalancer == nil {
					out.LoadBalancer = new(LoadBalancerStatus)
				}
				(*out.LoadBalancer).UnmarshalEasyJSON(in)
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
func easyjson4b13d882EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in ServiceStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"conditions\":"
		out.RawString(prefix[1:])
		if in.Conditions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Conditions {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if in.LoadBalancer != nil {
		const prefix string = ",\"loadBalancer\":"
		out.RawString(prefix)
		(*in.LoadBalancer).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4b13d882EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4b13d882EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4b13d882DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4b13d882DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
