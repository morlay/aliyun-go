package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddLiveAppSnapshotConfigRequest struct {
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

func (req *AddLiveAppSnapshotConfigRequest) Invoke(client *sdk.Client) (resp *AddLiveAppSnapshotConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "AddLiveAppSnapshotConfig", "", "")
	resp = &AddLiveAppSnapshotConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddLiveAppSnapshotConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
