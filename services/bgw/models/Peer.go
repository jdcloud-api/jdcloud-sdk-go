// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package models


type Peer struct {

    /* Peer的Id (Optional) */
    PeerId string `json:"peerId"`

    /* Peer的BGP状态，取值为：Up, Down (Optional) */
    PeerBgpStatus string `json:"peerBgpStatus"`

    /* Peer的ip信息详情 (Optional) */
    PeerIps PeerIps `json:"peerIps"`
}
