package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddLiveRecordVodConfigRequest struct {
	requests.RpcRequest
	AppName                    string `position:"Query" name:"AppName"`
	AutoCompose                string `position:"Query" name:"AutoCompose"`
	DomainName                 string `position:"Query" name:"DomainName"`
	CycleDuration              int    `position:"Query" name:"CycleDuration"`
	OwnerId                    int64  `position:"Query" name:"OwnerId"`
	ComposeVodTranscodeGroupId string `position:"Query" name:"ComposeVodTranscodeGroupId"`
	StreamName                 string `position:"Query" name:"StreamName"`
	VodTranscodeGroupId        string `position:"Query" name:"VodTranscodeGroupId"`
}

func (req *AddLiveRecordVodConfigRequest) Invoke(client *sdk.Client) (resp *AddLiveRecordVodConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "AddLiveRecordVodConfig", "live", "")
	resp = &AddLiveRecordVodConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddLiveRecordVodConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
