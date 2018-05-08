package crm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	ResultCode    common.String
	ResultMessage common.String
	Data          FindContacterData
}

type FindContacterData struct {
	ContacterId       common.Long
	KpId              common.Long
	CustomerId        common.Long
	ContacterName     common.String
	ContacterEmail    common.String
	ContacterMobile   common.String
	ContacterPosition common.String
}
