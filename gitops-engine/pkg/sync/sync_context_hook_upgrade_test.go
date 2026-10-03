package sync

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic/fake"

	synccommon "github.com/argoproj/argo-cd/gitops-engine/v3/pkg/sync/common"
	"github.com/argoproj/argo-cd/gitops-engine/v3/pkg/utils/kube"
	"github.com/argoproj/argo-cd/gitops-engine/v3/pkg/utils/kube/kubetest"
	testingutils "github.com/argoproj/argo-cd/gitops-engine/v3/pkg/utils/testing"
)

// A tracked resource can be converted into a hook that keeps the same name,
// for example a Job that gains the Sync hook annotation. The old object is
// pruned and the hook must still be created in the same operation.
func TestSync_TrackedResourceConvertedToSameNameHook(t *testing.T) {
	for _, hookType := range []synccommon.HookType{synccommon.HookTypeSync, synccommon.HookTypePostSync} {
		t.Run(string(hookType), func(t *testing.T) {
			live := testingutils.NewPod()
			live.SetNamespace(testingutils.FakeArgoCDNamespace)

			hookObj := testingutils.NewPod()
			hookObj.SetNamespace(testingutils.FakeArgoCDNamespace)
			hookObj = testingutils.Annotate(hookObj, synccommon.AnnotationKeyHook, string(hookType))
			hookObj = testingutils.Annotate(hookObj, synccommon.AnnotationKeyHookDeletePolicy, string(synccommon.HookDeletePolicyBeforeHookCreation))

			syncCtx := newTestSyncCtx(nil, WithPrune(true))
			syncCtx.resources = groupResources(ReconciliationResult{
				Live:   []*unstructured.Unstructured{live},
				Target: []*unstructured.Unstructured{nil},
			})
			syncCtx.hooks = []*unstructured.Unstructured{hookObj}
			syncCtx.dynamicIf = fake.NewSimpleDynamicClient(runtime.NewScheme(), live)

			// The first pass prunes the old, non-hook object and deletes it before
			// creating the hook, so the create itself happens on the next pass.
			syncCtx.Sync(context.Background())
			// The next reconciliation no longer sees the pruned object.
			syncCtx.resources = map[kube.ResourceKey]reconciledResource{}
			syncCtx.Sync(context.Background())

			phase, message, results := syncCtx.GetState()
			resourceOps, _ := syncCtx.resourceOps.(*kubetest.MockResourceOps)
			assert.NotEmpty(t, resourceOps.GetLastResourceCommand(kube.GetResourceKey(hookObj)),
				"the %s hook was never applied; phase=%s message=%q results=%+v", hookType, phase, message, results)
		})
	}
}
