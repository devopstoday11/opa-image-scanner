/*
Copyright 2018 The Kubernetes Authors.
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

package main

// import (
// 	"fmt"
// 	"strings"

// 	corev1 "k8s.io/api/core/v1"
// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// 	"k8s.io/api/admission/v1beta1"
// 	"k8s.io/klog"
// )

// // only allow pods to pull images from specific registry.
// func admitPods(ar v1beta1.AdmissionReview) *v1beta1.AdmissionResponse {
// 	klog.Info("admitting pods")
// 	podResource := metav1.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"}
// 	if ar.Request.Resource != podResource {
// 		err := fmt.Errorf("expect resource to be %s", podResource)
// 		klog.Error(err)
// 		return toAdmissionResponse(err)
// 	}

// 	raw := ar.Request.Object.Raw
// 	pod := corev1.Pod{}
// 	deserializer := codecs.UniversalDeserializer()
// 	if _, _, err := deserializer.Decode(raw, nil, &pod); err != nil {
// 		klog.Error(err)
// 		return toAdmissionResponse(err)
// 	}
// 	reviewResponse := v1beta1.AdmissionResponse{}
// 	reviewResponse.Allowed = true

// 	var msg string
// 	if v, ok := pod.Labels["webhook-e2e-test"]; ok {
// 		if v == "webhook-disallow" {
// 			reviewResponse.Allowed = false
// 			msg = msg + "the pod contains unwanted label; "
// 		}
// 		if v == "wait-forever" {
// 			reviewResponse.Allowed = false
// 			msg = msg + "the pod response should not be sent; "
// 			<-make(chan int) // Sleep forever - no one sends to this channel
// 		}
// 	}
// 	for _, container := range pod.Spec.Containers {
// 		if strings.Contains(container.Name, "webhook-disallow") {
// 			reviewResponse.Allowed = false
// 			msg = msg + "the pod contains unwanted container name; "
// 		}
// 	}
// 	if !reviewResponse.Allowed {
// 		reviewResponse.Result = &metav1.Status{Message: strings.TrimSpace(msg)}
// 	}
// 	return &reviewResponse
// }
