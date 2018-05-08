package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeSQLLogReportsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	EndTime              string `position:"Query" name:"EndTime"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeSQLLogReportsRequest) Invoke(client *sdk.Client) (resp *DescribeSQLLogReportsResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeSQLLogReports", "rds", "")
	resp = &DescribeSQLLogReportsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSQLLogReportsResponse struct {
	responses.BaseResponse
	RequestId        common.String
	TotalRecordCount common.Integer
	PageNumber       common.Integer
	PageRecordCount  common.Integer
	Items            DescribeSQLLogReportsItemList
}

type DescribeSQLLogReportsItem struct {
	ReportTime       common.String
	LatencyTopNItems DescribeSQLLogReportsLatencyTopNItemList
	QPSTopNItems     DescribeSQLLogReportsQPSTopNItemList
}

type DescribeSQLLogReportsLatencyTopNItem struct {
	SQLText         common.String
	AvgLatency      common.Long
	SQLExecuteTimes common.Long
}

type DescribeSQLLogReportsQPSTopNItem struct {
	SQLText         common.String
	SQLExecuteTimes common.Long
}

type DescribeSQLLogReportsItemList []DescribeSQLLogReportsItem

func (list *DescribeSQLLogReportsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLLogReportsItem)
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

type DescribeSQLLogReportsLatencyTopNItemList []DescribeSQLLogReportsLatencyTopNItem

func (list *DescribeSQLLogReportsLatencyTopNItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLLogReportsLatencyTopNItem)
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

type DescribeSQLLogReportsQPSTopNItemList []DescribeSQLLogReportsQPSTopNItem

func (list *DescribeSQLLogReportsQPSTopNItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLLogReportsQPSTopNItem)
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
