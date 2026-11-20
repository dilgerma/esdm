package rules_test

// eventModelingParents is the core model the event-modeling rule
// tests build on: everything a slice or screen can refer to, with
// no aggregate, as an event model needs none. Each test appends the
// event-modeling documents it checks, separated by `---`.
const eventModelingParents = `apiVersion: schema.esdm.io/core/v1
kind: domain
name: d
---
apiVersion: schema.esdm.io/core/v1
kind: bounded-context
name: bc
scope:
  domain: d
---
apiVersion: schema.esdm.io/core/v1
kind: actor
name: user
scope:
  domain: d
  boundedContext: bc
type: human
---
apiVersion: schema.esdm.io/core/v1
kind: dynamic-consistency-boundary
name: order-boundary
scope:
  domain: d
  boundedContext: bc
identifiedBy:
  - name: order-id
    source: generated
    generator: uuid
consults:
  - boundedContext: bc
    event: order-placed
    criteria: Events published by place-order.
---
apiVersion: schema.esdm.io/core/v1
kind: command
name: place-order
scope:
  domain: d
  boundedContext: bc
  dynamicConsistencyBoundary: order-boundary
data:
  type: object
publishes:
  - order-placed
---
apiVersion: schema.esdm.io/core/v1
kind: event
name: order-placed
scope:
  domain: d
  boundedContext: bc
data:
  type: object
---
apiVersion: schema.esdm.io/core/v1
kind: read-model
name: orders
scope:
  domain: d
  boundedContext: bc
schema:
  type: object
projections:
  - boundedContext: bc
    event: order-placed
    rule: Project order-placed into orders.
---
apiVersion: schema.esdm.io/core/v1
kind: query
name: orders-query
scope:
  domain: d
  boundedContext: bc
readModel: orders
result:
  type: object
---
apiVersion: schema.esdm.io/core/v1
kind: policy
name: shipper
scope:
  domain: d
deliveryGuarantee: at-most-once
handles:
  - boundedContext: bc
    event: order-placed
emits:
  - boundedContext: bc
    dynamicConsistencyBoundary: order-boundary
    command: place-order
---
apiVersion: schema.esdm.io/core/v1
kind: external-system
name: payment-provider
scope:
  domain: d
direction: inbound
`
