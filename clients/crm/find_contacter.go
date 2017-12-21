package crm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FindContacterRequest struct {
	requests.RpcRequest
	ContacterId int64 `position:"Query" name:"ContacterId"`
}

func (req *FindContacterRequest) Invoke(client *sdk.Client) (resp *FindContacterResponse, err error) {
	req.InitWithApiInfo("Crm", "2015-03-24", "FindContacter", "", "")
	resp = &FindContacterResponse{}
	err = client.DoAction(req, resp)
	return
}

type FindContacterResponse struct {
	responses.BaseResponse
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
