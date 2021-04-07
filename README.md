# Demo-Controller

## What does this controller do?

Not much. It watches for events related to lifetime of pods and logs messages about them.

## Building
You need `dep`. Get and install it here: [https://github.com/golang/dep](https://github.com/golang/dep). Then run,
```
# to fetch dependencies
dep ensure
# to build the whole thing
make
```

## Running
Make sure your `kubectl` is working. 

### Running as standalone binary
Just run `./demo-controller`. 

### Running as pod in a cluster
*  set `DOCKER_REPO` variable in [`Makefile`](Makefile) to point to a docker registry where you can store the image
*  run `make build-image` to build locally a docker image with your controller
*  run `make push-image` to push it to the registry
*  edit [`demo-controller.yaml`](demo-controller.yaml) and change `image: YOUR_URL:TAG` to pint to your image registry and the version tag you want to deploy
*  run `kubectl create -f demo-controller.yaml`

