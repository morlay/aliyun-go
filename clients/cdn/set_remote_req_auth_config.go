package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetRemoteReqAuthConfigRequest struct {
	requests.RpcRequest
	AuthPath      string `position:"Query" name:"AuthPath"`
	DomainName    string `position:"Query" name:"DomainName"`
	AuthEnable    string `position:"Query" name:"AuthEnable"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	TimeOut       string `position:"Query" name:"TimeOut"`
	AuthType      string `position:"Query" name:"AuthType"`
	AuthProvider  string `position:"Query" name:"AuthProvider"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	AuthCrash     string `position:"Query" name:"AuthCrash"`
	AuthAddr      string `position:"Query" name:"AuthAddr"`
	AuthFileType  string `position:"Query" name:"AuthFileType"`
}

func (req *SetRemoteReqAuthConfigRequest) Invoke(client *sdk.Client) (resp *SetRemoteReqAuthConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetRemoteReqAuthConfig", "", "")
	resp = &SetRemoteReqAuthConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetRemoteReqAuthConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
