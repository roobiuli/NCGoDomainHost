package models

import "strings"

type BuildProcess struct {
	User string
	Key string
	ClientIp string
	Domain string
	HostnameRecord string
	RecordType string
	RecAddress string
}

type NCRequester interface {
	Decode() ([]byte, error)
}


type NCReqBuilder interface {
	WithDomain(s string) *BuildProcess
	WithRecord(s string) *BuildProcess
	WithRecType(s string) *BuildProcess
	WithRecAddress(s string) *BuildProcess
	Build() *NCRequester
}

type NcSetRecordRequestBuilder struct {
	b *BuildProcess
}

func (r *NcSetRecordRequestBuilder) WithDomain(s string) *BuildProcess {
	r.b = &BuildProcess{Domain: s}
	return r.b
}
func (r *NcSetRecordRequestBuilder) WithRecord(s string) *BuildProcess {
	r.b.HostnameRecord = s
	return r.b
}
func (r *NcSetRecordRequestBuilder) WithRecType(s string) *BuildProcess {
	r.b.RecordType = s
	return r.b
}
func (r *NcSetRecordRequestBuilder) WithRecAddress(s string) *BuildProcess {
	r.b.RecAddress  = s
	return r.b
}
func (r *NcSetRecordRequestBuilder) Build() *NcSetDnsHostsRequest {
	return &NcSetDnsHostsRequest{
		NcGetDnsHostsRequest: NcGetDnsHostsRequest{
			ApiUser:  r.b.User,
			ApiKey:   r.b.Key,
			UserName: r.b.User,
			Command:  "namecheap.domains.dns.setHosts",
			ClientIp: "",
			Domain:   r.b.Domain,
			Tld:      strings.Split(r.b.Domain, ".")[1],
		},
		HostName:             r.b.HostnameRecord,
		RecordType:           r.b.RecordType,
		RecordAddress:        r.b.RecAddress,
		Ttl:                  "100",
	}
}



type NcGetRecordRequestBuilder struct {
	b *BuildProcess
}

func (r *NcGetRecordRequestBuilder) WithDomain(s string) *BuildProcess {
	r.b = &BuildProcess{Domain: s}
	return r.b
}
func (r *NcGetRecordRequestBuilder) WithRecord(s string) *BuildProcess {
	r.b.HostnameRecord = s
	return r.b
}
func (r *NcGetRecordRequestBuilder) WithRecType(s string) *BuildProcess {
	r.b.RecordType = s
	return r.b
}
func (r *NcGetRecordRequestBuilder) WithRecAddress(s string) *BuildProcess {
	r.b.RecAddress  = s
	return r.b
}
func (r *NcGetRecordRequestBuilder) Build() *NcGetDnsHostsRequest {
	return &NcGetDnsHostsRequest{
		ApiUser:  r.b.User,
		ApiKey:   r.b.Key,
		UserName: r.b.User,
		Command:  "namecheap.domains.dns.getHosts",
		ClientIp: "",
		Domain:   r.b.Domain,
		Tld:      strings.Split(r.b.Domain, ".")[1],
	}
}
