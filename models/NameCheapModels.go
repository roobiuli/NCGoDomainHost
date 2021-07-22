package models

import (
	"encoding/xml"
)


type NcResponse struct {
	XMLName xml.Name `xml:"ApiResponse"`
	Status string 	`xml:"Status,attr"`
	Xmlns	string `xml:"xmlns,attr"`
	RequestedCommand string `xml:"RequestedCommand"`
	Server string	`xml:"Server"`
	GMTtime string	`xml:"GMTTimeDifference"`
	ExecutionTime	string	`xml:"ExecutionTime"`

}

type CommandResponse struct {
	XMLName xml.Name `xml:"CommandResponse"`
	Type string	`xml:"Type,attr"`
}

type CommandGetDNS struct {
	CommandResponse
	DomainResult DomainDNSGetHostResult `xml:"DomainDNSGetHostsResult"`
}

type CommandSetDns struct {
	CommandResponse DomainDnsSetHostResult
}

type DomainDNSGetHostResult struct {
	XMLName xml.Name `xml:"DomainDNSGetHostsResult"`
	Domain string `xml:"Domain,attr"`
	IsUsingDNS	string	`xml:"IsUsingOurDns,attr"`
	Hosts []DnsRecord `xml:"Host"`
}

type DomainDnsSetHostResult struct {
	XMLName	string `xml:"DomainDNSSetHostsResult"`
	Domain string	`xml:"Domain,attr"`
	IsSuccess	string	`xml:"IsSuccess,attr"`
}

type DnsRecord struct {
	XMLName xml.Name `xml:"Host"`
	HostId	string	`xml:"HostId,attr"`
	Name string		`xml:"Name,attr"`
	Type string		`xml:"Type,attr"`
	Address string  `xml:"Address,attr"`
	MXPref	string	`xml:"MXPref,attr"`
	Ttl		string	`xml:"TTL,attr"`
}

type NcGetDnsHostsResponse struct {
	NcResponse
	CommandGetDNS
}

type NcSetDnsHostsResponse struct {
	NcResponse
	CommandSetDns
}

type NcSetDnsHostsRequest struct {
	NcGetDnsHostsRequest
	HostName string
	RecordType string
	RecordAddress string
	Ttl string
}

type NcGetDnsHostsRequest struct {
	ApiUser string `xml:"ApiUser"`
	ApiKey string  `xml:"ApiKey"`
	UserName string `xml:"UserName"`
	Command string	`xml:"Command"`
	ClientIp string  `xml:"ClientIp"`
	Domain string    `xml:"SLD"`
	Tld string		`xml:"TLD"`
}










