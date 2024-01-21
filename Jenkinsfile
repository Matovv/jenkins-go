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
        stage('Lint') {
            steps {
                script {
                    echo "Linting code..."
                    sh 'go get -u github.com/mgechev/revive'
                    // Replace 'golint' with the actual linter command you use
                    def lintExitCode = sh(script: 'revive ./...', returnStatus: true)
                    if (lintExitCode != 0) {
                        error('Linting failed!')
                    }
                }
            }
        }
        stage('Test') {
            steps {
                script {             
                    def testExitCode = sh(script: 'go test ./...', returnStatus: true)
                    if (testExitCode != 0) {
                        error('Tests failed!')
                    }
                }
            }
        }
        stage('Build') {
            steps {
                script {
                    echo "Building.."
                     def testExitCode = sh(script: 'go build -o myapp .', returnStatus: true)
                    if (testExitCode != 0) {
                        error('Build failed!')
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

    post {
        always {
            // Send email notification after build completes
            emailext (
                subject: "Build ${currentBuild.currentResult}",
                body: "Build ${currentBuild.currentResult}: Job ${env.JOB_NAME} - Build #${env.BUILD_NUMBER}\n\n${BUILD_URL}",
                recipientProviders: [culprits(), requestor()],
                attachLog: true,
                compressLog: true,
            )
        }
    }
}