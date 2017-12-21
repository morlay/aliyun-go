package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamRecordIndexFileRequest struct {
	requests.RpcRequest
	RecordId      string `position:"Query" name:"RecordId"`
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (req *DescribeLiveStreamRecordIndexFileRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamRecordIndexFileResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamRecordIndexFile", "", "")
	resp = &DescribeLiveStreamRecordIndexFileResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamRecordIndexFileResponse struct {
	responses.BaseResponse
	RequestId       string
	RecordIndexInfo DescribeLiveStreamRecordIndexFileRecordIndexInfo
}

type DescribeLiveStreamRecordIndexFileRecordIndexInfo struct {
	RecordId   string
	RecordUrl  string
	DomainName string
	AppName    string
	StreamName string
	OssObject  string
	StartTime  string
	EndTime    string
	Duration   float32
	Height     int
	Width      int
	CreateTime string
}
