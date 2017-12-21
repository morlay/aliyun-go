package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDrdsDBRequest struct {
	requests.RpcRequest
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (req *DescribeDrdsDBRequest) Invoke(client *sdk.Client) (resp *DescribeDrdsDBResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "DescribeDrdsDB", "", "")
	resp = &DescribeDrdsDBResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDrdsDBResponse struct {
	responses.BaseResponse
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
