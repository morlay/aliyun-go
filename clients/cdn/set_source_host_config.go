package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetSourceHostConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Enable        string `position:"Query" name:"Enable"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	BackSrcDomain string `position:"Query" name:"BackSrcDomain"`
}

func (req *SetSourceHostConfigRequest) Invoke(client *sdk.Client) (resp *SetSourceHostConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetSourceHostConfig", "", "")
	resp = &SetSourceHostConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetSourceHostConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
