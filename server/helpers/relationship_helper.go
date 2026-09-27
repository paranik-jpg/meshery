package helpers

import (
	"strings"

	"github.com/meshery/schemas/models/v1alpha3/relationship"
)

// NormalizeRelationshipDefinition lowercases relationship classification fields
// (Kind, RelationshipType, SubType) to prevent case-sensitive duplicate registrations.
func NormalizeRelationshipDefinition(r *relationship.RelationshipDefinition) {
	if r == nil {
		return
	}
	r.Kind = relationship.RelationshipDefinitionKind(strings.ToLower(string(r.Kind)))
	r.RelationshipType = strings.ToLower(r.RelationshipType)
	r.SubType = strings.ToLower(r.SubType)
}
