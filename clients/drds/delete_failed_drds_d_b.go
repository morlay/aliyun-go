package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteFailedDrdsDBRequest struct {
	requests.RpcRequest
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (req *DeleteFailedDrdsDBRequest) Invoke(client *sdk.Client) (resp *DeleteFailedDrdsDBResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "DeleteFailedDrdsDB", "", "")
	resp = &DeleteFailedDrdsDBResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteFailedDrdsDBResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
}
