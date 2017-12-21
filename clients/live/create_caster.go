package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateCasterRequest struct {
	requests.RpcRequest
	CasterTemplate string `position:"Query" name:"CasterTemplate"`
	NormType       int    `position:"Query" name:"NormType"`
	Period         int    `position:"Query" name:"Period"`
	SecurityToken  string `position:"Query" name:"SecurityToken"`
	CasterName     string `position:"Query" name:"CasterName"`
	ClientToken    string `position:"Query" name:"ClientToken"`
	ChargeType     string `position:"Query" name:"ChargeType"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	Version        string `position:"Query" name:"Version"`
}

func (req *CreateCasterRequest) Invoke(client *sdk.Client) (resp *CreateCasterResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "CreateCaster", "", "")
	resp = &CreateCasterResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateCasterResponse struct {
	responses.BaseResponse
	RequestId string
	CasterId  string
}
