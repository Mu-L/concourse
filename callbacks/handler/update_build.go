package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	TurbineBuilds "github.com/concourse/turbine/api/builds"
	"github.com/pivotal-golang/lager"

	"github.com/concourse/atc/builds"
)

func (handler *Handler) UpdateBuild(w http.ResponseWriter, r *http.Request) {
	buildIDStr := r.FormValue(":build")

	buildID, err := strconv.Atoi(buildIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var turbineBuild TurbineBuilds.Build
	if err := json.NewDecoder(r.Body).Decode(&turbineBuild); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log := handler.logger.Session("update-build", lager.Data{
		"id":      buildID,
		"status":  turbineBuild.Status,
		"inputs":  turbineBuild.Inputs,
		"outputs": turbineBuild.Outputs,
	})

	var status builds.Status

	switch turbineBuild.Status {
	case TurbineBuilds.StatusStarted:
		status = builds.StatusStarted
	case TurbineBuilds.StatusSucceeded:
		status = builds.StatusSucceeded
	case TurbineBuilds.StatusFailed:
		status = builds.StatusFailed
	case TurbineBuilds.StatusErrored:
		status = builds.StatusErrored
	case TurbineBuilds.StatusAborted:
		status = builds.StatusAborted
	default:
		log.Info("unknown-status")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Info("save-status")

	err = handler.buildDB.SaveBuildStatus(buildID, status)
	if err != nil {
		log.Error("failed-to-save-build-status", err)
		return
	}

	if status == builds.StatusStarted {
		err := handler.buildDB.SaveBuildStartTime(buildID, time.Now())
		if err != nil {
			log.Error("failed-to-save-build-start-time", err)
			return
		}
	} else {
		err := handler.buildDB.SaveBuildEndTime(buildID, time.Now())
		if err != nil {
			log.Error("failed-to-save-build-end-time", err)
			return
		}
	}

	build, err := handler.buildDB.GetBuild(buildID)
	if err != nil {
		log.Error("failed-to-get-build", err)
		return
	}

	if build.JobName == "" {
		// one-off build; don't care about its inputs or outputs
		return
	}

	switch turbineBuild.Status {
	case TurbineBuilds.StatusStarted:
		for _, input := range turbineBuild.Inputs {
			err = handler.buildDB.SaveBuildInput(buildID, vrFromInput(input))
			if err != nil {
				log.Error("failed-to-save-input", err)
			}
		}
	case TurbineBuilds.StatusSucceeded:
		explicitOutput := make(map[string]bool)

		for _, output := range turbineBuild.Outputs {
			err = handler.buildDB.SaveBuildOutput(buildID, vrFromOutput(output))
			if err != nil {
				log.Error("failed-to-save-output-version", err)
			}

			explicitOutput[output.Name] = true
		}

		for _, input := range turbineBuild.Inputs {
			if explicitOutput[input.Name] {
				continue
			}

			err = handler.buildDB.SaveBuildOutput(buildID, vrFromInput(input))
			if err != nil {
				log.Error("failed-to-save-output-version", err)
			}
		}
	}
}

func vrFromInput(input TurbineBuilds.Input) builds.VersionedResource {
	metadata := make([]builds.MetadataField, len(input.Metadata))
	for i, md := range input.Metadata {
		metadata[i] = builds.MetadataField{
			Name:  md.Name,
			Value: md.Value,
		}
	}

	return builds.VersionedResource{
		Name:     input.Name,
		Type:     input.Type,
		Source:   builds.Source(input.Source),
		Version:  builds.Version(input.Version),
		Metadata: metadata,
	}
}

// same as input, but type is different.
//
// :(
func vrFromOutput(output TurbineBuilds.Output) builds.VersionedResource {
	metadata := make([]builds.MetadataField, len(output.Metadata))
	for i, md := range output.Metadata {
		metadata[i] = builds.MetadataField{
			Name:  md.Name,
			Value: md.Value,
		}
	}

	return builds.VersionedResource{
		Name:     output.Name,
		Type:     output.Type,
		Source:   builds.Source(output.Source),
		Version:  builds.Version(output.Version),
		Metadata: metadata,
	}
}
