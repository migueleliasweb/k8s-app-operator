# CONCEPT

## APP CRD
```yaml
core:
    name: ""
    image: ""
    labels: {}
    annotations: {}
    envVars: {}
    cmd: []
    args: []

storage:
    ephemeral:
        writablePaths:
        - ""
    persistent: {} # TBD

capacity:
  rollout:
    maxUnavailable: "20%"
  replicas: 3
  autoscaling:
    enabled: true
    maxReplicas: 10
    targetCPUUtilizationPercentage: 80
    targetMemoryUtilizationPercentage: 80
  disruption:
    allowedDisruption: "20%"
  resources:
    limits:
      cpu: 300m
      memory: 512Mi
    requests:
      cpu: 300m
      memory: 512Mi

ingress:
#   enabled: false
  healthchecks:
    enabled: true
    http:
        livenessPath: /healthz
        livenessPort: 8080
        livenessPath: /readyz
        livenessPort: 8080
    exec: {}
  fromNetworkCIDR:
  - cidr: 10.10.0.0/24
    port: 1234
  fromK8sPods:
  - labels:
        foo: bar
    port: 1234
    namespace: ""

egress:
#   enabled: false
  toNetworkCIDR:
  - cidr: 10.10.0.0/24
    port: 1234
  toK8sPods:
  - labels:
        foo: bar
    port: 1234
    namespace: ""

identity:
  IAMRole: ""
```

## Compliementary CRDs

APPOVerrides and APPDefaults should use the same CRD specs as a normal APP

- APPOverrides: APP -> APPOverrides = Result
- APPDefaults: AppDefaults -> APP = Result
- GlobalAppOverrides: APP -> GlobalAppOverrides -> APPOverrides = Result
- GlobalAPPDefaults: GlobalAPPDefaults -> AppDefaults -> APP = Result

## Platform level CRDs

These CRDs will give full access to the underlying types spec:

```
DeploymentSpec
ServiceSpec
HPASpec
IngressSpec
EgressSpec
```

- APPSpec and GlobalAppSpec: Result -> (Global)APPSpec = NewResult
