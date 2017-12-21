package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetCasterConfigRequest struct {
	UrgentMaterialId string  `position:"Query" name:"UrgentMaterialId"`
	TranscodeConfig  string  `position:"Query" name:"TranscodeConfig"`
	Delay            float32 `position:"Query" name:"Delay"`
	SecurityToken    string  `position:"Query" name:"SecurityToken"`
	CasterName       string  `position:"Query" name:"CasterName"`
	CasterId         string  `position:"Query" name:"CasterId"`
	DomainName       string  `position:"Query" name:"DomainName"`
	OwnerId          int64   `position:"Query" name:"OwnerId"`
	Version          string  `position:"Query" name:"Version"`
	RecordConfig     string  `position:"Query" name:"RecordConfig"`
}

func (r SetCasterConfigRequest) Invoke(client *sdk.Client) (response *SetCasterConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetCasterConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "SetCasterConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SetCasterConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetCasterConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetCasterConfigResponse struct {
	RequestId string
	CasterId  string
}
