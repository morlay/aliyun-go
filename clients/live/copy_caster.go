package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CopyCasterRequest struct {
	requests.RpcRequest
	SrcCasterId   string `position:"Query" name:"SrcCasterId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterName    string `position:"Query" name:"CasterName"`
	ClientToken   string `position:"Query" name:"ClientToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (req *CopyCasterRequest) Invoke(client *sdk.Client) (resp *CopyCasterResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "CopyCaster", "", "")
	resp = &CopyCasterResponse{}
	err = client.DoAction(req, resp)
	return
}

type CopyCasterResponse struct {
	responses.BaseResponse
	RequestId string
	CasterId  string
}
