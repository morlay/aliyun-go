package domain_intl

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveSingleTaskForCreatingDnsHostRequest struct {
	requests.RpcRequest
	InstanceId string                                  `position:"Query" name:"InstanceId"`
	Ips        *SaveSingleTaskForCreatingDnsHostIpList `position:"Query" type:"Repeated" name:"Ip"`
	DnsName    string                                  `position:"Query" name:"DnsName"`
	Lang       string                                  `position:"Query" name:"Lang"`
}

func (req *SaveSingleTaskForCreatingDnsHostRequest) Invoke(client *sdk.Client) (resp *SaveSingleTaskForCreatingDnsHostResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveSingleTaskForCreatingDnsHost", "domain", "")
	resp = &SaveSingleTaskForCreatingDnsHostResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveSingleTaskForCreatingDnsHostResponse struct {
	responses.BaseResponse
	RequestId string
	TaskNo    string
}

type SaveSingleTaskForCreatingDnsHostIpList []string

func (list *SaveSingleTaskForCreatingDnsHostIpList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
