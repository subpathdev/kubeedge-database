module github.com/subpathdev/kubeedge-database

go 1.12

require (
	github.com/coreos/etcd v3.3.15+incompatible // indirect
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/go-chassis/paas-lager v1.1.0 // indirect
	github.com/go-openapi/jsonreference v0.19.2 // indirect
	github.com/go-openapi/strfmt v0.17.2 // indirect
	github.com/go-openapi/swag v0.19.5 // indirect
	github.com/gogo/protobuf v1.2.2-0.20190723190241-65acae22fc9d // indirect
	github.com/googleapis/gnostic v0.3.1 // indirect
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/jessevdk/go-flags v1.4.1-0.20181221193153-c0795c8afcf4
	github.com/kr/pty v1.1.8 // indirect
	github.com/kubeedge/beehive v0.0.0-20190809132808-14b9c1bfd040 // indirect
	github.com/kubeedge/kubeedge v1.0.1-0.20190830034131-1b69040f5f86
	github.com/lib/pq v1.0.0
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/marten-seemann/qtls v0.2.4 // indirect
	github.com/mattn/go-isatty v0.0.9 // indirect
	github.com/mattn/go-shellwords v1.0.6 // indirect
	github.com/ncw/swift v1.0.49 // indirect
	github.com/prometheus/procfs v0.0.4 // indirect
	github.com/quobyte/api v0.1.5 // indirect
	github.com/rogpeppe/go-internal v1.3.1 // indirect
	github.com/vmware/govmomi v0.20.2 // indirect
	github.com/yvasiyarov/gorelic v0.0.7 // indirect
	golang.org/x/crypto v0.0.0-20190820162420-60c769a6c586 // indirect
	google.golang.org/appengine v1.6.2 // indirect
	google.golang.org/grpc v1.22.2 // indirect
	gopkg.in/gcfg.v1 v1.2.3 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	honnef.co/go/tools v0.0.1-2019.2.2 // indirect
	k8s.io/apimachinery v0.0.0
	k8s.io/client-go v0.0.0
	k8s.io/heapster v1.2.0 // indirect
	k8s.io/klog v0.4.0
	k8s.io/kubernetes v1.15.3 // indirect
	k8s.io/utils v0.0.0-20190809000727-6c36bc71fc4a // indirect
)

replace (
	github.com/Sirupsen/logrus v1.0.5 => github.com/sirupsen/logrus v1.0.5
	github.com/Sirupsen/logrus v1.3.0 => github.com/Sirupsen/logrus v1.0.6
	github.com/Sirupsen/logrus v1.4.0 => github.com/sirupsen/logrus v1.0.6
	k8s.io/api v0.0.0 => k8s.io/api v0.0.0-20190720062849-3043179095b6
	k8s.io/apiextensions-apiserver v0.0.0 => k8s.io/apiextensions-apiserver v0.0.0-20190718185103-d1ef975d28ce // indirect
	k8s.io/apimachinery v0.0.0 => k8s.io/apimachinery v0.0.0-20190612165923-1799e75a0719
	k8s.io/apiserver v0.0.0 => k8s.io/apiserver v0.0.0-20190718184206-a1aa83af71a7
	k8s.io/cli-runtime v0.0.0 => k8s.io/cli-runtime v0.0.0-20190718185405-0ce9869d0015
	k8s.io/client-go v0.0.0 => k8s.io/client-go v0.0.0-20190718183610-8e956561bbf5 // indirect
	k8s.io/cloud-provider v0.0.0 => k8s.io/cloud-provider v0.0.0-20190718190308-f8e43aa19282 // indirect
	k8s.io/cluster-bootstrap v0.0.0 => k8s.io/cluster-bootstrap v0.0.0-20190718190146-f7b0473036f9
	k8s.io/code-generator v0.0.0 => k8s.io/code-generator v0.0.0-20190612165923-18da4a14b22b
	k8s.io/component-base v0.0.0 => k8s.io/component-base v0.0.0-20190718183727-0ececfbe9772
	k8s.io/cri-api v0.0.0 => k8s.io/cri-api v0.0.0-20190531030430-6117653b35f1
	k8s.io/csi-api v0.0.0 => k8s.io/csi-api v0.0.0-20190313123203-94ac839bf26c // indirect
	k8s.io/csi-translation-lib v0.0.0 => k8s.io/csi-translation-lib v0.0.0-20190718190424-bef8d46b95de
	k8s.io/gengo v0.0.0 => k8s.io/gengo v0.0.0-20190327210449-e17681d19d3a // indirect
	k8s.io/heapster => k8s.io/heapster v1.2.0-beta.1 // indirect
	k8s.io/klog => k8s.io/klog v0.4.0 // indirect
	k8s.io/kube-aggregator v0.0.0 => k8s.io/kube-aggregator v0.0.0-20190718184434-a064d4d1ed7a
	k8s.io/kube-controller-manager v0.0.0 => k8s.io/kube-controller-manager v0.0.0-20190718190030-ea930fedc880
	k8s.io/kube-openapi v0.0.0 => k8s.io/kube-openapi v0.0.0-20190718094010-3cf2ea392886 // indirect
	k8s.io/kube-proxy v0.0.0 => k8s.io/kube-proxy v0.0.0-20190718185641-5233cb7cb41e
	k8s.io/kube-scheduler v0.0.0 => k8s.io/kube-scheduler v0.0.0-20190718185913-d5429d807831
	k8s.io/kubelet v0.0.0 => k8s.io/kubelet v0.0.0-20190718185757-9b45f80d5747
	k8s.io/legacy-cloud-providers v0.0.0 => k8s.io/legacy-cloud-providers v0.0.0-20190718190548-039b99e58dbd
	k8s.io/metrics v0.0.0 => k8s.io/metrics v0.0.0-20190718185242-1e1642704fe6
	k8s.io/node-api v0.0.0 => k8s.io/node-api v0.0.0-20190717025432-9e6fdeee55cc // indirect
	k8s.io/repo-infra v0.0.0 => k8s.io/repo-infra v0.0.0-20181204233714-00fe14e3d1a3 // indirect
	k8s.io/sample-apiserver v0.0.0 => k8s.io/sample-apiserver v0.0.0-20190718184639-baafa86838c0
	k8s.io/utils v0.0.0 => k8s.io/utils v0.0.0-20190712204705-3dccf664f023
)
