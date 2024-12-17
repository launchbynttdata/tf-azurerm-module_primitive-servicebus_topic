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

variable "name" {
  description = "The name of the Service Bus Topic."
  type        = string
}

variable "namespace_id" {
  description = "The ID of the Service Bus Namespace."
  type        = string
}

variable "partitioning_enabled" {
  description = "Boolean flag which controls whether to enable partitioning on the Service Bus Topic."
  type        = bool
  default     = true
}

variable "max_size_in_megabytes" {
  description = "The maximum size of the topic in megabytes."
  type        = number
  default     = 1024
}

variable "status" {
  description = "The status of the Service Bus Topic."
  type        = string
  default     = "Active"
}

variable "support_ordering" {
  description = "Boolean flag which controls whether the topic supports ordering."
  type        = bool
  default     = false
}

variable "auto_delete_on_idle" {
  description = "The idle interval after which the topic is automatically deleted."
  type        = string
  default     = "P10675199DT2H48M5.4775807S"
}

variable "default_message_ttl" {
  description = "The TTL of messages sent to this topic if no TTL value is set on the message itself."
  type        = string
  default     = "P10675199DT2H48M5.4775807S"
}

variable "duplicate_detection_history_time_window" {
  description = "The duration of the duplicate detection history."
  type        = string
  default     = "PT10M"
}
