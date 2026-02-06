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

package cores

import "testing"

func TestDecodeContent(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"EmptyString", "", ""},
		{"ValidString", "Hello%20World", "Hello World"},
		{"ValidString2", "%D7%B4%CC%AC%B7%B5%BB%D8%B3%C9%B9%A6", "状态返回成功"},
		{"InvalidString", "Hello%2", "Hello%2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DecodeContent(tt.input)
			if result != tt.expected {
				t.Errorf("DecodeContent(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}
