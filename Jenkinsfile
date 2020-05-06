pipeline {
   agent any

   tools {
       go 'Golang 1.13.6'
   }

   stages {
      stage('Checkout') {
          steps {
              git url: 'https://github.com/ronaldoafonso/rpatterns'
          }
      }
      stage('Unit Test') {
          steps {
              sh 'go test -v ./...'
          }
      }
   }
}
