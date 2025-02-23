# Welcome App - DevOps Technical Test

Aplikasi sederhana yang dibangun menggunakan Golang dan Gin framework, di-deploy menggunakan Kubernetes (k3s) dengan automated CI/CD pipeline menggunakan GitHub Actions.

## Daftar Isi
- [Struktur Proyek](#struktur-proyek)
- [Pre-Required](#pre-required)
- [Menjalankan Aplikasi Secara Lokal](#menjalankan-aplikasi-secara-lokal)
- [API Endpoint](#api-endpoint)
- [Deployment Guide](#deployment-guide)

## Struktur Proyek

```
welcome-app/
├── .github/
│   └── workflows/
│       └── cicd.yml
├── cmd/
│   └── main.go
│─────── internal/
│       ├── handler/
│       │   └── welcome.go
│       └── router/
│           └── router.go
├── deployments/
│   ├── kubernetes/
│   │   ├── base/
│   │   │   ├── configmap.yaml
│   │   │   ├── secret.yaml
│   │   │   ├── deployment.yaml
│   │   │   ├── service.yaml
│   │   │   └── ingress.yaml
│   │   └── overlays/
│   │       ├── development/
│   │       │   ├── kustomization.yaml
│   │       │   └── patch-deployment.yaml
│   │       └── production/
│   │           ├── kustomization.yaml
│   │           └── patch-deployment.yaml
│   └── docker/
│       └── Dockerfile
├── GCP_Plan/
│   └── gcp-architecture.md
├── README.md
└── docker-compose.yaml
```

## Pre-Required

Untuk menjalankan aplikasi ini secara lokal, Anda memerlukan:

1. Go 1.23 atau lebih tinggi
2. Docker
3. kubectl
4. k3s (untuk deployment)
5. Github Actions

## Menjalankan Aplikasi Secara Lokal

1. Clone repository:
```bash
git clone https://github.com/mdafaardiansyah/ptest-sltr-devops.git
cd ptest-sltr-devops
```

2. Menjalankan aplikasi tanpa Docker:
```bash
go run cmd/main.go
```

3. Menjalankan dengan Docker:
```bash
docker build -t welcome-app:latest -f deployments/docker/Dockerfile .
docker run -p 5000:5000 welcome-app:latest
```

4. Menggunakan docker-compose:
```bash
docker compose up
```

## API Endpoint

Aplikasi ini memiliki beberapa endpoint:

1. Welcome Endpoint
    - URL: `/welcome/{nama}`
    - Method: GET
    - Parameter Path: nama (opsional)
    - Response:
        - Jika nama diisi: "Selamat datang {nama}"
        - Jika nama kosong: "Anonymous"
    - Contoh:
      ```bash
      # Dengan nama
      Local
      - curl http://localhost:5000/welcome/Muhammad_Dafa_Ardiansyah
      > Selamat datang Muhammad_Dafa_Ardiansyah
      
      Real URL Deployment:
      - curl https://welcome.ardidafa.glanze.site/Muhammad_Dafa_Ardiansyah
      > Selamat datang Muhammad_Dafa_Ardiansyah
 
      # Tanpa nama
      Local
      - curl http://localhost:5000/welcome
      > Anonymous
      
      Real URL Deployment:
      - curl https://welcome.ardidafa.glanze.site/
      > Anonymous
      ```

2. Health Check
    - URL: `/health`
    - Method: GET
    - Response: `{"status": "ok"}`

## Deployment Guide

### Setup Kubernetes (k3s)

1. Install k3s di server:
```bash
curl -sfL https://get.k3s.io | INSTALL_K3S_EXEC="--tls-san YOUR_SERVER_IP --node-external-ip YOUR_SERVER_IP" sh -
```

2. Simpan kubeconfig:
```bash
cat /etc/rancher/k3s/k3s.yaml | sed "s/127.0.0.1/YOUR_SERVER_IP/g"
```

3. Install cert-manager:
```bash
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.13.3/cert-manager.yaml
```

### Deploy Aplikasi

1. Apply Kubernetes manifests:
```bash
kubectl create namespace welcome-app
kubectl apply -f deployments/kubernetes/base/
```

2. Verifikasi deployment:
```bash
kubectl get pods -n welcome-app
kubectl get svc -n welcome-app
kubectl get ingress -n welcome-app
```

### Setup GitHub Actions

1. Tambahkan secrets berikut di repository GitHub:
    - DOCKERHUB_USERNAME: Username Docker Hub
    - DOCKERHUB_TOKEN: Token akses Docker Hub
    - KUBECONFIG: File kubeconfig ( base64 )
- generate KUBECONFIG :
```bash
"cat /etc/rancher/k3s/k3s.yaml | base64 -w 0"
```
2. Pilih Branch pada menu Workflow Dispatch > Release tag to build and deploy:
```bash
example tag: "v1.0.0"
```
## Troubleshooting

Jika mengalami masalah saat deployment:

1. Periksa status pod:
```bash
kubectl describe pod -n welcome-app
kubectl logs -n welcome-app <pod-name>
```

2. Periksa status ingress:
```bash
kubectl describe ingress -n welcome-app
```

3. Periksa status sertifikat SSL:
```bash
kubectl get certificate -n welcome-app
kubectl describe certificate -n welcome-app
```