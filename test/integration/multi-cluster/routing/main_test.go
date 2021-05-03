package main

import (
	"fmt"
	"github.com/solo-io/gloo-mesh/pkg/test/apps/context"
	echo2 "github.com/solo-io/gloo-mesh/pkg/test/apps/echo"
	gloo_mesh "github.com/solo-io/gloo-mesh/pkg/test/apps/gloo-mesh"
	"github.com/solo-io/gloo-mesh/pkg/test/common"
	"istio.io/istio/pkg/test/framework/components/echo"
	"istio.io/istio/pkg/test/framework/components/environment/kube"
	"istio.io/istio/pkg/test/framework/config"
	"istio.io/istio/pkg/test/framework/label"
	"istio.io/istio/pkg/test/framework/resource"
	"net/http"
	"os"
	"testing"

	"istio.io/istio/pkg/test/framework"
	"istio.io/istio/pkg/test/framework/components/istio"
)

var (
	i             istio.Instance
	deploymentCtx context.DeploymentContext
)

func TestMain(m *testing.M) {

	licenceKey := os.Getenv("GLOO_MESH_LICENSE_KEY")
	// get kube settings from command line
	config.Parse()
	kubeSettings, _ := kube.NewSettingsFromCommandLine()
	clusterKubeConfigs := make(map[string]string)

	// this is a hack to match the kube configs with the cluster names so we can match them when deploy happens
	for i, k := range kubeSettings.KubeConfig {
		clusterKubeConfigs[fmt.Sprintf("cluster-%d", i)] = k
	}

	framework.
		NewSuite(m).
		Label(label.Multicluster).
		RequireMinClusters(2).
		Setup(istio.Setup(&i, common.IstioSetupFunc("gloo-mesh-istio.yaml"))).
		Setup(gloo_mesh.Deploy(&deploymentCtx, &gloo_mesh.Config{
			ClusterKubeConfigs:                  clusterKubeConfigs,
			DeployControlPlaneToManagementPlane: true,
		},
			licenceKey)).
		Setup(echo2.DeployEchos(&deploymentCtx)).
		Run()
}

// Run the API tests
func TestRouting(t *testing.T) {
	framework.
		NewTest(t).
		Run(func(ctx framework.TestContext) {

			tgs := []common.TestGroup{
				{
					Name: "routing",
					Cases: []common.TestCase{
						{
							Name:        "prefix-1",
							Description: "HTTP/HTTPS prefix based routing",
							Test:        testPrefixMatch,
							Namespace:   deploymentCtx.EchoContext.AppNamespace.Name(),
							FileName:    "prefix-1.yaml",
						},
					},
				},
			}
			for _, tg := range tgs {
				tg.Run(ctx, t, &deploymentCtx)
			}
		})
}

// testPrefixMatch makes a call from frontend to backend application
func testPrefixMatch(ctx resource.Context, t *testing.T, deploymentCtx *context.DeploymentContext) {
	src := deploymentCtx.EchoContext.Deployments.GetOrFail(t, echo.Service("frontend"))

	backendHost := "backend." + deploymentCtx.EchoContext.AppNamespace.Name() + ".svc.cluster.local"

	src.CallOrFail(t, echo.CallOptions{
		Port: &echo.Port{
			Protocol:    "http",
			ServicePort: 8090,
		},
		Scheme:    "http",
		Address:   backendHost,
		Method:    http.MethodGet,
		Path:      "/route1",
		Count:     1,
		Validator: echo.ExpectOK(),
	})

	src.CallOrFail(t, echo.CallOptions{
		Port: &echo.Port{
			Protocol:    "http",
			ServicePort: 8090,
		},
		Scheme:    "http",
		Address:   backendHost,
		Method:    http.MethodGet,
		Path:      "/route2",
		Count:     1,
		Validator: echo.ExpectOK(),
	})

	src.CallOrFail(t, echo.CallOptions{
		Port: &echo.Port{
			Protocol:    "http",
			ServicePort: 8090,
		},
		Scheme:    "http",
		Address:   backendHost,
		Method:    http.MethodGet,
		Path:      "/bad-route",
		Count:     1,
		Validator: echo.ExpectCode("404"),
	})
}
