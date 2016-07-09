Basic application intended to demonstrate zero-downtime failover between
Cloud Foundry regions.

## Deployment

```
cf login
cf push -f manifests/us.yml
```

That should create an application in a Cloud Foundry provider. From that
manifest, the app should be named hello-go-web.

## Consistent Routing

We want to ensure that any Cloud Foundry provider we use will respond to
requests that look the same from our load-balancer. Assume our
load-balancer will be making requests to hello-go-web.example.com.

First, register the domain example.com in Cloud Foundry:

```
cf create-domain my-org example.com
```

Create a route so that requests to hello-go-web.example.com are recognised:

```
cf create-route my-space example.com  --hostname hello-go-web
```

Finally, map the route to your recently pushed application:

```
cf map-route hello-go-web example.com --hostname hello-go-web
```

Now you should be able to make a request to your application as so:

```
curl -H 'hello-go-web.example.com' http://hello-go-web.some.cloud-foundry.provider.com
```

Repeat that for as many Cloud Foundry providers as you wish, then configure
Fastly to use them as backends.

## Fastly

Other options are available, but Fastly is easy to get going and has a free
tier with TLS.

### Create a service

1. Host of hello-go-web.cfapps.io (if that's where your app is running on Cloud Foundry)
1. Domain of `foo.global.ssl.fastly.net` where `foo` is any available word

### Create a Healthcheck

1. Make requests to `/_status`

### Edit Hosts

1. Add the healthcheck in to your host

### Edit Settings

1. Add a Default Host of `hello-go-web.example.com`

### Add more hosts

1. Add a host for each Cloud Foundry provider that you've deployed the app to

### Test

Once that has been activated, you should be able to access
https://foo.global.ssl.fastly.net/ and see your service running.
