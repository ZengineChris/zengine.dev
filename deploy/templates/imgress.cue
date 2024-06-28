package templates

import (
	netv1 "k8s.io/api/networking/v1"
)

#Ingress: netv1.#Ingress & {
    #config:    #Config
	apiVersion: "networking.k8s.io/v1"
	kind:       "Ingress"
	metadata:   #config.metadata
    if #config.ingress.annotations != _|_ {
		metadata: annotations: #config.ingress.annotations
	}
    spec: netv1.#IngressSpec & {
        ingressClassName: "nginx"
        tls: [
            netv1.#IngressTLS & {
                hosts: ["zengine.dev"]
                secretName: "zengine-dev-tls"
            }
        ]
        rules: [
            netv1.#IngressRule &{
               host: "zengine.dev" 
               http: {
                    paths: [
                        {
                            path: "/"
                            pathType: "Prefix"
                            backend: {
                                service: {
                                    name: "website"
                                    port: {
                                        name: "http" 
                                    }
                                }
                            }
                        }
                    ]
                }
            }
        ]

	} 
} 
