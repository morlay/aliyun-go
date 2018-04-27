package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveSingleTaskForSynchronizingDnsHostRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	Lang       string `position:"Query" name:"Lang"`
}

func (req *SaveSingleTaskForSynchronizingDnsHostRequest) Invoke(client *sdk.Client) (resp *SaveSingleTaskForSynchronizingDnsHostResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveSingleTaskForSynchronizingDnsHost", "", "")
	resp = &SaveSingleTaskForSynchronizingDnsHostResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveSingleTaskForSynchronizingDnsHostResponse struct {
	responses.BaseResponse
	RequestId string
	TaskNo    string
}
