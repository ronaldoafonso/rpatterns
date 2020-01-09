pipeline {
    agent any
    stages {
        stage ("Unit Test") {
            steps {
                sh "go test -v ./..."
            }
        }
    }
}
