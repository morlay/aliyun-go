package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteDrdsDBRequest struct {
	requests.RpcRequest
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (req *DeleteDrdsDBRequest) Invoke(client *sdk.Client) (resp *DeleteDrdsDBResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "DeleteDrdsDB", "", "")
	resp = &DeleteDrdsDBResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteDrdsDBResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
}
