package teslamaxcompute

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetClusterInstanceRequest struct {
	requests.RpcRequest
	Cluster  string `position:"Query" name:"Cluster"`
	PageSize int    `position:"Query" name:"PageSize"`
	PageNum  int    `position:"Query" name:"PageNum"`
	Region   string `position:"Query" name:"Region"`
	Status   string `position:"Query" name:"Status"`
}

func (req *GetClusterInstanceRequest) Invoke(client *sdk.Client) (resp *GetClusterInstanceResponse, err error) {
	req.InitWithApiInfo("TeslaMaxCompute", "2018-01-04", "GetClusterInstance", "", "")
	resp = &GetClusterInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetClusterInstanceResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      GetClusterInstanceData
}

type GetClusterInstanceData struct {
	Total  common.Integer
	Detail GetClusterInstanceInstanceList
}

type GetClusterInstanceInstance struct {
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

type GetClusterInstanceInstanceList []GetClusterInstanceInstance

func (list *GetClusterInstanceInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetClusterInstanceInstance)
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
