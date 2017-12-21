package crm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyContacterRequest struct {
	ContacterId       int64  `position:"Query" name:"ContacterId"`
	ContacterName     string `position:"Query" name:"ContacterName"`
	ContacterEmail    string `position:"Query" name:"ContacterEmail"`
	ContacterMobile   string `position:"Query" name:"ContacterMobile"`
	ContacterPosition string `position:"Query" name:"ContacterPosition"`
}

func (r ModifyContacterRequest) Invoke(client *sdk.Client) (response *ModifyContacterResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyContacterRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Crm", "2015-03-24", "ModifyContacter", "", "")

	resp := struct {
		*responses.BaseResponse
		ModifyContacterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyContacterResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyContacterResponse struct {
	Success       bool
	ResultCode    string
	ResultMessage string
}
