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
        stage('Test') {
            steps {
                script {
                    // Use catchError to catch errors during 'go test'
                    catchError(buildResult: 'UNIT TEST FAILED!') {
                        echo "Testing.."
                        sh 'go test ./...'
                    }
                }
            }
        }
        stage('Build') {
            steps {
                script {
                    // Use catchError to catch errors during 'go test'
                    catchError(buildResult: 'BUILD FAILED!') {
                        echo "Building.."
                        sh 'go build -o myapp .'
                    }
                }
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