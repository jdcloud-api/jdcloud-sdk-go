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


type FirewallRule struct {

    /*  (Optional) */
    Products []string `json:"products"`

    /* The priority of the rule to allow control of processing order.
A lower number indicates high priority. If not provided, any rules with a priority will be sequenced before those without.
 (Optional) */
    Priority int `json:"priority"`

    /* Whether this firewall rule is currently paused. (Optional) */
    Paused bool `json:"paused"`

    /* Short reference tag to quickly select related rules. (Optional) */
    Ref string `json:"ref"`

    /*  (Optional) */
    Action_parameters Action_parameters `json:"action_parameters"`

    /* The action to apply to a matched request. Note that action "log" is only available for enterprise customers. (Optional) */
    Action string `json:"action"`

    /*  (Optional) */
    Filter Filter `json:"filter"`

    /* Firewall Rule identifier (Optional) */
    Id string `json:"id"`

    /* A description of the rule to help identify it. (Optional) */
    Description string `json:"description"`
}