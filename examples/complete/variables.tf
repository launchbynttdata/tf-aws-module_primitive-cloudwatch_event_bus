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

variable "logical_product_family" {
  description = "Logical product family for resource naming."
  type        = string
}

variable "logical_product_service" {
  description = "Logical product service for resource naming."
  type        = string
}

variable "class_env" {
  description = "Class environment for resource naming (e.g., dev, prod)."
  type        = string
}

variable "instance_env" {
  description = "Instance environment for resource naming."
  type        = string
}

variable "instance_resource" {
  description = "Instance resource for resource naming (numeric identifier)."
  type        = number
}

variable "resource_names_map" {
  description = "Map of resource types to naming configuration (name, max_length)."
  type = map(object({
    name       = string
    max_length = number
  }))
}

variable "description" {
  description = "Event bus description."
  type        = string
  default     = null
}

variable "event_source_name" {
  description = "Partner event source that the new event bus will be matched with. Must match name for partner event buses."
  type        = string
  default     = null
}

variable "dead_letter_config" {
  description = "Configuration details of the Amazon SQS queue for EventBridge to use as a dead-letter queue (DLQ)."
  type = object({
    arn = optional(string)
  })
  default = null
}

variable "tags" {
  description = "Map of tags to assign to the event bus."
  type        = map(string)
  default     = {}
}
