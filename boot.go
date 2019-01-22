package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/drycc/workflow-manager/config"
	"github.com/drycc/workflow-manager/data"
	"github.com/drycc/workflow-manager/handlers"
	"github.com/drycc/workflow-manager/jobs"
	"github.com/drycc/workflow-manager/k8s"
	"github.com/gorilla/mux"
	kcl "k8s.io/kubernetes/pkg/client/unversioned"
)

func main() {
	kubeClient, err := kcl.NewInCluster()
	if err != nil {
		log.Fatalf("Error creating new Kubernetes client (%s)", err)
	}
	apiClient, err := config.GetSwaggerClient(config.Spec.VersionsAPIURL)
	if err != nil {
		log.Fatalf("Error creating new swagger api client (%s)", err)
	}
	dryccK8sResources := k8s.NewResourceInterfaceNamespaced(kubeClient, config.Spec.DryccNamespace)
	clusterID := data.NewClusterIDFromPersistentStorage(dryccK8sResources.Secrets())
	installedDryccData := data.NewInstalledDryccData(dryccK8sResources)
	availableVersion := data.NewAvailableVersionsFromAPI(
		apiClient,
		config.Spec.VersionsAPIURL,
	)
	availableComponentVersion := data.NewLatestReleasedComponent(dryccK8sResources, availableVersion)

	pollDur := time.Duration(config.Spec.Polling) * time.Second
	// we want to do the following jobs according to our remote API interval:
	// 1. get latest stable drycc component versions
	// 2. send diagnostic data, if appropriate
	glvdPeriodic := jobs.NewGetLatestVersionDataPeriodic(
		installedDryccData,
		clusterID,
		availableVersion,
		availableComponentVersion,
		pollDur,
	)

	svPeriodic := jobs.NewSendVersionsPeriodic(
		apiClient,
		clusterID,
		dryccK8sResources,
		availableVersion,
		pollDur,
	)
	toDo := []jobs.Periodic{glvdPeriodic, svPeriodic}
	log.Printf("Starting periodic jobs at interval %s", pollDur)
	ch := jobs.DoPeriodic(toDo)
	defer close(ch)

	// Get a new router, with handler functions
	r := handlers.RegisterRoutes(mux.NewRouter(), availableVersion, dryccK8sResources)
	// Bind to a port and pass our router in
	hostStr := fmt.Sprintf(":%s", config.Spec.Port)
	log.Printf("Serving on %s", hostStr)
	if err := http.ListenAndServe(hostStr, r); err != nil {
		close(ch)
		log.Println("Unable to open up TLS listener")
		log.Fatal("ListenAndServe: ", err)
	}
}
