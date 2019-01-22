package data

import (
	"encoding/json"

	"github.com/drycc/workflow-manager/k8s"
	"github.com/drycc/workflow-manager/pkg/swagger/models"
)

var (
	daemonSetType  = "Daemon Set"
	deploymentType = "Deployment"
	rcType         = "Replication Controller"
)

const versionAnnotation = "component.drycc.cc/version"

// InstalledData is an interface for managing installed cluster metadata
type InstalledData interface {
	// will have a Get method to retrieve installed data
	Get() ([]byte, error)
}

// InstalledDryccData fulfills the InstalledData interface
type installedDryccData struct {
	k8sResources *k8s.ResourceInterfaceNamespaced
}

// NewInstalledDryccData returns a new InstalledDryccData using rcl as the rc.Lister implementation
func NewInstalledDryccData(ri *k8s.ResourceInterfaceNamespaced) InstalledData {
	return &installedDryccData{k8sResources: ri}
}

// Get method for InstalledDryccData
func (g *installedDryccData) Get() ([]byte, error) {
	var cluster models.Cluster
	deployments, err := k8s.GetDeployments(g.k8sResources.Deployments())
	if err != nil {
		return nil, err
	}
	for _, deployment := range deployments {
		component := models.ComponentVersion{
			Component: &models.Component{
				Name: deployment.Name,
				Type: &deploymentType,
			},
			Version: &models.Version{
				Version: deployment.Annotations[versionAnnotation],
				Data: &models.VersionData{
					Image: &deployment.Spec.Template.Spec.Containers[0].Image,
				},
			},
		}
		cluster.Components = append(cluster.Components, &component)
	}
	daemonSets, err := k8s.GetDaemonSets(g.k8sResources.DaemonSets())
	if err != nil {
		return nil, err
	}
	for _, daemonSet := range daemonSets {
		component := models.ComponentVersion{
			Component: &models.Component{
				Name: daemonSet.Name,
				Type: &daemonSetType,
			},
			Version: &models.Version{
				Version: daemonSet.Annotations[versionAnnotation],
				Data: &models.VersionData{
					Image: &daemonSet.Spec.Template.Spec.Containers[0].Image,
				},
			},
		}
		cluster.Components = append(cluster.Components, &component)
	}
	replicationControllers, err := k8s.GetReplicationControllers(g.k8sResources.ReplicationControllers())
	if err != nil {
		return nil, err
	}
	for _, rc := range replicationControllers {
		component := models.ComponentVersion{
			Component: &models.Component{
				Name: rc.Name,
				Type: &rcType,
			},
			Version: &models.Version{
				Version: rc.Annotations[versionAnnotation],
				Data: &models.VersionData{
					Image: &rc.Spec.Template.Spec.Containers[0].Image,
				},
			},
		}
		cluster.Components = append(cluster.Components, &component)
	}
	js, err := json.Marshal(cluster)
	if err != nil {
		return []byte{}, err
	}
	return js, nil
}
