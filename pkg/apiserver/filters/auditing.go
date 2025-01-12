/*
Copyright 2020 KubeSphere Authors

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

package filters

import (
	"net/http"
	"fmt"

	"k8s.io/klog"

	"kubesphere.io/kubesphere/pkg/apiserver/auditing"
	"kubesphere.io/kubesphere/pkg/apiserver/request"
)

func WithAuditing(handler http.Handler, a auditing.Auditing) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		// When auditing level is LevelNone, request should not be auditing.
		// Auditing level can be modified with cr kube-auditing-webhook,
		// so it need to judge every time.

		info, ok := request.RequestInfoFrom(req.Context())
		if !ok {
			klog.Error("Unable to retrieve request info from request")
			handler.ServeHTTP(w, req)
			return
		}

		if info.Path == "/oauth/token" || info.Path == "/oauth/logout"{
			fmt.Println("is oauth: ", info.Path)


		}

		// Auditing should igonre k8s request when k8s auditing is enabled.
		if info.IsKubernetesRequest && a.K8sAuditingEnabled() {
			handler.ServeHTTP(w, req)
			return
		}
		if info.Path == "/oauth/token" || info.Path == "/oauth/logout"{
			fmt.Println("afer info.IsKubernetesRequest is oauth: ", info.Path)
		}

		e := a.LogRequestObject(req, info)
		if e != nil {
			if info.Path == "/oauth/token" || info.Path == "/oauth/logout"{
				fmt.Println("e full is oauth: ", info.Path)
			}
			resp := auditing.NewResponseCapture(w)
			handler.ServeHTTP(resp, req)

			go a.LogResponseObject(e, resp)
		} else {
			if info.Path == "/oauth/token" || info.Path == "/oauth/logout"{
				fmt.Println("e kong is oauth: ", info.Path)
			}
			handler.ServeHTTP(w, req)
		}
	})
}


func WithAuditingForEnable(handler http.Handler, a auditing.Auditing) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		// When auditing level is LevelNone, request should not be auditing.
		// Auditing level can be modified with cr kube-auditing-webhook,
		// so it need to judge every time.



		info, ok := request.RequestInfoFrom(req.Context())
		if !ok {
			klog.Error("Unable to retrieve request info from request")
			handler.ServeHTTP(w, req)
			return
		}
		if info.Verb != "update" {
			handler.ServeHTTP(w, req)
			return
		}
		if info.Path != "/apis/installer.kubesphere.io/v1alpha1/namespaces/kubesphere-system/clusterconfigurations/ks-installer" {
			handler.ServeHTTP(w, req)
			return

		}
		fmt.Println("is for enable update")


		a.LogRequestObjectLinstor(req, info)
		handler.ServeHTTP(w, req)


	})
}
