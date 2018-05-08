package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId        common.String
	TotalRecordCount common.Integer
	PageNumber       common.Integer
	PageRecordCount  common.Integer
	Items            DescribeSQLLogReportListItemList
}

type DescribeSQLLogReportListItem struct {
	ReportTime       common.String
	LatencyTopNItems DescribeSQLLogReportListLatencyTopNItemList
	QPSTopNItems     DescribeSQLLogReportListQPSTopNItemList
}

type DescribeSQLLogReportListLatencyTopNItem struct {
	SQLText         common.String
	AvgLatency      common.Long
	SQLExecuteTimes common.Long
}

type DescribeSQLLogReportListQPSTopNItem struct {
	SQLText         common.String
	SQLExecuteTimes common.Long
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
