package domain_intl

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryDnsHostRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	Lang       string `position:"Query" name:"Lang"`
}

func (req *QueryDnsHostRequest) Invoke(client *sdk.Client) (resp *QueryDnsHostResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "QueryDnsHost", "domain", "")
	resp = &QueryDnsHostResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryDnsHostResponse struct {
	responses.BaseResponse
	RequestId   common.String
	DnsHostList QueryDnsHostDnsHostList
}

type QueryDnsHostDnsHost struct {
	DnsName common.String
	IpList  QueryDnsHostIpListList
}

type QueryDnsHostDnsHostList []QueryDnsHostDnsHost

func (list *QueryDnsHostDnsHostList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryDnsHostDnsHost)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type QueryDnsHostIpListList []common.String

func (list *QueryDnsHostIpListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
