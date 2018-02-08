package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateClusterRequest struct {
	requests.RpcRequest
	SccClusterId                string                        `position:"Query" name:"SccClusterId"`
	ImageId                     string                        `position:"Query" name:"ImageId"`
	EcsOrderManagerInstanceType string                        `position:"Query" name:"EcsOrderManagerInstanceType"`
	EhpcVersion                 string                        `position:"Query" name:"EhpcVersion"`
	AccountType                 string                        `position:"Query" name:"AccountType"`
	SecurityGroupId             string                        `position:"Query" name:"SecurityGroupId"`
	Description                 string                        `position:"Query" name:"Description"`
	KeyPairName                 string                        `position:"Query" name:"KeyPairName"`
	SecurityGroupName           string                        `position:"Query" name:"SecurityGroupName"`
	EcsOrderComputeInstanceType string                        `position:"Query" name:"EcsOrderComputeInstanceType"`
	ImageOwnerAlias             string                        `position:"Query" name:"ImageOwnerAlias"`
	VolumeType                  string                        `position:"Query" name:"VolumeType"`
	EcsOrderManagerCount        int                           `position:"Query" name:"EcsOrderManagerCount"`
	Password                    string                        `position:"Query" name:"Password"`
	EcsOrderLoginCount          int                           `position:"Query" name:"EcsOrderLoginCount"`
	ComputeSpotPriceLimit       string                        `position:"Query" name:"ComputeSpotPriceLimit"`
	VolumeProtocol              string                        `position:"Query" name:"VolumeProtocol"`
	OsTag                       string                        `position:"Query" name:"OsTag"`
	RemoteDirectory             string                        `position:"Query" name:"RemoteDirectory"`
	EcsOrderComputeCount        int                           `position:"Query" name:"EcsOrderComputeCount"`
	ComputeSpotStrategy         string                        `position:"Query" name:"ComputeSpotStrategy"`
	VSwitchId                   string                        `position:"Query" name:"VSwitchId"`
	Applications                *CreateClusterApplicationList `position:"Query" type:"Repeated" name:"Application"`
	EcsChargeType               string                        `position:"Query" name:"EcsChargeType"`
	HaEnable                    string                        `position:"Query" name:"HaEnable"`
	Name                        string                        `position:"Query" name:"Name"`
	SchedulerType               string                        `position:"Query" name:"SchedulerType"`
	VolumeId                    string                        `position:"Query" name:"VolumeId"`
	VolumeMountpoint            string                        `position:"Query" name:"VolumeMountpoint"`
	EcsOrderLoginInstanceType   string                        `position:"Query" name:"EcsOrderLoginInstanceType"`
	ZoneId                      string                        `position:"Query" name:"ZoneId"`
}

func (req *CreateClusterRequest) Invoke(client *sdk.Client) (resp *CreateClusterResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "CreateCluster", "ehs", "")
	resp = &CreateClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateClusterApplication struct {
	Tag string `name:"Tag"`
}

type CreateClusterResponse struct {
	responses.BaseResponse
	RequestId string
	ClusterId string
}

type CreateClusterApplicationList []CreateClusterApplication

func (list *CreateClusterApplicationList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateClusterApplication)
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
