// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package spec

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/erda-project/erda/pkg/parser/diceyml"
)

type ClusterPhase string

const (
	ClusterPhaseNone     ClusterPhase = ""
	ClusterPhaseInitJobs ClusterPhase = "InitJobs"
	ClusterPhaseCreating ClusterPhase = "Creating"
	ClusterPhaseUpdating ClusterPhase = "Updating"
	ClusterPhaseRunning  ClusterPhase = "Running"
	ClusterPhaseFailed   ClusterPhase = "Failed"
	ClusterPhasePending  ClusterPhase = "Pending"
)

type ComponentStatus string

const (
	ComponentStatusReady              ComponentStatus = "Ready"
	ComponentStatusDeploying          ComponentStatus = "Deploying"
	ComponentStatusUnReady            ComponentStatus = "Unready"
	ComponentStatusNeedCreateOrUpdate ComponentStatus = "NeedCreateOrUpdate"
)

type ClusterSize string

const (
	ClusterSizeTest ClusterSize = "test"
	ClusterSizeProd ClusterSize = "prod"
)

const(
	PANameHPA = "HPA"
	PANameVPA = "VPA"
)

type PATarget struct{
	PAKind                string
	TargetControllerKind  string
}

type DiceClusterList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Items             []DiceCluster `json:"items"`
}

type DiceCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status"`
}

type ClusterSpec struct {
	ResetStatus          bool        `json:"resetStatus"`
	EnableAutoScale      bool        `json:"enableAutoScale"`
	AddonConfigMap       string      `json:"addonConfigMap"`
	ClusterinfoConfigMap string      `json:"clusterinfoConfigMap"`
	PlatformDomain       string      `json:"platformDomain"`
	CookieDomain         string      `json:"cookieDomain"`
	Size                 ClusterSize `json:"size"`
	DiceCluster          string      `json:"diceCluster"`
	// collector, openapi
	MainPlatform map[string]string `json:"mainPlatform"`
	// key: dice-service-name(e.g. ui), value: domain
	// customDomain:
	//   ui: dice.terminus.io,*.terminus.io
	CustomDomain map[string]string `json:"customDomain"`
	// deployment affinity labels for specific dice-service
	// key: dice-service-name(e.g. gittar), value: label
	// e.g.
	// gittar: dice/gittar
	CustomAffinity map[string]string `json:"customAffinity"`

	InitJobs diceyml.Object `json:"initJobs"`

	Dice           diceyml.Object `json:"dice"`
	AddonPlatform  diceyml.Object `json:"addonPlatform"`
	Gittar         diceyml.Object `json:"gittar"`
	Pandora        diceyml.Object `json:"pandora"`
	DiceUI         diceyml.Object `json:"diceUI"`
	UC             diceyml.Object `json:"uc"`
	SpotAnalyzer   diceyml.Object `json:"spotAnalyzer"`
	SpotCollector  diceyml.Object `json:"spotCollector"`
	SpotDashboard  diceyml.Object `json:"spotDashboard"`
	SpotFilebeat   diceyml.Object `json:"spotFilebeat"`
	SpotStatus     diceyml.Object `json:"spotStatus"`
	SpotTelegraf   diceyml.Object `json:"spotTelegraf"`
	Tmc            diceyml.Object `json:"tmc"`
	FluentBit      diceyml.Object `json:"fluentBit"`
	Hepa           diceyml.Object `json:"hepa"`
	SpotMonitor    diceyml.Object `json:"spotMonitor"`
	Fdp            diceyml.Object `json:"fdp"`
	FdpUI          diceyml.Object `json:"fdpUI"`
	MeshController diceyml.Object `json:"meshController"`
}
type ClusterStatus struct {
	Phase      ClusterPhase               `json:"phase"`
	Conditions []Condition                `json:"conditions"`
	Components map[string]ComponentStatus `json:"components"`
}

type Condition struct {
	Reason         string `json:"reason"`
	TransitionTime string `json:"transitionTime"`
}
