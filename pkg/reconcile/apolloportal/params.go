package apolloportal

import (
	"apollo.io/apollo-operator/api/v1alpha1"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Params holds the reconciliation-specific parameters.
type Params struct {
	Client   client.Client
	Recorder record.EventRecorder
	Scheme   *runtime.Scheme
	Log      logr.Logger
	Instance v1alpha1.ApolloPortal
	//Config   config.Config
}
