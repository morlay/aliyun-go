package afs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetEarlyWarningRequest struct {
	requests.RpcRequest
	TimeEnd         string `position:"Query" name:"TimeEnd"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	WarnOpen        string `position:"Query" name:"WarnOpen"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Channel         string `position:"Query" name:"Channel"`
	Title           string `position:"Query" name:"Title"`
	TimeOpen        string `position:"Query" name:"TimeOpen"`
	TimeBegin       string `position:"Query" name:"TimeBegin"`
	Frequency       string `position:"Query" name:"Frequency"`
}

func (req *SetEarlyWarningRequest) Invoke(client *sdk.Client) (resp *SetEarlyWarningResponse, err error) {
	req.InitWithApiInfo("afs", "2018-01-12", "SetEarlyWarning", "", "")
	resp = &SetEarlyWarningResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetEarlyWarningResponse struct {
	responses.BaseResponse
	RequestId string
	BizCode   string
}
