package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamSnapshotInfoRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	Limit         int    `position:"Query" name:"Limit"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (r DescribeLiveStreamSnapshotInfoRequest) Invoke(client *sdk.Client) (response *DescribeLiveStreamSnapshotInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveStreamSnapshotInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamSnapshotInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveStreamSnapshotInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveStreamSnapshotInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveStreamSnapshotInfoResponse struct {
	RequestId                  string
	NextStartTime              string
	LiveStreamSnapshotInfoList DescribeLiveStreamSnapshotInfoLiveStreamSnapshotInfoList
}

type DescribeLiveStreamSnapshotInfoLiveStreamSnapshotInfo struct {
	OssEndpoint string
	OssBucket   string
	OssObject   string
	CreateTime  string
}

type DescribeLiveStreamSnapshotInfoLiveStreamSnapshotInfoList []DescribeLiveStreamSnapshotInfoLiveStreamSnapshotInfo

func (list *DescribeLiveStreamSnapshotInfoLiveStreamSnapshotInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamSnapshotInfoLiveStreamSnapshotInfo)
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
