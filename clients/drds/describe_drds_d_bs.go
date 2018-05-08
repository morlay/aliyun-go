package drds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDrdsDBsRequest struct {
	requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (req *DescribeDrdsDBsRequest) Invoke(client *sdk.Client) (resp *DescribeDrdsDBsResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "DescribeDrdsDBs", "", "")
	resp = &DescribeDrdsDBsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDrdsDBsResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Data      DescribeDrdsDBsDbList
}

type DescribeDrdsDBsDb struct {
	DbName     common.String
	Status     common.Integer
	CreateTime common.String
	Msg        common.String
	Mode       common.String
}

type DescribeDrdsDBsDbList []DescribeDrdsDBsDb

func (list *DescribeDrdsDBsDbList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDrdsDBsDb)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
