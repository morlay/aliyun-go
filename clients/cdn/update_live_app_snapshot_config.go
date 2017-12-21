package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateLiveAppSnapshotConfigRequest struct {
	requests.RpcRequest
	TimeInterval       int    `position:"Query" name:"TimeInterval"`
	OssBucket          string `position:"Query" name:"OssBucket"`
	AppName            string `position:"Query" name:"AppName"`
	SecurityToken      string `position:"Query" name:"SecurityToken"`
	DomainName         string `position:"Query" name:"DomainName"`
	OssEndpoint        string `position:"Query" name:"OssEndpoint"`
	SequenceOssObject  string `position:"Query" name:"SequenceOssObject"`
	OverwriteOssObject string `position:"Query" name:"OverwriteOssObject"`
	OwnerId            int64  `position:"Query" name:"OwnerId"`
}

func (req *UpdateLiveAppSnapshotConfigRequest) Invoke(client *sdk.Client) (resp *UpdateLiveAppSnapshotConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "UpdateLiveAppSnapshotConfig", "", "")
	resp = &UpdateLiveAppSnapshotConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateLiveAppSnapshotConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
