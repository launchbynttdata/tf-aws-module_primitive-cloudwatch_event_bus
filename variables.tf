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
  description = "Name of the new event bus. The names of custom event buses cannot contain the / character. To create a partner event bus, ensure that the name matches the event_source_name."
  type        = string

  validation {
    condition     = length(regexall("/", var.name)) == 0
    error_message = "name must not contain '/'."
  }
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

variable "kms_key_identifier" {
  description = "Identifier of the AWS KMS customer managed key for EventBridge to use for encrypting events. Can be the key ARN, KeyId, key alias, or key alias ARN."
  type        = string
  default     = null
}

variable "dead_letter_config" {
  description = "Configuration details of the Amazon SQS queue for EventBridge to use as a dead-letter queue (DLQ)."
  type = object({
    arn = string
  })
  default = null
}

variable "tags" {
  description = "Map of tags to assign to the event bus."
  type        = map(string)
  default     = {}
}
