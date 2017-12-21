package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamRecordIndexFilesRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (r DescribeLiveStreamRecordIndexFilesRequest) Invoke(client *sdk.Client) (response *DescribeLiveStreamRecordIndexFilesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveStreamRecordIndexFilesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamRecordIndexFiles", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveStreamRecordIndexFilesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveStreamRecordIndexFilesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveStreamRecordIndexFilesResponse struct {
	RequestId           string
	RecordIndexInfoList DescribeLiveStreamRecordIndexFilesRecordIndexInfoList
}

type DescribeLiveStreamRecordIndexFilesRecordIndexInfo struct {
	RecordId   string
	RecordUrl  string
	DomainName string
	AppName    string
	StreamName string
	OssObject  string
	StartTime  string
	EndTime    string
	Duration   float32
	Height     int
	Width      int
	CreateTime string
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
