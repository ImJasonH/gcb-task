/*
Copyright 2020 Google, Inc.

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

// Code generated by injection-gen. DO NOT EDIT.

package fake

import (
	context "context"

	externalversions "github.com/imjasonh/gcb-task/pkg/client/cloudbuild/informers/externalversions"
	fake "github.com/imjasonh/gcb-task/pkg/client/cloudbuild/injection/client/fake"
	factory "github.com/imjasonh/gcb-task/pkg/client/cloudbuild/injection/informers/factory"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
)

var Get = factory.Get

func init() {
	injection.Fake.RegisterInformerFactory(withInformerFactory)
}

func withInformerFactory(ctx context.Context) context.Context {
	c := fake.Get(ctx)
	opts := make([]externalversions.SharedInformerOption, 0, 1)
	if injection.HasNamespaceScope(ctx) {
		opts = append(opts, externalversions.WithNamespace(injection.GetNamespaceScope(ctx)))
	}
	return context.WithValue(ctx, factory.Key{},
		externalversions.NewSharedInformerFactoryWithOptions(c, controller.GetResyncPeriod(ctx), opts...))
}
