// Code generated by injection-gen. DO NOT EDIT.

package fake

import (
	"context"

	subscription "github.com/kyma-project/kyma/components/event-bus/client/generated/injection/informers/eventing/v1alpha1/subscription"
	fake "github.com/kyma-project/kyma/components/event-bus/client/generated/injection/informers/factory/fake"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
)

var Get = subscription.Get

func init() {
	injection.Fake.RegisterInformer(withInformer)
}

func withInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := fake.Get(ctx)
	inf := f.Eventing().V1alpha1().Subscriptions()
	return context.WithValue(ctx, subscription.Key{}, inf), inf.Informer()
}
