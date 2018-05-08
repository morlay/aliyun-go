package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateDrdsDBRequest struct {
	requests.RpcRequest
	Encode         string `position:"Query" name:"Encode"`
	Password       string `position:"Query" name:"Password"`
	DbName         string `position:"Query" name:"DbName"`
	RdsInstances   string `position:"Query" name:"RdsInstances"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (req *CreateDrdsDBRequest) Invoke(client *sdk.Client) (resp *CreateDrdsDBResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "CreateDrdsDB", "", "")
	resp = &CreateDrdsDBResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateDrdsDBResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
}
