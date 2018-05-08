package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLiveStreamSnapshotInfoRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	Limit         int    `position:"Query" name:"Limit"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (req *DescribeLiveStreamSnapshotInfoRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamSnapshotInfoResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamSnapshotInfo", "", "")
	resp = &DescribeLiveStreamSnapshotInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamSnapshotInfoResponse struct {
	responses.BaseResponse
	RequestId                  common.String
	NextStartTime              common.String
	LiveStreamSnapshotInfoList DescribeLiveStreamSnapshotInfoLiveStreamSnapshotInfoList
}

type DescribeLiveStreamSnapshotInfoLiveStreamSnapshotInfo struct {
	OssEndpoint common.String
	OssBucket   common.String
	OssObject   common.String
	CreateTime  common.String
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
