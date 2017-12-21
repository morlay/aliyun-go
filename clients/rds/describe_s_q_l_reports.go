package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSQLReportsRequest struct {
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

func (r DescribeSQLReportsRequest) Invoke(client *sdk.Client) (response *DescribeSQLReportsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeSQLReportsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeSQLReports", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeSQLReportsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeSQLReportsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeSQLReportsResponse struct {
	RequestId        string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeSQLReportsItemList
}

type DescribeSQLReportsItem struct {
	ReportTime       string
	LatencyTopNItems DescribeSQLReportsLatencyTopNItemList
	QPSTopNItems     DescribeSQLReportsQPSTopNItemList
}

type DescribeSQLReportsLatencyTopNItem struct {
	SQLText         string
	AvgLatency      int64
	SQLExecuteTimes int64
}

type DescribeSQLReportsQPSTopNItem struct {
	SQLText         string
	SQLExecuteTimes int64
}

type DescribeSQLReportsItemList []DescribeSQLReportsItem

func (list *DescribeSQLReportsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLReportsItem)
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

type DescribeSQLReportsLatencyTopNItemList []DescribeSQLReportsLatencyTopNItem

func (list *DescribeSQLReportsLatencyTopNItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLReportsLatencyTopNItem)
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

type DescribeSQLReportsQPSTopNItemList []DescribeSQLReportsQPSTopNItem

func (list *DescribeSQLReportsQPSTopNItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLReportsQPSTopNItem)
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
