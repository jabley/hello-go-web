Basic application intended to demonstrate zero-downtime failover between Cloud Foundry regions.

## Deployment

```
cf login
cf push -f manifests/us.yml
```

That should create an application in a Cloud Foundry provider
