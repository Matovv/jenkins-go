pipeline {
    /*
    agent { 
        
        node {
            label 'docker-agent-python'
        }
    }
    */
    agent any
    triggers {
        pollSCM '* * * * *'
    }
    tools {
        // Specify the Go installation by name
        go 'go'
    }
    stages {
        stage('Build') {
            script {
                    // Run 'go test'
                    def testExitCode = sh(script: 'go test ./...', returnStatus: true)

                    // Check the exit code of 'go test'
                    if (testExitCode != 0) {
                        error "Tests failed! Exiting build."
                    }

                    // Continue with other build steps if tests pass
                    echo "Tests passed!"
            }
                /*
            steps {
                echo "Building.."
                sh '''
                go test
                go build -o myapp .
                '''
            }
            */
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