package http

import (
	"bytes"
	"encoding/xml"
	"github.com/roobiuli/NameCheapGoDomain/models"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type NCclient struct {
	c http.Client
	apiuser string
	apikey string

}

func NewNcclient() *NCclient  {
	return &NCclient{
		c: http.Client{},
		apiuser: os.Getenv("APIUSER"),
		apikey: os.Getenv("APIKEY"),
	}
}

func (c *NCclient) SetDomainRecord(dom, record, recordtype, recordaddr string) (*models.NcSetDnsHostsResponse, error) {

	reqdata := &models.NcSetDnsHostsRequest{
		NcGetDnsHostsRequest: models.NcGetDnsHostsRequest{
			ApiUser:  c.apiuser,
			ApiKey:   c.apikey,
			UserName: c.apiuser,
			Command:  "namecheap.domains.dns.setHosts",
			ClientIp: "192.168.0.22",
			Domain:   dom,
			Tld:      strings.Split(dom,".")[1],
		},
		HostName:             record,
		RecordType:           recordtype,
		RecordAddress:        recordaddr,
		Ttl:                  "100",
	}

	reqbyetdata, err := xml.Marshal(&reqdata)

	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, "https://api.namecheap.com/xml.response", bytes.NewBuffer(reqbyetdata))

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "text/xml")

	resp, err := c.c.Do(req)

	if err != nil {
		return nil, err
	}

	var xmlresp models.NcSetDnsHostsResponse

	data, _ := ioutil.ReadAll(resp.Body)

	xml.Unmarshal(data, &xmlresp)

	return &xmlresp, nil

}

func (c *NCclient) GetDomainRecords(dom, record string) (*models.NcGetDnsHostsResponse, error) {
	reqdata := &models.NcGetDnsHostsRequest{
		ApiUser: c.apiuser,
		ApiKey:   c.apiuser,
		UserName: c.apiuser,
		Command:  "namecheap.domains.dns.getHosts",
		ClientIp: "192.168.0.22",
		Domain:   dom,
		Tld:      strings.Split(dom,".")[1],
	}
	reqdatabyte, err := xml.Marshal(&reqdata)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.namecheap.com/xml.response", bytes.NewBuffer(reqdatabyte))

	req.Header.Set("Contexnt-Type", "text/xml")

	resp, err := c.c.Do(req)

	if err != nil {
		return nil, err
	}

	var respdata models.NcGetDnsHostsResponse

	rawd, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	errr := xml.Unmarshal(rawd, &respdata)

	if errr != nil {
		return nil, errr
	}
	return &respdata, nil

}

