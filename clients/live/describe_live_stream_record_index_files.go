package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLiveStreamRecordIndexFilesRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int    `position:"Query" name:"PageSize"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	PageNum       int    `position:"Query" name:"PageNum"`
	StreamName    string `position:"Query" name:"StreamName"`
	Order         string `position:"Query" name:"Order"`
}

func (req *DescribeLiveStreamRecordIndexFilesRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamRecordIndexFilesResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStreamRecordIndexFiles", "live", "")
	resp = &DescribeLiveStreamRecordIndexFilesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamRecordIndexFilesResponse struct {
	responses.BaseResponse
	RequestId           common.String
	PageNum             common.Integer
	PageSize            common.Integer
	Order               common.String
	TotalNum            common.Integer
	TotalPage           common.Integer
	RecordIndexInfoList DescribeLiveStreamRecordIndexFilesRecordIndexInfoList
}

type DescribeLiveStreamRecordIndexFilesRecordIndexInfo struct {
	RecordId    common.String
	RecordUrl   common.String
	DomainName  common.String
	AppName     common.String
	StreamName  common.String
	OssBucket   common.String
	OssEndpoint common.String
	OssObject   common.String
	StartTime   common.String
	EndTime     common.String
	Duration    common.Float
	Height      common.Integer
	Width       common.Integer
	CreateTime  common.String
}

type DescribeLiveStreamRecordIndexFilesRecordIndexInfoList []DescribeLiveStreamRecordIndexFilesRecordIndexInfo

func (list *DescribeLiveStreamRecordIndexFilesRecordIndexInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamRecordIndexFilesRecordIndexInfo)
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
