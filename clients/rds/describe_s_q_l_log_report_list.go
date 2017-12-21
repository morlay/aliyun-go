package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeSQLLogReportListRequest struct {
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

func (req *DescribeSQLLogReportListRequest) Invoke(client *sdk.Client) (resp *DescribeSQLLogReportListResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeSQLLogReportList", "rds", "")
	resp = &DescribeSQLLogReportListResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSQLLogReportListResponse struct {
	responses.BaseResponse
	RequestId        string
	TotalRecordCount int
	PageNumber       int
	PageRecordCount  int
	Items            DescribeSQLLogReportListItemList
}

type DescribeSQLLogReportListItem struct {
	ReportTime       string
	LatencyTopNItems DescribeSQLLogReportListLatencyTopNItemList
	QPSTopNItems     DescribeSQLLogReportListQPSTopNItemList
}

type DescribeSQLLogReportListLatencyTopNItem struct {
	SQLText         string
	AvgLatency      int64
	SQLExecuteTimes int64
}

type DescribeSQLLogReportListQPSTopNItem struct {
	SQLText         string
	SQLExecuteTimes int64
}

type DescribeSQLLogReportListItemList []DescribeSQLLogReportListItem

func (list *DescribeSQLLogReportListItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLLogReportListItem)
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

type DescribeSQLLogReportListLatencyTopNItemList []DescribeSQLLogReportListLatencyTopNItem

func (list *DescribeSQLLogReportListLatencyTopNItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLLogReportListLatencyTopNItem)
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

type DescribeSQLLogReportListQPSTopNItemList []DescribeSQLLogReportListQPSTopNItem

func (list *DescribeSQLLogReportListQPSTopNItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSQLLogReportListQPSTopNItem)
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
