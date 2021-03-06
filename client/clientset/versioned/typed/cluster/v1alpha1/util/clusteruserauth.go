/*
Copyright 2020 AppsCode Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"context"
	"fmt"

	api "go.bytebuilders.dev/resource-model/apis/cluster/v1alpha1"
	cs "go.bytebuilders.dev/resource-model/client/clientset/versioned/typed/cluster/v1alpha1"

	jsonpatch "github.com/evanphx/json-patch"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
	kutil "kmodules.xyz/client-go"
)

func CreateOrPatchClusterUserAuth(ctx context.Context, c cs.ClusterV1alpha1Interface, meta metav1.ObjectMeta, transform func(in *api.ClusterUserAuth) *api.ClusterUserAuth, opts metav1.PatchOptions) (*api.ClusterUserAuth, kutil.VerbType, error) {
	cur, err := c.ClusterUserAuths().Get(ctx, meta.Name, metav1.GetOptions{})
	if kerr.IsNotFound(err) {
		klog.V(3).Infof("Creating ClusterUserAuth %s.", meta.Name)
		out, err := c.ClusterUserAuths().Create(ctx, transform(&api.ClusterUserAuth{
			TypeMeta: metav1.TypeMeta{
				Kind:       api.ResourceKindClusterUserAuth,
				APIVersion: api.SchemeGroupVersion.String(),
			},
			ObjectMeta: meta,
		}), metav1.CreateOptions{
			DryRun:       opts.DryRun,
			FieldManager: opts.FieldManager,
		})
		return out, kutil.VerbCreated, err
	} else if err != nil {
		return nil, kutil.VerbUnchanged, err
	}
	return PatchClusterUserAuth(ctx, c, cur, transform, opts)
}

func PatchClusterUserAuth(ctx context.Context, c cs.ClusterV1alpha1Interface, cur *api.ClusterUserAuth, transform func(*api.ClusterUserAuth) *api.ClusterUserAuth, opts metav1.PatchOptions) (*api.ClusterUserAuth, kutil.VerbType, error) {
	return PatchClusterUserAuthObject(ctx, c, cur, transform(cur.DeepCopy()), opts)
}

func PatchClusterUserAuthObject(ctx context.Context, c cs.ClusterV1alpha1Interface, cur, mod *api.ClusterUserAuth, opts metav1.PatchOptions) (*api.ClusterUserAuth, kutil.VerbType, error) {
	curJson, err := json.Marshal(cur)
	if err != nil {
		return nil, kutil.VerbUnchanged, err
	}

	modJson, err := json.Marshal(mod)
	if err != nil {
		return nil, kutil.VerbUnchanged, err
	}

	patch, err := jsonpatch.CreateMergePatch(curJson, modJson)
	if err != nil {
		return nil, kutil.VerbUnchanged, err
	}
	if len(patch) == 0 || string(patch) == "{}" {
		return cur, kutil.VerbUnchanged, nil
	}
	klog.V(3).Infof("Patching ClusterUserAuth %s with %s.", cur.Name, string(patch))
	out, err := c.ClusterUserAuths().Patch(ctx, cur.Name, types.MergePatchType, patch, opts)
	return out, kutil.VerbPatched, err
}

func TryUpdateClusterUserAuth(ctx context.Context, c cs.ClusterV1alpha1Interface, meta metav1.ObjectMeta, transform func(*api.ClusterUserAuth) *api.ClusterUserAuth, opts metav1.UpdateOptions) (result *api.ClusterUserAuth, err error) {
	attempt := 0
	err = wait.PollImmediate(kutil.RetryInterval, kutil.RetryTimeout, func() (bool, error) {
		attempt++
		cur, e2 := c.ClusterUserAuths().Get(ctx, meta.Name, metav1.GetOptions{})
		if kerr.IsNotFound(e2) {
			return false, e2
		} else if e2 == nil {
			result, e2 = c.ClusterUserAuths().Update(ctx, transform(cur.DeepCopy()), opts)
			return e2 == nil, nil
		}
		klog.Errorf("Attempt %d failed to update ClusterUserAuth %s due to %v.", attempt, cur.Name, e2)
		return false, nil
	})

	if err != nil {
		err = fmt.Errorf("failed to update ClusterUserAuth %s after %d attempts due to %v", meta.Name, attempt, err)
	}
	return
}

func UpdateClusterUserAuthStatus(
	ctx context.Context,
	c cs.ClusterV1alpha1Interface,
	meta metav1.ObjectMeta,
	transform func(*api.ClusterUserAuthStatus) *api.ClusterUserAuthStatus,
	opts metav1.UpdateOptions,
) (result *api.ClusterUserAuth, err error) {
	apply := func(x *api.ClusterUserAuth) *api.ClusterUserAuth {
		out := &api.ClusterUserAuth{
			TypeMeta:   x.TypeMeta,
			ObjectMeta: x.ObjectMeta,
			Spec:       x.Spec,
			Status:     *transform(x.Status.DeepCopy()),
		}
		return out
	}

	attempt := 0
	cur, err := c.ClusterUserAuths().Get(ctx, meta.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	err = wait.PollImmediate(kutil.RetryInterval, kutil.RetryTimeout, func() (bool, error) {
		attempt++
		var e2 error
		result, e2 = c.ClusterUserAuths().UpdateStatus(ctx, apply(cur), opts)
		if kerr.IsConflict(e2) {
			latest, e3 := c.ClusterUserAuths().Get(ctx, meta.Name, metav1.GetOptions{})
			switch {
			case e3 == nil:
				cur = latest
				return false, nil
			case kutil.IsRequestRetryable(e3):
				return false, nil
			default:
				return false, e3
			}
		} else if err != nil && !kutil.IsRequestRetryable(e2) {
			return false, e2
		}
		return e2 == nil, nil
	})

	if err != nil {
		err = fmt.Errorf("failed to update status of ClusterUserAuth %s after %d attempts due to %v", meta.Name, attempt, err)
	}
	return
}
