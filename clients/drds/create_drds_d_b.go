package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateDrdsDBRequest struct {
	Encode         string `position:"Query" name:"Encode"`
	Password       string `position:"Query" name:"Password"`
	DbName         string `position:"Query" name:"DbName"`
	RdsInstances   string `position:"Query" name:"RdsInstances"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (r CreateDrdsDBRequest) Invoke(client *sdk.Client) (response *CreateDrdsDBResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateDrdsDBRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Drds", "2017-10-16", "CreateDrdsDB", "", "")

	resp := struct {
		*responses.BaseResponse
		CreateDrdsDBResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateDrdsDBResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateDrdsDBResponse struct {
	RequestId string
	Success   bool
}
