package teslamaxcompute

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetUserInstanceRequest struct {
	requests.RpcRequest
	PageSize int    `position:"Query" name:"PageSize"`
	PageNum  int    `position:"Query" name:"PageNum"`
	Region   string `position:"Query" name:"Region"`
	User     string `position:"Query" name:"User"`
	Status   string `position:"Query" name:"Status"`
}

func (req *GetUserInstanceRequest) Invoke(client *sdk.Client) (resp *GetUserInstanceResponse, err error) {
	req.InitWithApiInfo("TeslaMaxCompute", "2018-01-04", "GetUserInstance", "", "")
	resp = &GetUserInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetUserInstanceResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      GetUserInstanceData
}

type GetUserInstanceData struct {
	Total  common.Integer
	Detail GetUserInstanceInstanceList
}

type GetUserInstanceInstance struct {
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

type GetUserInstanceInstanceList []GetUserInstanceInstance

func (list *GetUserInstanceInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetUserInstanceInstance)
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
