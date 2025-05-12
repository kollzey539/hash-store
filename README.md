# 🔐 Hash Service

A lightweight Go-based web service that stores and retrieves string values using their **SHA-256 hash**. Data is stored in an AWS S3 bucket and exposed via a simple HTTPS API.

---

## 🚀 Deployment

The service is designed to run in a **Kubernetes environment**, but can also be started locally or in a Docker container.

### Options:
- ✅ **Kubernetes** (recommended for production)
- ✅ **Docker** for containerized local testing
- ✅ **Direct local run** for development

---

## 🌐 Sample API Usage

### 📥 Store a string

```bash
curl -X POST https://hash-store.rigettidemo.com/hash \
  -H "Content-Type: application/json" \
  -d '{"data":"hello world"}'
```

### 📥 Response

```json
{
"hash": "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
}
```
You can then retrieve the original string using:
```bash
curl https://hash-store.rigettidemo.com/hash/b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
```
## ⚙️ Environment Configuration

The following environment variables **must be set** regardless of the deployment method:

| Variable               | Description                      |
|------------------------|----------------------------------|
| `AWS_ACCESS_KEY_ID`    | Your AWS access key              |
| `AWS_SECRET_ACCESS_KEY`| Your AWS secret key              |
| `AWS_REGION`           | AWS region (e.g. `us-east-1`)    |
| `S3_BUCKET_NAME`       | Name of the S3 bucket for storage |


## 🧭 In Kubernetes
```yaml
env:
  - name: AWS_ACCESS_KEY_ID
    valueFrom:
      secretKeyRef:
        name: aws-creds
        key: AWS_ACCESS_KEY_ID
```

## 🐳  In Docker
Pass variables via the docker run command:
```bash
docker run -e AWS_ACCESS_KEY_ID=xxx -e AWS_SECRET_ACCESS_KEY=xxx \
  -e AWS_REGION=us-east-1 -e S3_BUCKET_NAME=my-bucket hash-store
```
## 🧪  Local Development
Set .env file to the below
```env
AWS_ACCESS_KEY_ID=your_key
AWS_SECRET_ACCESS_KEY=your_secret
AWS_REGION=us-east-1
S3_BUCKET_NAME=my-bucket
```

Uncomment the following in your `main.go` to load local environment variables:
```go
// err := godotenv.Load()
// if err != nil {
//   log.Fatal("Error loading .env file")
// }
```
## 📦   Build & Run Locally
```bash
go build -o hash-store .
./hash-store
```

Then access it at `http://localhost:8080/hash`

## 📁 Endpoints

| Method | Endpoint                                                           | Description                       |
|--------|--------------------------------------------------------------------|-----------------------------------|
| POST   | `https://hash-store.rigettidemo.com/hash`                          | Stores a string, returns the hash |
| GET    | `https://hash-store.rigettidemo.com/hash/{hash}`                  | Retrieves original string         |

---

## ⚙️ CI/CD Pipeline

This project uses GitHub Actions to automate image builds and GitOps-based deployment via Argo CD.

### 📦 What It Does

1. **Builds** a Docker image from the Go application
2. **Tags** the image with:
   - `latest`
   - `build-${{ github.run_id }}`
3. **Pushes** both tags to the Amazon ECR repository
4. **Clones** the external values repository
5. **Updates** the `values.yaml` file with the new image tag
6. **Commits & pushes** the change to `main` branch of the values repo
7. Argo CD automatically picks up the updated values and syncs the deployment

### 🏗️ Image Repository

- **Registry:** `985883769551.dkr.ecr.us-east-1.amazonaws.com`
- **Repository:** `rigettidemo`

### 📁 Values File Updated

- Repository: [`values`](https://github.com/kollzey539/values)
- Path: `hash-store/values.yaml`
- Field Updated: `tag: build-<run_id>`

This ensures that every commit to `main` results in an updated image and fully automated deployment via GitOps.

---

## 👨‍💻 Maintainer

**Kolawole Olowoporoku**  
GitHub: [@kollzey539](https://github.com/kollzey539)

