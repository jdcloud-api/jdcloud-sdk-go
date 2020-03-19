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


type RedisKey struct {

    /* key名称  */
    Name string `json:"name"`

    /* key所在的db号  */
    Db int `json:"db"`

    /* string类型的key表示字节数，list类型的key表示列表长度，set或zset类型的key表示集合或有序集合的大小、hash类型的key表示字典的大小等等  */
    Size int `json:"size"`

    /* string、list、set、zset、hash五种类型  */
    KeyType string `json:"keyType"`

    /* key访问的频度 (Optional) */
    Frequency int `json:"frequency"`
}
