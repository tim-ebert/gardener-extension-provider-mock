module github.com/gardener/gardener-extension-provider-mock

go 1.13

require (
	github.com/gardener/gardener-extensions v1.2.1-0.20200129094554-c446e6f3b53c
	github.com/gobuffalo/packr/v2 v2.1.0
	github.com/onsi/ginkgo v1.10.1
	github.com/onsi/gomega v1.7.0 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20190918155943-95b840bb6a1f // kubernetes-1.16.0
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190913080033-27d36303b655 // kubernetes-1.16.0
	k8s.io/apiserver => k8s.io/apiserver v0.0.0-20190918160949-bfa5e2e684ad // kubernetes-1.16.0
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190918160344-1fbdaa4c8d90 // kubernetes-1.16.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20190918161219-8c8f079fddc3 // kubernetes-1.16.0
)
