package arms4finance

import (
	"encoding/json"

	"github.com/morlay/aliyun-go/core"
)

func (c *Arms4financeClient) WhereInDimQuery(req *WhereInDimQueryArgs) (resp *WhereInDimQueryResponse, err error) {
	resp = &WhereInDimQueryResponse{}
	err = c.Request("WhereInDimQuery", req, resp)
	return
}

type WhereInDimQueryDimensions struct {
	Key   string
	Value string
}
type WhereInDimQueryArgs struct {
	WhereInKey     string
	Measuress      WhereInDimQueryMeasuresList
	IntervalInSec  int
	DateStr        string
	IsDrillDown    core.Bool
	MinTime        int64
	DatasetId      int64
	WhereInValuess WhereInDimQueryWhereInValuesList
	MaxTime        int64
	Dimensionss    WhereInDimQueryDimensionsList
}
type WhereInDimQueryResponse struct {
	Data string
}

type WhereInDimQueryMeasuresList []string

func (list *WhereInDimQueryMeasuresList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type WhereInDimQueryWhereInValuesList []string

func (list *WhereInDimQueryWhereInValuesList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type WhereInDimQueryDimensionsList []WhereInDimQueryDimensions

func (list *WhereInDimQueryDimensionsList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]WhereInDimQueryDimensions)
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

func (c *Arms4financeClient) ARMSQueryDataSet(req *ARMSQueryDataSetArgs) (resp *ARMSQueryDataSetResponse, err error) {
	resp = &ARMSQueryDataSetResponse{}
	err = c.Request("ARMSQueryDataSet", req, resp)
	return
}

type ARMSQueryDataSetDimensions struct {
	Key   string
	Value string
}
type ARMSQueryDataSetArgs struct {
	Measuress     ARMSQueryDataSetMeasuresList
	IntervalInSec int
	DateStr       string
	IsDrillDown   core.Bool
	MinTime       int64
	DatasetId     int64
	MaxTime       int64
	Dimensionss   ARMSQueryDataSetDimensionsList
}
type ARMSQueryDataSetResponse struct {
	Data string
}

type ARMSQueryDataSetMeasuresList []string

func (list *ARMSQueryDataSetMeasuresList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type ARMSQueryDataSetDimensionsList []ARMSQueryDataSetDimensions

func (list *ARMSQueryDataSetDimensionsList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ARMSQueryDataSetDimensions)
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
