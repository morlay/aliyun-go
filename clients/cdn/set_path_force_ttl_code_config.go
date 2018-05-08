package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetPathForceTtlCodeConfigRequest struct {
	requests.RpcRequest
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	CodeString    string `position:"Query" name:"CodeString"`
	Path          string `position:"Query" name:"Path"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

func (req *SetPathForceTtlCodeConfigRequest) Invoke(client *sdk.Client) (resp *SetPathForceTtlCodeConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetPathForceTtlCodeConfig", "", "")
	resp = &SetPathForceTtlCodeConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetPathForceTtlCodeConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
