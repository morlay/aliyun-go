package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDrdsDBRequest struct {
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (r DescribeDrdsDBRequest) Invoke(client *sdk.Client) (response *DescribeDrdsDBResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDrdsDBRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Drds", "2017-10-16", "DescribeDrdsDB", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDrdsDBResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDrdsDBResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDrdsDBResponse struct {
	RequestId string
	Success   bool
	Data      DescribeDrdsDBData
}

type DescribeDrdsDBData struct {
	DbName     string
	Status     int
	CreateTime string
	Msg        string
	Mode       string
}
