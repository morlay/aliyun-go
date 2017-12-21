package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetPageCompressConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Enable        string `position:"Query" name:"Enable"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *SetPageCompressConfigRequest) Invoke(client *sdk.Client) (resp *SetPageCompressConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetPageCompressConfig", "", "")
	resp = &SetPageCompressConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetPageCompressConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
