// Copyright 2023 Trillian Authors. All Rights Reserved.
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

// render is a command to take a Claimant Model specified as yaml and
// output markdown representations of it.
package main

import (
	"flag"
	"os"

	claimant "github.com/google/trillian/docs/claimantmodel/experimental/cmd/render/internal"
	"gopkg.in/yaml.v2"
	"k8s.io/klog/v2"
)

var (
	dmf = flag.String("domain_model_file", "", "path to domain model yaml file")
	fmf = flag.String("full_model_file", "", "path to full model yaml file")
)

func main() {
	flag.Parse()

	if (len(*dmf) == 0) == (len(*fmf) == 0) {
		klog.Exitf("--domain_model_file OR --full_model_file are required")
	}

	if len(*dmf) > 0 {
		saved, err := os.ReadFile(*dmf)
		if err != nil {
			klog.Exitf("failed to read model: %v", err)
		}

		domain := claimant.Model{}
		if err := yaml.Unmarshal(saved, &domain); err != nil {
			klog.Exitf("failed to parse Model: %v", err)
		}
		handleSingleModel(domain)
	} else {
		saved, err := os.ReadFile(*fmf)
		if err != nil {
			klog.Exitf("failed to read model: %v", err)
		}

		models := claimant.Models{}
		if err := yaml.Unmarshal(saved, &models); err != nil {
			klog.Exitf("failed to parse Models: %v", err)
		}
		handleMultiModels(models)
	}
}

func handleSingleModel(domain claimant.Model) { _ = "STUB: not implemented"; return }

func handleMultiModels(models claimant.Models) { _ = "STUB: not implemented"; return }

func getGenerateDocs() string { _ = "STUB: not implemented"; return "" }
