package afs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribePersonMachineListRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
}

func (req *DescribePersonMachineListRequest) Invoke(client *sdk.Client) (resp *DescribePersonMachineListResponse, err error) {
	req.InitWithApiInfo("afs", "2018-01-12", "DescribePersonMachineList", "", "")
	resp = &DescribePersonMachineListResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribePersonMachineListResponse struct {
	responses.BaseResponse
	RequestId        string
	BizCode          string
	PersonMachineRes DescribePersonMachineListPersonMachineRes
}

type DescribePersonMachineListPersonMachineRes struct {
	HasConfiguration string
	PersonMachines   DescribePersonMachineListPersonMachineList
}

type DescribePersonMachineListPersonMachine struct {
	ConfigurationName   string
	Appkey              string
	ConfigurationMethod string
	ApplyType           string
	Scene               string
	LastUpdate          string
}

type DescribePersonMachineListPersonMachineList []DescribePersonMachineListPersonMachine

func (list *DescribePersonMachineListPersonMachineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePersonMachineListPersonMachine)
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
