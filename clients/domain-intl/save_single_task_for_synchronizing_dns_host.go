package domain_intl

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveSingleTaskForSynchronizingDnsHostRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	Lang       string `position:"Query" name:"Lang"`
}

func (req *SaveSingleTaskForSynchronizingDnsHostRequest) Invoke(client *sdk.Client) (resp *SaveSingleTaskForSynchronizingDnsHostResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveSingleTaskForSynchronizingDnsHost", "domain", "")
	resp = &SaveSingleTaskForSynchronizingDnsHostResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveSingleTaskForSynchronizingDnsHostResponse struct {
	responses.BaseResponse
	RequestId common.String
	TaskNo    common.String
}
