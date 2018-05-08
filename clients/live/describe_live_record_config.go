package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLiveRecordConfigRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int    `position:"Query" name:"PageSize"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	PageNum       int    `position:"Query" name:"PageNum"`
	StreamName    string `position:"Query" name:"StreamName"`
	Order         string `position:"Query" name:"Order"`
}

func (req *DescribeLiveRecordConfigRequest) Invoke(client *sdk.Client) (resp *DescribeLiveRecordConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveRecordConfig", "live", "")
	resp = &DescribeLiveRecordConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveRecordConfigResponse struct {
	responses.BaseResponse
	RequestId         common.String
	PageNum           common.Integer
	PageSize          common.Integer
	Order             common.String
	TotalNum          common.Integer
	TotalPage         common.Integer
	LiveAppRecordList DescribeLiveRecordConfigLiveAppRecordList
}

type DescribeLiveRecordConfigLiveAppRecord struct {
	DomainName       common.String
	AppName          common.String
	StreamName       common.String
	OssEndpoint      common.String
	OssBucket        common.String
	CreateTime       common.String
	StartTime        common.String
	EndTime          common.String
	OnDemond         common.Integer
	RecordFormatList DescribeLiveRecordConfigRecordFormatList
}

type DescribeLiveRecordConfigRecordFormat struct {
	Format               common.String
	OssObjectPrefix      common.String
	SliceOssObjectPrefix common.String
	CycleDuration        common.Integer
}

type DescribeLiveRecordConfigLiveAppRecordList []DescribeLiveRecordConfigLiveAppRecord

func (list *DescribeLiveRecordConfigLiveAppRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveRecordConfigLiveAppRecord)
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

type DescribeLiveRecordConfigRecordFormatList []DescribeLiveRecordConfigRecordFormat

func (list *DescribeLiveRecordConfigRecordFormatList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveRecordConfigRecordFormat)
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
