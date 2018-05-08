package imm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ConvertOfficeFormatRequest struct {
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
	StartPage       int64  `position:"Query" name:"StartPage"`
	MaxSheetCol     int64  `position:"Query" name:"MaxSheetCol"`
	TgtType         string `position:"Query" name:"TgtType"`
	NotifyEndpoint  string `position:"Query" name:"NotifyEndpoint"`
	SrcUri          string `position:"Query" name:"SrcUri"`
	TgtUri          string `position:"Query" name:"TgtUri"`
}

func (req *ConvertOfficeFormatRequest) Invoke(client *sdk.Client) (resp *ConvertOfficeFormatResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "ConvertOfficeFormat", "imm", "")
	resp = &ConvertOfficeFormatResponse{}
	err = client.DoAction(req, resp)
	return
}

type ConvertOfficeFormatResponse struct {
	responses.BaseResponse
	RequestId  string
	TaskId     string
	TgtLoc     string
	Status     string
	CreateTime string
	Percent    int
}
