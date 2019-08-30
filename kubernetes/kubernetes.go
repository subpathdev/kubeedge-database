package kubernetes

import (
	"github.com/kubeedge/kubeedge/cloud/pkg/apis/devices/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
)

type ResourceEventHandler struct {
	events chan watch.Event
}

func (r ResourceEventHandler) obj2Event(typ watch.EventType, obj interface{}) {
	eventObj, ok := obj.(runtime.Object)
	if !ok {
		klog.Errorf("unknown type: %T, ignore", obj)
		return
	}

	r.events <- watch.Event{Type: typ, Object: eventObj}
}

func (r ResourceEventHandler) OnAdd(obj interface{}) {
	klog.Infof("adding new device")
	r.obj2Event(watch.Added, obj)
}

func (r ResourceEventHandler) OnUpdate(oldObj, obj interface{}) {
	klog.Infof("modify device")
	r.obj2Event(watch.Modified, obj)
}

func (r ResourceEventHandler) OnDelete(obj interface{}) {
	klog.Infof("delete device")
	r.obj2Event(watch.Deleted, obj)
}

var kubernetesRestClient *rest.RESTClient

func createScheme(schem *runtime.Scheme) error {
	schem.AddKnownTypes(schema.GroupVersion{Group: "devices.kubeedge.io", Version: "v1alpha1"}, &v1alpha1.Device{}, &v1alpha1.DeviceList{})
	metav1.AddToGroupVersion(schem, schema.GroupVersion{Group: "devices.kubeedge.io", Version: "v1alpha1"})
	return nil
}

func Init(k8sServer, configPath string, event chan watch.Event) error {
	klog.Infof("init kubernetes connection")

	schem := runtime.NewScheme()
	schemaBuilder := runtime.NewSchemeBuilder(createScheme)

	if err := schemaBuilder.AddToScheme(schem); err != nil {
		klog.Errorf("can not build schema; err: %v\n", err)
		return err
	}

	conf, err := clientcmd.BuildConfigFromFlags(k8sServer, configPath)
	if err != nil {
		klog.Errorf("can not connect to kubernetes api server; err: %v\n", err)
		return err
	}

	conf.ContentType = runtime.ContentTypeJSON
	conf.APIPath = "/apis"
	conf.GroupVersion = &schema.GroupVersion{Group: "devices.kubeedge.io", Version: "v1alpha1"}
	conf.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: serializer.NewCodecFactory(schem)}

	kubernetesRestClient, err = rest.RESTClientFor(conf)
	if err != nil {
		klog.Errorf("can not create REST client; err: %v\n", err)
		return err
	}

	var dev v1alpha1.Device
	lw := cache.NewListWatchFromClient(kubernetesRestClient, "devices", "kubeedge", fields.Everything())
	si := cache.NewSharedInformer(lw, &dev, 0)
	reh := ResourceEventHandler{events: event}
	si.AddEventHandler(reh)

	klog.Infof("value of v1alpha1: %v\n", dev)

	stopNever := make(chan struct{})
	go si.Run(stopNever)

	return nil
}
