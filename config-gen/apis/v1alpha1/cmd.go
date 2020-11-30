/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"github.com/markbates/pkger"
	"github.com/spf13/cobra"
	"sigs.k8s.io/kustomize/kyaml/fn/framework"
	"sigs.k8s.io/kustomize/kyaml/kio"
)

// NewCommand returns a new cobra command
func NewCommand() *cobra.Command {
	kp := &KubebuilderProject{}

	return framework.TemplateCommand{
		API: kp,

		MergeResources: true, // apply additional inputs as patches

		// these are run before the templates
		PreProcessFilters: []kio.Filter{
			// run controller-gen libraries to generate configuration from code
			ControllerGenFilter{KubebuilderProject: kp},
			// inject generated certificates
			CertFilter{KubebuilderProject: kp},
		},

		// generate resources
		TemplatesFn: framework.TemplatesFromDir(pkger.Dir("/config-gen/templates/resources")),

		// patch resources
		PatchTemplatesFn: framework.PatchTemplatesFromDir(
			CRDPatchTemplate(kp),
			CertManagerPatchTemplate(kp),
			ControllerManagerPatchTemplate(kp),
		),

		// perform final modifications
		PostProcessFilters: []kio.Filter{
			// sort the resources
			ComponentFilter{KubebuilderProject: kp},
			SortFilter{KubebuilderProject: kp},
		},
	}.GetCommand()
}
