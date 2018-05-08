package drds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeRdsListRequest struct {
	requests.RpcRequest
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (req *DescribeRdsListRequest) Invoke(client *sdk.Client) (resp *DescribeRdsListResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "DescribeRdsList", "", "")
	resp = &DescribeRdsListResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeRdsListResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Data      DescribeRdsListRdsInstanceList
}

type DescribeRdsListRdsInstance struct {
	InstanceId       common.Integer
	InstanceName     common.String
	ConnectUrl       common.String
	Port             common.Integer
	InstanceStatus   common.String
	DbType           common.String
	ReadWeight       common.Integer
	ReadOnlyChildren DescribeRdsListChildList
}

type DescribeRdsListChild struct {
	InstanceId     common.String
	InstanceName   common.String
	ConnectUrl     common.String
	Port           common.Integer
	InstanceStatus common.String
	DbType         common.String
	ReadWeight     common.Integer
}

type DescribeRdsListRdsInstanceList []DescribeRdsListRdsInstance

func (list *DescribeRdsListRdsInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRdsListRdsInstance)
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

type DescribeRdsListChildList []DescribeRdsListChild

func (list *DescribeRdsListChildList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRdsListChild)
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
