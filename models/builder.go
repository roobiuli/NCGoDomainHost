package models


type BuildProcess struct {
	User string
	Key string
	ClientIp string
	Domain string
	HostnameRecord string
	RecordType string
	RecAddress string
}


type NCReqBuilder interface {
	WithDomain(s string) *BuildProcess
	WithRecord(s string) *BuildProcess
	WithRecType(s string) *BuildProcess
	WithRecAddress(s string) *BuildProcess
	Build() NCReqBuilder
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
func (r *NcSetRecordRequestBuilder) Build() NCReqBuilder {
	return r
}
