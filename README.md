# envoy-authz

## 原理
### istio-envoyfileter
```yaml
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: test-ingress
  namespace: istio-system
spec:
  workloadSelector:
    labels:
      istio: ingressgateway
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: ANY
        listener:
          filterChain:
            filter:
              name: "envoy.http_connection_manager"
      patch:
        operation: INSERT_BEFORE
        value:
          name: envoy.ext_authz
          config:
            grpc_service:
              # NOTE: *SHOULD* use envoy_grpc as ext_authz can use dynamic clusters and has connection pooling
              google_grpc:
                target_uri: envoy-authz.test:9999
                stat_prefix: ext_authz
              timeout: 0.2s
            failure_mode_allow: false
            with_request_body:
              max_request_bytes: 8192
              allow_partial_message: true
```
### envoy authz

### jwt