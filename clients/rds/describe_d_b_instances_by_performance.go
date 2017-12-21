package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDBInstancesByPerformanceRequest struct {
	Tag4value            string `position:"Query" name:"Tag.4.value"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Tag2key              string `position:"Query" name:"Tag.2.key"`
	Tag5key              string `position:"Query" name:"Tag.5.key"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Tag3key              string `position:"Query" name:"Tag.3.key"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tag5value            string `position:"Query" name:"Tag.5.value"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	Tags                 string `position:"Query" name:"Tags"`
	Tag1key              string `position:"Query" name:"Tag.1.key"`
	Tag1value            string `position:"Query" name:"Tag.1.value"`
	SortKey              string `position:"Query" name:"SortKey"`
	SortMethod           string `position:"Query" name:"SortMethod"`
	Tag2value            string `position:"Query" name:"Tag.2.value"`
	PageSize             int    `position:"Query" name:"PageSize"`
	Tag4key              string `position:"Query" name:"Tag.4.key"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	Tag3value            string `position:"Query" name:"Tag.3.value"`
	ProxyId              string `position:"Query" name:"ProxyId"`
}

func (r DescribeDBInstancesByPerformanceRequest) Invoke(client *sdk.Client) (response *DescribeDBInstancesByPerformanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDBInstancesByPerformanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstancesByPerformance", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDBInstancesByPerformanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDBInstancesByPerformanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDBInstancesByPerformanceResponse struct {
	RequestId        string
	PageNumber       int
	TotalRecordCount int
	PageRecordCount  int
	Items            DescribeDBInstancesByPerformanceDBInstancePerformanceList
}

type DescribeDBInstancesByPerformanceDBInstancePerformance struct {
	CPUUsage              string
	IOPSUsage             string
	DiskUsage             string
	SessionUsage          string
	DBInstanceId          string
	DBInstanceDescription string
}

type DescribeDBInstancesByPerformanceDBInstancePerformanceList []DescribeDBInstancesByPerformanceDBInstancePerformance

func (list *DescribeDBInstancesByPerformanceDBInstancePerformanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancesByPerformanceDBInstancePerformance)
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
