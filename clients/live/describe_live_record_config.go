package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveRecordConfigRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int    `position:"Query" name:"PageSize"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	PageNum       int    `position:"Query" name:"PageNum"`
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
	RequestId         string
	PageNum           int
	PageSize          int
	Order             string
	TotalNum          int
	TotalPage         int
	LiveAppRecordList DescribeLiveRecordConfigLiveAppRecordList
}

type DescribeLiveRecordConfigLiveAppRecord struct {
	DomainName       string
	AppName          string
	OssEndpoint      string
	OssBucket        string
	CreateTime       string
	RecordFormatList DescribeLiveRecordConfigRecordFormatList
}

type DescribeLiveRecordConfigRecordFormat struct {
	Format               string
	OssObjectPrefix      string
	SliceOssObjectPrefix string
	CycleDuration        int
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
