package cdn

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
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (req *DescribeLiveStreamRecordIndexFilesRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamRecordIndexFilesResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamRecordIndexFiles", "", "")
	resp = &DescribeLiveStreamRecordIndexFilesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamRecordIndexFilesResponse struct {
	responses.BaseResponse
	RequestId           common.String
	RecordIndexInfoList DescribeLiveStreamRecordIndexFilesRecordIndexInfoList
}

type DescribeLiveStreamRecordIndexFilesRecordIndexInfo struct {
	RecordId   common.String
	RecordUrl  common.String
	DomainName common.String
	AppName    common.String
	StreamName common.String
	OssObject  common.String
	StartTime  common.String
	EndTime    common.String
	Duration   common.Float
	Height     common.Integer
	Width      common.Integer
	CreateTime common.String
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
