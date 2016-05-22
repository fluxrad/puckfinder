Puckfinder
==========

Find drop-in hockey in your area

#### Install

```
go get github.com/fluxrad/puckfinder
```

#### Misc

`Dockerfile.ng` is used to bootstrap a new angular project. I created the `app` directory with the following:

````
docker build -f Dockerfile.ng -t ng .
docker run -v `pwd`:/puckfinder ng new puckfinder -sg -dir /puckfinder/app
```
