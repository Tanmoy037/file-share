package server

import (
    "fmt"
    "net/http"
    "os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "github.com/joho/godotenv"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    err := godotenv.Load()
    if err != nil {
        http.Error(w, fmt.Sprintf("Error loading .env file: %v", err), http.StatusBadRequest)
        return
    }

    region := os.Getenv("AWS_REGION")
    bucket := os.Getenv("S3_BUCKET")
    accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
    secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

    sess, err := session.NewSession(&aws.Config{
        Region: aws.String(region),
        Credentials: credentials.NewStaticCredentials(
            accessKeyID,
            secretAccessKey,
            "",
        ),
    }, nil)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error creating AWS session: %v", err), http.StatusBadRequest)
        return
    }

    svc := s3.New(sess)

    err = r.ParseMultipartForm(10 << 20) // 10 MB limit
    if err != nil {
        http.Error(w, fmt.Sprintf("Error parsing form data: %v", err), http.StatusBadRequest)
        return
    }

    file, header, err := r.FormFile("file")
    if err != nil {
        http.Error(w, fmt.Sprintf("Error retrieving file: %v", err), http.StatusBadRequest)
        return
    }
    defer file.Close()

    filename := header.Filename

    _, err = svc.PutObject(&s3.PutObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(filename),
        Body:   file,
        // ACL:    aws.String("public-read"),
    })
    if err != nil {
        http.Error(w, fmt.Sprintf("Error uploading file: %v", err), http.StatusBadRequest)
        return
    }

    publicURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, region, filename)

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(publicURL))
}
