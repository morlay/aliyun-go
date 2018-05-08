package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeFlowProjectClusterSettingRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
}

func (req *DescribeFlowProjectClusterSettingRequest) Invoke(client *sdk.Client) (resp *DescribeFlowProjectClusterSettingResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeFlowProjectClusterSetting", "", "")
	resp = &DescribeFlowProjectClusterSettingResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeFlowProjectClusterSettingResponse struct {
	responses.BaseResponse
	RequestId    common.String
	GmtCreate    common.Long
	GmtModified  common.Long
	ProjectId    common.String
	ClusterId    common.String
	DefaultUser  common.String
	DefaultQueue common.String
	UserList     DescribeFlowProjectClusterSettingUserListList
	QueueList    DescribeFlowProjectClusterSettingQueueListList
	HostList     DescribeFlowProjectClusterSettingHostListList
}

type DescribeFlowProjectClusterSettingUserListList []common.String

func (list *DescribeFlowProjectClusterSettingUserListList) UnmarshalJSON(data []byte) error {
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

type DescribeFlowProjectClusterSettingQueueListList []common.String

func (list *DescribeFlowProjectClusterSettingQueueListList) UnmarshalJSON(data []byte) error {
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

type DescribeFlowProjectClusterSettingHostListList []common.String

func (list *DescribeFlowProjectClusterSettingHostListList) UnmarshalJSON(data []byte) error {
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
