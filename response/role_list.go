/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package response

type RoleList struct {
	Response

	rs `json:"result"`
}

type rs struct {
	//是否还有更多数据。
	HasMore bool `json:"hasMore"`

	RoleGroups []roleGroup `json:"list"`
}

type roleGroup struct {
	//角色组Id
	GroupId int `json:"groupId"`

	//角色组名称
	Name string `json:"name"`

	Roles []role `json:"roles"`
}

type role struct {
	//角色Id
	Id int `json:"id"`

	//角色名称
	Name string `json:"name"`
}
