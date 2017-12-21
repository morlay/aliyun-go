package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyCasterComponentRequest struct {
	ComponentId       string `position:"Query" name:"ComponentId"`
	ImageLayerContent string `position:"Query" name:"ImageLayerContent"`
	CasterId          string `position:"Query" name:"CasterId"`
	ComponentLayer    string `position:"Query" name:"ComponentLayer"`
	ComponentName     string `position:"Query" name:"ComponentName"`
	OwnerId           int64  `position:"Query" name:"OwnerId"`
	Version           string `position:"Query" name:"Version"`
	ComponentType     string `position:"Query" name:"ComponentType"`
	SecurityToken     string `position:"Query" name:"SecurityToken"`
	Effect            string `position:"Query" name:"Effect"`
	TextLayerContent  string `position:"Query" name:"TextLayerContent"`
}

func (r ModifyCasterComponentRequest) Invoke(client *sdk.Client) (response *ModifyCasterComponentResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyCasterComponentRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "ModifyCasterComponent", "", "")

	resp := struct {
		*responses.BaseResponse
		ModifyCasterComponentResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyCasterComponentResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyCasterComponentResponse struct {
	RequestId   string
	ComponentId string
}
