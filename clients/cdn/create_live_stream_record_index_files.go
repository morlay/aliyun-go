package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateLiveStreamRecordIndexFilesRequest struct {
	requests.RpcRequest
	OssBucket     string `position:"Query" name:"OssBucket"`
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OssEndpoint   string `position:"Query" name:"OssEndpoint"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	OssObject     string `position:"Query" name:"OssObject"`
}

func (req *CreateLiveStreamRecordIndexFilesRequest) Invoke(client *sdk.Client) (resp *CreateLiveStreamRecordIndexFilesResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "CreateLiveStreamRecordIndexFiles", "", "")
	resp = &CreateLiveStreamRecordIndexFilesResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateLiveStreamRecordIndexFilesResponse struct {
	responses.BaseResponse
	RequestId  string
	RecordInfo CreateLiveStreamRecordIndexFilesRecordInfo
}

type CreateLiveStreamRecordIndexFilesRecordInfo struct {
	RecordId   string
	RecordUrl  string
	Duration   float32
	Height     int
	Width      int
	CreateTime string
}
