package crm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FindContacterRequest struct {
	ContacterId int64 `position:"Query" name:"ContacterId"`
}

func (r FindContacterRequest) Invoke(client *sdk.Client) (response *FindContacterResponse, err error) {
	req := struct {
		*requests.RpcRequest
		FindContacterRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Crm", "2015-03-24", "FindContacter", "", "")

	resp := struct {
		*responses.BaseResponse
		FindContacterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.FindContacterResponse

	err = client.DoAction(&req, &resp)
	return
}

type FindContacterResponse struct {
	Success       bool
	ResultCode    string
	ResultMessage string
	Data          FindContacterData
}

type FindContacterData struct {
	ContacterId       int64
	KpId              int64
	CustomerId        int64
	ContacterName     string
	ContacterEmail    string
	ContacterMobile   string
	ContacterPosition string
}
