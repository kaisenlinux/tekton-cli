/*
Copyright 2021 The Tekton Authors

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

package v1beta1

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/tektoncd/pipeline/pkg/apis/validate"
	"github.com/tektoncd/triggers/pkg/apis/config"
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"knative.dev/pkg/apis"
	"knative.dev/pkg/webhook/resourcesemantics"
)

// paramsRegexp captures TriggerTemplate parameter names $(tt.params.NAME)
var paramsRegexp = regexp.MustCompile(`\$\(tt.params.(?P<var>[_a-zA-Z][_a-zA-Z0-9.-]*)\)`)

var _ resourcesemantics.VerbLimited = (*TriggerTemplate)(nil)

// SupportedVerbs returns the operations that validation should be called for
func (t *TriggerTemplate) SupportedVerbs() []admissionregistrationv1.OperationType {
	return []admissionregistrationv1.OperationType{admissionregistrationv1.Create, admissionregistrationv1.Update}
}

// Validate validates a TriggerTemplate.
func (t *TriggerTemplate) Validate(ctx context.Context) *apis.FieldError {
	errs := validate.ObjectMetadata(t.GetObjectMeta()).ViaField("metadata")
	return errs.Also(t.Spec.validate(ctx).ViaField("spec"))
}

// revive:disable:unused-parameter

// Validate validates a TriggerTemplateSpec.
func (s *TriggerTemplateSpec) validate(ctx context.Context) (errs *apis.FieldError) {
	if equality.Semantic.DeepEqual(s, &TriggerTemplateSpec{}) {
		errs = errs.Also(apis.ErrMissingField(apis.CurrentField))
	}
	if len(s.ResourceTemplates) == 0 {
		errs = errs.Also(apis.ErrMissingField("resourceTemplates"))
	}
	errs = errs.Also(validateResourceTemplates(s.ResourceTemplates).ViaField("resourceTemplates"))
	errs = errs.Also(verifyParamDeclarations(s.Params, s.ResourceTemplates).ViaField("resourceTemplates"))
	return errs
}

func validateResourceTemplates(templates []TriggerResourceTemplate) (errs *apis.FieldError) {
	for i, trt := range templates {
		if err := config.EnsureAllowedType(trt.RawExtension); err != nil {
			if runtime.IsMissingVersion(err) {
				errs = errs.Also(apis.ErrMissingField(fmt.Sprintf("[%d].apiVersion", i)))
			}
			if runtime.IsMissingKind(err) {
				errs = errs.Also(apis.ErrMissingField(fmt.Sprintf("[%d].kind", i)))
			}
			if runtime.IsNotRegisteredError(err) {
				errStr := err.Error()
				if inSchemeIdx := strings.Index(errStr, " in scheme"); inSchemeIdx > -1 {
					// not registered error messages currently include the scheme variable location in your file,
					// which can of course change if you move the location of the variable in your file.
					// So will filter it out here to facilitate our unit testing, as the scheme location is not
					// useful for our purposes.
					errStr = errStr[:inSchemeIdx]
				}
				errs = errs.Also(apis.ErrInvalidValue(
					errStr,
					fmt.Sprintf("[%d]", i)))
			}
			// we allow structural errors because of param substitution
		}
	}
	return errs
}

// Verify every param in the ResourceTemplates is declared with a ParamSpec
func verifyParamDeclarations(params []ParamSpec, templates []TriggerResourceTemplate) *apis.FieldError {
	declaredParamNames := sets.NewString()
	for _, param := range params {
		declaredParamNames.Insert(param.Name)
	}
	for i, template := range templates {
		// Get all params in the template $(tt.params.NAME)
		templateParams := paramsRegexp.FindAllSubmatch(template.RawExtension.Raw, -1)
		for _, templateParam := range templateParams {
			templateParamName := string(templateParam[1])
			if !declaredParamNames.Has(templateParamName) {
				fieldErr := apis.ErrInvalidValue(
					fmt.Sprintf("undeclared param '$(tt.params.%s)'", templateParamName),
					fmt.Sprintf("[%d]", i),
				)
				fieldErr.Details = fmt.Sprintf("'$(tt.params.%s)' must be declared in spec.params", templateParamName)
				return fieldErr
			}
		}
	}

	return nil
}
