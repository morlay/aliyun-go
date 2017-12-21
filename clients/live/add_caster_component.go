package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddCasterComponentRequest struct {
	ImageLayerContent string `position:"Query" name:"ImageLayerContent"`
	CasterId          string `position:"Query" name:"CasterId"`
	ComponentLayer    string `position:"Query" name:"ComponentLayer"`
	ComponentName     string `position:"Query" name:"ComponentName"`
	OwnerId           int64  `position:"Query" name:"OwnerId"`
	Version           string `position:"Query" name:"Version"`
	ComponentType     string `position:"Query" name:"ComponentType"`
	SecurityToken     string `position:"Query" name:"SecurityToken"`
	LocationId        string `position:"Query" name:"LocationId"`
	Effect            string `position:"Query" name:"Effect"`
	TextLayerContent  string `position:"Query" name:"TextLayerContent"`
}

func (r AddCasterComponentRequest) Invoke(client *sdk.Client) (response *AddCasterComponentResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddCasterComponentRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "AddCasterComponent", "", "")

	resp := struct {
		*responses.BaseResponse
		AddCasterComponentResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddCasterComponentResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddCasterComponentResponse struct {
	RequestId   string
	ComponentId string
}
