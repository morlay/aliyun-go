package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type PhotoProcessRequest struct {
	requests.RpcRequest
	NotifyTopicName string `position:"Query" name:"NotifyTopicName"`
	NotifyEndpoint  string `position:"Query" name:"NotifyEndpoint"`
	Project         string `position:"Query" name:"Project"`
	ExternalID      string `position:"Query" name:"ExternalID"`
	SrcUri          string `position:"Query" name:"SrcUri"`
	Style           string `position:"Query" name:"Style"`
	TgtUri          string `position:"Query" name:"TgtUri"`
}

func (req *PhotoProcessRequest) Invoke(client *sdk.Client) (resp *PhotoProcessResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "PhotoProcess", "imm", "")
	resp = &PhotoProcessResponse{}
	err = client.DoAction(req, resp)
	return
}

type PhotoProcessResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TaskId     common.String
	TgtLoc     common.String
	Status     common.String
	CreateTime common.String
	Percent    common.Integer
}
