package concourse

import (
	"encoding/json"
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/go-concourse/concourse/internal"
	"github.com/tedsuo/rata"
	"net/http"
	"net/url"
)

func (team *team) Resource(pipelineRef atc.PipelineRef, resourceName string) (atc.Resource, bool, error) {
	params := rata.Params{
		"pipeline_name": pipelineRef.Name,
		"resource_name": resourceName,
		"team_name":     team.Name(),
	}

	var resource atc.Resource
	err := team.connection.Send(internal.Request{
		RequestName: atc.GetResource,
		Params:      params,
		Query:       pipelineRef.QueryParams(),
	}, &internal.Response{
		Result: &resource,
	})
	switch err.(type) {
	case nil:
		return resource, true, nil
	case internal.ResourceNotFoundError:
		return resource, false, nil
	default:
		return resource, false, err
	}
}

func (team *team) ListResources(pipelineRef atc.PipelineRef) ([]atc.Resource, error) {
	params := rata.Params{
		"pipeline_name": pipelineRef.Name,
		"team_name":     team.Name(),
	}

	var resources []atc.Resource
	err := team.connection.Send(internal.Request{
		RequestName: atc.ListResources,
		Params:      params,
		Query:       pipelineRef.QueryParams(),
	}, &internal.Response{
		Result: &resources,
	})

	return resources, err
}


func (team *team) ClearResourceCache(pipelineRef atc.PipelineRef, ResourceName string, version atc.Version) (int64, error) {
	params := rata.Params{
		"team_name":     team.Name(),
		"pipeline_name": pipelineRef.Name,
		"resource_name": ResourceName,
	}

	queryParams := url.Values{}

	if version != nil {
		jsonBytes, err := json.Marshal(version)
		if err != nil {
			return 0, err
		}
		params["version"] = string(jsonBytes)
	}

	var crcResponse atc.ClearResourceCacheResponse
	responseHeaders := http.Header{}
	response := internal.Response{
		Headers: &responseHeaders,
		Result:  &crcResponse,
	}
	err := team.connection.Send(internal.Request{
		RequestName: atc.ClearResourceCache,
		Params:      params,
		Query:       merge(queryParams, pipelineRef.QueryParams()),
	}, &response)

	if err != nil {
		return 0, err
	} else {
		return crcResponse.CachesRemoved, nil
	}
}