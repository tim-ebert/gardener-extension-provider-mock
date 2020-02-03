module github.com/gardener/gardener-extension-provider-mock

go 1.13

require (
	github.com/gardener/gardener v0.35.1-0.20200128130120-5b69a02f511a
	github.com/gardener/gardener-extensions v1.2.1-0.20200129094554-c446e6f3b53c
	github.com/gobuffalo/packr/v2 v2.1.0
	github.com/onsi/ginkgo v1.10.1
	github.com/onsi/gomega v1.7.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	k8s.io/apimachinery v0.0.0-20191016060620-86f2f1b9c076
	k8s.io/component-base v0.0.0-20190918160511-547f6c5d7090
	sigs.k8s.io/controller-runtime v0.4.0
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20190918155943-95b840bb6a1f // kubernetes-1.16.0
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190913080033-27d36303b655 // kubernetes-1.16.0
	k8s.io/apiserver => k8s.io/apiserver v0.0.0-20190918160949-bfa5e2e684ad // kubernetes-1.16.0
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190918160344-1fbdaa4c8d90 // kubernetes-1.16.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20190918161219-8c8f079fddc3 // kubernetes-1.16.0
)
