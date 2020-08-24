# k8shpa
# Projeto Deploy Continuo
Parabéns! Você chegou a nossa fase final onde você consolidará seus conhecimentos participando de todo ciclo de deploy de uma aplicação.

Nessa fase, o objetivo principal será fazer o deploy automático da aplicação já desenvolvida em Go Lang nas fases passadas.

As regras do processo são:
- Quando qualquer push ou uma PR for relizada no Github em um branch diferente do Master, o processo de CI tem que ser executado.
- Quando um merge ou um push entrarem no branch Master, o processo de CI/CD deve ser chamado, fazendo assim o deploy de forma automática no Kubernetes.

Faça o upload de todo código em um repositório, incluindo os do Cloudbuild, para que possamos fazer a avaliação.

Boa sorte e conte sempre conosco!


# Imagem no Docker Hub
https://hub.docker.com/repository/docker/samysilveira/go-hpa


# CI/CD pipeline trigger log
Step #3 - "CD - Ajustando os arquivos do k8s":             cpu: "100m"
DONE
PUSH
Finished Step #4 - "CD - Atualizando k8s"
Step #4 - "CD - Atualizando k8s": deployment.apps/go-hpa configured
Step #4 - "CD - Atualizando k8s": Running: kubectl apply -f deployment-new.yaml
Step #4 - "CD - Atualizando k8s": kubeconfig entry generated for cluster-docker-gohpa.
Step #4 - "CD - Atualizando k8s": Fetching cluster endpoint and auth data.
Step #4 - "CD - Atualizando k8s": Running: gcloud container clusters get-credentials --project="vsamycitest" --zone="us-central1-c" "cluster-docker-gohpa"
Step #4 - "CD - Atualizando k8s": Already have image (with digest): gcr.io/cloud-builders/kubectl
Starting Step #4 - "CD - Atualizando k8s"
Finished Step #3 - "CD - Ajustando os arquivos do k8s"
Step #3 - "CD - Ajustando os arquivos do k8s":           limits:
Step #3 - "CD - Ajustando os arquivos do k8s":             cpu: "50m"
Step #3 - "CD - Ajustando os arquivos do k8s":           requests:
Step #3 - "CD - Ajustando os arquivos do k8s":         resources:
Step #3 - "CD - Ajustando os arquivos do k8s":         - containerPort: 8080
Step #3 - "CD - Ajustando os arquivos do k8s":         ports:
Step #3 - "CD - Ajustando os arquivos do k8s":         image: gcr.io/vsamycitest/go-hpa:582677a
Step #3 - "CD - Ajustando os arquivos do k8s":       - name: go-hpa
Step #3 - "CD - Ajustando os arquivos do k8s":       containers:
Step #3 - "CD - Ajustando os arquivos do k8s":     spec:
Step #3 - "CD - Ajustando os arquivos do k8s":         app: go-hpa
Step #3 - "CD - Ajustando os arquivos do k8s":       labels:
Step #3 - "CD - Ajustando os arquivos do k8s":     metadata:
Step #3 - "CD - Ajustando os arquivos do k8s":   template:
Step #3 - "CD - Ajustando os arquivos do k8s":       app: go-hpa
Step #3 - "CD - Ajustando os arquivos do k8s":     matchLabels:
Step #3 - "CD - Ajustando os arquivos do k8s":   selector:
Step #3 - "CD - Ajustando os arquivos do k8s":   replicas: 1
Step #3 - "CD - Ajustando os arquivos do k8s": spec:
Step #3 - "CD - Ajustando os arquivos do k8s":   name: go-hpa
Step #3 - "CD - Ajustando os arquivos do k8s": metadata: 
Step #3 - "CD - Ajustando os arquivos do k8s": kind: Deployment
Step #3 - "CD - Ajustando os arquivos do k8s": apiVersion: apps/v1
Step #3 - "CD - Ajustando os arquivos do k8s": Already have image (with digest): gcr.io/cloud-builders/gcloud
Starting Step #3 - "CD - Ajustando os arquivos do k8s"
Finished Step #2 - "CD - Salvando Imagem"
Step #2 - "CD - Salvando Imagem": 582677a: digest: sha256:c44c413bdb35f3a50dc896ef0a754581e309fc9a472510234ea29248a2a7eeb6 size: 528
Step #2 - "CD - Salvando Imagem": 71d50ec58f93: Pushed
Step #2 - "CD - Salvando Imagem": 71d50ec58f93: Preparing
Step #2 - "CD - Salvando Imagem": The push refers to repository [gcr.io/vsamycitest/go-hpa]
Step #2 - "CD - Salvando Imagem": Already have image (with digest): gcr.io/cloud-builders/docker
Starting Step #2 - "CD - Salvando Imagem"
Finished Step #1 - "CD - Construindo Imagem"
Step #1 - "CD - Construindo Imagem": Successfully tagged gcr.io/vsamycitest/go-hpa:582677a
Step #1 - "CD - Construindo Imagem": Successfully built d26f4c862b1a
Step #1 - "CD - Construindo Imagem":  ---> d26f4c862b1a
Step #1 - "CD - Construindo Imagem": Removing intermediate container e19b447b7dd1
Step #1 - "CD - Construindo Imagem":  ---> Running in e19b447b7dd1
Step #1 - "CD - Construindo Imagem": Step 9/9 : CMD ["gohpa"]
Step #1 - "CD - Construindo Imagem":  ---> 5e20dcde5d82
Step #1 - "CD - Construindo Imagem": Removing intermediate container b32aa90bd366
Step #1 - "CD - Construindo Imagem":  ---> Running in b32aa90bd366
Step #1 - "CD - Construindo Imagem": Step 8/9 : EXPOSE 8080
Step #1 - "CD - Construindo Imagem":  ---> 5918dcc3da69
Step #1 - "CD - Construindo Imagem": Step 7/9 : COPY --from=builder /go/src/gohpa/gohpa /usr/bin/gohpa
Step #1 - "CD - Construindo Imagem":  ---> 
Step #1 - "CD - Construindo Imagem": Step 6/9 : FROM scratch
Step #1 - "CD - Construindo Imagem":  ---> e1bad9542884
Step #1 - "CD - Construindo Imagem": Removing intermediate container ebed31bb4904
Step #1 - "CD - Construindo Imagem":  ---> Running in ebed31bb4904
Step #1 - "CD - Construindo Imagem": Step 5/9 : RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gohpa .
Step #1 - "CD - Construindo Imagem":  ---> 44f71aa50d7f
Step #1 - "CD - Construindo Imagem": Removing intermediate container 6ca205536aff
Step #1 - "CD - Construindo Imagem":  ---> Running in 6ca205536aff
Step #1 - "CD - Construindo Imagem": Step 4/9 : WORKDIR /go/src/gohpa
Step #1 - "CD - Construindo Imagem":  ---> 2c23e6531f89
Step #1 - "CD - Construindo Imagem": Step 3/9 : COPY ./src/gohpa/main.go /go/src/gohpa/main.go
Step #1 - "CD - Construindo Imagem":  ---> 881b2afa7def
Step #1 - "CD - Construindo Imagem": Removing intermediate container 1e9f7b4b1d0b
Step #1 - "CD - Construindo Imagem":  ---> Running in 1e9f7b4b1d0b
Step #1 - "CD - Construindo Imagem": Step 2/9 : ARG APP_VERSION=0.1.0
Step #1 - "CD - Construindo Imagem":  ---> 1a87ceb1ace5
Step #1 - "CD - Construindo Imagem": Status: Downloaded newer image for golang:alpine
Step #1 - "CD - Construindo Imagem": Digest: sha256:73182a0a24a1534e31ad9cc9e3a4bb46bb030a883b26eda0a87060f679b83607
Step #1 - "CD - Construindo Imagem": alpine: Pulling from library/golang
Step #1 - "CD - Construindo Imagem": Step 1/9 : FROM golang:alpine as builder

