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
                    echo "Testing.."
                    sh 'go test ./...' || error('Tests failed!')
                
                }
            }
        }
        stage('Build') {
            steps {
                script {
                    echo "Building.."
                    sh 'go build -o myapp .' || error('Build failed!')
                    
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