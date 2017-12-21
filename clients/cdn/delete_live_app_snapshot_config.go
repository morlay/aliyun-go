package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteLiveAppSnapshotConfigRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteLiveAppSnapshotConfigRequest) Invoke(client *sdk.Client) (resp *DeleteLiveAppSnapshotConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DeleteLiveAppSnapshotConfig", "", "")
	resp = &DeleteLiveAppSnapshotConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteLiveAppSnapshotConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
