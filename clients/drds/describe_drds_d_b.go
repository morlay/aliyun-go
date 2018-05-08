package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Success   bool
	Data      DescribeDrdsDBData
}

type DescribeDrdsDBData struct {
	DbName     common.String
	Status     common.Integer
	CreateTime common.String
	Msg        common.String
	Mode       common.String
}
