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
  description = "The name of the Redis Cache Access Policy. Changing this forces a new Redis Cache Access Policy to be created."
  type        = string
  nullable    = false
}

variable "redis_cache_id" {
  description = "The ID of the Redis Cache. Changing this forces a new Redis Cache Access Policy to be created."
  type        = string
  nullable    = false
}

variable "permissions" {
  description = "Permissions that are going to be assigned to this Redis Cache Access Policy."
  type        = string
  nullable    = false
}
