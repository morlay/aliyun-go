package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetFileTypeForceTtlCodeConfigRequest struct {
	requests.RpcRequest
	FileType      string `position:"Query" name:"FileType"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	CodeString    string `position:"Query" name:"CodeString"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

func (req *SetFileTypeForceTtlCodeConfigRequest) Invoke(client *sdk.Client) (resp *SetFileTypeForceTtlCodeConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetFileTypeForceTtlCodeConfig", "", "")
	resp = &SetFileTypeForceTtlCodeConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetFileTypeForceTtlCodeConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
