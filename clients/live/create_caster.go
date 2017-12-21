package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateCasterRequest struct {
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

func (r CreateCasterRequest) Invoke(client *sdk.Client) (response *CreateCasterResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateCasterRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "CreateCaster", "", "")

	resp := struct {
		*responses.BaseResponse
		CreateCasterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateCasterResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateCasterResponse struct {
	RequestId string
	CasterId  string
}
