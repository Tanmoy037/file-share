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
	err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file:", err)
        return
    }

    // Get environment variables
    region := os.Getenv("AWS_REGION")
    bucket := os.Getenv("S3_BUCKET")
    accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
    secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

    // Configure the AWS session
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String(region),
        Credentials: credentials.NewStaticCredentials(
            accessKeyID,
            secretAccessKey,
            "",
        ),
    }, nil)
    if err != nil {
        fmt.Println("Error creating AWS session:", err)
        return
    }

    svc := s3.New(sess)

    filePath := "Donuts.png"
    key := "Donuts.png"

    file, err := os.Open(filePath)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error opening file:", err)
        return
    }
    defer file.Close()

    // Upload the file and set it to public-read
    _, err = svc.PutObject(&s3.PutObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(key),
        Body:   file,
        ACL:    aws.String("public-read"), // Make the file publicly accessible
    })
    if err != nil {
        fmt.Println("Error uploading file:", err)
        return
    }

    // Construct the public URL
    publicURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, region, key)

    fmt.Println("File uploaded successfully!")
    fmt.Println("Public URL:", publicURL)
}


// func DownloadHandler


