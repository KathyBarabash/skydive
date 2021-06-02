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

func easyjson86cb25fbDecodeGithubComSkydiveProjectSkydiveTopologyProbesOvnOvnmodel(in *jlexer.Lexer, out *LoadBalancer) {
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
		case "HealthCheck":
			if in.IsNull() {
				in.Skip()
				out.HealthCheck = nil
			} else {
				in.Delim('[')
				if out.HealthCheck == nil {
					if !in.IsDelim(']') {
						out.HealthCheck = make([]string, 0, 4)
					} else {
						out.HealthCheck = []string{}
					}
				} else {
					out.HealthCheck = (out.HealthCheck)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.HealthCheck = append(out.HealthCheck, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "IPPortMappings":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.IPPortMappings = make(map[string]string)
				} else {
					out.IPPortMappings = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v3 string
					v3 = string(in.String())
					(out.IPPortMappings)[key] = v3
					in.WantComma()
				}
				in.Delim('}')
			}
		case "Name":
			out.Name = string(in.String())
		case "Protocol":
			if in.IsNull() {
				in.Skip()
				out.Protocol = nil
			} else {
				in.Delim('[')
				if out.Protocol == nil {
					if !in.IsDelim(']') {
						out.Protocol = make([]string, 0, 4)
					} else {
						out.Protocol = []string{}
					}
				} else {
					out.Protocol = (out.Protocol)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Protocol = append(out.Protocol, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "SelectionFields":
			if in.IsNull() {
				in.Skip()
				out.SelectionFields = nil
			} else {
				in.Delim('[')
				if out.SelectionFields == nil {
					if !in.IsDelim(']') {
						out.SelectionFields = make([]string, 0, 4)
					} else {
						out.SelectionFields = []string{}
					}
				} else {
					out.SelectionFields = (out.SelectionFields)[:0]
				}
				for !in.IsDelim(']') {
					var v5 string
					v5 = string(in.String())
					out.SelectionFields = append(out.SelectionFields, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Vips":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Vips = make(map[string]string)
				} else {
					out.Vips = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v6 string
					v6 = string(in.String())
					(out.Vips)[key] = v6
					in.WantComma()
				}
				in.Delim('}')
			}
		case "ExternalIDsMeta":
			(out.ExternalIDsMeta).UnmarshalEasyJSON(in)
		case "IPPortMappingsMeta":
			(out.IPPortMappingsMeta).UnmarshalEasyJSON(in)
		case "VipsMeta":
			(out.VipsMeta).UnmarshalEasyJSON(in)
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
func easyjson86cb25fbEncodeGithubComSkydiveProjectSkydiveTopologyProbesOvnOvnmodel(out *jwriter.Writer, in LoadBalancer) {
	out.RawByte('{')
	first := true
	_ = first
	if in.UUID != "" {
		const prefix string = ",\"UUID\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.UUID))
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
			v7First := true
			for v7Name, v7Value := range in.ExternalIDs {
				if v7First {
					v7First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v7Name))
				out.RawByte(':')
				out.String(string(v7Value))
			}
			out.RawByte('}')
		}
	}
	if len(in.HealthCheck) != 0 {
		const prefix string = ",\"HealthCheck\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v8, v9 := range in.HealthCheck {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	if len(in.IPPortMappings) != 0 {
		const prefix string = ",\"IPPortMappings\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v10First := true
			for v10Name, v10Value := range in.IPPortMappings {
				if v10First {
					v10First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v10Name))
				out.RawByte(':')
				out.String(string(v10Value))
			}
			out.RawByte('}')
		}
	}
	if in.Name != "" {
		const prefix string = ",\"Name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if len(in.Protocol) != 0 {
		const prefix string = ",\"Protocol\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v11, v12 := range in.Protocol {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	if len(in.SelectionFields) != 0 {
		const prefix string = ",\"SelectionFields\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v13, v14 := range in.SelectionFields {
				if v13 > 0 {
					out.RawByte(',')
				}
				out.String(string(v14))
			}
			out.RawByte(']')
		}
	}
	if len(in.Vips) != 0 {
		const prefix string = ",\"Vips\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v15First := true
			for v15Name, v15Value := range in.Vips {
				if v15First {
					v15First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v15Name))
				out.RawByte(':')
				out.String(string(v15Value))
			}
			out.RawByte('}')
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
	if len(in.IPPortMappingsMeta) != 0 {
		const prefix string = ",\"IPPortMappingsMeta\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.IPPortMappingsMeta).MarshalEasyJSON(out)
	}
	if len(in.VipsMeta) != 0 {
		const prefix string = ",\"VipsMeta\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.VipsMeta).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LoadBalancer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson86cb25fbEncodeGithubComSkydiveProjectSkydiveTopologyProbesOvnOvnmodel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LoadBalancer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson86cb25fbEncodeGithubComSkydiveProjectSkydiveTopologyProbesOvnOvnmodel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LoadBalancer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson86cb25fbDecodeGithubComSkydiveProjectSkydiveTopologyProbesOvnOvnmodel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LoadBalancer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson86cb25fbDecodeGithubComSkydiveProjectSkydiveTopologyProbesOvnOvnmodel(l, v)
}
