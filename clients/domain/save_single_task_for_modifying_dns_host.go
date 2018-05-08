package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveSingleTaskForModifyingDnsHostRequest struct {
	requests.RpcRequest
	InstanceId string                                   `position:"Query" name:"InstanceId"`
	Ips        *SaveSingleTaskForModifyingDnsHostIpList `position:"Query" type:"Repeated" name:"Ip"`
	DnsName    string                                   `position:"Query" name:"DnsName"`
	Lang       string                                   `position:"Query" name:"Lang"`
}

func (req *SaveSingleTaskForModifyingDnsHostRequest) Invoke(client *sdk.Client) (resp *SaveSingleTaskForModifyingDnsHostResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveSingleTaskForModifyingDnsHost", "", "")
	resp = &SaveSingleTaskForModifyingDnsHostResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveSingleTaskForModifyingDnsHostResponse struct {
	responses.BaseResponse
	RequestId common.String
	TaskNo    common.String
}

type SaveSingleTaskForModifyingDnsHostIpList []string

func (list *SaveSingleTaskForModifyingDnsHostIpList) UnmarshalJSON(data []byte) error {
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
