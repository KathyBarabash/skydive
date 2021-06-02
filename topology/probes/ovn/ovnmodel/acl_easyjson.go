// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package ovnmodel

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

func easyjsonD215af56DecodeGithubComSkydiveProjectSkydiveTopologyProbesOvnOvnmodel(in *jlexer.Lexer, out *ACL) {
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
		case "UUID":
			out.UUID = string(in.String())
		case "Action":
			out.Action = string(in.String())
		case "Direction":
			out.Direction = string(in.String())
		case "ExternalIDs":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.ExternalIDs = make(map[string]string)
				} else {
					out.ExternalIDs = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 string
					v1 = string(in.String())
					(out.ExternalIDs)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "Log":
			out.Log = bool(in.Bool())
		case "Match":
			out.Match = string(in.String())
		case "Meter":
			if in.IsNull() {
				in.Skip()
				out.Meter = nil
			} else {
				in.Delim('[')
				if out.Meter == nil {
					if !in.IsDelim(']') {
						out.Meter = make([]string, 0, 4)
					} else {
						out.Meter = []string{}
					}
				} else {
					out.Meter = (out.Meter)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.Meter = append(out.Meter, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				in.Delim('[')
				if out.Name == nil {
					if !in.IsDelim(']') {
						out.Name = make([]string, 0, 4)
					} else {
						out.Name = []string{}
					}
				} else {
					out.Name = (out.Name)[:0]
				}
				for !in.IsDelim(']') {
					var v3 string
					v3 = string(in.String())
					out.Name = append(out.Name, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Priority":
			out.Priority = int(in.Int())
		case "Severity":
			if in.IsNull() {
				in.Skip()
				out.Severity = nil
			} else {
				in.Delim('[')
				if out.Severity == nil {
					if !in.IsDelim(']') {
						out.Severity = make([]string, 0, 4)
					} else {
						out.Severity = []string{}
					}
				} else {
					out.Severity = (out.Severity)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Severity = append(out.Severity, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "ExternalIDsMeta":
			(out.ExternalIDsMeta).UnmarshalEasyJSON(in)
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
func easyjsonD215af56EncodeGithubComSkydiveProjectSkydiveTopologyProbesOvnOvnmodel(out *jwriter.Writer, in ACL) {
	out.RawByte('{')
	first := true
	_ = first
	if in.UUID != "" {
		const prefix string = ",\"UUID\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.UUID))
	}
	if in.Action != "" {
		const prefix string = ",\"Action\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Action))
	}
	if in.Direction != "" {
		const prefix string = ",\"Direction\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Direction))
	}
	if len(in.ExternalIDs) != 0 {
		const prefix string = ",\"ExternalIDs\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v5First := true
			for v5Name, v5Value := range in.ExternalIDs {
				if v5First {
					v5First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v5Name))
				out.RawByte(':')
				out.String(string(v5Value))
			}
			out.RawByte('}')
		}
	}
	if in.Log {
		const prefix string = ",\"Log\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Log))
	}
	if in.Match != "" {
		const prefix string = ",\"Match\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Match))
	}
	if len(in.Meter) != 0 {
		const prefix string = ",\"Meter\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v6, v7 := range in.Meter {
				if v6 > 0 {
					out.RawByte(',')
				}
				out.String(string(v7))
			}
			out.RawByte(']')
		}
	}
	if len(in.Name) != 0 {
		const prefix string = ",\"Name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v8, v9 := range in.Name {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	if in.Priority != 0 {
		const prefix string = ",\"Priority\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Priority))
	}
	if len(in.Severity) != 0 {
		const prefix string = ",\"Severity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v10, v11 := range in.Severity {
				if v10 > 0 {
					out.RawByte(',')
				}
				out.String(string(v11))
			}
			out.RawByte(']')
		}
	}
	if len(in.ExternalIDsMeta) != 0 {
		const prefix string = ",\"ExternalIDsMeta\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.ExternalIDsMeta).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ACL) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD215af56EncodeGithubComSkydiveProjectSkydiveTopologyProbesOvnOvnmodel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ACL) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD215af56EncodeGithubComSkydiveProjectSkydiveTopologyProbesOvnOvnmodel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ACL) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD215af56DecodeGithubComSkydiveProjectSkydiveTopologyProbesOvnOvnmodel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ACL) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD215af56DecodeGithubComSkydiveProjectSkydiveTopologyProbesOvnOvnmodel(l, v)
}
