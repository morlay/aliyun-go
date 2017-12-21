package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateLiveAppSnapshotConfigRequest struct {
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

func (r UpdateLiveAppSnapshotConfigRequest) Invoke(client *sdk.Client) (response *UpdateLiveAppSnapshotConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateLiveAppSnapshotConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "UpdateLiveAppSnapshotConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateLiveAppSnapshotConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateLiveAppSnapshotConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateLiveAppSnapshotConfigResponse struct {
	RequestId string
}
