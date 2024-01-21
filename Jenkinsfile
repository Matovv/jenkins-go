pipeline {
    agent any
    triggers {
        pollSCM '* * * * *'
    }
    tools {
        // Specify the Go installation by name
        go 'Go'
    }
    stages {
        stage('Build') {
            steps {
              script {
                    // Use catchError to catch errors during 'go test'
                    catchError(buildResult: 'UNSTABLE') {
                        echo "Building.."
                        sh 'go test ./...'
                        sh 'go build -o myapp .'
                }
              }
        }
        stage('Test') {
            steps {
                echo "Testing.."
                sh '''
                ./myapp
                '''
            }
        }
        stage('Deliver') {
            steps {
                echo 'Deliver....'
                sh '''
                echo "doing delivery stuff.."
                '''
            }
        }
    }
}