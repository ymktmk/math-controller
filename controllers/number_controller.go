/*
Copyright 2022.

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

package controllers

import (
	"context"
	"fmt"
	"math"

	"github.com/pkg/errors"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	mathv1beta1 "github.com/ymktmk/math-controller/api/v1beta1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

type NumberReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=math.ymktmk.github.io,resources=numbers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=math.ymktmk.github.io,resources=numbers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=math.ymktmk.github.io,resources=numbers/finalizers,verbs=update

func (r *NumberReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	obj := mathv1beta1.Number{}
	// namespaceのNumberリソースを探す
	if err := r.Client.Get(ctx, req.NamespacedName, &obj); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, errors.WithStack(err)
	}

	{
		// manifestのvalueを受け取ってFizzbuzzする
		obj.Status.FizzBuzz = fizzbuzz(obj.Spec.Value)
		obj.Status.IsSquare = isSquare(obj.Spec.Value)
		logger.Info("New Status", "status", obj.Status)
		if err := r.Status().Update(ctx, &obj); err != nil {
			return ctrl.Result{}, errors.WithStack(err)
		}
	}

	return ctrl.Result{}, nil
}

func fizzbuzz(n int64) string {
	if n%15 == 0 {
		return "FizzBuzz"
	}
	if n%3 == 0 {
		return "Fizz"
	}
	if n%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprintf("%d", n)
}

func isSquare(num int64) bool {
	sqrt := int64(math.Floor(math.Sqrt(float64(num))))
	return sqrt*sqrt == num
}

func (r *NumberReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&mathv1beta1.Number{}).
		Complete(r)
}
