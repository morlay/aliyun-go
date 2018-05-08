package teslamaxcompute

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetQuotaInstanceRequest struct {
	requests.RpcRequest
	Cluster   string `position:"Query" name:"Cluster"`
	PageSize  int    `position:"Query" name:"PageSize"`
	QuotaId   string `position:"Query" name:"QuotaId"`
	PageNum   int    `position:"Query" name:"PageNum"`
	Region    string `position:"Query" name:"Region"`
	QuotaName string `position:"Query" name:"QuotaName"`
	Status    string `position:"Query" name:"Status"`
}

func (req *GetQuotaInstanceRequest) Invoke(client *sdk.Client) (resp *GetQuotaInstanceResponse, err error) {
	req.InitWithApiInfo("TeslaMaxCompute", "2018-01-04", "GetQuotaInstance", "", "")
	resp = &GetQuotaInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetQuotaInstanceResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      GetQuotaInstanceData
}

type GetQuotaInstanceData struct {
	Total  common.Integer
	Detail GetQuotaInstanceInstanceList
}

type GetQuotaInstanceInstance struct {
	Project         common.String
	InstanceId      common.String
	Status          common.String
	UserAccount     common.String
	NickName        common.String
	Cluster         common.String
	RunTime         common.String
	CpuUsed         common.Long
	CpuRequest      common.Long
	CpuUsedTotal    common.Long
	CpuUsedRatioMax common.Float
	CpuUsedRatioMin common.Float
	MemUsed         common.Long
	MemRequest      common.Long
	MemUsedTotal    common.Long
	MemUsedRatioMax common.Float
	MemUsedRatioMin common.Float
	TaskType        common.String
	SkynetId        common.String
	QuotaName       common.String
	QuotaId         common.Integer
}

type GetQuotaInstanceInstanceList []GetQuotaInstanceInstance

func (list *GetQuotaInstanceInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetQuotaInstanceInstance)
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
