package teslamaxcompute

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetProjectInstanceRequest struct {
	requests.RpcRequest
	PageSize int    `position:"Query" name:"PageSize"`
	Project  string `position:"Query" name:"Project"`
	PageNum  int    `position:"Query" name:"PageNum"`
	Region   string `position:"Query" name:"Region"`
	Status   string `position:"Query" name:"Status"`
}

func (req *GetProjectInstanceRequest) Invoke(client *sdk.Client) (resp *GetProjectInstanceResponse, err error) {
	req.InitWithApiInfo("TeslaMaxCompute", "2018-01-04", "GetProjectInstance", "", "")
	resp = &GetProjectInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetProjectInstanceResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      GetProjectInstanceData
}

type GetProjectInstanceData struct {
	Total  common.Integer
	Detail GetProjectInstanceInstanceList
}

type GetProjectInstanceInstance struct {
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

type GetProjectInstanceInstanceList []GetProjectInstanceInstance

func (list *GetProjectInstanceInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetProjectInstanceInstance)
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
