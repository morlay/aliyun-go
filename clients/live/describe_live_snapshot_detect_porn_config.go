package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveSnapshotDetectPornConfigRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int    `position:"Query" name:"PageSize"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	PageNum       int    `position:"Query" name:"PageNum"`
	Order         string `position:"Query" name:"Order"`
}

func (req *DescribeLiveSnapshotDetectPornConfigRequest) Invoke(client *sdk.Client) (resp *DescribeLiveSnapshotDetectPornConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveSnapshotDetectPornConfig", "", "")
	resp = &DescribeLiveSnapshotDetectPornConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveSnapshotDetectPornConfigResponse struct {
	responses.BaseResponse
	RequestId                        string
	PageNum                          int
	PageSize                         int
	Order                            string
	TotalNum                         int
	TotalPage                        int
	LiveSnapshotDetectPornConfigList DescribeLiveSnapshotDetectPornConfigLiveSnapshotDetectPornConfigList
}

type DescribeLiveSnapshotDetectPornConfigLiveSnapshotDetectPornConfig struct {
	DomainName  string
	AppName     string
	OssEndpoint string
	OssBucket   string
	OssObject   string
	Interval    int
}

type DescribeLiveSnapshotDetectPornConfigLiveSnapshotDetectPornConfigList []DescribeLiveSnapshotDetectPornConfigLiveSnapshotDetectPornConfig

func (list *DescribeLiveSnapshotDetectPornConfigLiveSnapshotDetectPornConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveSnapshotDetectPornConfigLiveSnapshotDetectPornConfig)
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
