package crm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateContacterRequest struct {
	KpId              int64  `position:"Query" name:"KpId"`
	ContacterName     string `position:"Query" name:"ContacterName"`
	ContacterEmail    string `position:"Query" name:"ContacterEmail"`
	ContacterMobile   string `position:"Query" name:"ContacterMobile"`
	ContacterPosition string `position:"Query" name:"ContacterPosition"`
}

func (r CreateContacterRequest) Invoke(client *sdk.Client) (response *CreateContacterResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateContacterRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Crm", "2015-03-24", "CreateContacter", "", "")

	resp := struct {
		*responses.BaseResponse
		CreateContacterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateContacterResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateContacterResponse struct {
	Success       bool
	ResultCode    string
	ResultMessage string
	Data          CreateContacterData
}

type CreateContacterData struct {
	ContacterId int64
}
