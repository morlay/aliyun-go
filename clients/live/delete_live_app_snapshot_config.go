package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteLiveAppSnapshotConfigRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteLiveAppSnapshotConfigRequest) Invoke(client *sdk.Client) (response *DeleteLiveAppSnapshotConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteLiveAppSnapshotConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DeleteLiveAppSnapshotConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteLiveAppSnapshotConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteLiveAppSnapshotConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteLiveAppSnapshotConfigResponse struct {
	RequestId string
}
