package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddLiveAppSnapshotConfigRequest struct {
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

func (r AddLiveAppSnapshotConfigRequest) Invoke(client *sdk.Client) (response *AddLiveAppSnapshotConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddLiveAppSnapshotConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "AddLiveAppSnapshotConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		AddLiveAppSnapshotConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddLiveAppSnapshotConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddLiveAppSnapshotConfigResponse struct {
	RequestId string
}
