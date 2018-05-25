package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetOfficeConversionTaskRequest struct {
	requests.RpcRequest
	Project string `position:"Query" name:"Project"`
	TaskId  string `position:"Query" name:"TaskId"`
}

func (req *GetOfficeConversionTaskRequest) Invoke(client *sdk.Client) (resp *GetOfficeConversionTaskResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "GetOfficeConversionTask", "imm", "")
	resp = &GetOfficeConversionTaskResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetOfficeConversionTaskResponse struct {
	responses.BaseResponse
	RequestId       common.String
	TgtType         common.String
	Status          common.String
	Percent         common.Integer
	PageCount       common.Integer
	TaskId          common.String
	TgtUri          common.String
	ImageSpec       common.String
	NotifyTopicName common.String
	NotifyEndpoint  common.String
	ExternalID      common.String
	CreateTime      common.String
	FinishTime      common.String
	SrcUri          common.String
	FailDetail      GetOfficeConversionTaskFailDetail
}

type GetOfficeConversionTaskFailDetail struct {
	Code common.String
}
