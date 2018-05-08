package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveSnapshotDetectPornConfig", "live", "")
	resp = &DescribeLiveSnapshotDetectPornConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveSnapshotDetectPornConfigResponse struct {
	responses.BaseResponse
	RequestId                        common.String
	PageNum                          common.Integer
	PageSize                         common.Integer
	Order                            common.String
	TotalNum                         common.Integer
	TotalPage                        common.Integer
	LiveSnapshotDetectPornConfigList DescribeLiveSnapshotDetectPornConfigLiveSnapshotDetectPornConfigList
}

type DescribeLiveSnapshotDetectPornConfigLiveSnapshotDetectPornConfig struct {
	DomainName  common.String
	AppName     common.String
	OssEndpoint common.String
	OssBucket   common.String
	OssObject   common.String
	Interval    common.Integer
	Scenes      DescribeLiveSnapshotDetectPornConfigSceneList
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

type DescribeLiveSnapshotDetectPornConfigSceneList []common.String

func (list *DescribeLiveSnapshotDetectPornConfigSceneList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
