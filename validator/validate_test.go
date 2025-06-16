package validator

import (
	"fmt"
	"testing"

	"github.com/gofrs/uuid"

	"github.com/meshery/meshkit/models/catalog/v1alpha1"
	"github.com/meshery/schemas/models/v1beta1"
	"github.com/meshery/schemas/models/v1beta1/category"
	"github.com/meshery/schemas/models/v1beta1/component"
	"github.com/meshery/schemas/models/v1beta1/model"
	"github.com/meshery/schemas/models/v1beta1/pattern"
)

type ValidationCases struct {
	Path     string
	Resource interface{}

	ShouldPass bool
}

func TestValidator(t *testing.T) {
	tests := []ValidationCases{
		{
			Path: "design",
			Resource: &pattern.PatternFile{
				Name:          "test-design",
				SchemaVersion: v1beta1.DesignSchemaVersion,
				Components: []*component.ComponentDefinition{
					{
						Id:          uuid.FromStringOrNil("test-component-id"),
						DisplayName: "Test Component",
						Component: component.Component{
							Kind:    "Deployment",
							Version: "v1",
						},
						Model: model.ModelDefinition{
							Name: "Kubernetes",
						},
					},
				},
			},
			ShouldPass: true,
		},
		{
			Path: "catalog_data",
			Resource: v1alpha1.CatalogData{
				PublishedVersion: "v.10.9",
				ContentClass:     "sdsds",
				Compatibility: []v1alpha1.CatalogDataCompatibility{
					"kubernetes",
				},
				PatternCaveats: "NA",
				PatternInfo:    "NA",
				Type:           v1alpha1.CatalogDataType("Dployment"),
			},
			ShouldPass: false,
		},
		{
			Path: "models",
			Resource: model.ModelDefinition{

				SchemaVersion: v1beta1.ModelSchemaVersion,
				Version:       "1.0.0",

				Category: category.CategoryDefinition{
					Name: "test",
				},
				Model: model.Model{
					Version: "1.0.0",
				},
				Status:      "",
				DisplayName: "",
				Description: "",
			},
			ShouldPass: false,
		},
	}

	for _, test := range tests {
		t.Run("validaion", func(_t *testing.T) {
			schema, err := GetSchemaFor(test.Path)
			if err != nil {
				t.Errorf("%v", err)

			}

			err = Validate(schema, test.Resource)
			fmt.Println(err)
			if test.ShouldPass && err != nil {
				t.Errorf("test failed for %s, got %s, want %t, error: %v", test.Path, "false", test.ShouldPass, err)

			} else if !test.ShouldPass && err == nil {
				t.Errorf("test failed for %s, got %s, want %t error: %v", test.Path, "true", !test.ShouldPass, err)
			}

		})
	}
}
