package drds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDrdsDBsRequest struct {
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (r DescribeDrdsDBsRequest) Invoke(client *sdk.Client) (response *DescribeDrdsDBsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDrdsDBsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Drds", "2017-10-16", "DescribeDrdsDBs", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDrdsDBsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDrdsDBsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDrdsDBsResponse struct {
	RequestId string
	Success   bool
	Data      DescribeDrdsDBsDbList
}

type DescribeDrdsDBsDb struct {
	DbName     string
	Status     int
	CreateTime string
	Msg        string
	Mode       string
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
