package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLogicDBInstanceTopologyRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLogicDBInstanceTopologyRequest) Invoke(client *sdk.Client) (resp *DescribeLogicDBInstanceTopologyResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeLogicDBInstanceTopology", "rds", "")
	resp = &DescribeLogicDBInstanceTopologyResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLogicDBInstanceTopologyResponse struct {
	responses.BaseResponse
	DBInstanceId          int
	DBInstanceName        string
	DBInstanceStatus      int
	DBInstanceStatusDesc  string
	DBInstanceConnType    string
	DBInstanceDescription string
	Engine                string
	EngineVersion         string
	Items                 DescribeLogicDBInstanceTopologyLogicDBInstanceParameterList
}

type DescribeLogicDBInstanceTopologyLogicDBInstanceParameter struct {
	DBInstanceID          int
	DBInstanceName        string
	DBInstanceStatus      int
	DBInstanceStatusDesc  string
	DBInstanceConnType    string
	DBInstanceDescription string
	Engine                string
	EngineVersion         string
	CharacterType         string
}

type DescribeLogicDBInstanceTopologyLogicDBInstanceParameterList []DescribeLogicDBInstanceTopologyLogicDBInstanceParameter

func (list *DescribeLogicDBInstanceTopologyLogicDBInstanceParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLogicDBInstanceTopologyLogicDBInstanceParameter)
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
