package drds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeRdsListRequest struct {
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (r DescribeRdsListRequest) Invoke(client *sdk.Client) (response *DescribeRdsListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeRdsListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Drds", "2017-10-16", "DescribeRdsList", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeRdsListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeRdsListResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeRdsListResponse struct {
	RequestId string
	Success   bool
	Data      DescribeRdsListRdsInstanceList
}

type DescribeRdsListRdsInstance struct {
	InstanceId       int
	InstanceName     string
	ConnectUrl       string
	Port             int
	InstanceStatus   string
	DbType           string
	ReadWeight       int
	ReadOnlyChildren DescribeRdsListChildList
}

type DescribeRdsListChild struct {
	InstanceId     string
	InstanceName   string
	ConnectUrl     string
	Port           int
	InstanceStatus string
	DbType         string
	ReadWeight     int
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
