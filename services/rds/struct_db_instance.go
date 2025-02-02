package rds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// DBInstance is a nested struct in rds response
type DBInstance struct {
	Engine                       string                                     `json:"Engine" xml:"Engine"`
	VpcName                      string                                     `json:"VpcName" xml:"VpcName"`
	VpcCloudInstanceId           string                                     `json:"VpcCloudInstanceId" xml:"VpcCloudInstanceId"`
	DBInstanceNetType            string                                     `json:"DBInstanceNetType" xml:"DBInstanceNetType"`
	DBInstanceClass              string                                     `json:"DBInstanceClass" xml:"DBInstanceClass"`
	InstanceNetworkType          string                                     `json:"InstanceNetworkType" xml:"InstanceNetworkType"`
	DBInstanceType               string                                     `json:"DBInstanceType" xml:"DBInstanceType"`
	SwitchWeight                 int                                        `json:"SwitchWeight" xml:"SwitchWeight"`
	DBInstanceId                 string                                     `json:"DBInstanceId" xml:"DBInstanceId"`
	VpcId                        string                                     `json:"VpcId" xml:"VpcId"`
	DBInstanceStorageType        string                                     `json:"DBInstanceStorageType" xml:"DBInstanceStorageType"`
	DedicatedHostZoneIdForMaster string                                     `json:"DedicatedHostZoneIdForMaster" xml:"DedicatedHostZoneIdForMaster"`
	DedicatedHostIdForLog        string                                     `json:"DedicatedHostIdForLog" xml:"DedicatedHostIdForLog"`
	TipsLevel                    int                                        `json:"TipsLevel" xml:"TipsLevel"`
	EngineVersion                string                                     `json:"EngineVersion" xml:"EngineVersion"`
	PayType                      string                                     `json:"PayType" xml:"PayType"`
	Tips                         string                                     `json:"Tips" xml:"Tips"`
	DedicatedHostZoneIdForSlave  string                                     `json:"DedicatedHostZoneIdForSlave" xml:"DedicatedHostZoneIdForSlave"`
	TempDBInstanceId             string                                     `json:"TempDBInstanceId" xml:"TempDBInstanceId"`
	ZoneId                       string                                     `json:"ZoneId" xml:"ZoneId"`
	DedicatedHostZoneIdForLog    string                                     `json:"DedicatedHostZoneIdForLog" xml:"DedicatedHostZoneIdForLog"`
	DedicatedHostNameForSlave    string                                     `json:"DedicatedHostNameForSlave" xml:"DedicatedHostNameForSlave"`
	ConnectionMode               string                                     `json:"ConnectionMode" xml:"ConnectionMode"`
	LockMode                     string                                     `json:"LockMode" xml:"LockMode"`
	DedicatedHostIdForSlave      string                                     `json:"DedicatedHostIdForSlave" xml:"DedicatedHostIdForSlave"`
	LockReason                   string                                     `json:"LockReason" xml:"LockReason"`
	Category                     string                                     `json:"Category" xml:"Category"`
	InsId                        int                                        `json:"InsId" xml:"InsId"`
	GuardDBInstanceId            string                                     `json:"GuardDBInstanceId" xml:"GuardDBInstanceId"`
	DBInstanceDescription        string                                     `json:"DBInstanceDescription" xml:"DBInstanceDescription"`
	CreateTime                   string                                     `json:"CreateTime" xml:"CreateTime"`
	GeneralGroupName             string                                     `json:"GeneralGroupName" xml:"GeneralGroupName"`
	DestroyTime                  string                                     `json:"DestroyTime" xml:"DestroyTime"`
	DedicatedHostIdForMaster     string                                     `json:"DedicatedHostIdForMaster" xml:"DedicatedHostIdForMaster"`
	DedicatedHostNameForLog      string                                     `json:"DedicatedHostNameForLog" xml:"DedicatedHostNameForLog"`
	RegionId                     string                                     `json:"RegionId" xml:"RegionId"`
	ResourceGroupId              string                                     `json:"ResourceGroupId" xml:"ResourceGroupId"`
	ExpireTime                   string                                     `json:"ExpireTime" xml:"ExpireTime"`
	MutriORsignle                bool                                       `json:"MutriORsignle" xml:"MutriORsignle"`
	DedicatedHostGroupId         string                                     `json:"DedicatedHostGroupId" xml:"DedicatedHostGroupId"`
	DedicatedHostGroupName       string                                     `json:"DedicatedHostGroupName" xml:"DedicatedHostGroupName"`
	VSwitchId                    string                                     `json:"VSwitchId" xml:"VSwitchId"`
	MasterInstanceId             string                                     `json:"MasterInstanceId" xml:"MasterInstanceId"`
	DBInstanceStatus             string                                     `json:"DBInstanceStatus" xml:"DBInstanceStatus"`
	ReplicateId                  string                                     `json:"ReplicateId" xml:"ReplicateId"`
	ConnectionString             string                                     `json:"ConnectionString" xml:"ConnectionString"`
	DedicatedHostNameForMaster   string                                     `json:"DedicatedHostNameForMaster" xml:"DedicatedHostNameForMaster"`
	AutoUpgradeMinorVersion      string                                     `json:"AutoUpgradeMinorVersion" xml:"AutoUpgradeMinorVersion"`
	DeletionProtection           bool                                       `json:"DeletionProtection" xml:"DeletionProtection"`
	ReadOnlyDBInstanceIds        ReadOnlyDBInstanceIdsInDescribeDBInstances `json:"ReadOnlyDBInstanceIds" xml:"ReadOnlyDBInstanceIds"`
}
