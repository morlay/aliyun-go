package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteDrdsDBRequest struct {
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (r DeleteDrdsDBRequest) Invoke(client *sdk.Client) (response *DeleteDrdsDBResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteDrdsDBRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Drds", "2017-10-16", "DeleteDrdsDB", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteDrdsDBResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteDrdsDBResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteDrdsDBResponse struct {
	RequestId string
	Success   bool
}
