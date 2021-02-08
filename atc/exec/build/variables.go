package build

import (
	"github.com/concourse/concourse/vars"
)

type Variables struct {
	parentScope vars.Variables
	localVars   vars.StaticVariables
	tracker     *vars.Tracker
}

func NewVariables(tracker *vars.Tracker) *Variables {
	return &Variables{
		localVars: vars.StaticVariables{},
		tracker:   tracker,
	}
}

func (v *Variables) Get(ref vars.Reference) (interface{}, bool, error) {
	if ref.Source == "." {
		val, found, err := v.localVars.Get(ref)
		if found || err != nil {
			return val, found, err
		}
	}

	if v.parentScope != nil {
		val, found, err := v.parentScope.Get(ref)
		if found || err != nil {
			return val, found, err
		}
	}

	return nil, false, nil
}

func (v *Variables) NewScope(tracker *vars.Tracker) *Variables {
	return &Variables{
		parentScope: v,
		localVars:   vars.StaticVariables{},
		tracker:     tracker,
	}
}

func (v *Variables) SetVar(source, name string, val interface{}, redact bool) {
	v.localVars[name] = val

	if redact {
		v.tracker.Track(vars.Reference{Source: source, Path: name}, val)
	}
}

func (v *Variables) RedactionEnabled() bool {
	return v.tracker.Enabled
}
