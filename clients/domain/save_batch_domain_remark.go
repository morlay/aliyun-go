package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveBatchDomainRemarkRequest struct {
	requests.RpcRequest
	InstanceIds string `position:"Query" name:"InstanceIds"`
	Remark      string `position:"Query" name:"Remark"`
	Lang        string `position:"Query" name:"Lang"`
}

func (req *SaveBatchDomainRemarkRequest) Invoke(client *sdk.Client) (resp *SaveBatchDomainRemarkResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveBatchDomainRemark", "", "")
	resp = &SaveBatchDomainRemarkResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveBatchDomainRemarkResponse struct {
	responses.BaseResponse
	RequestId common.String
}
