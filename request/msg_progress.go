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

package request

type MessageProgress struct {
	//发送消息时使用的微应用的ID
	AgentId int `json:"agent_id" validate:"required"`

	//发送消息时钉钉返回的任务ID。
	TaskId int `json:"task_id" validate:"required"`
}

type messageProgressBuilder struct {
	mp *MessageProgress
}

func NewMessageProgress(taskId int) *messageProgressBuilder {
	return &messageProgressBuilder{mp: &MessageProgress{TaskId: taskId}}
}

func (b *messageProgressBuilder) SetAgentId(agentId int) *messageProgressBuilder {
	b.mp.AgentId = agentId
	return b
}

func (b *messageProgressBuilder) Build() *MessageProgress {
	return b.mp
}
