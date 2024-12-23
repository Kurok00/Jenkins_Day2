pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                // Checkout the code from the repository
                git 'https://github.com/Kurok00/Jenkins_Day2.git'
            }
        }
        stage('Build') {
            agent {
                docker {
                    image 'golang:1.23.3-alpine'
                    args '-v /var/jenkins_home/go:/go'
                }
            }
            steps {
                // Build the Go application
                sh 'go build -o main .'
            }
        }
        stage('Test') {
            agent {
                docker {
                    image 'golang:1.23.3-alpine'
                    args '-v /var/jenkins_home/go:/go'
                }
            }
            steps {
                // Run tests
                sh 'go test ./...'
            }
        }
        stage('Docker Build') {
            steps {
                // Build the Docker image
                sh 'docker build -t your-image-name .'
            }
        }
        stage('Deploy') {
            steps {
                // Deploy the Docker image
                sh 'docker run -d -p 4001:4001 your-image-name'
            }
        }
    }
    post {
        always {
            // Clean up workspace
            cleanWs()
        }
    }
}