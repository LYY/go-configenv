package configenv

import (
	"github.com/kylelemons/go-gypsy/yaml"
)

// Get retrieves a list from the file specified by a string of the same
// format as that expected by Child.  If the final node is not a List or node content is not a Scalar, Get
// will return an error.
func yamlGetList(f *yaml.File, spec string) ([]string, error) {
	cnode, err := yaml.Child(f.Root, spec)
	if err != nil {
		return nil, err
	}

	if cnode == nil {
		return nil, &yaml.NodeNotFound{
			Full: spec,
			Spec: spec,
		}
	}

	lst, ok := cnode.(yaml.List)
	if !ok {
		return nil, &yaml.NodeTypeMismatch{
			Full:     spec,
			Spec:     spec,
			Token:    "$",
			Expected: "yaml.List",
			Node:     cnode,
		}
	}

	rst := make([]string, len(lst))
	for idx, node := range lst {
		scalar, ok := node.(yaml.Scalar)
		if !ok {
			return nil, &yaml.NodeTypeMismatch{
				Full:     spec,
				Spec:     spec,
				Token:    "$",
				Expected: "yaml.Scalar",
				Node:     node,
			}
		}
		val := scalar.String()
		rst[idx] = val
	}
	return rst, nil

}
