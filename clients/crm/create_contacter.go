package crm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateContacterRequest struct {
	requests.RpcRequest
	KpId              int64  `position:"Query" name:"KpId"`
	ContacterName     string `position:"Query" name:"ContacterName"`
	ContacterEmail    string `position:"Query" name:"ContacterEmail"`
	ContacterMobile   string `position:"Query" name:"ContacterMobile"`
	ContacterPosition string `position:"Query" name:"ContacterPosition"`
}

func (req *CreateContacterRequest) Invoke(client *sdk.Client) (resp *CreateContacterResponse, err error) {
	req.InitWithApiInfo("Crm", "2015-03-24", "CreateContacter", "", "")
	resp = &CreateContacterResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateContacterResponse struct {
	responses.BaseResponse
	Success       bool
	ResultCode    string
	ResultMessage string
	Data          CreateContacterData
}

type CreateContacterData struct {
	ContacterId int64
}
