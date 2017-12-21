package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteFailedDrdsDBRequest struct {
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (r DeleteFailedDrdsDBRequest) Invoke(client *sdk.Client) (response *DeleteFailedDrdsDBResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteFailedDrdsDBRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Drds", "2017-10-16", "DeleteFailedDrdsDB", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteFailedDrdsDBResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteFailedDrdsDBResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteFailedDrdsDBResponse struct {
	RequestId string
	Success   bool
}
