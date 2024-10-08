// Code generated by timoni.
// Note that this file must have no imports and all values must be concrete.

@if(!debug)

package main

// Placeholder for user-supplied values.
values: {
	image: {
		repository: "ghcr.io/zenginedev/website"
		tag:        "0.0.4"
		digest:     ""
	}

	pod: {
		annotations: "cluster-autoscaler.kubernetes.io/safe-to-evict": "true"
		imagePullSecrets: [{
			name: "regcred"
		}]
	}

    service: type: "ClusterIP"


	ingress: {
		annotations: {
            "cert-manager.io/cluster-issuer": "letsencrypt"
            "kubernetes.io/tls-acme": "true"
        }
	}

	resources: {
		limits: {
			cpu:    "100m"
			memory: "128Mi"
		}
	}
}
