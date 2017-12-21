package crm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FindContacterTestRequest struct {
	ContacterId int64 `position:"Query" name:"ContacterId"`
}

func (r FindContacterTestRequest) Invoke(client *sdk.Client) (response *FindContacterTestResponse, err error) {
	req := struct {
		*requests.RpcRequest
		FindContacterTestRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Crm", "2015-03-24", "FindContacterTest", "", "")

	resp := struct {
		*responses.BaseResponse
		FindContacterTestResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.FindContacterTestResponse

	err = client.DoAction(&req, &resp)
	return
}

type FindContacterTestResponse struct {
	Success       bool
	ResultCode    string
	ResultMessage string
	Data          FindContacterTestData
}

type FindContacterTestData struct {
	ContacterId       int64
	KpId              int64
	CustomerId        int64
	ContacterName     string
	ContacterEmail    string
	ContacterMobile   string
	ContacterPosition string
}
