package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeParametersRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeParametersRequest) Invoke(client *sdk.Client) (resp *DescribeParametersResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeParameters", "rds", "")
	resp = &DescribeParametersResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeParametersResponse struct {
	responses.BaseResponse
	RequestId         common.String
	Engine            common.String
	EngineVersion     common.String
	ConfigParameters  DescribeParametersDBInstanceParameterList
	RunningParameters DescribeParametersDBInstanceParameterList
}

type DescribeParametersDBInstanceParameter struct {
	ParameterName        common.String
	ParameterValue       common.String
	ParameterDescription common.String
}

type DescribeParametersDBInstanceParameterList []DescribeParametersDBInstanceParameter

func (list *DescribeParametersDBInstanceParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeParametersDBInstanceParameter)
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
