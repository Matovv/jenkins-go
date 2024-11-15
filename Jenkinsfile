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
                     echo "Running tests with coverage..."
                    def testExitCode = sh(script: 'go test ./... -coverprofile=coverage.out', returnStatus: true)
                    if (testExitCode != 0) {
                        error('Tests failed!')
                    }
                    // Check the coverage percentage
                    def coveragePercentageRaw = sh(script: 'go tool cover -func=coverage.out | grep total | awk \'{print $3}\'', returnStdout: true).trim()
                    echo "Raw Coverage: ${coveragePercentageRaw}"
                    // Remove the percentage sign and convert to a float
                    def coveragePercentage = coveragePercentageRaw.replaceAll('%', '').toFloat()
                    echo "Coverage: ${coveragePercentage}%"
                    // Fail the build if coverage is below 90%
                    if (coveragePercentage < 90.0) {
                        error("Test coverage is below 90% (${coveragePercentage}%).")
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