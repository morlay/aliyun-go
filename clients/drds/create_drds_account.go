package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateDrdsAccountRequest struct {
	requests.RpcRequest
	Password       string `position:"Query" name:"Password"`
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	UserName       string `position:"Query" name:"UserName"`
}

func (req *CreateDrdsAccountRequest) Invoke(client *sdk.Client) (resp *CreateDrdsAccountResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "CreateDrdsAccount", "", "")
	resp = &CreateDrdsAccountResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateDrdsAccountResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
}
