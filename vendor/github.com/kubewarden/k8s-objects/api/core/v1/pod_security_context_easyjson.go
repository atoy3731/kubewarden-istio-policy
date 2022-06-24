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

func easyjson6769e492DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *PodSecurityContext) {
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
		case "fsGroup":
			out.FsGroup = int64(in.Int64())
		case "fsGroupChangePolicy":
			out.FsGroupChangePolicy = string(in.String())
		case "runAsGroup":
			out.RunAsGroup = int64(in.Int64())
		case "runAsNonRoot":
			out.RunAsNonRoot = bool(in.Bool())
		case "runAsUser":
			out.RunAsUser = int64(in.Int64())
		case "seLinuxOptions":
			if in.IsNull() {
				in.Skip()
				out.SeLinuxOptions = nil
			} else {
				if out.SeLinuxOptions == nil {
					out.SeLinuxOptions = new(SELinuxOptions)
				}
				easyjson6769e492DecodeGithubComKubewardenK8sObjectsApiCoreV11(in, out.SeLinuxOptions)
			}
		case "seccompProfile":
			if in.IsNull() {
				in.Skip()
				out.SeccompProfile = nil
			} else {
				if out.SeccompProfile == nil {
					out.SeccompProfile = new(SeccompProfile)
				}
				easyjson6769e492DecodeGithubComKubewardenK8sObjectsApiCoreV12(in, out.SeccompProfile)
			}
		case "supplementalGroups":
			if in.IsNull() {
				in.Skip()
				out.SupplementalGroups = nil
			} else {
				in.Delim('[')
				if out.SupplementalGroups == nil {
					if !in.IsDelim(']') {
						out.SupplementalGroups = make([]int64, 0, 8)
					} else {
						out.SupplementalGroups = []int64{}
					}
				} else {
					out.SupplementalGroups = (out.SupplementalGroups)[:0]
				}
				for !in.IsDelim(']') {
					var v1 int64
					v1 = int64(in.Int64())
					out.SupplementalGroups = append(out.SupplementalGroups, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "sysctls":
			if in.IsNull() {
				in.Skip()
				out.Sysctls = nil
			} else {
				in.Delim('[')
				if out.Sysctls == nil {
					if !in.IsDelim(']') {
						out.Sysctls = make([]*Sysctl, 0, 8)
					} else {
						out.Sysctls = []*Sysctl{}
					}
				} else {
					out.Sysctls = (out.Sysctls)[:0]
				}
				for !in.IsDelim(']') {
					var v2 *Sysctl
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(Sysctl)
						}
						easyjson6769e492DecodeGithubComKubewardenK8sObjectsApiCoreV13(in, v2)
					}
					out.Sysctls = append(out.Sysctls, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "windowsOptions":
			if in.IsNull() {
				in.Skip()
				out.WindowsOptions = nil
			} else {
				if out.WindowsOptions == nil {
					out.WindowsOptions = new(WindowsSecurityContextOptions)
				}
				easyjson6769e492DecodeGithubComKubewardenK8sObjectsApiCoreV14(in, out.WindowsOptions)
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
func easyjson6769e492EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in PodSecurityContext) {
	out.RawByte('{')
	first := true
	_ = first
	if in.FsGroup != 0 {
		const prefix string = ",\"fsGroup\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.FsGroup))
	}
	if in.FsGroupChangePolicy != "" {
		const prefix string = ",\"fsGroupChangePolicy\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FsGroupChangePolicy))
	}
	if in.RunAsGroup != 0 {
		const prefix string = ",\"runAsGroup\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.RunAsGroup))
	}
	if in.RunAsNonRoot {
		const prefix string = ",\"runAsNonRoot\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.RunAsNonRoot))
	}
	if in.RunAsUser != 0 {
		const prefix string = ",\"runAsUser\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.RunAsUser))
	}
	if in.SeLinuxOptions != nil {
		const prefix string = ",\"seLinuxOptions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson6769e492EncodeGithubComKubewardenK8sObjectsApiCoreV11(out, *in.SeLinuxOptions)
	}
	if in.SeccompProfile != nil {
		const prefix string = ",\"seccompProfile\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson6769e492EncodeGithubComKubewardenK8sObjectsApiCoreV12(out, *in.SeccompProfile)
	}
	{
		const prefix string = ",\"supplementalGroups\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.SupplementalGroups == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.SupplementalGroups {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.Int64(int64(v4))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"sysctls\":"
		out.RawString(prefix)
		if in.Sysctls == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Sysctls {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					easyjson6769e492EncodeGithubComKubewardenK8sObjectsApiCoreV13(out, *v6)
				}
			}
			out.RawByte(']')
		}
	}
	if in.WindowsOptions != nil {
		const prefix string = ",\"windowsOptions\":"
		out.RawString(prefix)
		easyjson6769e492EncodeGithubComKubewardenK8sObjectsApiCoreV14(out, *in.WindowsOptions)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PodSecurityContext) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6769e492EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PodSecurityContext) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6769e492EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PodSecurityContext) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6769e492DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PodSecurityContext) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6769e492DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
func easyjson6769e492DecodeGithubComKubewardenK8sObjectsApiCoreV14(in *jlexer.Lexer, out *WindowsSecurityContextOptions) {
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
		case "gmsaCredentialSpec":
			out.GmsaCredentialSpec = string(in.String())
		case "gmsaCredentialSpecName":
			out.GmsaCredentialSpecName = string(in.String())
		case "hostProcess":
			out.HostProcess = bool(in.Bool())
		case "runAsUserName":
			out.RunAsUserName = string(in.String())
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
func easyjson6769e492EncodeGithubComKubewardenK8sObjectsApiCoreV14(out *jwriter.Writer, in WindowsSecurityContextOptions) {
	out.RawByte('{')
	first := true
	_ = first
	if in.GmsaCredentialSpec != "" {
		const prefix string = ",\"gmsaCredentialSpec\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.GmsaCredentialSpec))
	}
	if in.GmsaCredentialSpecName != "" {
		const prefix string = ",\"gmsaCredentialSpecName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.GmsaCredentialSpecName))
	}
	if in.HostProcess {
		const prefix string = ",\"hostProcess\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.HostProcess))
	}
	if in.RunAsUserName != "" {
		const prefix string = ",\"runAsUserName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RunAsUserName))
	}
	out.RawByte('}')
}
func easyjson6769e492DecodeGithubComKubewardenK8sObjectsApiCoreV13(in *jlexer.Lexer, out *Sysctl) {
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
		case "name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
			}
		case "value":
			if in.IsNull() {
				in.Skip()
				out.Value = nil
			} else {
				if out.Value == nil {
					out.Value = new(string)
				}
				*out.Value = string(in.String())
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
func easyjson6769e492EncodeGithubComKubewardenK8sObjectsApiCoreV13(out *jwriter.Writer, in Sysctl) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		if in.Value == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Value))
		}
	}
	out.RawByte('}')
}
func easyjson6769e492DecodeGithubComKubewardenK8sObjectsApiCoreV12(in *jlexer.Lexer, out *SeccompProfile) {
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
		case "localhostProfile":
			out.LocalhostProfile = string(in.String())
		case "type":
			if in.IsNull() {
				in.Skip()
				out.Type = nil
			} else {
				if out.Type == nil {
					out.Type = new(string)
				}
				*out.Type = string(in.String())
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
func easyjson6769e492EncodeGithubComKubewardenK8sObjectsApiCoreV12(out *jwriter.Writer, in SeccompProfile) {
	out.RawByte('{')
	first := true
	_ = first
	if in.LocalhostProfile != "" {
		const prefix string = ",\"localhostProfile\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.LocalhostProfile))
	}
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Type == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Type))
		}
	}
	out.RawByte('}')
}
func easyjson6769e492DecodeGithubComKubewardenK8sObjectsApiCoreV11(in *jlexer.Lexer, out *SELinuxOptions) {
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
		case "level":
			out.Level = string(in.String())
		case "role":
			out.Role = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "user":
			out.User = string(in.String())
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
func easyjson6769e492EncodeGithubComKubewardenK8sObjectsApiCoreV11(out *jwriter.Writer, in SELinuxOptions) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Level != "" {
		const prefix string = ",\"level\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Level))
	}
	if in.Role != "" {
		const prefix string = ",\"role\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Role))
	}
	if in.Type != "" {
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	if in.User != "" {
		const prefix string = ",\"user\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.User))
	}
	out.RawByte('}')
}