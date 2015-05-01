//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package zero_width_non_joiner

import (
	"regexp"

	"github.com/khlieng/name_pending/Godeps/_workspace/src/github.com/blevesearch/bleve/analysis"
	"github.com/khlieng/name_pending/Godeps/_workspace/src/github.com/blevesearch/bleve/analysis/char_filters/regexp_char_filter"
	"github.com/khlieng/name_pending/Godeps/_workspace/src/github.com/blevesearch/bleve/registry"
)

const Name = "zero_width_spaces"

var zeroWidthNonJoinerRegexp = regexp.MustCompile(`\x{200C}`)

func CharFilterConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.CharFilter, error) {
	replaceBytes := []byte(" ")
	return regexp_char_filter.NewRegexpCharFilter(zeroWidthNonJoinerRegexp, replaceBytes), nil
}

func init() {
	registry.RegisterCharFilter(Name, CharFilterConstructor)
}