Step #1 - "CD - Construindo Imagem": Sending build context to Docker daemon   7.47MB
Step #1 - "CD - Construindo Imagem": Already have image (with digest): gcr.io/cloud-builders/docker
Starting Step #1 - "CD - Construindo Imagem"
Finished Step #0 - "CI - Rodando Teste"
Step #0 - "CI - Rodando Teste": ok  	gohpa	0.003s
Step #0 - "CI - Rodando Teste": Running: go test gohpa
Step #0 - "CI - Rodando Teste": Documentation at https://github.com/GoogleCloudPlatform/cloud-builders/blob/master/go/README.md
Step #0 - "CI - Rodando Teste": 
Step #0 - "CI - Rodando Teste":                 ***** END OF NOTICE *****
Step #0 - "CI - Rodando Teste": 
Step #0 - "CI - Rodando Teste": For details, visit https://hub.docker.com/_/golang.
Step #0 - "CI - Rodando Teste": 
Step #0 - "CI - Rodando Teste": multiple platforms, are maintained by the Docker Community.
Step #0 - "CI - Rodando Teste": Alternative official `golang` images, including multiple versions across
Step #0 - "CI - Rodando Teste": 
Step #0 - "CI - Rodando Teste":                    ***** NOTICE *****
Step #0 - "CI - Rodando Teste": 
Step #0 - "CI - Rodando Teste": Already have image (with digest): gcr.io/cloud-builders/go
Starting Step #0 - "CI - Rodando Teste"
BUILD
Operation completed over 1 objects/3.7 MiB.                                      
/ [1 files][  3.7 MiB/  3.7 MiB]                                                
/ [0 files][    0.0 B/  3.7 MiB]                                                
Copying gs://251590220208.cloudbuild-source.googleusercontent.com/582677ac7e761d0d81f50b25ab35554211acbe16-3c85ff77-4ef5-4cb1-a7dd-9ebd8ad5a4d9.tar.gz#1598292552183941...
Fetching storage object: gs://251590220208.cloudbuild-source.googleusercontent.com/582677ac7e761d0d81f50b25ab35554211acbe16-3c85ff77-4ef5-4cb1-a7dd-9ebd8ad5a4d9.tar.gz#1598292552183941
FETCHSOURCE

starting build "11606116-bd3e-41ee-b15e-56213c5867eb"