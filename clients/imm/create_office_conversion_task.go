package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateOfficeConversionTaskRequest struct {
	requests.RpcRequest
	ImageSpec       string `position:"Query" name:"ImageSpec"`
	SrcType         string `position:"Query" name:"SrcType"`
	NotifyTopicName string `position:"Query" name:"NotifyTopicName"`
	ModelId         string `position:"Query" name:"ModelId"`
	Project         string `position:"Query" name:"Project"`
	ExternalID      string `position:"Query" name:"ExternalID"`
	MaxSheetRow     int64  `position:"Query" name:"MaxSheetRow"`
	MaxSheetCount   int64  `position:"Query" name:"MaxSheetCount"`
	EndPage         int64  `position:"Query" name:"EndPage"`
	SheetOnePage    string `position:"Query" name:"SheetOnePage"`
	Password        string `position:"Query" name:"Password"`
	StartPage       int64  `position:"Query" name:"StartPage"`
	MaxSheetCol     int64  `position:"Query" name:"MaxSheetCol"`
	TgtType         string `position:"Query" name:"TgtType"`
	NotifyEndpoint  string `position:"Query" name:"NotifyEndpoint"`
	SrcUri          string `position:"Query" name:"SrcUri"`
	TgtUri          string `position:"Query" name:"TgtUri"`
}

func (req *CreateOfficeConversionTaskRequest) Invoke(client *sdk.Client) (resp *CreateOfficeConversionTaskResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "CreateOfficeConversionTask", "imm", "")
	resp = &CreateOfficeConversionTaskResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateOfficeConversionTaskResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TaskId     common.String
	TgtLoc     common.String
	Status     common.String
	CreateTime common.String
	Percent    common.Integer
}
